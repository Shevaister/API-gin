package main

import (
	"API"
	"API/pkg/handler"
	"API/pkg/repository"
	"API/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to initialise configs: %s", err.Error())
	}

	db, err := repository.NewSQLDB(repository.Config{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "",
		DBName:   "commentsdb",
		//SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialise database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(API.Server)
	if err := srv.Run(viper.GetString(viper.GetString("8000")), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
