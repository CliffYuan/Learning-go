package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) printPoint()  {
	fmt.Println("point:",p.name)
}

func (p field) print()  {
	fmt.Println(p.name)
}

//todo for range
func main()  {

	fmt.Println("---------------------range point---------------------------")
	//points
	data:=[]*field{{"one"},{"two"},{"three"}}

	for _,v := range data{
		go v.print();
		go v.printPoint()
	}

	time.Sleep(3*time.Second)

	fmt.Println("---------------------range values---------------------------")
	//


	data2:=[]field{{"one"},{"two"},{"three"}}

	for _,v := range data2{
		go v.print();
		go v.printPoint()
	}

	time.Sleep(3*time.Second)

}
