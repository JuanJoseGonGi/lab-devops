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
			return
		}
	}()

	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestServer_StartFileServer(t *testing.T) {
	server := NewServer(":8080")
	func() {
		time.Sleep(1 * time.Second)

		err := server.Shutdown(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
	}()

	err := server.StartFileServer()
	if err != nil {
		t.Error(err)
		return
	}
}
