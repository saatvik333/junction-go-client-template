package utils

import (
	"log"
	"os"

	cryptoSecp256k1 "github.com/cometbft/cometbft/crypto/secp256k1"
)

func WriteToFile(filePath string, data []byte) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadPublicKey(filePath string) cryptoSecp256k1.PubKey {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	//secpPrivKey := secp256k1.PrivKeyFromBytes(data)
	//pk := secpPrivKey.PubKey().SerializeCompressed()

	return cryptoSecp256k1.PubKey(data)
}

func LoadPublicKeyByte(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func LoadPrivateKey(filePath string) cryptoSecp256k1.PrivKey {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return cryptoSecp256k1.PrivKey(data)
}
