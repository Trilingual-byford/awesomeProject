package main

import (
	"awesomeProject/handlers"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Init http")
	logger := log.New(os.Stdout, "awesome-api", log.LstdFlags)
	//hello := handlers.NewHello(logger)
	products := handlers.NewProducts(logger)
	//http.HandleFunc("/",hello.ServeHTTP)
	router := mux.NewRouter()
	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", products.GetProducts)
	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/product", products.AddProduct)
	postRouter.Use(products.MiddleWareValidateProduct)
	putRouter := router.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/product/{id:[0-9]+}", products.UpdateProduct)
	//router.Handle("/product",products.GetProducts)
	s := &http.Server{
		Addr:        ":9090",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()
	mChan := make(chan os.Signal)
	signal.Notify(mChan, os.Interrupt)
	signal.Notify(mChan, os.Kill)
	// Block until a signal is received.
	sig := <-mChan
	log.Println("Got signal:", sig)
	log.Println("Recieved terminate,graceful shutdown", sig)
	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeout)
}
