package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort          = "8081"
	inventoryAddress  = "localhost:50051"
	paymentAddress    = "localhost:50052"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
	PendingPayment    = "PENDING_PAYMENT"
	PAID              = "PAID"
	CANCELLED         = "CANCELLED"
)

type OrdersStorage struct {
	mu     sync.RWMutex
	orders map[string]*ordersV1.Order
}

func NewOrdersStorage() *OrdersStorage {
	return &OrdersStorage{
		orders: make(map[string]*ordersV1.Order),
	}
}

func (s *OrdersStorage) Create(req *ordersV1.CreateOrderRequest, totalPrice float32) *ordersV1.Order {
	order := &ordersV1.Order{
		OrderUUID:  uuid.NewString(),
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		TotalPrice: totalPrice,
		Status:     PendingPayment,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders[order.OrderUUID] = order
	return order
}

func (s *OrdersStorage) GetByUuid(uuid string) *ordersV1.Order {
	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[uuid]
	if !ok {
		return nil
	}
	return order
}

type OrdersHandler struct {
	storage         *OrdersStorage
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient
}

func NewOrdersHandler(storage *OrdersStorage, iClient inventoryV1.InventoryServiceClient, pClient paymentV1.PaymentServiceClient) *OrdersHandler {
	return &OrdersHandler{storage, iClient, pClient}
}

func (h *OrdersHandler) CreateOrder(ctx context.Context, req *ordersV1.CreateOrderRequest) (ordersV1.CreateOrderRes, error) {
	resParts, err := h.inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{Filter: &inventoryV1.PartsFilter{Uuids: req.PartUuids}})
	if err != nil {
		return nil, err
	}
	if len(resParts.Parts) != len(req.PartUuids) {
		return nil, errors.New("найдены не все запчасти, проверьте uuid деталей")
	}
	var totalPrice float32

	for _, p := range resParts.Parts {
		totalPrice += p.Price
	}

	o := h.storage.Create(req, totalPrice)

	return &ordersV1.CreateOrderResponse{OrderUUID: o.OrderUUID, TotalPrice: o.TotalPrice}, nil
}

func (h *OrdersHandler) GetOrderByUuid(ctx context.Context, params ordersV1.GetOrderByUuidParams) (ordersV1.GetOrderByUuidRes, error) {
	order := h.storage.GetByUuid(params.OrderUUID)
	if order == nil {
		return &ordersV1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: "Заказ с uuid " + params.OrderUUID + " не найден!",
		}, nil
	}
	return order, nil
}

func (h *OrdersHandler) PayOrder(ctx context.Context, req *ordersV1.PayOrderRequest, params ordersV1.PayOrderParams) (ordersV1.PayOrderRes, error) {
	order := h.storage.GetByUuid(params.OrderUUID)
	if order == nil {
		return &ordersV1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: "Заказ с uuid " + params.OrderUUID + " не найден!",
		}, nil
	}

	resp, err := h.paymentClient.PayOrder(ctx, &paymentV1.PayOrderRequest{UserUuid: order.UserUUID, OrderUuid: order.OrderUUID, PaymentMethod: paymentV1.PaymentMethod(paymentV1.PaymentMethod_value[req.PaymentMethod])})
	if err != nil {
		return nil, err
	}

	order.Status = PAID
	order.TransactionUUID = ordersV1.NewOptString(resp.TransactionUuid)
	order.PaymentMethod = ordersV1.NewOptString(req.PaymentMethod)
	return &ordersV1.PayOrderResponse{TransactionUUID: order.TransactionUUID.Value}, nil
}

func (h *OrdersHandler) CancelOrder(ctx context.Context, params ordersV1.CancelOrderParams) (ordersV1.CancelOrderRes, error) {
	order := h.storage.GetByUuid(params.OrderUUID)
	if order == nil {
		return &ordersV1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: "Заказ с uuid " + params.OrderUUID + " не найден!",
		}, nil
	}

	if order.Status == PAID {
		return &ordersV1.ConflictError{}, nil
	}

	order.Status = CANCELLED
	return &ordersV1.CancelOrderNoContent{}, nil
}

func (h *OrdersHandler) NewError(_ context.Context, err error) *ordersV1.GenericErrorStatusCode {
	return &ordersV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: ordersV1.GenericError{
			Code:    ordersV1.NewOptInt(http.StatusInternalServerError),
			Message: ordersV1.NewOptString(err.Error()),
		},
	}
}

func main() {
	// создает коннект до микросервиса inventory
	connInventory, err := grpc.NewClient(inventoryAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}

	// Закрываем соединение после отключения order сервиса, чтобы не было зависшего соединение, которое уже не нужно
	defer func() {
		if cerr := connInventory.Close(); cerr != nil {
			log.Printf("failed to close connect: %v", cerr)
		}
	}()

	// создаем коннект до микросервиса payment
	connPayment, err := grpc.NewClient(paymentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}

	// Закрываем соединение после отключения order сервиса, чтобы не было зависшего соединение, которое уже не нужно
	defer func() {
		if cerr := connInventory.Close(); cerr != nil {
			log.Printf("failed to close connect: %v", cerr)
		}
	}()

	// оборачивает коннект чтобы у него были методы gRPC сервера, чтобы можно было вызывать просто как метод структуры в go
	inventoryClient := inventoryV1.NewInventoryServiceClient(connInventory)
	// оборачивает коннект чтобы у него были методы gRPC сервера, чтобы можно было вызывать просто как метод структуры в go
	paymentClient := paymentV1.NewPaymentServiceClient(connPayment)

	storage := NewOrdersStorage()
	ordersHandler := NewOrdersHandler(storage, inventoryClient, paymentClient)

	ordersServer, err := ordersV1.NewServer(ordersHandler)
	if err != nil {
		log.Printf("ошибка создания OpenAPI: %v", err)
		return
	}
	// Инициализируем роутер Chi
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", ordersServer)

	// Запускаем HTTP-сервер
	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout, // Защита от Slowloris атак - тип DDoS-атаки, при которой
		// атакующий умышленно медленно отправляет HTTP-заголовки, удерживая соединения открытыми и истощая
		// пул доступных соединений на сервере. ReadHeaderTimeout принудительно закрывает соединение,
		// если клиент не успел отправить все заголовки за отведенное время.
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
