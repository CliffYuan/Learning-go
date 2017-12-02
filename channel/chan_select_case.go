package main

import (
	"fmt"
	"time"
)

func main()  {
	inch := make(chan int)
	outch := make(chan  int)


	go func() {

		var in <-chan int = inch
		var out chan<-int
		var val int

		fmt.Println("-------------------------------------------------1")
		for{
			select{
			case out <- val:    // todo first invoke this ???  val=0 ???
				out=nil
				in=inch
				fmt.Println("--------------------out")
			case val= <-in:
				out=outch
				in=nil
				fmt.Println("--------------------in")
			}
		}

	}()


	go func() {
		fmt.Println("-------------------------------------------------2")

		for r :=range outch{
			fmt.Println("result:",r)
		}

	}()

	time.Sleep(0)
	inch <- 1

	fmt.Println("---------------------- inch <- 1 finish  ---------------------------")

	inch <- 2
	time.Sleep(3*time.Second)

}
