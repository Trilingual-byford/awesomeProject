package main

import (
	"awesomeProject/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
func main() {
	fmt.Println("Init http")
	logger := log.New(os.Stdout, "awesome-api", log.LstdFlags)
	hello := handlers.NewHello(logger)
	products := handlers.NewProducts(logger)
	//http.HandleFunc("/",hello.ServeHTTP)
	mux := http.NewServeMux()
	mux.Handle("/",hello)
	mux.Handle("/product",products)
	s := &http.Server{
		Addr:        ":9090",
		Handler:     mux,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		err :=s.ListenAndServe()
		if err!=nil{
			logger.Fatal(err)
		}
	}()
	mChan := make(chan os.Signal)
	signal.Notify(mChan,os.Interrupt)
	signal.Notify(mChan,os.Kill)
	// Block until a signal is received.
	sig := <-mChan
	log.Println("Got signal:", sig)
	log.Println("Recieved terminate,graceful shutdown",sig)
	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeout)
}
