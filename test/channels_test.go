package test

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
