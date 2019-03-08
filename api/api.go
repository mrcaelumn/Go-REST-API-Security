//go:generate goagen bootstrap -d github.com/mrcaelumn/Go-REST-API-Security/api/design

package api

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
	"github.com/mrcaelumn/Go-REST-API-Security/api/app"
	"github.com/mrcaelumn/Go-REST-API-Security/api/controller"
	"github.com/tylerb/graceful"
)

func Run(ctx context.Context, listener net.Listener, log log15.Logger) error {
	// Create service
	service := goa.New("go-rest-security")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "getToken" controller
	c1 := controller.NewActionController(service)
	app.MountActionController(service, c1)
	// Mount "swagger" controller
	c2 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "version" controller
	c3 := controller.NewVersionController(service)
	app.MountVersionController(service, c3)

	var err error
	// Start service
	server := &graceful.Server{
		NoSignalHandling: true,
		Server: &http.Server{
			Handler: service.Mux,
		},
	}

	c := make(chan error, 1)
	go func() {
		c <- server.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		server.Stop(time.Duration(3) * time.Second)
		<-server.StopChan()
		// draining the channel
		<-c
	case err := <-c:
		if err != nil {
			log15.Error(err.Error())
		}
	}

	return err
}
