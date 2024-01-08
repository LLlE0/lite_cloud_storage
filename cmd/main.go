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
	//init service
	services := service.NewService(viper.GetString("ip"), viper.GetString("port"))
	//init db instance
	DBInstance := dbworker.NewDBInstance()
	//init server
	srv := new(service.Server)
	//create handler
	handlers := handler.NewHandler(services, srv, DBInstance)
	//run the application
	services.RunApp()
	if err := srv.Run(viper.GetString("port"), viper.GetString("ip"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server: %s", err.Error())
	}
}

// handle config file
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
