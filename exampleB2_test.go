//package main

package simhash_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
)

func testitB() {
	hashes := make([]uint64, len(doc2))
	sh := simhash.NewSimhash()
	for i, d := range doc2 {
		hashes[i] = sh.BuildSimhash(string(d), simhash.Doc2words)
		fmt.Printf("Simhash of '%s': %x\n", d, hashes[i])
	}

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", doc2[0], doc2[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", doc2[0], doc2[2], simhash.Compare(hashes[0], hashes[2]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", doc2[0], doc2[3], simhash.Compare(hashes[0], hashes[3]))
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleCompareB() {
	doc2 = [][]byte{
		[]byte("Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
		[]byte("2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic"),
		[]byte("2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic"),
		[]byte("2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic"),
	}
	testitB()

	fmt.Println("================")
	doc2 = [][]byte{
		[]byte("2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic"),
		[]byte("2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic"),
		[]byte("Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
		[]byte("2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic"),
	}
	testitB()

	// Output:
	// Simhash of 'Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic': 1832cd1ce6eb263e
	// Simhash of '2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic': 834df3eb7eb3ebe
	// Simhash of '2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic': 1832df0ee4eb3e39
	// Simhash of '2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic': 8b6df0ee7eb2f3c
	// Comparison of `Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic` and `2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic`: 13
	// Comparison of `Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic` and `2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic`: 10
	// Comparison of `Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic` and `2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic`: 11
	// ================
	// Simhash of '2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic': 1832df0ee4eb3e39
	// Simhash of '2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic': 8b6df0ee7eb2f3c
	// Simhash of 'Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic': 1832cd1ce6eb263e
	// Simhash of '2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic': 834df3eb7eb3ebe
	// Comparison of `2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic` and `2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic`: 9
	// Comparison of `2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic` and `Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic`: 10
	// Comparison of `2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic` and `2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic`: 13
}
