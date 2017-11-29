package main

import (
	"log"
	"fmt"
)

func main()  {


	naturals:=make(chan int)

	squares:=make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			naturals <- i
		}

		close(naturals)

	}()

	go func() {
		for {
		  n,ok:= <- naturals
		  if(!ok){
		  	log.Println("naturals close")
		  	break
		  }

		  squares <- n

		}

		close(squares)

	}()

	for  {
		data,ok:=<- squares
		if(!ok){
			log.Println("squares close")
			break
		}
		log.Println("squares ",data)
	}


	fmt.Printf("all is close")
}
