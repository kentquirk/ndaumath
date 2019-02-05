package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Ndau) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int64
		zb0001, err = dc.ReadInt64()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Ndau(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Ndau) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt64(int64(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Ndau) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt64(o, int64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Ndau) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int64
		zb0001, bts, err = msgp.ReadInt64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Ndau(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Ndau) Msgsize() (s int) {
	s = msgp.Int64Size
	return
}
