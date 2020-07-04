package main

import (
	"os/signal"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/whuangz/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	ph := handlers.NewProducs(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err : = s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	sig := <-sigChan
	l.Println("Recieved terminate, gratefully shutdonw", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
