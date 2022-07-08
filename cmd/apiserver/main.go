package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dvd-denis/legko-server/internal/app/apiserver"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	godotenv.Load()
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	config.BindAddr = ":" + os.Getenv("PORT")
	config.Store.DatabaseURL = os.Getenv("DATABASE_URL")

	s := apiserver.New(config)
	go func() {
		if err := s.Start(); err != nil {
			logrus.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Stop(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}
