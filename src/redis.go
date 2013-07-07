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

	c.Cmd("set", "foo1", "bar1")
	c.Cmd("set", "foo2", "bar2")
	c.Cmd("set", "foo3", "bar3")
	mgetRep, err := c.Cmd("mget", "foo1", "foo2", "foo3").List()
	for i, _ := range mgetRep {
		fmt.Printf("%d => %s\n", i, mgetRep[i])
	}

	c.Append("set", "k1", "v1")
	c.Append("get", "k1")
	setRep = c.GetReply()
	errHandler(err)
	getRep, err = c.GetReply().Str()
	errHandler(err)
	fmt.Println(setRep)
	fmt.Println(getRep)

	// Looks like the radix client doesn't have support for transactions or pubsub, so i'll skip those.
}
