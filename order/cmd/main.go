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
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	inventoryV1 "github.com/rocket-crm/inventory"
)

const (
	httpPort          = "8081"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
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

type OrdersHandler struct {
	storage *OrdersStorage
}

func NewOrdersHandler(storage *OrdersStorage) *OrdersHandler {
	return &OrdersHandler{storage}
}

func (h *OrdersHandler) CreateOrder(ctx context.Context, req *ordersV1.CreateOrderRequest) (ordersV1.CreateOrderRes, error) {
	inventoryV1.
	return nil, nil
}

func (h *OrdersHandler) GetOrderByUuid(ctx context.Context, params ordersV1.GetOrderByUuidParams) (ordersV1.GetOrderByUuidRes, error) {
	return nil, nil
}

func (h *OrdersHandler) PayOrder(ctx context.Context, req *ordersV1.PayOrderRequest, params ordersV1.PayOrderParams) (ordersV1.PayOrderRes, error) {
	return nil, nil
}

func (h *OrdersHandler) CancelOrder(ctx context.Context, params ordersV1.CancelOrderParams) (ordersV1.CancelOrderRes, error) {
	return nil, nil
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
	storage := NewOrdersStorage()
	ordersHandler := NewOrdersHandler(storage)

	ordersServer, err := ordersV1.NewServer(ordersHandler)
	if err != nil {
		log.Fatalf("ошибка создания OpenAPI: %v", err)
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
