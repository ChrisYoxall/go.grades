package main

import (
	"context"
	"fmt"
	"grades/log"
	"grades/portal"
	"grades/registry"
	"grades/service"
	stlog "log"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceURL = fmt.Sprintf("http://%v:%v", host, port)
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(), r, host, port, portal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	// Try to get the log service
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Log service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)

		// Trying to use log service
		//log.Debug("Log service found.")
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal service.")
}
