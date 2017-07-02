// package main

package simhashUTF_test

import (
	"fmt"

	"github.com/go-dedup/simhash/simhashUTF"
)

// to show the full code in GoDoc
type dummy struct {
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_output() {
	var docs = [][]byte{
		[]byte("this is a test phrase"),
		[]byte("this is a test phrass"),
		[]byte("these are test phrases"),
		[]byte("foo bar"),
	}

	hashes := make([]uint64, len(docs))
	sh := simhashUTF.NewSimhash()
	for i, d := range docs {
		hashes[i] = sh.GetSimhash(sh.NewWordFeatureSet(d))
		fmt.Printf("Simhash of '%s': %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], sh.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], sh.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[3], sh.Compare(hashes[0], hashes[3]))

	// Output:
	// Simhash of 'this is a test phrase': 8c3a5f7e9ecb3f35
	// Simhash of 'this is a test phrass': 8c3a5f7e9ecb3f21
	// Simhash of 'these are test phrases': ddfdbf7fbfaffb1d
	// Simhash of 'foo bar': d8dbe7186bad3db3
	// Comparison of `this is a test phrase` and `this is a test phrass`: 2
	// Comparison of `this is a test phrase` and `these are test phrases`: 22
	// Comparison of `this is a test phrase` and `foo bar`: 29

}
