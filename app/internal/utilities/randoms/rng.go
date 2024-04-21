package randoms

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

func UniformCryptoFloat() float64 {
	if DEBUG_MODE {
		return mrand.Float64()
	} else {

		maxInt := big.NewInt(math.MaxInt64)

		randomBigInt, err := crand.Int(crand.Reader, maxInt)
		if err != nil {
			log.Panicln(errors.Wrap(err, "Could not call RNG for UniformCrypto"))
		}
		randomBigFloat := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(randomBigInt), big.NewFloat(0).SetInt(maxInt))

		floated, _ := randomBigFloat.Float64()
		return floated
	}
}

func UniformIntN(n int) int {
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
