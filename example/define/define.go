package define

import (
	"github.com/AkiOuma/atfield"
	"github.com/AkiOuma/atfield/example/a"
	"github.com/AkiOuma/atfield/example/b"
)

var _ = atfield.LinkStruct(a.User{}, b.User{}).
	LinkField(a.User{}.UserName, b.User{}.Name).
	LinkField(a.User{}.Id, b.User{}.ID).
	LinkField(a.User{}.Address, b.User{}.Addr)
