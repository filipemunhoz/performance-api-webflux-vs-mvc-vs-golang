package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"performance-go/handlers"
	"time"
)

func main() {

	cl := log.New(os.Stdout, "client", log.LstdFlags)
	pl := log.New(os.Stdout, "product-api", log.LstdFlags)

	ch := handlers.NewClient(cl)
	ph := handlers.NewProducts(pl)

	sm := http.NewServeMux()
	sm.Handle("/products", ph)
	sm.Handle("/performance-go", ch)

	s := &http.Server{
		Addr:         ":8083",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {

		err := s.ListenAndServe()
		if err != nil {
			cl.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	cl.Println("Receive terminate, shuting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
