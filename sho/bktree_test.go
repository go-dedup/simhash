// Copyright 2015, Yahoo Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sho_test

import (
	"fmt"
	"testing"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/sho"
)

func ExampleNewOracle() {
	// just add 0 and 1.
	oracle := sho.NewOracle()
	for i := uint64(1); i < 2; i++ {
		oracle.See(i)
	}
	r := uint8(2)
	fmt.Printf("Within distance of %d,\n", r)
	for i := uint64(0); i < 30; i++ {
		fmt.Printf("Has the oracle seen anything closed to %02d (%08b)? %t\n", i, i, oracle.Seen(i, r))
	}
	// Output:
	// Within distance of 2,
	// Has the oracle seen anything closed to 00 (00000000)? true
	// Has the oracle seen anything closed to 01 (00000001)? true
	// Has the oracle seen anything closed to 02 (00000010)? true
	// Has the oracle seen anything closed to 03 (00000011)? true
	// Has the oracle seen anything closed to 04 (00000100)? true
	// Has the oracle seen anything closed to 05 (00000101)? true
	// Has the oracle seen anything closed to 06 (00000110)? false
	// Has the oracle seen anything closed to 07 (00000111)? false
	// Has the oracle seen anything closed to 08 (00001000)? true
	// Has the oracle seen anything closed to 09 (00001001)? true
	// Has the oracle seen anything closed to 10 (00001010)? false
	// Has the oracle seen anything closed to 11 (00001011)? false
	// Has the oracle seen anything closed to 12 (00001100)? false
	// Has the oracle seen anything closed to 13 (00001101)? false
	// Has the oracle seen anything closed to 14 (00001110)? false
	// Has the oracle seen anything closed to 15 (00001111)? false
	// Has the oracle seen anything closed to 16 (00010000)? true
	// Has the oracle seen anything closed to 17 (00010001)? true
	// Has the oracle seen anything closed to 18 (00010010)? false
	// Has the oracle seen anything closed to 19 (00010011)? false
	// Has the oracle seen anything closed to 20 (00010100)? false
	// Has the oracle seen anything closed to 21 (00010101)? false
	// Has the oracle seen anything closed to 22 (00010110)? false
	// Has the oracle seen anything closed to 23 (00010111)? false
	// Has the oracle seen anything closed to 24 (00011000)? false
	// Has the oracle seen anything closed to 25 (00011001)? false
	// Has the oracle seen anything closed to 26 (00011010)? false
	// Has the oracle seen anything closed to 27 (00011011)? false
	// Has the oracle seen anything closed to 28 (00011100)? false
	// Has the oracle seen anything closed to 29 (00011101)? false
}

func BenchmarkOracleSee(b *testing.B) {
	oracle := sho.NewOracle()
	for i := 0; i < b.N; i++ {
		// for i := uint64(1); i < 10000; i++ {
		oracle.See(uint64(i))
		// }
	}
}

func BenchmarkOracleSeen(b *testing.B) {
	oracle := sho.NewOracle()
	for i := uint64(1); i < 1000000; i++ {
		oracle.See(i)
	}
	b.ResetTimer()
	r := uint8(2)
	for i := 0; i < b.N; i++ {
		oracle.Seen(uint64(i), r)
	}
}

