package main

import (
	"SCA/internal/server"
	"SCA/internal/storage"
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"
)

func main() {
	config := newConfig()
	addr := net.JoinHostPort(config.host, config.port)
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	storage := storage.New(ctx)
	err := storage.Conn()
	if err != nil {
		panic(err)
	}
	defer storage.Close()

	server := server.New(addr, ctx, storage)
	err = server.Conn()
	if err != nil {
		panic(err)
	}
	defer server.Close()

	<-ctx.Done()

	log.Println("exit")
}