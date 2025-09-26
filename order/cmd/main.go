package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	orderAPI "github.com/rocket-crm/order/internal/api/order/v1"
	orderRepository "github.com/rocket-crm/order/internal/repository/order"
	orderService "github.com/rocket-crm/order/internal/service/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort          = "8080"
	inventoryAddress  = "localhost:50051"
	paymentAddress    = "localhost:50052"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

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

	repository := orderRepository.NewRepository()
	service := orderService.NewService(repository, inventoryClient, paymentClient)
	api := orderAPI.NewAPI(service)

	ordersServer, err := ordersV1.NewServer(api)
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
