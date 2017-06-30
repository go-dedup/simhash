
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

## {{.Name}} - Go simhash package

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

#### > {{cat "example_test.go" | color "go"}}

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
