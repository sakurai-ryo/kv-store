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

func InitTransactionLog() error {
	var err error

	logger, err := NewFileTransactionLogger(LOG_FILE)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrorCreateLogger.Error(), err)
	}

	events, errors := logger.ReadEvents()
	e, ok := Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case EventDelete:
				err = Delete(e.Key)
			case EventPut:
				err = Put(e.Key, e.Value)
			}
		}
	}

	logger.Run()

	return err
}

func InitServer() {
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
