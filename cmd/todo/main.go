package main

import (
	"context"
	"go-todo/internal/handler"
	"go-todo/internal/repository"
	"go-todo/internal/service"
	"go-todo/pkg/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(
		repository.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	)
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
			logrus.Fatalf("failed to run server: %s", err.Error())
		}
	}()

	logrus.Println("Todo app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Todo app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("local")

	return viper.ReadInConfig()
}
