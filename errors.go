package main

import "errors"

var (
	ErrorNoSuchKey              = errors.New("no such key")
	ErrorDataLimitExceeded      = errors.New("the maximum data size is 4KB")
	ErrorTransactionNumberOrder = errors.New("transaction numbers out of sequence")
	ErrorLogFileOpen            = errors.New("cannot open log file")
	ErrorLogParse               = errors.New("input parse error")
	ErrorLogRead                = errors.New("transaction log read failure")
	ErrorCreateLogger           = errors.New("failed to create event logger")
)
