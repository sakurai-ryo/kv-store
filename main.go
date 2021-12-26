package main

import "log"

func main() {
	logger, err := InitFileTransactionLogger()
	if err != nil {
		log.Fatalf("failed to replay transaction log: %v\n", err)
	}
	handler := NewHandler(logger)
	StartServer(handler)
}
