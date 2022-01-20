package main

import (
	"context"
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
	flag.StringVar(&configPath, "config-path", "configs/config.env", "path to apiserver config file")
}

func main() {
	flag.Parse()

	loadConfig, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("error init configuration file: %s", err)
	}

	router := gin.Default()

	//Initialize composites

	postgreSQLComposite, err := composites.NewPostgreSQLComposite(context.TODO(), loadConfig)

	userComposite, err := composites.NewUserComposite(postgreSQLComposite)
	if err != nil {
		log.Fatal("user composite failed: ", err)
	}
	userComposite.Handler.Register(router)

	srv := handlers.NewServer(loadConfig.Listen.Port, router)
	if err := srv.Run(); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}
