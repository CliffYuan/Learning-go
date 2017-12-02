package main

import "fmt"

func main()  {
	c:=make(chan int)

	select{
	case c <- getInt():
		fmt.Println("---------chan case invoke-----------")
	case c <- 12:
		fmt.Println("---------12 case invoke-----------")
	default:
		fmt.Println("---------default invoke-----------")
	}


}

func getInt() int {
	fmt.Println("invoke getInt")
	return 100
}