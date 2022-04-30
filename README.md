# ATField - (Auto Translate field)

a module for generate struct converting file

# Usage
full code in `example`

first, install atfield
```bash
go get github.com/AkiOuma/atfield/cmd/atfield
```

assume we have to struct name `User` in package a and b
```go
// example/a.go
type User struct {
	Id           int64
	UserName     string
	Address      *c.Address
	RegisterTime time.Time
	Labels       map[string]*c.Label
	PhoneNumber  []c.PhoneNumber
	BirthDay     c.Date
	School       string
}


// example/b.go
type User struct {
	ID           int32
	Name         []byte
	Addr         *c.Address
	RegisterTime time.Time
	PhoneNumber  []c.PhoneNumber
	Labels       map[string]*c.Label
	BirthDay     c.Date
}

```

we want to genenrate convert file in `example/define`, we need to create a new go file in `example/define` as follow: 
```go
// example/define
import (
	"github.com/AkiOuma/atfield"
	"github.com/AkiOuma/atfield/example/a"
	"github.com/AkiOuma/atfield/example/b"
)

var _ = atfield.LinkStruct(a.User{}, b.User{}).
	LinkField(a.User{}.UserName, b.User{}.Name).
	LinkField(a.User{}.Id, b.User{}.ID).
	LinkField(a.User{}.Address, b.User{}.Addr)

```

enter directory `example/define`, and run atfield in terminal:
```bash
$ atfield
```

then convert file will be generated in `example/define` named `atfield.go`

```go
// Code generated by atfield. DO NOT EDIT.

//+build !atfield

package define

import (
	"github.com/AkiOuma/atfield/example/a"
	"github.com/AkiOuma/atfield/example/b"
)

func AUserToBUser(source *a.User) *b.User {
	result := &b.User{}
	result.BirthDay = source.BirthDay
	result.ID = int32(source.Id)
	result.Name = []byte(source.UserName)
	result.Addr = source.Address
	result.RegisterTime = source.RegisterTime
	result.Labels = source.Labels
	result.PhoneNumber = source.PhoneNumber
	return result
}

func BulkAUserToBUser(source []*a.User) []*b.User {
	result := make([]*b.User, 0, len(source))
	for _, v := range source {
		result = append(result, AUserToBUser(v))
	}
	return result
}

func BUserToAUser(source *b.User) *a.User {
	result := &a.User{}
	result.Labels = source.Labels
	result.BirthDay = source.BirthDay
	result.Id = int64(source.ID)
	result.UserName = string(source.Name)
	result.Address = source.Addr
	result.RegisterTime = source.RegisterTime
	result.PhoneNumber = source.PhoneNumber
	return result
}

func BulkBUserToAUser(source []*b.User) []*a.User {
	result := make([]*a.User, 0, len(source))
	for _, v := range source {
		result = append(result, BUserToAUser(v))
	}
	return result
}

```