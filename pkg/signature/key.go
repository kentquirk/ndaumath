package signature

import (
	"encoding"
	"fmt"

	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

// A Key is a public or private key which knows about its algorithm
//
// This is most useful when abstracting over what might be a public or a
// private key. To recover the concrete instance, consider a typeswitch:
//
// switch key := keyI.(type) {
// case PublicKey:
//     ...
// case PrivateKey:
//     ...
// }
//
// Key includes several other interfaces to ensure consistent marshalling and
// unmarshalling in both binary and text formats
type Key interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	fmt.Stringer
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Sizer

	KeyBytes() []byte
	ExtraBytes() []byte
	Algorithm() Algorithm
	Truncate()
	Zeroize()
}

// IsPublic is true when the supplied Key is public
func IsPublic(k Key) bool {
	_, ok := k.(*PublicKey)
	return ok
}

// IsPrivate is true when the supplied Key is private
func IsPrivate(k Key) bool {
	_, ok := k.(*PrivateKey)
	return ok
}

// ParseKey attempts to parse the given text as a Key of the proper type
func ParseKey(text string) (Key, error) {
	switch {
	case MaybePrivate(text):
		priv := new(PrivateKey)
		err := priv.UnmarshalText([]byte(text))
		return priv, err
	case MaybePublic(text):
		pub := new(PublicKey)
		err := pub.UnmarshalText([]byte(text))
		return pub, err
	default:
		return nil, errors.New("text does not appear to be an ndau key")
	}
}
