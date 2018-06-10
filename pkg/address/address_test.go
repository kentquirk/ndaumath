package address

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func getKinds() []Kind {
	return []Kind{
		KindUser,
		KindNdau,
		KindExchange,
		KindEndowment,
	}
}

func TestArbitraryAddressesAreValid(t *testing.T) {
	kinds := getKinds()
	for i := 0; i < 16; i++ {
		key := make([]byte, 32)
		_, err := rand.Read(key)
		require.NoError(t, err)

		t.Run(string(i), func(t *testing.T) {
			t.Log("key", fmt.Sprintf("%x", key))
			address, err := Generate(kinds[i&3], key)
			require.NoError(t, err)
			t.Log("address", address)
			require.NoError(t, Validate(address))
		})
	}
}

func TestKnownKeyGeneratesKnownValue(t *testing.T) {
	key := make([]byte, 16)
	for i := byte(0); i < 16; i++ {
		key[i] = i
	}

	address, err := Generate(KindUser, key)
	require.NoError(t, err)
	require.Equal(t, "ndadprx764ciigti8d8whtw2kct733r85qvjukhqhke3dka4", address)
}

func BenchmarkGeneration(b *testing.B) {
	key := make([]byte, 32)
	kinds := getKinds()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		_, err := rand.Read(key)
		if err != nil {
			b.FailNow()
		}
		b.StartTimer()

		_, err = Generate(kinds[n&3], key)
		if err != nil {
			b.FailNow()
		}
	}
}