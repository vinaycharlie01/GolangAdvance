package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Cat struct {
	name string
}

func (c *Cat) Speak() string {
	// fmt.Println(c.name)
	return c.name
}

type Animal2 interface {
	Eating() string
}

type Dog struct {
	Animal
	Animal2
	name string
}

func (d *Dog) Speak() string {
	// fmt.Println(d.name)
	return d.name

}

func (d *Dog) Eating() string {
	return d.name
}

type Animals struct {
	Animal Animal
}

func (a *Animals) HowSpeak(b Animal) string {
	return b.Speak()
	// b.Speak()
	// a.Animal.Speak()
	// a.Animal.Speak()
	// a.Animal.Speak()
}

func (a *Animals) Eat(b Animal2) string {
	return "fhfhfhfjfn"
}

func main() {
	c := Animals{}
	a := &Cat{name: "meooooow...!"}
	b := &Dog{name: "Bow..........!"}
	// fmt.Println(b.Eating())
	// c.Animal = a
	fmt.Println(c.HowSpeak(a))

	// c.Animal = b
	fmt.Println(c.HowSpeak(b))
	fmt.Println(b.Eating())
	fmt.Println(c.Eat(b))

	// big.NewInt(a).And()
}

// package main

// import "fmt"

// type Animal interface {
// 	Speak() string
// 	Details() string
// 	Eating() string
// 	ModityDetail(age int, name string) string
// }

// type Cat struct {
// 	name  string
// 	age   int
// 	sound string
// 	food  string
// }

// func (c *Cat) Speak() string {
// 	return c.sound
// }

// func (c *Cat) Details() string {
// 	return fmt.Sprintf("%v: %v", c.name, c.age)
// }

// func (c *Cat) Eating() string {
// 	return c.food
// }

// func (c *Cat) ModityDetail(age int, name string) string {
// 	c.age = age
// 	c.name = name
// 	return fmt.Sprintf("%v: %v", c.name, c.age)
// }

// type Animals struct {
// 	Animal Animal
// }

// func (c *Animals) Speak(c1 Animal) string {
// 	return c1.Speak()
// }

// func (c *Animals) Details(c1 Animal) string {
// 	return fmt.Sprintf("%v: %v", c1.Details())
// }

// func (c *Animals) Eating(c1 Animal) string {
// 	return c1.Eating()
// }

// type Dog struct {
// 	name  string
// 	age   int
// 	sound string
// 	food  string
// }

// func (c *Dog) Speak() string {
// 	return c.sound
// }

// func (c *Dog) Details() string {
// 	return fmt.Sprintf("%v: %v", c.name, c.age)
// }

// func (c *Dog) Eating() string {
// 	return c.food
// }

// func (c *Dog) ModityDetail(age int, name string) string {
// 	c.age = age
// 	c.name = name
// 	return fmt.Sprintf("%v: %v", c.name, c.age)
// }

// func main() {
// 	c := &Animals{}
// 	// fmt.Println(c.Speak(&Dog{sound: "Bow"}))
// 	// b := Cat{"cat1", 2, "Meowwww...!", "Milk"}
// 	// b1 := Dog{"Dog1", 3, "Bow", "chikem"}
// 	// fmt.Println(c.Speak(&b))

// }
