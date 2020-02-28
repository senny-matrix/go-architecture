package mongo

import "github.com/senny-matrix/go-architecture"
type Db map[int]architecture.User

func (m Db) Save(n int, u architecture.User)  {
	m[n] = u
}

func (m Db) Retrieve(n int) architecture.User {
	return m[n]
}
