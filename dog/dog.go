package dog

import "go-demo/food"

type Dog struct {
	Name string `json:"name"`
}

func New(name string) *Dog {
	return &Dog{Name: name}
}

func (d Dog) Eat(food food.Food) string {
	switch food.Name {
	case "meat":
		return "Dog like meat, yum yum!"
	case "":
		return "Dog is hungry, please give me some food!"
	default:
		return "Dog doesn't like that food."
	}
}

func (d Dog) Speak() string {
	return "woof woof woof!"
}