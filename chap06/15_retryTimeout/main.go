package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(context.Context) error) {
	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully.")
			return
		}

		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1:", ctx.Err())
			return
		}

		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)

		select {
		case <-ctx.Done():
			fmt.Println("time expited 2:", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("try again.")
		}
	}
}

func check(ctx context.Context) error {
	// <-ctx.Done()
	// return ctx.Err()
	return errors.New("some error")
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(4*time.Second))
	defer cancel()
	retryTimeout(ctx, time.Duration(2*time.Second), check)
}
