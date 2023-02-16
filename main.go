package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	config2 "severyanochka/internals/config"
	"severyanochka/internals/server"
)

func main() {
	config := config2.LoadDevConfig()
	ctx, cancel := context.WithCancel(context.Background())

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)

	srv := server.NewServer(ctx, config)

	go func() {
		osCall := <-s
		log.Println("Server stoped: ", osCall)
		srv.Shutdown()
		cancel()
	}()

	srv.Start()
}
