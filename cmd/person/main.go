package main

import (
	"fmt"
	"github.com/senny-matrix/go-architecture"
	"github.com/senny-matrix/go-architecture/storage/mongo"
	"github.com/senny-matrix/go-architecture/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := architecture.Person{
		First: "Senny",
	}

	p2 := architecture.Person{
		First: "Rogers",
	}

	ps := architecture.NewPersonService(dbm)

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)

	fmt.Println(ps.Get( 1))
	fmt.Println(ps.Get( 2))

	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)

	fmt.Println(ps.Get( 1))
	fmt.Println(ps.Get( 2))
}
