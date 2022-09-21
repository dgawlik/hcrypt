package main

import (
	"math/rand"
	"testing"
)

func TestHomomorphicPropertiesA(t *testing.T) {
	key := hexToVector("a291a728727ac647a53193be9583c504")

	st := bootstrapA(key)

	for i := 0; i < 100_000; i++ {
		a := numberToVector(rand.Uint64())
		b := numberToVector(rand.Uint64())

		leftOr := st.encrypt(a).or(st.encrypt(b))
		rightOr := st.encrypt(a.or(b))

		leftAnd := st.encrypt(a).and(st.encrypt(b))
		rightAnd := st.encrypt(a.and(b))

		leftXor := st.encrypt(a).xor(st.encrypt(b))
		rightXor := st.encrypt(a.xor(b))

		if (leftOr != rightOr) || (leftAnd != rightAnd) || (leftXor != rightXor) {
			t.Fail()
		}
	}

}

func TestEncryptDecyptA(t *testing.T) {
	key := hexToVector("a291a728727ac647a53193be9583c504")

	st := bootstrapA(key)

	for i := 0; i < 100_000; i++ {
		A := numberToVector(rand.Uint64())
		encA := st.encrypt(A)
		decA := st.decrypt(encA)

		if A != decA || A == encA {
			t.Fail()
		}
	}
}
