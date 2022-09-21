package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
)

type U128 [16]byte

func numberToVector(x uint64) U128 {
	return *(*U128)(decode(numberToHex(x)))
}

func hexToVector(x string) U128 {
	return *(*U128)(decode(x))
}

func (a U128) xor(b U128) U128 {
	var c U128

	for i := 0; i < 16; i++ {
		c[i] = a[i] ^ b[i]
	}

	return c
}

func (a U128) and(b U128) U128 {
	var c U128

	for i := 0; i < 16; i++ {
		c[i] = a[i] & b[i]
	}

	return c
}

func (a U128) or(b U128) U128 {
	var c U128

	for i := 0; i < 16; i++ {
		c[i] = a[i] | b[i]
	}

	return c
}

func (a U128) add(b U128) U128 {
	var c U128

	carry := 0
	for i := 0; i < 16; i++ {
		s := int(a[i]) + int(b[i]) + carry
		carry = s / 256
		c[i] = byte(s)
	}

	return c
}

func (a U128) sub(b U128) U128 {
	var c U128

	var bigA big.Int
	bigA.SetBytes(a[:])

	var bigB big.Int
	bigB.SetBytes(b[:])

	var bigC big.Int

	bigC.Sub(&bigB, &bigA)
	bigC.FillBytes(c[:])

	return c
}

func (a U128) nonZero() bool {

	for i := 0; i < 128; i++ {
		if HasBit(a[i/8], i%8) {
			return true
		}
	}

	return false
}

func decode(text string) []byte {
	v, _ := hex.DecodeString(text)
	return v
}

func encode(payload []byte) string {

	var buf bytes.Buffer

	for _, v := range payload {
		buf.WriteString(fmt.Sprintf("%02x", v))
	}

	return buf.String()
}

func numberToHex(num uint64) string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%032x", num))

	return buf.String()
}

func hexToNumber(text string) uint64 {
	buf, _ := hex.DecodeString(text)
	return binary.BigEndian.Uint64(buf[8:])
}
