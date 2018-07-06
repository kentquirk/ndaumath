package eai

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *RSRow) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	err = z.Duration.DecodeMsg(dc)
	if err != nil {
		return
	}
	{
		var zb0002 string
		zb0002, err = dc.ReadString()
		if err != nil {
			return
		}
		z.Rate = parseRateString(zb0002)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RSRow) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = z.Duration.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteString((Rate).toString(z.Rate))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RSRow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o, err = z.Duration.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendString(o, (Rate).toString(z.Rate))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RSRow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	bts, err = z.Duration.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	{
		var zb0002 string
		zb0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		z.Rate = parseRateString(zb0002)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RSRow) Msgsize() (s int) {
	s = 1 + z.Duration.Msgsize() + msgp.StringPrefixSize + len((Rate).toString(z.Rate))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RTRow) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	err = z.From.DecodeMsg(dc)
	if err != nil {
		return
	}
	{
		var zb0002 string
		zb0002, err = dc.ReadString()
		if err != nil {
			return
		}
		z.Rate = parseRateString(zb0002)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RTRow) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = z.From.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteString((Rate).toString(z.Rate))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RTRow) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o, err = z.From.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendString(o, (Rate).toString(z.Rate))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RTRow) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	bts, err = z.From.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	{
		var zb0002 string
		zb0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		z.Rate = parseRateString(zb0002)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RTRow) Msgsize() (s int) {
	s = 1 + z.From.Msgsize() + msgp.StringPrefixSize + len((Rate).toString(z.Rate))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Rate) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			return
		}
		(*z) = parseRateString(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Rate) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString((Rate).toString(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Rate) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, (Rate).toString(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Rate) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		(*z) = parseRateString(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Rate) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len((Rate).toString(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RateSlice) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RateSlice, zb0002)
	}
	for zb0001 := range *z {
		var zb0003 uint32
		zb0003, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		err = (*z)[zb0001].Duration.DecodeMsg(dc)
		if err != nil {
			return
		}
		{
			var zb0004 string
			zb0004, err = dc.ReadString()
			if err != nil {
				return
			}
			(*z)[zb0001].Rate = parseRateString(zb0004)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z RateSlice) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0005 := range z {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = z[zb0005].Duration.EncodeMsg(en)
		if err != nil {
			return
		}
		err = en.WriteString((Rate).toString(z[zb0005].Rate))
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RateSlice) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0005 := range z {
		// array header, size 2
		o = append(o, 0x92)
		o, err = z[zb0005].Duration.MarshalMsg(o)
		if err != nil {
			return
		}
		o = msgp.AppendString(o, (Rate).toString(z[zb0005].Rate))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RateSlice) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RateSlice, zb0002)
	}
	for zb0001 := range *z {
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		bts, err = (*z)[zb0001].Duration.UnmarshalMsg(bts)
		if err != nil {
			return
		}
		{
			var zb0004 string
			zb0004, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001].Rate = parseRateString(zb0004)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RateSlice) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0005 := range z {
		s += 1 + z[zb0005].Duration.Msgsize() + msgp.StringPrefixSize + len((Rate).toString(z[zb0005].Rate))
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RateTable) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RateTable, zb0002)
	}
	for zb0001 := range *z {
		var zb0003 uint32
		zb0003, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		err = (*z)[zb0001].From.DecodeMsg(dc)
		if err != nil {
			return
		}
		{
			var zb0004 string
			zb0004, err = dc.ReadString()
			if err != nil {
				return
			}
			(*z)[zb0001].Rate = parseRateString(zb0004)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z RateTable) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0005 := range z {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = z[zb0005].From.EncodeMsg(en)
		if err != nil {
			return
		}
		err = en.WriteString((Rate).toString(z[zb0005].Rate))
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RateTable) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0005 := range z {
		// array header, size 2
		o = append(o, 0x92)
		o, err = z[zb0005].From.MarshalMsg(o)
		if err != nil {
			return
		}
		o = msgp.AppendString(o, (Rate).toString(z[zb0005].Rate))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RateTable) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RateTable, zb0002)
	}
	for zb0001 := range *z {
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		bts, err = (*z)[zb0001].From.UnmarshalMsg(bts)
		if err != nil {
			return
		}
		{
			var zb0004 string
			zb0004, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001].Rate = parseRateString(zb0004)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RateTable) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0005 := range z {
		s += 1 + z[zb0005].From.Msgsize() + msgp.StringPrefixSize + len((Rate).toString(z[zb0005].Rate))
	}
	return
}
