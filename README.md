
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
- [Credits](#credits)
- [Similar Projects](#similar-projects)

## simhash - Go simhash package

`simhash` is a [Go](http://golang.org/) implementation of Charikar's [simhash](http://www.cs.princeton.edu/courses/archive/spring04/cos598B/bib/CharikarEstim.pdf) algorithm.

`simhash` is a hash with the useful property that similar documents produce similar hashes.
Therefore, if two documents are similar, the Hamming-distance between the simhash of the
documents will be small.

This package currently just implements the simhash algorithm. Future work will make use of this
package to enable quickly identifying near-duplicate documents within a large collection of
documents.

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

## Credits

- [mfonda/simhash](https://github.com/mfonda/simhash) forked source

## Similar Projects

All the following similar projects have been considered before adopting [mfonda/simhash](https://github.com/mfonda/simhash) instead.

- [dgryski/go-simstore](https://github.com/dgryski/go-simstore) One of the earliest but ["_not very promising_"](https://groups.google.com/forum/#!msg/golang-nuts/E9UVskCnSJc/gm7KF27LnI0J)
- [AllenDang/simhash](https://github.com/AllenDang/simhash) Ported from C# code, but don't like its interface
- [yanyiwu/gosimhash](https://github.com/yanyiwu/gosimhash) For Chinese only. Don't like keeping two packages for the same purpose, and don't like its dependency on "结巴"中文分词 approach
