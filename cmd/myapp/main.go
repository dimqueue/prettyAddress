package main

import (
	"context"
	"fmt"
	"privatekey/internal/generator"
	"privatekey/internal/parallel"
)

func findNumber(ctx context.Context) (int, bool) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return 0, false
		default:
			if i == 42 {
				return i, true
			}
		}
	}
}

func main() {

	//generator.GenerateSequenceN()

	//fmt.Println(flag,key)
	//fmt.Printf("private scalar D = %x\n", key.D)
	//parallel.RunParallel(generator.GenerateSequenceN, 6)

	n, ok := parallel.RunParallel(generator.GenerateSequenceN, 5)
	if ok {
		fmt.Println("found", n)
	} else {
		fmt.Println("not found")
	}

}
