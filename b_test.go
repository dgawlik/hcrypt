package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHomomorphicPropertiesB(t *testing.T) {
	key := rand.Uint64()

	st := bootstrapB(key)

	for i := 0; i < 100_000; i++ {
		a := rand.Uint64()
		b := rand.Uint64()

		left := st.encrypt(a) + st.encrypt(b)
		right := st.encrypt(a + b)

		if left != right {
			fmt.Println(left)
			fmt.Println(right)
			t.Fail()
		}
	}

}

func TestEncryptDecyptB(t *testing.T) {
	key := rand.Uint64()

	st := bootstrapB(key)

	for i := 0; i < 100_000; i++ {
		A := rand.Uint64()
		encA := st.encrypt(A)
		decA := st.decrypt(encA)

		if A != decA || A == encA {
			t.Fail()
		}
	}
}
