package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vadim-dmitriev/sittme/common"

	"github.com/vadim-dmitriev/sittme/app"
	"github.com/valyala/fasthttp"
)

func main() {
	cfg, err := common.NewConfig()
	if err != nil {
		log.Printf("Can`t get config: %s", err.Error())
		return
	}

	srv := app.New(cfg)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt)
	signal.Notify(quitCh, syscall.SIGTERM)

	server := fasthttp.Server{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		Handler:      srv.Handler(),
	}

	serverPort := cfg.Server.Port
	log.Printf("Service started on localhost:%d", serverPort)

	listenAndServeErrChan := make(chan error, 1)

	go func() {
		listenAndServeErrChan <- server.ListenAndServe(fmt.Sprintf("localhost:%d", serverPort))
	}()

	select {
	case err := <-listenAndServeErrChan:
		log.Printf("Server error: %s", err.Error())

	case <-quitCh:
		log.Printf("Gracefully terminating server...")
		if err := server.Shutdown(); err != nil {
			panic(err)
		}
	}

	log.Printf("Server stoped.")

}
