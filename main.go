package main

import (
	"context"
	"go-microservices/handlers"
	"go-microservices/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := models.NewValidation()

	productsHandler := handlers.NewProducts(l, v)

	// using Gorilla serve mux router
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandler.GetAllProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.DeleteProduct)

	// handling timeouts and graceful shutdown
	server := &http.Server{
		Addr:         "8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// to prevent blocking
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// to shutdown gracefully, waiting max 30 sec for current operations to complete
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
