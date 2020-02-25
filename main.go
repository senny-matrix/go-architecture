package main

import "fmt"

type person struct {
	first string
}

func (p person) speak()  {
	fmt.Println("From a person - this is my name", p.first)
}

// Any TYPE that has the methods specified by an interface
// is also of the interface type
type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	// In Go a VALUE can be of more than one TYPE
	// in this example, p1 is both TYPE person and TYPE human
	var x human
	x = p1
	fmt.Printf("%T\n", x)
}
