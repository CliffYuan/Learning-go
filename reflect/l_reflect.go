package main

import (
	"reflect"
	"fmt"
)

func main() {
	//basic operate
	basic()

	//iterator operate



}



func field_for()  {


}

func basic()  {
	var x float64=3.4
	fmt.Println("type:",reflect.TypeOf(x))

	v:=reflect.ValueOf(x)
	show(v);

	fmt.Println("v canset:",v.CanSet())

	fmt.Println("point type:",reflect.TypeOf(&x));
	k:=reflect.ValueOf(&x);
	fmt.Println("k canset:",k.CanSet())
	k=k.Elem();
	fmt.Println("k canset:",k.CanSet())
	k.SetFloat(5.0)
	fmt.Println(k.Interface())
	fmt.Println(k)
}

func show(v reflect.Value) {
	fmt.Println("-----------show start-----------------")
	fmt.Println("value:",v)
	fmt.Println("type:",v.Type())
	fmt.Println("kind:",v.Kind())
	fmt.Println("value:",v.Float())

	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n",v.Interface())

	y:=v.Interface().(float64)
	fmt.Println("convert:",y);
	fmt.Println("-----------show end-----------------")
}


