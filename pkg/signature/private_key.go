package signature

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


import (
	"bytes"
	"fmt"
	"strings"

	"github.com/tinylib/msgp/msgp"
)

// PrivateKeyPrefix always prefixes Ndau private keys in text serialization
const PrivateKeyPrefix = "npvt"

// MaybePrivate provides a fast way to check whether a string looks like
// it might be an ndau private key.
//
// To get a definitive answer as to whether something is a private key, one
// must attempt to deserialize it using UnmarshalText and check the error
// value. That takes some work; it's faster to use this to get a first impression.
//
// This function will allow some false positives, but no false negatives:
// some values for which it returns `true` may not be actual valid keys,
// but no values for which it returns `false` will return actual valid keys.
func MaybePrivate(s string) bool {
	return strings.HasPrefix(s, PrivateKeyPrefix)
}

// ensure that PrivateKey implements export interfaces
var _ Key = (*PrivateKey)(nil)

// A PrivateKey is the private half of a keypair
type PrivateKey struct{ keyBase }

// RawPrivateKey creates a PrivateKey from raw data
//
// This is unsafe and subject to only minimal type-checking; it should
// normally be avoided.
func RawPrivateKey(al Algorithm, key, extra []byte) (*PrivateKey, error) {
	if key == nil {
		key = []byte{}
	}
	if extra == nil {
		extra = []byte{}
	}
	pk := PrivateKey{keyBase{
		algorithm: al,
		key:       key,
		extra:     extra,
	}}
	if len(key) != pk.Size() {
		return nil, fmt.Errorf("wrong private key length: have %d, want %d", len(key), pk.Size())
	}
	return &pk, nil
}

// Size returns the size of this key
func (key PrivateKey) Size() int {
	return key.Algorithm().PrivateKeySize()
}

// Sign the supplied message
func (key PrivateKey) Sign(message []byte) Signature {
	al := key.Algorithm()
	return Signature{
		algorithm: al,
		data:      al.Sign(key.key, message),
	}
}

// Unmarshal unmarshals the serialized bytes into the PrivateKey pointer
func (key *PrivateKey) Unmarshal(serialized []byte) error {
	err := key.keyBase.Unmarshal(serialized)
	if err == nil {
		if len(key.key) != key.Size() {
			err = fmt.Errorf("Wrong size private key: expect len %d, have %d", key.Size(), len(key.key))
		}
	}
	return err
}

// UnmarshalMsg implements msgp.Unmarshaler
func (key *PrivateKey) UnmarshalMsg(in []byte) (leftover []byte, err error) {
	leftover, err = key.keyBase.UnmarshalMsg(in)
	if err == nil {
		if len(key.key) != key.Size() {
			err = fmt.Errorf("Wrong size signature: expect len %d, have %d", key.Size(), len(key.key))
		}
	}
	return leftover, err
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
// Msgsize implements msgp.Sizer
//
// This method was copy-pasted from the IdentifiedData Msgsize implementation,
// as fundamentally a PrivateKey gets serialized as an IdentifiedData, and so should
// have the same size.
func (key *PrivateKey) Msgsize() (s int) {
	s = 1 + msgp.Uint8Size + msgp.BytesPrefixSize + len(key.key)
	return
}

// MarshalText implements encoding.TextMarshaler.
//
// PublicKeys encode like Keys, with the addition of a human-readable prefix
// for easy identification.
func (key PrivateKey) MarshalText() ([]byte, error) {
	bytes, err := key.keyBase.MarshalText()
	bytes = append([]byte(PrivateKeyPrefix), bytes...)
	return bytes, err
}

// UnmarshalText implements encoding.TextUnmarshaler
func (key *PrivateKey) UnmarshalText(text []byte) error {
	expectPrefix := []byte(PrivateKeyPrefix)
	lep := len(expectPrefix)
	if !bytes.Equal(expectPrefix, text[:lep]) {
		return fmt.Errorf("private key must begin with %q; got %q", PrivateKeyPrefix, text[:lep])
	}
	err := key.keyBase.UnmarshalText(text[lep:])
	if err == nil {
		if len(key.key) != key.Size() {
			err = fmt.Errorf("Wrong size key: expect len %d, have %d", key.Size(), len(key.key))
		}
	}
	return err
}

// FullString returns the key's human-readable serialization
func (key PrivateKey) FullString() string {
	return key.keyBase.FullString(PrivateKeyPrefix)
}

// String returns a shorthand for the key's data
//
// This returns the first 8 characters of the text serialization,
// an ellipsis, then the final 4 characters of the text serialization.
// Total output size is constant at 15 characters.
//
// This destructively truncates the key, but it is a useful format for
// humans.
func (key PrivateKey) String() string {
	return key.keyBase.String(PrivateKeyPrefix)
}

// Truncate removes all extra data from this key.
//
// This is a destructive operation which cannot be undone; make copies
// first if you need to.
func (key *PrivateKey) Truncate() {
	key.extra = nil
}

// Zeroize removes all data from this key
//
// This is a destructive operation which cannot be undone; make copies
// first if you need to.
func (key *PrivateKey) Zeroize() {
	if key == nil {
		return
	}
	key.keyBase.Zeroize()
}

// MarshalString is like MarshalText, but to a string
func (key *PrivateKey) MarshalString() (string, error) {
	// Why doesn't MarshalText produce a string anyway?
	t, err := key.MarshalText()
	if t == nil {
		t = []byte{}
	}
	return string(t), err
}

// ParsePrivateKey parses a string representation of a private key, if possible
func ParsePrivateKey(s string) (*PrivateKey, error) {
	key := new(PrivateKey)
	err := key.UnmarshalText([]byte(s))
	return key, err
}
