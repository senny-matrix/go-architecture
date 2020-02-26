package main

import "fmt"

type person struct {
	first string
}

func (p person) speak()  {
	fmt.Println("From a person - this is my name", p.first)
}

type secretAgent struct {
	person
	ltk	bool
}

func (sa secretAgent) speak()  {
	fmt.Println("I am a secret agent - this is my name", sa.first)
}


// Any TYPE that has the methods specified by an interface
// is also of the interface type
type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}
func main() {
	p1 := person{
		first: "Miss MoneyPenny",
	}
	
	sa1 := secretAgent{
		person: person{
			first: "James Bond",
		},
		ltk:    true,
	}

	//fmt.Printf("%T\n", p1)

	// In Go a VALUE can be of more than one TYPE
	// in this example, p1 is both TYPE person and TYPE human
	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()

	println("________________________________")
	foo(x)
	foo(y)
	foo(p1)
	foo(sa1)
}
