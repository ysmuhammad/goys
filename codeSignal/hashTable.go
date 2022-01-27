package main

import "fmt"

const arr = 5

// Hashtable with type of array
type HashTable struct {
	arr [arr]*Bucket
}

// Bucket, type of each element in hashtable
type Bucket struct {
	head *Node
}

// Node, type of each element in Bucket
type Node struct {
	val  string
	next *Node
}

// Initialize hashtable with bucket
func initializeBucket() *HashTable {
	a := &HashTable{}
	for i := range a.arr {
		a.arr[i] = &Bucket{}
	}
	return a
}

// Hashing alg
func hash(s string) int {
	final := 0
	for _, v := range s {
		final += int(v)
	}
	return final % arr
}

// Insert hashtable
// Search hashtable
// Delete hashtable

// Insert bucket
func (b *Bucket) insertBucket(s string) *Bucket {
	new := &Bucket{}
	new.head = 
}

// Search bucket
// Delete bucket

func main() {
	test := initializeBucket()
	fmt.Println(test)

}