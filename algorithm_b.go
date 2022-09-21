package main

import "lukechampine.com/uint128"

type StateB struct {
	Prime uint128.Uint128
}

func bootstrapB(key uint64) StateB {
	k := uint128.From64(key)

	if k.Mod(uint128.From64(2)) == uint128.Zero {
		k = k.Div(uint128.From64(2))
	}

	return StateB{k}
}

func gcdExtended(a, b uint128.Uint128, x, y *uint128.Uint128) uint128.Uint128 {

	if a == uint128.Zero {
		*x = uint128.From64(0)
		*y = uint128.From64(1)
		return b
	}

	var x1, y1 uint128.Uint128
	gcd := gcdExtended(b.Mod(a), a, &x1, &y1)

	*x = y1.SubWrap(b.Div(a).MulWrap(x1))
	*y = x1

	return gcd
}

func modInverse(a uint64) uint64 {
	M := uint128.Uint128{0, 1}
	var x, y uint128.Uint128

	gcdExtended(uint128.From64(a), M, &x, &y)
	return x.Mod(M).AddWrap(M).Mod(M).Lo

}

func (st *StateB) encrypt(x uint64) uint64 {
	invG := modInverse(st.Prime.Lo)
	for i := 0; i < 10; i++ {
		x *= invG
	}
	return x
}

func (st *StateB) decrypt(x uint64) uint64 {
	for i := 0; i < 10; i++ {
		x *= st.Prime.Lo
	}
	return x
}
