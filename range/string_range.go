package main

import "fmt"

func main()  {

	data:="abc"
	for i,j:=range data{
		fmt.Println(i,"=",j)

		fmt.Printf("%#x ---kkk",j)

		fmt.Println("\r-----------")
	}


}
