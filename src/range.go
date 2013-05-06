package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5}
	sum := 0

	for _, val := range nums {
		sum += val
	}
	fmt.Println("sum:", sum)

	for key, val := range nums {
		if val == 3 {
			fmt.Println("3 is at index:", key)
		}
	}

	kvs := map[string]string{"a": "apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	for index, rune := range "go" {
		fmt.Println(index, rune)
	}
}