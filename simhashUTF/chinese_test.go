package simhashUTF_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/sho"
	"github.com/go-dedup/simhash/simhashUTF"

	"golang.org/x/text/unicode/norm"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_Chinese_output() {
	var docs = [][]byte{
		[]byte("当山峰没有棱角的时候"),
		[]byte("当山谷没有棱角的时候"),
		[]byte("棱角的时候"),
		[]byte("你妈妈喊你回家吃饭哦，回家罗回家罗"),
		[]byte("你妈妈叫你回家吃饭啦，回家罗回家罗"),
	}

	// Code starts

	oracle := sho.NewOracle()
	r := uint8(3)
	hashes := make([]uint64, len(docs))
	sh := simhashUTF.NewUTFSimhash(norm.NFKC)
	for i, d := range docs {
		hashes[i] = sh.GetSimhash(sh.NewUnicodeWordFeatureSet(d, norm.NFC))
		hash := hashes[i]
		if oracle.Seen(hash, r) {
			fmt.Printf("=: Simhash of %x for '%s' ignored.\n", hash, d)
		} else {
			oracle.See(hash)
			fmt.Printf("+: Simhash of %x for '%s' added.\n", hash, d)
		}
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[3], simhash.Compare(hashes[0], hashes[3]))

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[3], docs[4], simhash.Compare(hashes[0], hashes[1]))

	// Code ends

	// Output:
	// +: Simhash of a5edea16c0c7a180 for '当山峰没有棱角的时候' added.
	// +: Simhash of 2e285bd230856c9 for '当山谷没有棱角的时候' added.
	// +: Simhash of 53ecd232f2383dee for '棱角的时候' added.
	// +: Simhash of e4e6edb1f89fa9ff for '你妈妈喊你回家吃饭哦，回家罗回家罗' added.
	// +: Simhash of ffe1e5ffffd7b9e7 for '你妈妈叫你回家吃饭啦，回家罗回家罗' added.
	// Comparison of `当山峰没有棱角的时候` and `当山谷没有棱角的时候`: 41
	// Comparison of `当山峰没有棱角的时候` and `棱角的时候`: 32
	// Comparison of `当山峰没有棱角的时候` and `你妈妈喊你回家吃饭哦，回家罗回家罗`: 27
	// Comparison of `你妈妈喊你回家吃饭哦，回家罗回家罗` and `你妈妈叫你回家吃饭啦，回家罗回家罗`: 41
}
