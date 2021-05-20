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

	"github.com/cauabernardino/bucket-list/config"
	"github.com/cauabernardino/bucket-list/db"
	"github.com/cauabernardino/bucket-list/handlers"
)

func main() {
	config.LoadEnvs()
	addr := fmt.Sprintf(":%d", config.API_PORT)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Conn.Close()

	httpHandler := handlers.NewHandler(db)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)

	log.Printf("Started server on %s", addr)

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

// Stop handles timeout of server
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
