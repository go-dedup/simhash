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
