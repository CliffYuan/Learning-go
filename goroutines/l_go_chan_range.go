package main

import "fmt"

func main()  {

	naturals:=make(chan int)

	squares:=make(chan int)


	go func() {

		for i := 0; i < 100; i++{
			naturals<-i
		}

		close(naturals)

	}()

	go func() {

		for v:=range naturals {
			squares<-v
		}

		close(squares)

	}()

	for v:=range squares{
		fmt.Println(v)
	}


}
