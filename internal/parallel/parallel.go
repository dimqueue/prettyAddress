package parallel

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"sync"
)

//add a structure for initializing parameters for execution

func preRun() int {
	var n int
	fmt.Scanln(&n)
	return n
}

func RunParallel(fn func(context.Context, int) (*ecdsa.PrivateKey, bool), workers int) (*ecdsa.PrivateKey, bool) {
	n := preRun()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan *ecdsa.PrivateKey, 1)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			if v, ok := fn(ctx, n); ok {
				select {
				case out <- v:
					cancel()
				default:
				}
			}
		}()
	}
	go func() { wg.Wait(); close(out) }()
	v, ok := <-out
	return v, ok
}
