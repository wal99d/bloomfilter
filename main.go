package main

import (
	"fmt"

	"github.com/wal99d/bloomfilter/internals"
)

func main() {
	filter := internals.NewBloomFilter(10000, []internals.HashFunc{internals.NewHashFunc()})

	filter.Add("A")
	filter.Add("B")
	filter.Add("C")

	fmt.Println("Check if A exists:", filter.Check("A"))
	fmt.Println("Check if D exists:", filter.Check("D"))
	filter.Remove("C")
	fmt.Println("C is removed")
	fmt.Println("Check if C exists:", filter.Check("C"))
}
