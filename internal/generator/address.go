package generator

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func GenerateAddress() (*ecdsa.PrivateKey, common.Address) {
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
	//fmt.Println(address)

	hash := crypto.Keccak256(publicKeyBytes[1:])
	//fmt.Println(hexutil.Encode(hash[12:]))
	//fmt.Println(hash[12:])
	hash = hash[12:]
	return privateKey, address
}
