package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type EventType byte

const (
	_                     = iota
	EventDelete EventType = iota
	EventPut
)

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

type FileTransactionLogger struct {
	events       chan<- Event
	errors       <-chan error
	lastSequence uint64
	file         io.ReadWriteCloser
}

func NewFileTransactionLogger(filename string) (*FileTransactionLogger, error) {
	// ReadWrite mode, append only, if not exist, create file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorLogFileOpen.Error(), err)
	}

	return &FileTransactionLogger{
		file: file,
	}, nil
}

func (l *FileTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}

func (l *FileTransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.lastSequence++
			_, err := fmt.Fprintf(l.file, LOG_FORMAT, l.lastSequence, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

func mappingLogToEvent(e *Event, str string) error {
	line := strings.Split(str, "\t")
	sequence, err := strconv.ParseUint(line[0], 10, 64)
	if err != nil {
		return err
	}
	ty, err := strconv.Atoi(line[1])
	if err != nil {
		return err
	}
	e.Sequence, e.EventType, e.Key, e.Value = sequence, EventType(ty), line[2], strings.Join(line[2:], " ")
	return nil
}

func (l *FileTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	scanner := bufio.NewScanner(l.file)
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		var e Event
		defer close(outEvent)
		defer close(outError)

		for scanner.Scan() {
			//if _, err := fmt.Sscanf(line, LOG_FORMAT, &e.Sequence, &e.EventType, &e.Key, &e.Value); err != nil {
			//	outError <- fmt.Errorf("%s: %w", ErrorLogParse, err)
			//	return
			//}
			if err := mappingLogToEvent(&e, scanner.Text()); err != nil {
				outError <- err
				return
			}

			if l.lastSequence >= e.Sequence {
				outError <- ErrorTransactionNumberOrder
			}
			l.lastSequence = e.Sequence

			outEvent <- e
		}

		if err := scanner.Err(); err != nil {
			outError <- fmt.Errorf("%s: %w", ErrorLogRead, err)
			return
		}
	}()

	return outEvent, outError
}
