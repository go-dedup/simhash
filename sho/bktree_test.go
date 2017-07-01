// Copyright 2015, Yahoo Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sho

import (
	"fmt"
	"testing"
)

func ExampleNewOracle() {
	// just add 0 and 1.
	oracle := NewOracle()
	for i := uint64(1); i < 2; i++ {
		oracle.See(i)
	}
	r := uint8(2)
	fmt.Printf("Within distance of %d,\n", r)
	for i := uint64(0); i < 30; i++ {
		fmt.Printf("Has the oracle seen anything closed to %02d (%08b)? %t\n", i, i, oracle.Seen(i, r))
	}
	// Output:
	// Within distance of 2,
	// Has the oracle seen anything closed to 00 (00000000)? true
	// Has the oracle seen anything closed to 01 (00000001)? true
	// Has the oracle seen anything closed to 02 (00000010)? true
	// Has the oracle seen anything closed to 03 (00000011)? true
	// Has the oracle seen anything closed to 04 (00000100)? true
	// Has the oracle seen anything closed to 05 (00000101)? true
	// Has the oracle seen anything closed to 06 (00000110)? false
	// Has the oracle seen anything closed to 07 (00000111)? false
	// Has the oracle seen anything closed to 08 (00001000)? true
	// Has the oracle seen anything closed to 09 (00001001)? true
	// Has the oracle seen anything closed to 10 (00001010)? false
	// Has the oracle seen anything closed to 11 (00001011)? false
	// Has the oracle seen anything closed to 12 (00001100)? false
	// Has the oracle seen anything closed to 13 (00001101)? false
	// Has the oracle seen anything closed to 14 (00001110)? false
	// Has the oracle seen anything closed to 15 (00001111)? false
	// Has the oracle seen anything closed to 16 (00010000)? true
	// Has the oracle seen anything closed to 17 (00010001)? true
	// Has the oracle seen anything closed to 18 (00010010)? false
	// Has the oracle seen anything closed to 19 (00010011)? false
	// Has the oracle seen anything closed to 20 (00010100)? false
	// Has the oracle seen anything closed to 21 (00010101)? false
	// Has the oracle seen anything closed to 22 (00010110)? false
	// Has the oracle seen anything closed to 23 (00010111)? false
	// Has the oracle seen anything closed to 24 (00011000)? false
	// Has the oracle seen anything closed to 25 (00011001)? false
	// Has the oracle seen anything closed to 26 (00011010)? false
	// Has the oracle seen anything closed to 27 (00011011)? false
	// Has the oracle seen anything closed to 28 (00011100)? false
	// Has the oracle seen anything closed to 29 (00011101)? false
}

func BenchmarkOracleSee(b *testing.B) {
	oracle := NewOracle()
	for i := 0; i < b.N; i++ {
		// for i := uint64(1); i < 10000; i++ {
		oracle.See(uint64(i))
		// }
	}
}

func BenchmarkOracleSeen(b *testing.B) {
	oracle := NewOracle()
	for i := uint64(1); i < 1000000; i++ {
		oracle.See(i)
	}
	b.ResetTimer()
	r := uint8(2)
	for i := 0; i < b.N; i++ {
		oracle.Seen(uint64(i), r)
	}
}
