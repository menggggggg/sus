package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/menggggggg/sus/pkg/app"
	"github.com/menggggggg/sus/tools"
	log "github.com/sirupsen/logrus"
)

var help = flag.Bool("h", false, "show help info")
var version = flag.Bool("v", false, "show version info")

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("server name: %s, version: %s", tools.Name, tools.Version)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	_app := app.New()

	errChan := make(chan error)
	go func() {
		errChan <- _app.Run()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		if err != nil {
			log.Error("run server faild, error: ", err)
		}
	case sig := <-signals:
		log.WithField("signal", sig.String()).Error("receive signal")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		log.WithField("error", _app.Close(ctx)).Error("sus server closed")
	}
}
