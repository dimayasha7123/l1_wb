package main

import (
	"fmt"
	"reflect"
)

/*
Задание:
	Разработать программу, которая в рантайме способна определить тип
	переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var x interface{}

	x = 4
	//x = "asdrf"
	//x = false
	//x = make(chan int)
	//x = struct {
	//	A int
	//	B int
	//}{3, 4}

	fmt.Println(x)

	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	case bool:
		fmt.Println("x is bool")
	case chan int:
		fmt.Println("x is chan int")
	default:
		fmt.Println("what is x...")
	}

	// also we can use specific type in .(type)
	_, ok := x.(int)
	if ok {
		fmt.Println("x is int")
	} else {
		fmt.Println("x is not int")
	}

	// also we can use %T in fmt
	strType := fmt.Sprintf("%T", x)
	fmt.Println(strType)

	// also we can use reflect.TypeOf()..., but don't
	fmt.Println(reflect.TypeOf(x).Kind())
	fmt.Println(reflect.TypeOf(x).Kind() == reflect.Int)
}
