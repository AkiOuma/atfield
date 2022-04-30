package a

import (
	"time"

	"github.com/AkiOuma/atfield/example/c"
)

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
