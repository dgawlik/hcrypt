package main

import (
	"encoding/binary"
	"math/rand"
)

type BitIndex [128]int

type StateA struct {
	Lookup        BitIndex
	InverseLookup BitIndex
}

func bootstrapA(key U128) StateA {
	var lookup BitIndex
	var invLookup BitIndex

	for i := 0; i < 128; i++ {
		lookup[i] = i
	}

	source := rand.NewSource(int64(int64(binary.BigEndian.Uint64(key[:16]))))
	rng := rand.New(source)
	count := 100_000
	for count > 0 {
		i, j := rng.Intn(128), rng.Intn(128)
		lookup[i], lookup[j] = lookup[j], lookup[i]
		count--
	}

	for i := 0; i < 128; i++ {
		invLookup[lookup[i]] = i
	}

	return StateA{
		lookup,
		invLookup,
	}
}

func SetBit(num byte, pos int) byte {
	return num | (1 << pos)
}

func ClearBit(num byte, pos int) byte {
	return num & byte(^(1 << pos))
}

func HasBit(num byte, pos int) bool {
	val := num & (1 << pos)
	return (val > 0)
}

func (st *StateA) permute(num U128) U128 {
	var yi U128

	for i := 0; i < 128; i++ {
		j := st.Lookup[i]

		if HasBit(num[i/8], i%8) {
			yi[j/8] |= SetBit(yi[j/8], j%8)
		}
	}

	return yi
}

func (st *StateA) inversePermute(num U128) U128 {
	var yi U128

	for i := 0; i < 128; i++ {
		j := st.InverseLookup[i]

		if HasBit(num[i/8], i%8) {
			yi[j/8] = SetBit(yi[j/8], j%8)
		}
	}

	return yi
}

func (st *StateA) encrypt(x U128) U128 {
	return st.permute(x)
}

func (st *StateA) decrypt(x U128) U128 {
	return st.inversePermute(x)
}
