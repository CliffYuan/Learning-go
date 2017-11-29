package main

import (
	"time"
	"fmt"
)

func main()  {
	go spinner(10*time.Millisecond)

	const c  = 45

	fibC:=fib(c)

	fmt.Printf("fib %d is %d",c,fibC)


}

func spinner(duration time.Duration)  {
	for _,r :=range "#$%" {
		fmt.Printf("\r%c",r)
		time.Sleep(duration)
	}
}

func fib(n int) int {
	if(n<2){
		return n;
	}
	return fib(n-1)+fib(n-2)

}