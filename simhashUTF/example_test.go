// package main

package simhashUTF_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/simhashUTF"
	"golang.org/x/text/unicode/norm"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_output() {
	hashes := make([]uint64, len(docs))
	sh := simhashUTF.NewUTFSimhash(norm.NFKC)
	for i, d := range docs {
		hashes[i] = sh.GetSimhash(sh.NewWordFeatureSet(d))
		fmt.Printf("Simhash of '%s': %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[3], simhash.Compare(hashes[0], hashes[3]))

	// Output:
	// Simhash of 'la fin d'un bel après-midi d'été': 58dbbd1fefab774a
	// Simhash of 'bonne après-midi': fadfbfbfdf8e7b7f
	// Simhash of 'Bonjour': ac5261af4fdd5252
	// Simhash of 'Bonsoir': fb42ceaf7cda4905
	// Comparison of `la fin d'un bel après-midi d'été` and `bonne après-midi`: 18
	// Comparison of `la fin d'un bel après-midi d'été` and `Bonjour`: 28
	// Comparison of `la fin d'un bel après-midi d'été` and `Bonsoir`: 34
}

var docs = [][]byte{
	[]byte("la fin d'un bel après-midi d'été"),
	[]byte("bonne après-midi"),
	[]byte("Bonjour"),
	[]byte("Bonsoir"),
}
