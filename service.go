package architecture

import "fmt"

// Person is how the architecture package stores a person
type Person struct {
	First string
}

// Accessor is how to store / retrieve a person
// When retrieving a person, if they do not exist, it return the zero value
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

type PersonService struct {
	a Accessor
}

func (ps PersonService) Get(n int)(Person, error){
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no Person with n of %d", n)
	}
	return p, nil
}

func Put(a Accessor, n int, p Person)  {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}