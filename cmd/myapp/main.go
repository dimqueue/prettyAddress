package main

import (
	"privatekey/internal/generator"
)

//	func GenerateSeveralWorker() {
//		addressChan := make(chan common.Address)
//		privateKeyChan := make(chan []byte)
//		for i := 0; i < 5; i++ {
//			go GenerateSequenceN(addressChan, privateKeyChan)
//		}
//
// }

func main() {
	//GenerateSeveralWorker()
	generator.GenerateSequenceN()
	//var wg sync.WaitGroup
	//context.Context()
}
