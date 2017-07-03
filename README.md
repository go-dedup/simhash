
# simhash

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-dedup/simhash?status.svg)](http://godoc.org/github.com/go-dedup/simhash)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-dedup/simhash)](https://goreportcard.com/report/github.com/go-dedup/simhash)
[![travis Status](https://travis-ci.org/go-dedup/simhash.svg?branch=master)](https://travis-ci.org/go-dedup/simhash)

## TOC
- [simhash - Go simhash package](#simhash---go-simhash-package)
  - [Design principle](#design-principle)
- [Installation](#installation)
- [Usage](#usage)
- [API](#api)
  - [> example_test.go](#-example_testgo)
- [Purpose](#purpose)
- [Versions](#versions)
  - [> simhashUTF/chinese_test.go](#-simhashutfchinese_testgo)
  - [> simhashCJK/example_test.go](#-simhashcjkexample_testgo)
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

### Design principle

The design principle of these packages follows the ["Unix philosophy"](https://en.wikipedia.org/wiki/Unix_philosophy): "_Do One Thing and Do It Well_". Thus the storing & checking, and different language handling are available in different building blocks, and can be added on request, or substituted at will, keeping the size of the core code minimum.

Thus, you can use exactly what you want to use without being forced to accept a huge package with features you don't want.

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
//package main

package simhash_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_output() {
	hashes := make([]uint64, len(docs))
	sh := simhash.NewSimhash()
	for i, d := range docs {
		hashes[i] = sh.GetSimhash(sh.NewWordFeatureSet(d))
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

var docs = [][]byte{
	[]byte("this is a test phrase"),
	[]byte("this is a test phrass"),
	[]byte("these are test phrases"),
	[]byte("foo bar"),
}
```

All patches welcome.

# Purpose

A few more words on the similarity checking and near-duplicate detection. The best article I found explaining it clearly is:

**Near-Duplicate Detection**  
https://moz.com/devblog/near-duplicate-detection/

This article, from the Moz Developer Blog, explained in details and in graph that,

- Why Does Duplication Matter
- What to Do About It, and
- How to Identify Duplication

and it went on to explain the different algorithms to do so.

Among the algorithms that solve the problem the best, one is `MinHash`, which is [the first that I tried](https://github.com/go-dedup/deduper), but found it to be bloated, cumbersome to use, and not working as I expected. The other one is `SimHash`, which is what all these are about. `SimHash` is designed by Google. It is simple, straightforward, thus very efficient and powerful. I like it very much, and should have used it in the first place.

FYI, [this is why I needed and looked for such similarity checking and near-duplicate detection algorithms in the first place](http://godoc.org/github.com/go-dedup/simhash/sho#example-package--Output) -- in the world that we cannot avoid the rule-breakers and spammers, at least we can use technologies to get rid of them for ourselves.

# Versions

Having forked from [mfonda/simhash](https://github.com/mfonda/simhash), `go-dedup/simhash` has been through a serious of interface changes. Detailed documents of such changes, and the reasons behind it, also how to use the original (v1) design API [can be found here](https://github.com/go-dedup/simhash/wiki/Version-2). 

The key characteristics of current design are,

- most of `simhash` related functions are provided as method(/member) functions of `SimhashBase` type(/class), as oppose to package functions before.
- and also very importantly, the `UnicodeWordFeatureSet` related functions no longer exist in core code any more, because
- the language-specific handling have been refactored out to a thin language handling layer.
- the goal of version 2 is to have different languages to have a unified user interface (API).

Such modular approach (v2 design) helps to reduce and limit the size of the core code, while make it easy to extend the core function with easy to use building blocks.

The added bonus is that, the original (v1) design does not support Chinese very well:

#### > simhashUTF/chinese_test.go
```go
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
```

[The result of similarity checking on Chinese text is very bad](https://github.com/go-dedup/simhash/blob/master/simhashUTF/chinese_test.go#L55-L58). But thanks to version 2's architecture, it is very easy to extend `simhash` to deal with Chinese:

#### > simhashCJK/example_test.go
```go
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
```

With the above, now the problem has been fix. [Check the result here](https://github.com/go-dedup/simhash/blob/master/simhashCJK/example_test.go#L35-L37).

## Credits

- [mfonda/simhash](https://github.com/mfonda/simhash) forked source

The most high quality open-source Go simhash implementation available. it is even [used internally by Yahoo Inc](https://github.com/yahoo/gryffin/tree/master/html-distance):

[![Yahoo Inc](https://avatars3.githubusercontent.com/u/16574?v=3&s=200)](https://github.com/yahoo)


## Similar Projects

All the following similar projects have been considered before adopting [mfonda/simhash](https://github.com/mfonda/simhash) instead.

- [dgryski/go-simstore](https://github.com/dgryski/go-simstore) One of the earliest but ["_not very promising_"](https://groups.google.com/forum/#!msg/golang-nuts/E9UVskCnSJc/gm7KF27LnI0J)
- [AllenDang/simhash](https://github.com/AllenDang/simhash) Ported from C# code, but don't like its interface
- [yanyiwu/gosimhash](https://github.com/yanyiwu/gosimhash) For Chinese only. Don't like keeping two packages for the same purpose, and don't like its dependency on "结巴"中文分词 approach
