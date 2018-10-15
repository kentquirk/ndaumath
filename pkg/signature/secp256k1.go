package signature

import (
	randP "crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/btcsuite/btcd/btcec"
	"github.com/oneiro-ndev/ndaumath/pkg/key"
)

// Secp256k1 is the HD Key type defined in [ek].
//
// Never edit this; it would be a const if go were smarter
//
// [ek]: https://github.com/oneiro-ndev/ndaumath/blob/master/pkg/key/extendedkey.go
var Secp256k1 = secp256k1{}

type secp256k1 struct{}

// static assert that secp256k1 is an Algorithm
var _ Algorithm = (*secp256k1)(nil)

// PublicKeySize is the size in bytes of this algorithm's public keys
func (secp256k1) PublicKeySize() int {
	return btcec.PubKeyBytesLenCompressed
}

// PrivateKeySize is the size in bytes of this algorithm's private keys
func (secp256k1) PrivateKeySize() int {
	return btcec.PrivKeyBytesLen
}

// SignatureSize is the size in bytes of this algorithm's signatures
func (secp256k1) SignatureSize() int {
	// strictly speaking, this is only occasionally true: the actual signature
	// size depends on the encoded size of arbitrarily-sized bigints, plus
	// some structure. The min size for single byte encodings is 8, and the
	// constant is 6; picking 70 means that the bigints can be up to 32
	// bytes each, which seems plausible.
	return 70
}

// Generate creates a new keypair
func (secp256k1) Generate(rand io.Reader) (public, private []byte, err error) {
	// first thing: if rand is not set, set it
	if rand == nil {
		rand = randP.Reader
	}
	// generate a seed of the recommended size
	seed := make([]byte, key.RecommendedSeedLen)
	_, err = io.ReadFull(rand, seed)
	if err != nil {
		return
	}

	var privateEK *key.ExtendedKey
	privateEK, err = key.NewMaster(seed)
	if err != nil {
		return
	}

	private = privateEK.Bytes()
	public = privateEK.PubKeyBytes()

	return
}

// btcec documentation claims that we're not supposed to sign full messages,
// but instead pre-hash the message and sign the hash. That sounds weird
// to me, but who am I to disobey the documentation?
func hash(message []byte) []byte {
	h := sha256.Sum256(message)
	return h[:]
}

// Sign signs the message with privateKey and returns a signature
func (secp256k1) Sign(private, message []byte) []byte {
	ecPriv, _ := btcec.PrivKeyFromBytes(btcec.S256(), private)
	sig, err := ecPriv.Sign(hash(message))
	if err != nil {
		// errors happen deterministically, if the computed value for R or S
		// happens to equal 0 for a given private key and message.
		// this is very unlikely, which makes updating this interface
		// (with all attendent work updating usage downstream)
		// a fairly low priority.
		panic(err) // TOOD: update Algorithm signature to allow signature to fail without panic
	}
	return sig.Serialize()
}

// Verify verifies a message's signature
//
// Return true if the signature is valid
func (secp256k1) Verify(public, message, signature []byte) bool {
	pub, err := btcec.ParsePubKey(public, btcec.S256())
	if err != nil {
		// if the public key can't be parsed, signature is invalid
		return false
	}
	sig, err := btcec.ParseSignature(signature, btcec.S256())
	if err != nil {
		// if the signature can't be parsed, signature is invalid
		return false
	}

	return sig.Verify(hash(message), pub)
}
