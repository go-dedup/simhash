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
func ExampleBuildSimhash() {
	hashes := make([]uint64, len(docs))
	sh := simhashUTF.NewUTFSimhash(norm.NFKC)
	for i, d := range docs {
		hashes[i] = sh.BuildSimhash(string(d), simhash.Doc2words)
		fmt.Printf("Simhash of '%s': %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[3], simhash.Compare(hashes[0], hashes[3]))

	// Output:
	// Simhash of 'la fin d'un bel après-midi d'été': 8a73bd4c862137a8
	// Simhash of 'bonne après-midi': cef7bd7ec38ff5ac
	// Simhash of 'Bonjour': ac5261af4fdd5252
	// Simhash of 'Bonsoir': fb42ceaf7cda4905
	// Comparison of `la fin d'un bel après-midi d'été` and `bonne après-midi`: 19
	// Comparison of `la fin d'un bel après-midi d'été` and `Bonjour`: 35
	// Comparison of `la fin d'un bel après-midi d'été` and `Bonsoir`: 41
}

// Note the output is different from the original:

// Simhash of 'la fin d'un bel après-midi d'été': 58dbbd1fefab774a
// Simhash of 'bonne après-midi': fadfbfbfdf8e7b7f
// Simhash of 'Bonjour': ac5261af4fdd5252
// Simhash of 'Bonsoir': fb42ceaf7cda4905
// Comparison of `la fin d'un bel après-midi d'été` and `bonne après-midi`: 18
// Comparison of `la fin d'un bel après-midi d'été` and `Bonjour`: 28
// Comparison of `la fin d'un bel après-midi d'été` and `Bonsoir`: 34

var exampleB_dummy bool
