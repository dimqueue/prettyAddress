package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"prettyAddress/internal/generator"
	"prettyAddress/internal/parallel"
)

func main() {

	//generator.GenerateSequenceN()

	//fmt.Println(flag,key)
	//fmt.Printf("private scalar D = %x\n", key.D)
	//parallel.RunParallel(generator.GenerateSequenceN, 6)

	privateKey, ok := parallel.RunParallel(generator.GenerateSequenceN, 5)
	if ok {
		fmt.Println("Private key: ", hexutil.Encode(crypto.FromECDSA(privateKey)))
	} else {
		fmt.Println("not found")
	}

}
