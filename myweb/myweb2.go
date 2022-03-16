package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//http://pub.funshion.com:1341/service
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	start := time.Now()
	port := "1341"
	srv := http.Server{
		Addr:    ":" + port,
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/", handler)

	//gracefully shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		data := <-sigint
		log.Printf("received signal: " + data.String())
		log.Printf("start to shutdown...")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("HTTP server Shutdown: %v", err)
		}
	}()

	//listen
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	log.Printf("listen on %s", port)
	log.Printf("time elapse: %d ms", time.Now().Sub(start).Milliseconds())

	//serve
	err = srv.Serve(ln)

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	} else {
		log.Printf("successfully shutdown: %v", err)
	}
}
