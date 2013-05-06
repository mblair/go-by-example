package main

import "fmt"

type name struct {
	first, last string
}

func (n *name) fullName() string {
	return n.first + " " + n.last
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	n := name{first: "Matthew", last: "Blair"}
	np := &n
	fmt.Println("My full name is", np.fullName())
}