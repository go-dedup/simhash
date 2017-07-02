package simhashEng

import (
	"github.com/go-dedup/simhash"
)

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

type SimhashEngT struct {
	simhash.SimhashT
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewSimhash makes a new Simhash
func NewSimhash() *SimhashEngT {
	return &SimhashEngT{}
}
