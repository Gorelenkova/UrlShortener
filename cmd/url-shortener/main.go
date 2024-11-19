package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shorter/internal/config"
	"url-shorter/internal/lib/logger/sl"
	"url-shorter/internal/storage/sqlite"
)
const(
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
    cfg := config.MustLoad()
    fmt.Printf("Loaded config: %+v\n", cfg)
	log := setupLoger(cfg.Env)
	log.Info("starting url shortener", slog.String("env",cfg.Env))
	log.Debug("debug message are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil{
		log.Error("failed to init storage",sl.Err(err))
		os.Exit(1)
	}
	_ = storage
}

func setupLoger(env string) *slog.Logger{
	var log *slog.Logger
	switch env{
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelDebug}),)
	case envDev: 
	log = slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelDebug}),)
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelInfo}),)
	}
	return log
}