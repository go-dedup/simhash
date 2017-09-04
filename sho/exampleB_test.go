package sho_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/sho"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleFind() {
	var docs = []string{
		"GNU - Python Standard Library (2001)",
		"(eBook) GNU - Python Standard Library 2001",
		"Python Standard Library",
		"GNU - 2001 - Python Standard Library",
	}

	// Code starts

	oracle := sho.NewOracle()
	sh := simhash.NewSimhash()
	r := uint8(18)
	for _, d := range docs {
		hash := sh.BuildSimhash(d, simhash.Doc2words)
		if oracle.Seen(hash, r) {
			fmt.Printf("=: Simhash of %x for '%s' ignored.\n", hash, d)
		} else {
			oracle.See(hash)
			fmt.Printf("+: Simhash of %x for '%s' added.\n", hash, d)
		}
	}

	fmt.Println("================")
	oracle = sho.NewOracle()
	// r = uint8(8)
	for _, d := range docs {
		fmt.Printf(">: %v\n", simhash.Doc2words(d))
		hash := sh.BuildSimhash(d, simhash.Doc2words)
		if h, nd, seen := oracle.Find(hash, r); seen == true {
			fmt.Printf("=: Simhash of %x ignored for %x (%d).\n", hash, h, nd)
		} else {
			oracle.See(hash)
			fmt.Printf("+: Simhash of %x added.\n", hash)
		}
	}

	// fmt.Println("================")
	// for _, d := range docs {
	// 	hash := sh.BuildSimhash(d, simhash.Doc2words)
	// 	if n := oracle.Search(hash, r); len(n) > 0 {
	// 		fmt.Printf("!: Similiar found for %x (%v).\n", hash, n)
	// 	}
	// }

	// Code ends

	// Output:
	// +: Simhash of 55d4263ae1a6e6d6 for 'GNU - Python Standard Library (2001)' added.
	// =: Simhash of d5d6363ef9e6e6d7 for '(eBook) GNU - Python Standard Library 2001' ignored.
	// =: Simhash of 55b47e2af1a4a4d2 for 'Python Standard Library' ignored.
	// =: Simhash of 55d4263ae1a6e6d6 for 'GNU - 2001 - Python Standard Library' ignored.
	// ================
	// >: [gnu python standard library 2001]
	// +: Simhash of 55d4263ae1a6e6d6 added.
	// >: [ebook gnu python standard library 2001]
	// =: Simhash of d5d6363ef9e6e6d7 ignored for 55d4263ae1a6e6d6 (8).
	// >: [python standard library]
	// =: Simhash of 55b47e2af1a4a4d2 ignored for 55d4263ae1a6e6d6 (11).
	// >: [gnu 2001 python standard library]
	// =: Simhash of 55d4263ae1a6e6d6 ignored for 55d4263ae1a6e6d6 (0).
}

// ================
// !: Similiar found for f733d6ea421279e ([{5617743756152481750 12}]).
// !: Similiar found for 55b47e2af1a4a4d2 ([{5617743756152481750 17}]).
// !: Similiar found for 5df77f2ef1e3afde ([{5617743756152481750 12}]).
