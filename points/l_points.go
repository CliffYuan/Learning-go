package main

import "fmt"

type person struct{
	age uint8
	name string
}

func main()  {
	p1:=person{18,"xnd"}

	fmt.Println("person p1 address:",&p1)
	fmt.Println("person p1 value:",p1)

	//传值-----------------------------------
	changeName(p1)

	fmt.Println("person p1 address:",&p1)
	fmt.Println("person p1 value:",p1)

	fmt.Println("---------------------------------------------------")

	//传地址-----------------------------------
	changeNamePoint(&p1)


	fmt.Println("person p1 address:",&p1)
	fmt.Println("person p1 value:",p1)


	fmt.Println("------------------------new---------------------------")
	//new
	p2:=new(person)
	p2.name="new_xnd"
	p2.age=1

	fmt.Println("person p2 address:",&p2)
	fmt.Println("person p2 value:",p2)
	fmt.Println("person p2 value:",*p2)

	changeName(*p2)

	fmt.Println("-----------------------change after----------------------------")

	fmt.Println("person p2 address:",&p2)
	fmt.Println("person p2 value:",p2)
	fmt.Println("person p2 value:",*p2)

	fmt.Println("---------------------------------------------------")


}

func changeName(p person)  {
	fmt.Println("changeName person p1 address:",&p)
	fmt.Println("changeName person p1 value:",p)

	p.age=9
	p.name="xiaoniudu"

	fmt.Println("changeName person p1 address:",&p)
	fmt.Println("changeName person p1 value:",p)

}

func changeNamePoint(p *person)  {
	fmt.Println("changeName person p1 address:",p)
	fmt.Println("changeName person p1 address:",&p)
	fmt.Println("changeName person p1 value:",*p)

	p.age=88
	p.name="xiaoniudu8888"

	fmt.Println("changeName person p1 address:",p)
	fmt.Println("changeName person p1 address:",&p)
	fmt.Println("changeName person p1 value:",*p)

}
