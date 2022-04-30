package b

import (
	"time"

	"github.com/AkiOuma/atfield/example/c"
)

type User struct {
	ID           int32
	Name         []byte
	Addr         *c.Address
	RegisterTime time.Time
	PhoneNumber  []c.PhoneNumber
	Labels       map[string]*c.Label
	BirthDay     c.Date
}
