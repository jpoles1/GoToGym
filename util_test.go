package main

import (
	"errors"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	loadEnv()
}

func TestErrCheck(t *testing.T) {
	errCheck("Testing Err Check Function!", errors.New("Test error"))
}

func TestSendAlert(t *testing.T) {
	sendAlert(true, "Testing high importance alerts!", "Testing high importance alerts!")
}
