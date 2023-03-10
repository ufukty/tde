package utilities

import (
	crand "crypto/rand"
	"log"
	"math"
	"math/big"
	mrand "math/rand"

	"github.com/pkg/errors"
)

const DEBUG_MODE = false

// func init() {
// 	mrand.Seed(time.Now().UnixNano())
// }

func URandFloatForCrypto() float64 {
	if DEBUG_MODE {
		return mrand.Float64()
	} else {

		maxInt := big.NewInt(math.MaxInt64)

		randomBigInt, err := crand.Int(crand.Reader, maxInt)
		if err != nil {
			log.Panicln(errors.Wrap(err, "Could not call RNG for URandFloatForCrypto"))
		}
		randomBigFloat := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(randomBigInt), big.NewFloat(0).SetInt(maxInt))

		floated, _ := randomBigFloat.Float64()
		return floated
	}
}

func URandIntN(n int) int {
	if DEBUG_MODE {
		return mrand.Intn(n)
	} else {
		maxInt := big.NewInt(int64(n))
		randomBigInt, err := crand.Int(crand.Reader, maxInt)
		if err != nil {
			log.Panicln(errors.Wrap(err, "Could not call RNG for URandIntN"))
		}
		return int(randomBigInt.Int64())
	}
}

func Pick[T any](s []T) *T {
	if len(s) == 0 {
		panic("Cannot Pick from empty slice")
	}
	return &s[URandIntN(len(s))]
}

func PickExcept[T comparable](s []T, e []T) *T {
	if len(s) == 0 {
		panic("Cannot Pick from empty slice")
	}
	cleaned := SliceExceptItems(s, e)
	if len(cleaned) == 0 {
		panic("Exceptions remove every item in Pick slice")
	}
	return Pick(cleaned)
}

func Coin() bool {
	return *Pick([]bool{true, false})
}
