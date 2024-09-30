package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/farhanfatur/grpc-node-to-go/controller"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	logger := log.New(os.Stdout, "|service-go|", log.LstdFlags)
	err = godotenv.Load("../../common/config/.env")
	if err != nil {
		log.Fatal(err)
	}

	productController := controller.NewProduct(logger)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/login", productController.Login)

	host := os.Getenv("GO_HOST")
	port := os.Getenv("GO_PORT")
	// fmt.Println(host, port)

	server := &http.Server{
		Addr:         host + ":" + port,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		logger.Printf("Starting server on port %s:%s\n", host, port)
		err = server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	signal := <-signalChan
	logger.Println("Received terminate signal, shutdown: ", signal)

	tContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tContext)
}
