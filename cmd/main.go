package main

import (
	"context"
	"fmt"
	goods "goods-square/internal"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	// "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	server, err := goods.NewServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to prepare server, %s", err.Error())
		return
	}
	fmt.Printf("start server\n")
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "closed server, %s", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	fmt.Printf("SIGNAL %d received\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown, %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "server shutdown completed")
}
