package main

import (
	"sort"
	"fmt"
)

func main() {

	//
	sort_IntSlice();

	sort_Ints();


}

func sort_IntSlice()  {
	sint:=sort.IntSlice{1,5,6,8,2}

	sort.Sort(sint)

	fmt.Println(sint)
}

func sort_Ints()  {
	sint:=[]int {1,10,3}
	sort.Ints(sint)

	fmt.Println(sint)
}
