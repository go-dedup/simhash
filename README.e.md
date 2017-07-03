
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

#### > {{cat "example_test.go" | color "go"}}

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

#### > {{cat "simhashUTF/chinese_test.go" | color "go"}}

[The result of similarity checking on Chinese text is very bad](https://github.com/go-dedup/simhash/blob/master/simhashUTF/chinese_test.go#L55-L58). But thanks to version 2's architecture, it is very easy to extend `simhash` to deal with Chinese:

#### > {{cat "simhashCJK/example_test.go" | color "go"}}

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
