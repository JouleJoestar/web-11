package main

import (
	"flag"
	"log"
	"web-11/internal/count/api"
	"web-11/internal/count/config"
	"web-11/internal/count/provider"
	"web-11/internal/count/usecase"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "../../configs/count_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, use)

	log.Printf("Сервер запущен на %s\n", srv.Address)
	srv.Run()
}
