package main

import (
	"github.com/LLlE0/lite_cloud_storage/pkg/dbworker"
	"github.com/LLlE0/lite_cloud_storage/pkg/handler"
	"github.com/LLlE0/lite_cloud_storage/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error while initializing the config file: %s", err.Error())
	}

	services := service.NewService(viper.GetString("ip"), viper.GetString("port"))
	DBInstance := dbworker.NewDBInstance()
	srv := new(service.Server)
	handlers := handler.NewHandler(services, srv, DBInstance)
	services.RunApp()
	if err := srv.Run(viper.GetString("port"), viper.GetString("ip"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
