package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println()

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// TIL from golang.org post on slices
	// slices refer to the underlying array's storage.
	// slices are pointer, length and capacity triples.
	a := []int{1, 2, 3, 4, 5}
	b := a[:4]
	fmt.Println("b's len:", len(b))
  a[2] = 8
  fmt.Println("b[2]:", b[2])
	b = b[:cap(b)]
	fmt.Println("b's len:", len(b))
}
