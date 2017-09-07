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
	r, err := GetResources(cfg)
	if err != nil {
		log.Printf("ABORTING: Failed to get resources: %v\n", err)
		return
	}

	// Creating my own server var to have access to server.Shutdown()
	m, err := MakeMux(cfg, r)
	if err != nil {

	}
	server := &http.Server{Addr: cfg.HTTPPort, Handler: m}
	r.Log.Record("Starting server on port %s", cfg.HTTPPort)
	dn := make(chan byte)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				r.Log.ServerErr("Listen and Serve Error: %v", err)
			}
			dn <- 0
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ch:
		r.Log.NewLine()
		r.Log.Record("Termination signal recieved, stopping server...")
		ctx := context.TODO()
		err := server.Shutdown(ctx)
		if err != nil {
			r.Log.ServerErr("shutdown failure: %v", err)
		}
	case <-dn:
		r.Log.NewLine()
		r.Log.Record("Exiting program...")
	}
}
