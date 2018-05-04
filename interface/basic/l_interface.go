package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side*s.side
}

type Rectangle struct {
	length,width float32
}

func (s Rectangle) Area() float32 {
	return s.length*s.width
}


func main() {

	//
	type nullInterface interface{}
	var x int=5
	var n nullInterface
	n=x
	fmt.Println(n)
	n=3.5
	fmt.Println(n)
	fmt.Println("------------------------------------")

	//
	r:=Rectangle{5,4}
	q:=&Square{5}

	shapes:=[]Shaper{r,q}//q  must be not ptr

	for n,_:=range shapes{
		fmt.Println("Shape Details:",shapes[n])
		fmt.Println("Area of this shape is:",shapes[n].Area())
	}

}
