package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vinibgoulart/sqleasy/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go cli.Exec(ctx, &waitGroup)

	closeChannel := make(chan os.Signal, 1)
	signal.Notify(closeChannel, syscall.SIGINT, syscall.SIGTERM)

	<-closeChannel
	cancel()

	waitGroup.Wait()
}
