package main

import (
	"demoApi/internal/app/adapters/handlers"
	"demoApi/internal/app/composites"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to apiserver config file")
}

func initConfig() error {
	flag.Parse()
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config initialize: %s", err.Error())
	}

	router := gin.Default()

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
