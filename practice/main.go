package main

import "fmt"

type Person struct {
	a string
}

func (b *Person) Name() {
	fmt.Println(b.a)
}

type Aminal struct {
	a string
}

func (b *Aminal) Name() {
	fmt.Println(b.a)
}

func main() {

	a := Person{"Vinay"}
	a.Name()

	b := Aminal{"Cat"}
	b.Name()

}
