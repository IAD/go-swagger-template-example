package main

import (
	"log"
	"os"

	"github.com/IAD/go-swagger-template-example/internal/service"
	"github.com/go-openapi/swag"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)

	logger := logrus.NewEntry(logrus.New())

	port, err := swag.ConvertInt64(os.Getenv("PORT"))
	if err != nil {
		logrus.Errorf("unable to parse port. Err: %s", err)
	}

	if port == 0 {
		port = 8080
	}

	server, err := service.PrepareServer(port, logger)
	if err != nil {
		logrus.Fatalf("unable to prepare service. Err %s", err)
	}

	log.Printf("Starting server on the port %v", port)

	err = server.Serve()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
