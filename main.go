package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

type service struct{}

func main() {
	srv := new(service)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt)
	signal.Notify(quitCh, syscall.SIGTERM)

	server := fasthttp.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      newRouter(srv).Handler,
	}

	log.Printf("Service started")
	go server.ListenAndServe(":8080")

	// Reciving signal
	<-quitCh
	log.Printf("Gracefully terminating service...")
	if err := server.Shutdown(); err != nil {
		panic(err)
	}
	// TODO: wait for goroutines done -> use waitGroup
}
