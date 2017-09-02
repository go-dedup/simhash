// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"fmt"
	"testing"
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
	sh := NewSimhash()
	text := []byte("here's a test string.")
	fs := sh.NewWordFeatureSet(text)
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
	sh := NewSimhash()
	text := []byte("a a abc abc test test string.")
	fs := sh.NewWordFeatureSet(text)
	fmt.Printf("%#v\n", fs)
	actual := fs.GetFeatures()
	fmt.Printf("%#v\n", actual)
	fmt.Printf("%v\n", sh.Fingerprint(sh.Vectorize(actual)))
	// Output:
	// &simhash.WordFeatureSet{B:[]uint8{0x61, 0x20, 0x61, 0x20, 0x61, 0x62, 0x63, 0x20, 0x61, 0x62, 0x63, 0x20, 0x74, 0x65, 0x73, 0x74, 0x20, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e}}
	// []simhash.Feature{simhash.feature{sum:0xaf63bd4c8601b7be, weight:1}, simhash.feature{sum:0xaf63bd4c8601b7be, weight:1}, simhash.feature{sum:0xd8dcca186bafadcb, weight:1}, simhash.feature{sum:0xd8dcca186bafadcb, weight:1}, simhash.feature{sum:0x8c093f7e9fccbf69, weight:1}, simhash.feature{sum:0x8c093f7e9fccbf69, weight:1}, simhash.feature{sum:0x9926dcde0a17d48e, weight:1}}
	// 10108821242876116971
}

func TestGetFeatures(t *testing.T) {
	actual := DoGetFeatures([]byte("test string"), boundaries)
	expected := []Feature{
		NewFeature([]byte("test")),
		NewFeature([]byte("string"))}

	if len(actual) != len(expected) {
		t.Errorf("DoGetFeatures returned wrong number of features")
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
