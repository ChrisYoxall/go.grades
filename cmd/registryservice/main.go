package main

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

func main() {
	registry.SetupRegistryService()
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		sigCh := make(chan os.Signal)
		defer close(sigCh)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		fmt.Println("Registry service started. Press ctrl+c to stop.")
		<-sigCh
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
