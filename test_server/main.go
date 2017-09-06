package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := GetConfig()
	if err != nil {
		log.Printf("ABORTING: Failed to get configuration options: %v\n", err)
		return
	}

	if LOG, err = NewLogger(cfg); err != nil {
		log.Printf("ABORTING: Failed to initialize logging: %v\n", err)
		return
	}
	// Creating my own server var to have access to server.Shutdown()
	m := MakeMux(cfg)
	server := &http.Server{Addr: cfg.HTTPPort, Handler: m}
	LOG.Record("Starting server on port %s", cfg.HTTPPort)
	dn := make(chan byte)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				LOG.ServerErr("Listen and Serve Error: %v", err)
			}
			dn <- 0
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ch:
		LOG.NewLine()
		LOG.Record("Termination signal recieved, stopping server...")
		ctx := context.TODO()
		err := server.Shutdown(ctx)
		if err != nil {
			LOG.ServerErr("shutdown failure: %v", err)
		}
	case <-dn:
		LOG.NewLine()
		LOG.Record("Exiting program...")
	}
}
