package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var methods = map[int]string{
	0: http.MethodGet,
	1: http.MethodHead,
	2: http.MethodPost,
	3: http.MethodPut,
	4: http.MethodPatch,
	5: http.MethodDelete,
	6: http.MethodOptions,
	7: http.MethodTrace,
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	serverURL, ok := os.LookupEnv("SERVER_URL")
	if !ok {
		serverURL = "http://127.0.0.1:8080"
	}

	rps := 10
	if rpsEnv, ok := os.LookupEnv("RPS"); ok {
		var err error
		rps, err = strconv.Atoi(rpsEnv)
		if err != nil {
			log.Fatal(err)
		}
	}

	client := http.Client{
		Timeout: 1 * time.Minute,
	}

	t := time.NewTicker(1 * time.Second)

loop:
	for {
		for i := 0; i < rps; i++ {
			go func() {
				t := time.NewTimer(time.Millisecond * time.Duration(rand.Intn(1000)))
				select {
				case <-ctx.Done():
					t.Stop()
					return
				case <-t.C:
				}

				req, err := http.NewRequestWithContext(ctx, methods[rand.Intn(len(methods))], serverURL, nil)
				if err != nil {
					log.Fatal(err)
				}

				client.Do(req)
			}()
		}
		select {
		case <-ctx.Done():
			t.Stop()
			break loop
		case <-t.C:
		}
	}
}
