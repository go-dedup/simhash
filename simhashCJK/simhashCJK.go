////////////////////////////////////////////////////////////////////////////
// Package: simhashCJK
// Purpose: simhash language-specific handling for CJK
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

// simhashCJK -- simhash language-specific handling for CJK.
//
// This package is provided to showcase how easy it is to extend the simhash's language-specific handling functionality, even when CJK handling is dramatically different than Western Unicode handling.
//
package simhashCJK

import "github.com/go-dedup/simhash"

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

type SimhashCJK struct {
	simhash.SimhashBase
}

// CJKWordFeatureSet is a feature set in which each word is a feature,
// all equal weight.
type CJKWordFeatureSet struct {
	simhash.WordFeatureSet
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewSimhash makes a new Simhash
func NewSimhash() *SimhashCJK {
	return &SimhashCJK{}
}

func (st *SimhashCJK) NewWordFeatureSet(b []byte) *CJKWordFeatureSet {
	return st.NewCJKWordFeatureSet(b)
}

func (st *SimhashCJK) NewCJKWordFeatureSet(b []byte) *CJKWordFeatureSet {
	fs := &CJKWordFeatureSet{simhash.WordFeatureSet{b}}
	fs.Normalize()
	return fs
}

// // NewCJKSimhash makes a new SimhashCJK
// func NewCJKSimhash(_f norm.Form) *SimhashCJK {}

// Returns a []Feature representing each word in the byte slice
func (w *CJKWordFeatureSet) GetFeatures() []simhash.Feature {
	words := string(w.B)
	// fmt.Printf("%#v\n", words)
	i := 0
	features := make([]simhash.Feature, len([]rune(words)))
	for _, w := range words {
		features[i] = simhash.NewFeature([]byte(string(w)))
		i++
	}
	return features
}
