package main

import (
	"github.com/TotrazOrstO/Imapro/pkg/config"
	"github.com/TotrazOrstO/Imapro/pkg/database"
	"log"
	"net"
	"net/http"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Инициализация компонентов
	postgresDB, err := database.InitPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("Error initializing Postgres: %v", err)
	}

	userRepository := repository.NewPostgresUserRepository(postgresDB)
	userUsecase := usecase.NewUserUsecase(userRepository)

	// Инициализация и запуск HTTP сервера
	httpServer := http.NewServer(userUsecase)
	httpHandler := httpServer.SetupRoutes()

	go func() {
		log.Println("Starting HTTP server on :8080")
		log.Fatal(http.ListenAndServe(":8080", httpHandler))
	}()

	// Инициализация и запуск gRPC сервера
	grpcServer := grpc.NewServer(userUsecase)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on :9090: %v", err)
	}

	log.Println("Starting gRPC server on :9090")
	log.Fatal(grpcServer.Serve(lis))
}
