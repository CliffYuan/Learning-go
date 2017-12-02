package main

import "fmt"

//struct
type person struct {
	age int
	name string
}

//func
func main()  {

	varstruct := person{12,"xnd"}

	//map
	varmap := map[string]int{"one":1}
	//varmap := make(map[string]int)

	//slice
	varslice:=make([]int,6)

	//arrays
	vararray:=[3]int{1,2,3}

	fmt.Println("-------------------------------------before out",varstruct,varmap,varslice,vararray)

	change(varstruct,varmap,varslice,vararray)

	fmt.Println("-------------------------------------after out",varstruct,varmap,varslice,vararray)

}

func change(s person,m map[string]int ,sli []int, arr [3]int)  {

	fmt.Println("---before in",s,m,sli,arr)

	s.name="xiaoniudu"

	m["one"]=88

	sli[0]=88

	arr[0]=88

	fmt.Println("---after in",s,m,sli,arr)
}
