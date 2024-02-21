package main

import (
	"fmt"
	// "math/rand"
	// "time"
	"crypto/rand"
	"math/big"
)

// func rollDiceWithMath() int {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	rng := rand.New(source)
// 	return rng.Intn(6) + 1
// }

func rollDiceWithCrypto() int {
	result, _ := rand.Int(rand.Reader, big.NewInt(6))
	return int(result.Int64()) + 1
}

func main() {
	// result := rollDiceWithMath()
	// fmt.Printf("You rolled a: %d\n", result)

	result := rollDiceWithCrypto()
	fmt.Printf("You rolled a: %d\n", result)
}
