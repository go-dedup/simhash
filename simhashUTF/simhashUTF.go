package simhashUTF

import (
	"github.com/go-dedup/simhash"
)

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

type SimhashUTFT struct {
	simhash.SimhashT
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewSimhash makes a new Simhash
func NewSimhash() *SimhashUTFT {
	return &SimhashUTFT{}
}
