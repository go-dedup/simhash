// Copyright 2015, Yahoo Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// sho -- SimHash Oracle, checks if a fingerprint is similar to existing ones. It uses BK Tree (Burkhard and Keller) for storing and verifying if a fingerprint is closed to a set of fingerprint within a defined proximity distance.
//
// Distance is the hamming distance of the fingerprints.
//
// It is forked from from Yahoo Inc's github.com/yahoo/gryffin/html-distance, which was meant to be a go library for computing the proximity of the HTML pages. It's plan was to implementate the similiarity fingerprint of Charikar's simhash, and also:
//
// - Since fingerprint is of size 64 (inherited from hash/fnv), Similiarity is defined as 1 - d / 64.
//
// - In normal scenario, similarity > 95% (i.e. d<3) could be considered as duplicated html pages.
package sho

import (
	"github.com/go-dedup/simhash"
)

// Oracle answers the query if a fingerprint has been seen.
type Oracle struct {
	fingerprint uint64      // node value.
	nodes       [65]*Oracle // leaf nodes
}

type Nigh struct {
	H uint64 // hash
	D uint8  // distance
}

// NewOracle return an oracle that could tell if the fingerprint has been seen or not.
func NewOracle() *Oracle {
	return newNode(0)
}

func newNode(f uint64) *Oracle {
	return &Oracle{fingerprint: f}
}

// Distance return the similarity distance between two fingerprint.
func Distance(a, b uint64) uint8 {
	return simhash.Compare(a, b)
}

// See asks the oracle to see the fingerprint.
func (n *Oracle) See(f uint64) *Oracle {
	d := Distance(n.fingerprint, f)

	if d == 0 {
		// current node with same fingerprint.
		return n
	}

	// the target node is already set,
	if c := n.nodes[d]; c != nil {
		return c.See(f)
	}

	n.nodes[d] = newNode(f)
	return n.nodes[d]
}

// Seen asks the oracle if anything closed to the fingerprint in a range (r) is seen before.
func (n *Oracle) Seen(f uint64, r uint8) bool {
	_, _, seen := n.Find(f, r)
	return seen
}

// Find searches the oracle for one closed to the fingerprint in a range (r).
func (n *Oracle) Find(f uint64, r uint8) (uint64, uint8, bool) {
	d := Distance(n.fingerprint, f)
	if d < r {
		return n.fingerprint, d, true
	}

	k := d - r
	if k < 1 {
		k = 1
	}
	for ; k <= d+r; k++ {
		if k > 64 {
			break
		}
		if c := n.nodes[k]; c != nil {
			//print(f, " ", k, " ", d, " > ", c.fingerprint, "\n")
			if h, nd, seen := c.Find(f, r); seen == true {
				return h, nd, seen
			}
		}
	}
	return 0, 0, false
}

// Search searches the oracle for all fingerprints within the range (r).
func (n *Oracle) Search(f uint64, r uint8) []Nigh {
	matches := []Nigh{}
	// calculate the distance
	d := Distance(n.fingerprint, f)
	// if dist is less than tolerance value add it to similar matches
	if d < r && d != 0 {
		matches = append(matches, Nigh{n.fingerprint, d})
	}

	// iterate over the rest havinng tolerane in range (dist-TOL , dist+TOL)
	k := d - r
	if k < 1 {
		k = 1
	}
	for ; k <= d+r; k++ {
		if k > 64 {
			break
		}
		if c := n.nodes[k]; c != nil {
			t := c.Search(f, r)
			matches = append(matches, t...)
		}
	}
	return matches
}
