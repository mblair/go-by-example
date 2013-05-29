package main

import "strconv"
import "fmt"

func main() {
	p := fmt.Println

	f, _ := strconv.ParseFloat("1.234", 64)
	p(f)

	i, _ := strconv.ParseInt("123", 16, 64)
	p(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	u, e := strconv.ParseUint("-789", 0, 64)
	p(u)
	p(e)

	k, _ := strconv.Atoi("135")
	p(k)

	_, e = strconv.Atoi("wat")
	 p(e)
}
