package simhashUTF

import (
	"bytes"
	"regexp"

	"github.com/go-dedup/simhash"
	"golang.org/x/text/unicode/norm"
)

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

type SimhashUTFT struct {
	simhash.SimhashT
	f norm.Form
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var unicodeBoundaries = regexp.MustCompile(`[\pL-_']+`)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewSimhash makes a new Simhash
func NewSimhash() *SimhashUTFT {
	return &SimhashUTFT{}
}

// NewUTFSimhash makes a new SimhashUTF
func NewUTFSimhash(_f norm.Form) *SimhashUTFT {
	return &SimhashUTFT{f: _f}
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

func (st *SimhashUTFT) NewWordFeatureSet(b []byte) *UnicodeWordFeatureSet {
	return st.NewUnicodeWordFeatureSet(b, st.f)
}

func (st *SimhashUTFT) NewUnicodeWordFeatureSet(b []byte, f norm.Form) *UnicodeWordFeatureSet {
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
