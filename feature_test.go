// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"fmt"
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestNewFeature(t *testing.T) {
	expected := uint64(8811532157352841348)
	f := NewFeature([]byte("test string"))

	if f.Weight() != 1 {
		t.Errorf("feature.Weight(): expected 1, actual %d", f.Weight())
	}

	if f.Sum() != expected {
		t.Errorf("feature.Sum(): expected %d, actual %d", expected, f.Sum())
	}
}

func TestNewFeatureWithWeight(t *testing.T) {
	weight := 10
	expected := uint64(8811532157352841348)
	f := NewFeatureWithWeight([]byte("test string"), weight)

	if f.Weight() != weight {
		t.Errorf("feature.Weight(): expected %d, actual %d", weight, f.Weight())
	}

	if f.Sum() != expected {
		t.Errorf("feature.Sum(): expected %d, actual %d", expected, f.Sum())
	}
}

func TestFeatureSet(t *testing.T) {
	text := []byte("here's a test string.")
	fs := NewWordFeatureSet(text)
	expected := []Feature{
		NewFeature([]byte("here's")),
		NewFeature([]byte("a")),
		NewFeature([]byte("test")),
		NewFeature([]byte("string")),
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

func ExampleNewWordFeatureSet() {
	text := []byte("a a abc abc test test string.")
	fs := NewWordFeatureSet(text)
	fmt.Printf("%#v\n", fs)
	actual := fs.GetFeatures()
	fmt.Printf("%#v\n", actual)
	// Output:
	// &simhash.WordFeatureSet{b:[]uint8{0x61, 0x20, 0x61, 0x20, 0x61, 0x62, 0x63, 0x20, 0x61, 0x62, 0x63, 0x20, 0x74, 0x65, 0x73, 0x74, 0x20, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e}}
	// []simhash.Feature{simhash.feature{sum:0xaf63bd4c8601b7be, weight:1}, simhash.feature{sum:0xaf63bd4c8601b7be, weight:1}, simhash.feature{sum:0xd8dcca186bafadcb, weight:1}, simhash.feature{sum:0xd8dcca186bafadcb, weight:1}, simhash.feature{sum:0x8c093f7e9fccbf69, weight:1}, simhash.feature{sum:0x8c093f7e9fccbf69, weight:1}, simhash.feature{sum:0x9926dcde0a17d48e, weight:1}}
}

func TestUnicodeWordFeatureSet(t *testing.T) {
	text := []byte("la fin d'un bel après-midi d'été")
	fs := NewUnicodeWordFeatureSet(text, norm.NFKC)
	expected := []Feature{
		NewFeature([]byte("la")),
		NewFeature([]byte("fin")),
		NewFeature([]byte("d'un")),
		NewFeature([]byte("bel")),
		NewFeature([]byte("après-midi")),
		NewFeature([]byte("d'été")),
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

func ExampleNewUnicodeWordFeatureSet_inWestern() {
	text := []byte("la fin d'un bel après-midi d'été")
	fs := NewUnicodeWordFeatureSet(text, norm.NFKC)
	fmt.Printf("%#v\n", fs)
	actual := fs.GetFeatures()
	fmt.Printf("%#v\n", actual)
	// Output:
	// &simhash.UnicodeWordFeatureSet{b:[]uint8{0x6c, 0x61, 0x20, 0x66, 0x69, 0x6e, 0x20, 0x64, 0x27, 0x75, 0x6e, 0x20, 0x62, 0x65, 0x6c, 0x20, 0x61, 0x70, 0x72, 0xc3, 0xa8, 0x73, 0x2d, 0x6d, 0x69, 0x64, 0x69, 0x20, 0x64, 0x27, 0xc3, 0xa9, 0x74, 0xc3, 0xa9}, f:2}
	// []simhash.Feature{simhash.feature{sum:0x8325c07b4eb2548, weight:1}, simhash.feature{sum:0xd8cbc5186ba13198, weight:1}, simhash.feature{sum:0x15cdbd7eed98cfab, weight:1}, simhash.feature{sum:0xd8d9a1186bad324a, weight:1}, simhash.feature{sum:0x3adb901f8c8a7b5e, weight:1}, simhash.feature{sum:0x7e8f29c36ffb774e, weight:1}}
}

func ExampleNewUnicodeWordFeatureSet_inChinese() {
	text := []byte("当山峰没有棱角的时候")
	fs := NewUnicodeWordFeatureSet(text, norm.NFKC)
	fmt.Printf("%#v\n", fs)
	actual := fs.GetFeatures()
	fmt.Printf("%#v\n", actual)
	// Output:
	// &simhash.UnicodeWordFeatureSet{b:[]uint8{0xe5, 0xbd, 0x93, 0xe5, 0xb1, 0xb1, 0xe5, 0xb3, 0xb0, 0xe6, 0xb2, 0xa1, 0xe6, 0x9c, 0x89, 0xe6, 0xa3, 0xb1, 0xe8, 0xa7, 0x92, 0xe7, 0x9a, 0x84, 0xe6, 0x97, 0xb6, 0xe5, 0x80, 0x99}, f:2}
	// []simhash.Feature{simhash.feature{sum:0xa5edea16c0c7a180, weight:1}}
}

func TestGetFeatures(t *testing.T) {
	actual := getFeatures([]byte("test string"), boundaries)
	expected := []Feature{NewFeature([]byte("test")), NewFeature([]byte("string"))}

	if len(actual) != len(expected) {
		t.Errorf("getFeatures returned wrong number of features")
	}

	for i := 0; i < len(actual); i++ {
		if actual[i].Sum() != expected[i].Sum() {
			t.Errorf("feature.Sum(): expected %d, actual %d", expected[i].Sum(), actual[i].Sum())
		}
		if actual[i].Weight() != expected[i].Weight() {
			t.Errorf("feature.Weight(): expected %d, actual %d", expected[i].Weight(), actual[i].Weight())
		}
	}
}
