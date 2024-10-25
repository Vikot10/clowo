package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	milvus "github.com/milvus-io/milvus-sdk-go/v2/client"
	"go.uber.org/zap"

	"github.com/Vikot10/clowo/internal/config"
	"github.com/Vikot10/clowo/internal/logger"
)

func main() {
	errCfg := config.Load()
	if errCfg != nil {
		log.Printf("config load error: %v", errCfg)
		os.Exit(1)
	}

	logger.Init(config.Get().Debug)

	err := run()
	if err != nil {
		logger.Error("run error: %v", zap.Error(err))
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	client, err := milvus.NewClient(ctx, milvus.Config{
		Address: config.Get().Milvus.Address,
	})
	if err != nil {
		return err
	}
	defer client.Close()

	client.HasCollection(context.Background(), "YOUR_COLLECTION_NAME")
}
