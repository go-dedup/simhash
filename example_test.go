package simhash_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
)

// for standalone test, change package to main and the next func def to,
// func main() {
func Example_output() {
	var docs = [][]byte{
		[]byte("this is a test phrase"),
		[]byte("this is a test phrass"),
		[]byte("foo bar"),
	}

	hashes := make([]uint64, len(docs))
	for i, d := range docs {
		hashes[i] = simhash.Simhash(simhash.NewWordFeatureSet(d))
		fmt.Printf("Simhash of %s: %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))

	// Output:
	// Simhash of this is a test phrase: 8c3a5f7e9ecb3f35
	// Simhash of this is a test phrass: 8c3a5f7e9ecb3f21
	// Simhash of foo bar: d8dbe7186bad3db3
	// Comparison of `this is a test phrase` and `this is a test phrass`: 2
	// Comparison of `this is a test phrase` and `foo bar`: 29

}
