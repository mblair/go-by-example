package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
		nextInt := intSeq() // the function that intSeq returns here is a closure. it has "closed over" the state of `i`, so successive calls to intSeq increment i.

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // but new calls to `nextInt` effectively reset the state of `i`.

	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}