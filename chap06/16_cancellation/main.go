package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func channelCancellation(stop <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-stop:
			cancel()
		case <-ctx.Done():
		}
	}()

	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.ardanlabs.com/blog/index.xml", nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		sb, _ := ioutil.ReadAll(res.Body)
		fmt.Println("Response:", string(sb))
		return nil
	}(ctx)
}

func main() {
	stop := make(chan struct{})
	channelCancellation(stop)
}
