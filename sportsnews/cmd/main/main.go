package main

import (
	"context"

	"github.com/think-free/sportsnews/lib/logging"
	"github.com/think-free/sportsnews/sportsnews/internal/api"
	"github.com/think-free/sportsnews/sportsnews/internal/cliparams"
	"github.com/think-free/sportsnews/sportsnews/internal/config"
	"github.com/think-free/sportsnews/sportsnews/internal/database"
)

func main() {
	ctx := context.Background()
	c := config.New()
	cp := cliparams.New()

	logging.Init(cp.LogLevel)

	db := database.New(ctx, cp, c)
	srv := api.New(ctx, cp, db)
	err := srv.Run()
	if err != nil {
		logging.L(ctx).Fatal("error running service")
	}
}
