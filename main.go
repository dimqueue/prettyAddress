package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
)

func GenerateAddress() common.Address {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println(hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fmt.Println("Pubkey")
	//fmt.Println(hexutil.Encode(publicKeyBytes))
	//fmt.Println(publicKeyBytes)
	//fmt.Println("address")
	fmt.Println(address)

	hash := crypto.Keccak256(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash[12:]))
	fmt.Println(hash[12:])
	hash = hash[12:]
	return address
}

func GenerateSequenceN() {
	var n int
	fmt.Scanln(&n)
	for {
		var flag = false
		address := GenerateAddress()
		//address := []byte{255, 255, 255}
		//addressBytes := address.Bytes()
		firstNumber := (address[0] & 0xF0) >> 4
		fmt.Println("first number:")
		fmt.Println(firstNumber)
		numByte := (n + n%2) / 2
		fmt.Println("num byte:")
		fmt.Println(numByte)
		for i := 0; i < numByte; i++ {

			fmt.Println("address[0] & 0xF0) >> 4")
			fmt.Println((address[0] & 0xF0) >> 4)
			fmt.Println("address[0] & 0xF0")
			fmt.Println((address[0] & 0xF))

			if i == numByte-1 {
				if n%2 == 0 {
					if (address[i]&0xF0)>>4 == firstNumber && address[i]&0xF == firstNumber {
						fmt.Println("finaly: ", strings.ToUpper(address.String()))
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

func main() {
	GenerateSequenceN()
}
