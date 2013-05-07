package main

import "time"
import "fmt"

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	fmt.Println(<-timer1.C)
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	fmt.Println(timer2.Stop())
}