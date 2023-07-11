package main

import "fmt"

/*
Задание:
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры
	Human (аналог наследования)
*/

func main() {
	a := Action{
		Human: Human{
			Name: "Dima",
			Age:  20,
		},
		Type: "HardWork",
	}

	a.Do()
	a.Work()
	a.Do()

	fmt.Printf("%#v\n", a)
}

type Human struct {
	Name string
	Age  int
}

func (h *Human) Work() {
	fmt.Println("Working...")
	h.Age++
}

type Action struct {
	Human
	Type string
}

func (a *Action) Do() {
	a.Work()
	fmt.Printf("Do %s, now human is %d years old\n", a.Type, a.Age)
}
