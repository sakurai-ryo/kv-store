package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	router := NewRouter()

	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(PORT), 10),
		Handler: router,
	}

	// Graceful Shutdown
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
	}()

	fmt.Printf("start receiving at :%d\n", PORT)
	log.Fatal(server.ListenAndServe())
}
