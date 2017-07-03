////////////////////////////////////////////////////////////////////////////
// Package: simhashEng
// Purpose: simhash language-specific handling for English
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

// simhashEng -- simhash language-specific handling for English.
//
// This package is provided to showcase how easy it is to extend the simhash's language-specific handling functionality.
//
package simhashEng

import (
	"github.com/go-dedup/simhash"
)

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

type SimhashEng struct {
	simhash.SimhashBase
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewSimhash makes a new Simhash
func NewSimhash() *SimhashEng {
	return &SimhashEng{}
}
