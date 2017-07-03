////////////////////////////////////////////////////////////////////////////
// Package: simhashUTF
// Purpose: simhash language-specific handling for UTF
// Authors: Tong Sun (c) 2017, All rights reserved
//          Matthew Fonda (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// simhashUTF -- simhash language-specific handling for UTF.
//
// This package is to refactor the Unicode handling code from the original (v1) design out to this thin language handling layer, which showcases how easy it is to extend the simhash's language-specific handling functionality.
//
// Such modular approach (v2 design) helps to reduce and limit the size of the core code, while make it easy to extend the core function as well.
//
package simhashUTF

import (
	"bytes"
	"regexp"

	"github.com/go-dedup/simhash"
	"golang.org/x/text/unicode/norm"
)

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

type SimhashUTF struct {
	simhash.SimhashBase
	f norm.Form
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var unicodeBoundaries = regexp.MustCompile(`[\pL-_']+`)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewSimhash makes a new Simhash
func NewSimhash() *SimhashUTF {
	return &SimhashUTF{}
}

// NewUTFSimhash makes a new SimhashUTF
func NewUTFSimhash(_f norm.Form) *SimhashUTF {
	return &SimhashUTF{f: _f}
}

// UnicodeWordFeatureSet is a feature set in which each word is a feature,
// all equal weight.
//
// See: http://blog.golang.org/normalization
// See: https://groups.google.com/forum/#!topic/golang-nuts/YyH1f_qCZVc
type UnicodeWordFeatureSet struct {
	simhash.WordFeatureSet
	f norm.Form
}

func (st *SimhashUTF) NewWordFeatureSet(b []byte) *UnicodeWordFeatureSet {
	return st.NewUnicodeWordFeatureSet(b, st.f)
}

func (st *SimhashUTF) NewUnicodeWordFeatureSet(b []byte, f norm.Form) *UnicodeWordFeatureSet {
	fs := &UnicodeWordFeatureSet{simhash.WordFeatureSet{b}, f}
	fs.Normalize()
	return fs
}

func (w *UnicodeWordFeatureSet) Normalize() {
	b := bytes.ToLower(w.f.Append(nil, w.B...))
	w.WordFeatureSet.B = b
}

// Returns a []Feature representing each word in the byte slice
func (w *UnicodeWordFeatureSet) GetFeatures() []simhash.Feature {
	return simhash.DoGetFeatures(w.B, unicodeBoundaries)
}
