package main

import (
	"flag"
	"log"
	"web-11/internal/auth/api"
	"web-11/internal/auth/config"
	"web-11/internal/auth/provider"
	"web-11/internal/auth/usecase"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "../../configs/auth_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	uc := usecase.NewUsecase(prv)

	srv := api.NewServer(cfg.IP, cfg.Port, uc)

	log.Printf("Сервер Auth запущен на %s\n", srv.Address)
	if err := srv.Router.Start(srv.Address); err != nil {
		log.Fatal(err)
	}
}
