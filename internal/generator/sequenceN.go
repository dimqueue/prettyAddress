package generator

import (
	"fmt"
	"strings"
)

func GenerateSequenceN() {
	var n int
	fmt.Scanln(&n)

	for {
		var flag = false
		address := generator.GenerateAddress()
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
						fmt.Println("finaly: ", address.String()) //strings.ToUpper(address.String())
						flag = true
					}
				} else if (address[i]&0xF0)>>4 == firstNumber {
					fmt.Println("finaly2: ", strings.ToUpper(address.String()))
					flag = true
				}
				break
			} else if (address[i]&0xF0)>>4 != firstNumber || address[i]&0xF != firstNumber {
				break
			}
		}
		if flag {
			break
		}
	}
}
