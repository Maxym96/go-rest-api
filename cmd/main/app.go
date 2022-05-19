package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-app/internal/config"
	event "rest-app/internal/event/db"
	"rest-app/internal/user"
	"rest-app/package/client/postgresql"
	"rest-app/package/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("created router")
	router := httprouter.New()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.Background(), cfg.Storage)
	if err != nil {
		logger.Fatal(err)
	}
	repository := event.NewRepository(postgreSQLClient, logger)

	/*	e := event2.Event{
		ID:          "08f41ec7-c4de-4e62-b8e9-0cb6702859bf",
		Name:        "Test111111",
		Description: "Test-Test1111111",
		DateAndTime: "2022-08-08 00:00:00",
	}*/

	err = repository.Delete(context.Background(), "e93d8f5a-eec2-4e64-a16f-c0393f2f663a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)

}
func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	log.Println("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detected app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")

		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)

	} else {
		logger.Info("listen tcp ")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Fatal(server.Serve(listener))

}
