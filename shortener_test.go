package shortener

import (
	"math/rand"
	"testing"
)

func TestLib(T *testing.T) {
	maxId := AliasGenerator.MaxId()

	for i := int64(0); i < 100; i++ {
		r := rand.Int63n(maxId)
		k := AliasGenerator.Encode(r)
		t := AliasGenerator.Decode(k)

		if t != r {
			T.Errorf("error when processing %d, %s, %d", r, k, t)
		}
	}
}
