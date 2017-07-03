// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhashUTF

import (
	"fmt"
	"testing"

	"github.com/go-dedup/simhash"
	"golang.org/x/text/unicode/norm"
)

func TestUnicodeWordFeatureSet(t *testing.T) {
	sh := NewSimhash()
	text := []byte("la fin d'un bel après-midi d'été")
	fs := sh.NewUnicodeWordFeatureSet(text, norm.NFKC)
	expected := []simhash.Feature{
		simhash.NewFeature([]byte("la")),
		simhash.NewFeature([]byte("fin")),
		simhash.NewFeature([]byte("d'un")),
		simhash.NewFeature([]byte("bel")),
		simhash.NewFeature([]byte("après-midi")),
		simhash.NewFeature([]byte("d'été")),
	}

	actual := fs.GetFeatures()

	for i := 0; i < len(actual); i++ {
		if actual[i].Sum() != expected[i].Sum() {
			t.Errorf("feature.Sum(): expected %d, actual %d", expected[i].Sum(), actual[i].Sum())
		}
		if actual[i].Weight() != expected[i].Weight() {
			t.Errorf("feature.Weight(): expected %d, actual %d", expected[i].Weight(), actual[i].Weight())
		}
	}
}

func ExampleSimhashUTF_NewUnicodeWordFeatureSet_inWestern() {
	sh := NewSimhash()
	text := []byte("la fin d'un bel après-midi d'été")
	fs := sh.NewUnicodeWordFeatureSet(text, norm.NFKC)
	fmt.Printf("%#v\n", fs)
	actual := fs.GetFeatures()
	fmt.Printf("%#v\n", actual)
	// Output:
	// &simhashUTF.UnicodeWordFeatureSet{WordFeatureSet:simhash.WordFeatureSet{B:[]uint8{0x6c, 0x61, 0x20, 0x66, 0x69, 0x6e, 0x20, 0x64, 0x27, 0x75, 0x6e, 0x20, 0x62, 0x65, 0x6c, 0x20, 0x61, 0x70, 0x72, 0xc3, 0xa8, 0x73, 0x2d, 0x6d, 0x69, 0x64, 0x69, 0x20, 0x64, 0x27, 0xc3, 0xa9, 0x74, 0xc3, 0xa9}}, f:2}
	// []simhash.Feature{simhash.feature{sum:0x8325c07b4eb2548, weight:1}, simhash.feature{sum:0xd8cbc5186ba13198, weight:1}, simhash.feature{sum:0x15cdbd7eed98cfab, weight:1}, simhash.feature{sum:0xd8d9a1186bad324a, weight:1}, simhash.feature{sum:0x3adb901f8c8a7b5e, weight:1}, simhash.feature{sum:0x7e8f29c36ffb774e, weight:1}}
}

func ExampleSimhashUTF_NewUnicodeWordFeatureSet_inChinese() {
	sh := NewSimhash()
	text := []byte("当山峰没有棱角的时候")
	fs := sh.NewUnicodeWordFeatureSet(text, norm.NFKC)
	fmt.Printf("%#v\n", fs)
	actual := fs.GetFeatures()
	fmt.Printf("%#v\n", actual)
	// Output:
	// &simhashUTF.UnicodeWordFeatureSet{WordFeatureSet:simhash.WordFeatureSet{B:[]uint8{0xe5, 0xbd, 0x93, 0xe5, 0xb1, 0xb1, 0xe5, 0xb3, 0xb0, 0xe6, 0xb2, 0xa1, 0xe6, 0x9c, 0x89, 0xe6, 0xa3, 0xb1, 0xe8, 0xa7, 0x92, 0xe7, 0x9a, 0x84, 0xe6, 0x97, 0xb6, 0xe5, 0x80, 0x99}}, f:2}
	// []simhash.Feature{simhash.feature{sum:0xa5edea16c0c7a180, weight:1}}
}
