// Package bloomfilter implements a simple bloom filter using hashing functions.

// For a full guide visit https://github.com/wal99d/bloomfilter
// package main

// import (
// 	"fmt"

// 	"github.com/wal99d/bloomfilter"
// )

// func main() {
// 	filter := bloomfilter.NewBloomFilter(10000, []bloomfilter.HashFunc{bloomfilter.NewHashFunc()})

// 	filter.Add("A")
// 	filter.Add("B")
// 	filter.Add("C")

// 	fmt.Println("Check if A exists:", filter.Check("A"))
// 	fmt.Println("Check if D exists:", filter.Check("D"))
// 	filter.Remove("C")
// 	fmt.Println("C is removed")
// 	fmt.Println("Check if C exists:", filter.Check("C"))
// }

package bloomfilter
