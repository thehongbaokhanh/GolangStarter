package main

import (
	"fmt"
	"go-demo/cat"
	"go-demo/dog"
	"go-demo/food"
)

type Animal interface {
	Speak() string
	Eat(food food.Food) string
}


func speak(a Animal) {
	fmt.Println(a.Speak())
}

func eat(a Animal, f food.Food) {
	fmt.Println(a.Eat(f))
}

func main() {
	jack := dog.New("Jack")
	tom := cat.New("Tom")
	fish := food.Food{Name: "fish"}
	meat := food.Food{Name: "meat"}
	noFood := food.Food{Name: ""}
	fmt.Println("Đây là tiếng kêu của con chó " + jack.Name + ": ")
	speak(jack)
	fmt.Println("Đây là tiếng kêu của con mèo " + tom.Name + ": ")
	speak(tom)
	eat(tom, fish)
	eat(tom, meat)
	eat(tom, noFood)
	eat(jack, fish)
	eat(jack, meat)
	eat(jack, noFood)
}
