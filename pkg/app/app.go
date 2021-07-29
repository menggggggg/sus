package app

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type App struct {
	app *http.Server
}

func New() *App {
	return &App{
		app: &http.Server{
			Addr:    "",
			Handler: nil,
		},
	}
}

func (a *App) Run() error {
	// https
	// if "CertFile" != "" && "KeyFile" != "" {
	// 	log.Infof("shorten url server listen at %s with tls", ":8080")
	// 	return a.app.ListenAndServeTLS("CertFile", "KeyFile")
	// }

	// http
	log.Infof("shorten url server listen at %s", ":8080")
	return a.app.ListenAndServe()
}

func (a *App) Close(ctx context.Context) error {
	return a.app.Shutdown(ctx)
}
