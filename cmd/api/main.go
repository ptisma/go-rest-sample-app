package main

import (
	"github.com/joho/godotenv"
	"github.com/ptisma/go-rest-sample-app/cmd/api/router"
	"github.com/ptisma/go-rest-sample-app/pkg/application"
	"github.com/ptisma/go-rest-sample-app/pkg/exithandler"
	"github.com/ptisma/go-rest-sample-app/pkg/logger"
	"github.com/ptisma/go-rest-sample-app/pkg/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Info.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		logger.Error.Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server at %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			logger.Error.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
