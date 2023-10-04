// Package service provides generic functionality to start a web service and shut it down gracefully.
package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Start(ctx context.Context, serviceName string, port string, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, serviceName, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName string, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		sigCh := make(chan os.Signal)
		defer close(sigCh)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		fmt.Printf("%v started. Pres ctrl+c to stop.\n", serviceName)
		<-sigCh

		err := srv.Shutdown(ctx)
		if err != nil {
			panic(err)
		}
	}()

	return ctx
}
