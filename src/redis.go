package main

import "github.com/fzzy/radix/redis"
// import "time"
import "fmt"
import "time"
import "os"

func errHandler(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
	errHandler(err)
	defer c.Close()

	fmt.Println(c)

	setRep := c.Cmd("set", "foo", "bar")
	errHandler(setRep.Err)
	fmt.Println(setRep)

	getRep, err := c.Cmd("get", "foo").Str()
	errHandler(err)
	fmt.Println(getRep)
}
