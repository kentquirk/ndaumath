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
	"encoding"
	"fmt"

	"github.com/ndau/ndaumath/pkg/b32"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

// ensure that keyBase implements msgp marshal types
var _ msgp.Marshaler = (*keyBase)(nil)
var _ msgp.Unmarshaler = (*keyBase)(nil)
var _ msgp.Sizer = (*keyBase)(nil)

// ensure that keyBase implements text encoding interfaces
var _ encoding.TextMarshaler = (*keyBase)(nil)
var _ encoding.TextUnmarshaler = (*keyBase)(nil)

// A keyBase is a byte slice with known algorithm type
type keyBase struct {
	algorithm Algorithm
	key       []byte
	extra     []byte
}

func (key keyBase) pack() ([]byte, error) {
	if len(key.key) > 0xff { // capacity of single byte
		return nil, errors.New("can't pack keys of length > 0xff")
	}
	out := make([]byte, 1+len(key.key)+len(key.extra))
	out[0] = byte(len(key.key))
	split := 1 + len(key.key)
	copied := copy(out[1:split], key.key)
	if copied != len(key.key) {
		return nil, errors.New("pack: failed to copy full key data")
	}
	copied = copy(out[split:], key.extra)
	if copied != len(key.extra) {
		return nil, errors.New("pack: failed to copy full extra data")
	}
	return out, nil
}

func (key *keyBase) unpack(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	lk := int(data[0])
	split := 1 + lk
	if len(data) < split {
		return errors.New("can't unpack: too few bytes")
	}
	key.key = make([]byte, lk)
	copied := copy(key.key, data[1:split])
	if copied != lk {
		return errors.New("unpack: failed to copy full key data")
	}
	le := len(data) - split
	if le < 0 {
		panic("programming error in unpack")
	}
	key.extra = make([]byte, le)
	copied = copy(key.extra, data[split:])
	if copied != le {
		return errors.New("unpack: failed to copy full extra data")
	}
	return nil
}

// Marshal marshals the key into a serialized binary format
// which includes a type byte for the algorithm.
func (key keyBase) Marshal() (serialized []byte, err error) {
	data, err := key.pack()
	if err != nil {
		return nil, err
	}
	return marshal(key.Algorithm(), data)
}

// Unmarshal unmarshals the serialized binary data into the supplied key instance
func (key *keyBase) Unmarshal(serialized []byte) error {
	al, data, err := unmarshal(serialized)
	if err != nil {
		return err
	}
	key.algorithm = al
	err = key.unpack(data)
	return err
}

// MarshalMsg implements msgp.Marshaler
func (key keyBase) MarshalMsg(in []byte) (out []byte, err error) {
	out, err = key.Marshal()
	if err == nil {
		out = append(in, out...)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (key *keyBase) UnmarshalMsg(in []byte) (leftover []byte, err error) {
	var al Algorithm
	var db []byte

	al, db, leftover, err = unmarshalWithLeftovers(in)
	if err != nil {
		return
	}

	key.algorithm = al
	err = key.unpack(db)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
// Msgsize implements msgp.Sizer
//
// This method was copy-pasted from the IdentifiedData Msgsize implementation,
// as fundamentally a keyBase gets serialized as an IdentifiedData, and so should
// have the same size.
func (key *keyBase) Msgsize() (s int) {
	s = 1 + msgp.Uint8Size + msgp.BytesPrefixSize + len(key.key)
	return
}

// MarshalText implements encoding.TextMarshaler
//
// This marshaller uses a custom b32 encoding which is case-insensitive and
// lacks certain confusing pairs, for ease of human-friendly handling.
// For the same reason, it embeds a checksum, so it's easy to tell whether
// or not it was received correctly.
func (key keyBase) MarshalText() ([]byte, error) {
	bytes, err := key.Marshal()
	if err != nil {
		return nil, err
	}
	bytes = AddChecksum(bytes)
	return []byte(b32.Encode(bytes)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (key *keyBase) UnmarshalText(text []byte) error {
	bytes, err := b32.Decode(string(text))
	if err != nil {
		return err
	}
	var checksumOk bool
	bytes, checksumOk = CheckChecksum(bytes)
	if !checksumOk {
		return errors.New("key unmarshal failure: bad checksum")
	}
	return key.Unmarshal(bytes)
}

// KeyBytes returns the key's data
func (key keyBase) KeyBytes() []byte {
	if len(key.key) == 0 {
		return []byte{}
	}
	return key.key
}

// ExtraBytes returns any extra data
func (key keyBase) ExtraBytes() []byte {
	if len(key.extra) == 0 {
		return []byte{}
	}
	return key.extra
}

// Algorithm returns the key's algorithm
func (key keyBase) Algorithm() Algorithm {
	if key.algorithm == nil {
		return Null
	}
	return key.algorithm
}

// FullString returns the full text serialization of the key's data
//
// It suppresses errors in favor of a human-readable error message.
func (key keyBase) FullString(prefix string) string {
	// we can't deal with errors in this function, so let's just ignore the
	// error value and hope that we got at least something sensible back
	text, _ := key.MarshalText()
	if len(text) == 0 {
		return "<unmarshalable>"
	}
	return prefix + string(text)
}

// String returns a shorthand for the key's data
//
// This returns the first 8 characters of the text serialization,
// an ellipsis, then the final 4 characters of the text serialization.
// Total output size is constant at 15 characters.
//
// This destructively truncates the key, but it is a useful format for
// humans.
func (key keyBase) String(prefix string) string {
	s := key.FullString(prefix)
	if len(s) <= 15 {
		return s
	}
	return fmt.Sprintf("%s...%s", s[:8], s[len(s)-4:])
}

// Truncate removes all extra data from this key.
//
// This is a destructive operation which cannot be undone; make copies
// first if you need to.
func (key *keyBase) Truncate() {
	key.extra = nil
}

// Zeroize removes all data from this key
//
// This is a destructive operation which cannot be undone; make copies
// first if you need to.
func (key *keyBase) Zeroize() {
	key.algorithm = nil
	key.key = nil
	key.extra = nil
}

// IsZero is true when this key is the zero value
func (key keyBase) IsZero() bool {
	return (key.algorithm == nil || key.algorithm == Null) && len(key.key) == 0 && len(key.extra) == 0
}
