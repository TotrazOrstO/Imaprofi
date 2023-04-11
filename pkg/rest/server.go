package rest

import (
	"fmt"
	"github.com/TotrazOrstO/Imapro/pkg/config"
	"net/http"

	"github.com/TotrazOrstO/Imapro/pkg/cache"
	"github.com/TotrazOrstO/Imapro/pkg/centrifugo"
	"github.com/TotrazOrstO/Imapro/pkg/database"
	"github.com/TotrazOrstO/Imapro/pkg/grpc"
	"github.com/TotrazOrstO/Imapro/pkg/rabbitmq"
	"github.com/go-chi/chi"
)

func Run() error {
	router := chi.NewRouter()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("could not load config: %v", err)
	}

	// Инициализация компонентов
	db, err := database.InitPostgres(cfg.Postgres)
	if err != nil {
		return err
	}
	defer db.Close()

	cache, err := cache.InitRedis(cfg.Redis)
	if err != nil {
		return err
	}
	defer cache.Close()

	mq, err := rabbitmq.InitRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		return err
	}
	defer mq.Close()

	centrifugoClient, err := centrifugo.InitClient(cfg.CentrifugoConfig)
	if err != nil {
		return err
	}

	grpcClient, err := grpc.InitClient(cfg.GRPCServer)
	if err != nil {
		return err
	}

	// Регистрация обработчиков
	router.Get("/some-route", func(w http.ResponseWriter, r *http.Request) {
		// Ваш код здесь
	})

	return http.ListenAndServe(":8080", router)
}
