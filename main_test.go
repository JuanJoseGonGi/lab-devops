package main

import (
	"context"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	server := NewServer(":8080")

	go func() {
		time.Sleep(1 * time.Second)

		err := server.Shutdown(context.Background())
		if err != nil {
			t.Error(err)
		}
	}()

	err := server.Start()
	if err != nil {
		t.Error(err)
	}
}
