// Package service provides generic functionality to start a web service and shut it down gracefully.
package service

import (
	"context"
	"fmt"
	"grades/registry"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Start(ctx context.Context, reg registry.Registration, host string, port string, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)

	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host string, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	// Start the service in a separate goroutine, so it doesn't block. Handle case where service doesn't start.
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// Handle graceful shutdowns.
	go func() {
		sigCh := make(chan os.Signal)
		defer close(sigCh)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		fmt.Printf("%v started at http://%v:%v. Press ctrl+c to stop.\n", serviceName, host, port)
		<-sigCh

		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		err = srv.Shutdown(ctx)
		if err != nil {
			panic(err)
		}
	}()

	return ctx
}
