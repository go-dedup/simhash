// package main

package simhashCJK_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/simhashCJK"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_output() {
	hashes := make([]uint64, len(docs))
	sh := simhashCJK.NewSimhash()
	for i, d := range docs {
		fs := sh.NewWordFeatureSet(d)
		// fmt.Printf("%#v\n", fs)
		// actual := fs.GetFeatures()
		// fmt.Printf("%#v\n", actual)
		hashes[i] = sh.GetSimhash(fs)
		fmt.Printf("Simhash of '%s': %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[3], simhash.Compare(hashes[0], hashes[3]))

	// Output:
	// Simhash of '当山峰没有棱角的时候': d7185f186a2eea5a
	// Simhash of '当山谷没有棱角的时候': d71a5f186a2eea5a
	// Simhash of '棱角的时候': d71a5f186a2ffa52
	// Simhash of '你妈妈喊你回家吃饭哦，回家罗回家罗': d71bf7186a32b9f0
	// Comparison of `当山峰没有棱角的时候` and `当山谷没有棱角的时候`: 1
	// Comparison of `当山峰没有棱角的时候` and `棱角的时候`: 4
	// Comparison of `当山峰没有棱角的时候` and `你妈妈喊你回家吃饭哦，回家罗回家罗`: 16
}

var docs = [][]byte{
	[]byte("当山峰没有棱角的时候"),
	[]byte("当山谷没有棱角的时候"),
	[]byte("棱角的时候"),
	[]byte("你妈妈喊你回家吃饭哦，回家罗回家罗"),
}
