package parallel

import (
	"context"
	"crypto/ecdsa"
	"sync"
)

func RunParallel(fn func(context.Context) (*ecdsa.PrivateKey, bool), workers int) (*ecdsa.PrivateKey, bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan *ecdsa.PrivateKey, 1)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			if v, ok := fn(ctx); ok {
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
