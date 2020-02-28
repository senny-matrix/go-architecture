package main

import "fmt"

type hotDog int

func (h hotDog) cook(){
	fmt.Println("I am cooking!")
}

type hotFood interface {
	cook()
}

func main()  {
	var x hotDog = 42
	var y hotFood
	y = x
	z := hotFood(y)
	fmt.Println(z)
}