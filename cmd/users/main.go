package main

import (
	"fmt"

	"github.com/senny-matrix/go-architecture"
	"github.com/senny-matrix/go-architecture/storage/harddrive"
)
func main() {
	//storage := mongo{}
	storage := harddrive.Db{}

	u1 := architecture.User{
		First: "Rogers",
	}

	u2 := architecture.User{
		First: "Dyness",
	}

	architecture.Put(storage, 1, u1)
	architecture.Put(storage, 2, u2)

	fmt.Println(architecture.Get(storage, 1))
	fmt.Println(architecture.Get(storage, 2))

}
