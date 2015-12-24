package main
import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main() {
	secret := "bgvyzdsv"
	fiveZeros := calculateHash(secret, 5)
	sixZeros := calculateHash(secret, 6)
	fmt.Printf("Five zeros: %v\n", fiveZeros)
	fmt.Printf("Six zeros: %v", sixZeros)
}

func calculateHash(secret string, numberOfZeros int) int {
	var number int
	var zeros string
	hexHash := "123456"

	for i := 0; i < numberOfZeros; i++ {
		zeros += "0"
	}

	for hexHash[:numberOfZeros] != zeros {
		number++
		hash := md5.Sum([]byte(fmt.Sprintf("%v%v", secret, number)))
		hexHash = hex.EncodeToString(hash[:])
	}

	return number
}