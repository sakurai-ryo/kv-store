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

func InitFileTransactionLogger() (*FileTransactionLogger, error) {
	var err error

	logger, err := NewFileTransactionLogger(LOG_FILE)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorCreateLogger.Error(), err)
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
	if err != nil {
		return nil, err
	}

	logger.Run()

	return logger, nil
}

func StartServer(handler StoreHandler) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	router := NewRouter(handler)

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
