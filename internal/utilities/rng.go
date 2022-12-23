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
		log.Panicln(errors.Wrap(err, "Could not call RNG for URandFloatForCrypto"))
	}
	randomBigFloat := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(randomBigInt), big.NewFloat(0).SetInt(maxInt))

	floated, _ := randomBigFloat.Float64()
	return floated
}

func URandIntN(n int) int {
	maxInt := big.NewInt(int64(n))
	randomBigInt, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		log.Panicln(errors.Wrap(err, "Could not call RNG for URandIntN"))
	}
	return int(randomBigInt.Int64())
}

func Pick[T any](values []T) *T {
	length := len(values)
	rnd := int(URandFloatForCrypto() * float64(length))
	return &values[rnd]
}

func Coin() bool {
	return *Pick([]bool{true, false})
}
