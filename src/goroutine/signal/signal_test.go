package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignal(t *testing.T) {
	sigRecv := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv, sigs...)
	for sig := range sigRecv {
		fmt.Printf("Received a signal: %s\n", sig)
	}
}
