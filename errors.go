package main

import "errors"

var (
	ErrorNoSuchKey         = errors.New("no such key")
	ErrorDataLimitExceeded = errors.New("the maximum data size is 4KB")
)
