package main

import "fmt"

type Human struct {
	name   string
	weight float32
}

type Food struct {
	weight float32
	name   string
	isExist bool
}

func (h Human) eat(f Food) {
	weight := h.weight + f.weight
	fmt.Println(h.name, "eat", f.name)
	fmt.Println("New weight:", weight)
}	

func main() {
	Adam := Human{"Tom", 50}
	apple := Food{10, "apple", true}
	Adam.eat(apple)
}
