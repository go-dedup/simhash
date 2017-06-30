
# simhash

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-dedup/simhash?status.svg)](http://godoc.org/github.com/go-dedup/simhash)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-dedup/simhash)](https://goreportcard.com/report/github.com/go-dedup/simhash)
[![travis Status](https://travis-ci.org/go-dedup/simhash.svg?branch=master)](https://travis-ci.org/go-dedup/simhash)

## TOC
- [simhash - Go simhash package](#simhash---go-simhash-package)
- [Installation](#installation)
- [Usage](#usage)
- [API](#api)
  - [> example_test.go](#-example_testgo)
- [TODO](#todo)
  - [> chinese_test.go](#-chinese_testgo)
- [Credits](#credits)
- [Similar Projects](#similar-projects)

## simhash - Go simhash package

`simhash` is a [Go](http://golang.org/) implementation of Charikar's [simhash](http://www.cs.princeton.edu/courses/archive/spring04/cos598B/bib/CharikarEstim.pdf) algorithm.

`simhash` is a hash with the useful property that similar documents produce similar hashes.
Therefore, if two documents are similar, the Hamming-distance between the simhash of the
documents will be small.

This package only implements the simhash algorithm. To make use of this
package to enable quickly identifying near-duplicate documents within a large collection of
documents, check out the `sho` (SimHash Oracle) package at [github.com/go-dedup/simhash/sho](https://github.com/go-dedup/simhash/tree/master/sho). It has a simple [API](https://github.com/go-dedup/simhash/tree/master/sho#api) that is easy to use. 

# Installation

```
go get github.com/go-dedup/simhash
```

# Usage

Using `simhash` first requires tokenizing a document into a set of features (done through the
`FeatureSet` interface). This package provides an implementation, `WordFeatureSet`, which breaks
tokenizes the document into individual words. Better results are possible here, and future work
will go towards this.

# API

Example usage:

#### > example_test.go
```go
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
		[]byte("these are test phrases"),
		[]byte("foo bar"),
	}

	hashes := make([]uint64, len(docs))
	for i, d := range docs {
		hashes[i] = simhash.Simhash(simhash.NewWordFeatureSet(d))
		fmt.Printf("Simhash of '%s': %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[3], simhash.Compare(hashes[0], hashes[3]))

	// Output:
	// Simhash of 'this is a test phrase': 8c3a5f7e9ecb3f35
	// Simhash of 'this is a test phrass': 8c3a5f7e9ecb3f21
	// Simhash of 'these are test phrases': ddfdbf7fbfaffb1d
	// Simhash of 'foo bar': d8dbe7186bad3db3
	// Comparison of `this is a test phrase` and `this is a test phrass`: 2
	// Comparison of `this is a test phrase` and `these are test phrases`: 22
	// Comparison of `this is a test phrase` and `foo bar`: 29

}
```

All patches welcome.

# TODO

It does not support Chinese very well:

#### > chinese_test.go
```go
package simhash_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/sho"

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
	for i, d := range docs {
		hashes[i] = simhash.Simhash(simhash.NewUnicodeWordFeatureSet(d, norm.NFC))
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
```

All patches welcome.

## Credits

- [mfonda/simhash](https://github.com/mfonda/simhash) forked source

The most high quality open-source Go simhash implementation available. it is even [used internally by Yahoo Inc](https://github.com/yahoo/gryffin/tree/master/html-distance):

[![Yahoo Inc](https://avatars3.githubusercontent.com/u/16574?v=3&s=200)](https://github.com/yahoo)


## Similar Projects

All the following similar projects have been considered before adopting [mfonda/simhash](https://github.com/mfonda/simhash) instead.

- [dgryski/go-simstore](https://github.com/dgryski/go-simstore) One of the earliest but ["_not very promising_"](https://groups.google.com/forum/#!msg/golang-nuts/E9UVskCnSJc/gm7KF27LnI0J)
- [AllenDang/simhash](https://github.com/AllenDang/simhash) Ported from C# code, but don't like its interface
- [yanyiwu/gosimhash](https://github.com/yanyiwu/gosimhash) For Chinese only. Don't like keeping two packages for the same purpose, and don't like its dependency on "结巴"中文分词 approach
