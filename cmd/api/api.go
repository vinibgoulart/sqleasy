package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/vinibgoulart/sqleasy/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go http.ServerInit(ctx, &waitGroup)

	closeChannel := make(chan os.Signal, 1)
	signal.Notify(closeChannel, syscall.SIGINT, syscall.SIGTERM)

	<-closeChannel
	cancel()

	waitGroup.Wait()
}
