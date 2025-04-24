package generator

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// GenerateSequenceN generates an address with a beginning of N identical characters
func GenerateSequenceN(ctx context.Context, n int) (*ecdsa.PrivateKey, bool) {
	//var n int
	//fmt.Scanln(&n)
	var res *ecdsa.PrivateKey
	var address common.Address

	for {
		select {
		case <-ctx.Done():
			return nil, false
		default:
			//var flag = false
			res, address = GenerateAddress()
			firstNumber := (address[0] & 0xF0) >> 4
			//fmt.Println("first number:")
			//fmt.Println(firstNumber)
			numByte := (n + n%2) / 2
			//fmt.Println("num byte:")
			//fmt.Println(numByte)
			for i := 0; i < numByte; i++ {
				//fmt.Println("address[0] & 0xF0) >> 4")
				//fmt.Println((address[0] & 0xF0) >> 4)
				//fmt.Println("address[0] & 0xF0")
				//fmt.Println((address[0] & 0xF))
				if i == numByte-1 {
					if n%2 == 0 {
						if (address[i]&0xF0)>>4 == firstNumber && address[i]&0xF == firstNumber {
							fmt.Println("Address found: ", address.String()) //strings.ToUpper(address.String())
							//flag = true
							return res, true
						}
					} else if (address[i]&0xF0)>>4 == firstNumber {
						fmt.Println("Address found: ", strings.ToUpper(address.String()))
						//flag = true
						return res, true

					}
					break
				} else if (address[i]&0xF0)>>4 != firstNumber || address[i]&0xF != firstNumber {
					break
				}
			}
			//if flag {
			//	break
			//}
		}
	}
	//return false, res
}