func ExampleSearch_cars() {
	docs := [][]byte{
		[]byte("2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic"),
		[]byte("2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic"),
		[]byte("Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
		[]byte("2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic"),
	}
	testOracle(docs, 9)
	// Output:
	// 590700557005172541: [{626808552180887356 7} {590779825190022718 8}]
	// 626808552180887356: [{590700557005172541 7} {590779825190022718 7}]
	// 1743672742259666494: [{590779825190022718 6}]
	// 590779825190022718: [{590700557005172541 8} {1743672742259666494 6} {626808552180887356 7}]
}

func ExampleSearch_filesA() {
	docs := [][]byte{
		[]byte("test/sim/Audio Book - The Grey Coloured Bunnie.mp3"),
		[]byte("test/sim/GNU - Python Standard Library (2001).rar"),
		[]byte("test/sim/PopupTest.java"),
		[]byte("test/sim/(eBook) GNU - Python Standard Library 2001.pdf"),
		[]byte("test/sim/Python Standard Library.zip"),
		[]byte("test/sim/GNU - 2001 - Python Standard Library.pdf"),
		[]byte("test/sim/LayoutTest.java"),
		[]byte("test/sim/ColoredGrayBunny.ogg"),
	}
	testOracle(docs, 12)
	// Output:
	// 15608818089684999448: []
	// 15895225339941416922: [{15895313301139093466 7}]
	// 16760638039811334009: []
	// 15896913088138765851: [{15896913090289004511 9}]
	// 15895313301139093466: [{15895225339941416922 7}]
	// 15896913090289004511: [{15896913088138765851 9} {15895225339941416922 6} {15895313301139093466 11}]
	// 18345905108715307007: []
	// 15902139077912948570: []
}

func ExampleSearch_filesB() {
	docs := [][]byte{
		[]byte("test/sim/LayoutTest.java"),
		[]byte("test/sim/ColoredGrayBunny.ogg"),
		[]byte("test/sim/GNU - 2001 - Python Standard Library.pdf"),
		[]byte("test/sim/PopupTest.java"),
		[]byte("test/sim/(eBook) GNU - Python Standard Library 2001.pdf"),
		[]byte("test/sim/GNU - Python Standard Library (2001).rar"),
		[]byte("test/sim/Python Standard Library.zip"),
		[]byte("test/sim/Audio Book - The Grey Coloured Bunnie.mp3"),
	}
	testOracle(docs, 12)
	// Output:
	// 18345905108715307007: []
	// 15902139077912948570: []
	// 15896913090289004511: [{15896913088138765851 9} {15895313301139093466 11}]
	// 16760638039811334009: []
	// 15896913088138765851: [{15896913090289004511 9}]
	// 15895225339941416922: [{15896913090289004511 6} {15895313301139093466 7}]
	// 15895313301139093466: [{15896913090289004511 11}]
	// 15608818089684999448: []
}

func ExampleSearch_filesS() {
	docs := [][]byte{
		[]byte("test/sim/Audio Book - The Grey Coloured Bunnie.mp3"),
		[]byte("test/sim/ColoredGrayBunny.ogg"),
		[]byte("test/sim/(eBook) GNU - Python Standard Library 2001.pdf"),
		[]byte("test/sim/GNU - 2001 - Python Standard Library.pdf"),
		[]byte("test/sim/GNU - Python Standard Library (2001).rar"),
		[]byte("test/sim/LayoutTest.java"),
		[]byte("test/sim/PopupTest.java"),
		[]byte("test/sim/Python Standard Library.zip"),
	}
	testOracle(docs, 12)
	// Output:
	// 15608818089684999448: []
	// 15902139077912948570: []
	// 15896913088138765851: [{15896913090289004511 9}]
	// 15896913090289004511: [{15896913088138765851 9} {15895313301139093466 11}]
	// 15895225339941416922: [{15896913090289004511 6} {15895313301139093466 7}]
	// 18345905108715307007: []
	// 16760638039811334009: []
	// 15895313301139093466: [{15896913090289004511 11}]
}

func testOracle(docs [][]byte, r uint8) {
	oracle := sho.NewOracle()
	sh := simhash.NewSimhash()
	hashes := make([]uint64, len(docs))
	for i, d := range docs {
		hashes[i] = sh.GetSimhash(sh.NewWordFeatureSet(d))
		oracle.See(hashes[i])
	}

	for _, h := range hashes {
		fmt.Printf("%v: %v\n", h, oracle.Search(h, r))
		//fmt.Printf("\n\t%v\n", oracle.SearchAll(h, r))
	}
}
