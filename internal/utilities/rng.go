package utilities

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"

	"github.com/pkg/errors"
)

func URandFloatForCrypto() float64 {
	maxInt := big.NewInt(math.MaxInt64)

	randomBigInt, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		log.Panicln(errors.Wrap(err, "Could not call RNG for Roulette Wheel Selection"))
	}
	randomBigFloat := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(randomBigInt), big.NewFloat(0).SetInt(maxInt))

	floated, _ := randomBigFloat.Float64()
	return floated
}

func Pick[T any](values []T) *T {
	length := len(values)
	rnd := int(URandFloatForCrypto() * float64(length))
	return &values[rnd]
}

func Coin() bool {
	return *Pick([]bool{true, false})
}
