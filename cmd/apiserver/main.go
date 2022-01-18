package main

import (
	"demoApi/internal/app/adapters/handlers"
	"demoApi/internal/app/composites"
	"demoApi/internal/app/config"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to apiserver config file")
}

func main() {
	flag.Parse()

	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("error init configuration file: %s", err)
	}

	router := gin.New()

	userComposite, err := composites.NewUserComposite()
	if err != nil {
		log.Fatal("user composite failed: ", err)
	}
	userComposite.Handler.Register(router)

	srv := handlers.NewServer("8080", router)
	if err := srv.Run(); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}
