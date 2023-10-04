package main

import (
	"context"
	"fmt"
	"grades/log"
	"grades/service"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	port := "4000"

	ctx, err := service.Start(context.Background(), "Log Service", port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
