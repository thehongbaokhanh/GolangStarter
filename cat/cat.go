package cat

import "go-demo/food"

type Cat struct {
	Name string `json:"name"`
}

func New(name string) *Cat {
	return &Cat{Name: name}
}

func (c Cat) Eat(food food.Food) string {
	switch food.Name {
	case "fish":
		return "Cat like fish, yum yum!"
	case "":
		return "Cat is hungry, please give me some food!"
	default:
		return "Cat doesn't like that food."
	}
}

func (c Cat) Speak() string {
	return "meow meow meow!"
}