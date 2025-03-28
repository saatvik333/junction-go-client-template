package utils

import (
	"crypto/rand"
	"math/big"
	math "math/rand"
	"strconv"
	"time"
)


func Monikergenerate() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz") // Available characters for the random word.
	moniker := make([]rune, 5)

	// Initialize random seed for non-cryptographic randomness.
	math.Seed(time.Now().UnixNano())

	// Generate random 5-letter word
	for i := range moniker {
		// Generate a random index in the range of available characters.
		randomIndex := math.Intn(len(letters)) // Non-cryptographic randomness
		moniker[i] = letters[randomIndex]
	}

	// Generate a random 4-digit number.
	fourDigitNumber := strconv.Itoa(math.Intn(10000)) // Generates a number between 0 and 9999.

	// Generate a random 1-digit number.
	oneDigitNumber := strconv.Itoa(math.Intn(10)) // Generates a number between 0 and 9.

	// Concatenate the random moniker with the numbers.
	return string(moniker) + "_" + fourDigitNumber + "-" + oneDigitNumber
}

func KeysGenerateAndSupply(count int) ([]string, []uint64) {
	prefix := "eth1"
	allowedChars := "0123456789abcdef"

	keys := make([]string, count)
	supply := make([]uint64, count)

	// Generate the keys and supply values
	for i := 0; i < count; i++ {
		var randomString string
		for j := 0; j < 40; j++ {
			randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allowedChars))))
			randomChar := allowedChars[randomIndex.Int64()]
			randomString += string(randomChar)
		}
		
		keys[i] = prefix + randomString

		randomSupply, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		supply[i] = randomSupply.Uint64() 
	}

	return keys, supply
}
func DaGen() string {
	return []string{"eigenda", "celestia", "avail"}[math.Intn(3)]
}