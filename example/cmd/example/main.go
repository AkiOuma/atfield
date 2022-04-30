package main

import (
	"log"
	"time"

	"github.com/AkiOuma/atfield/example/a"
	"github.com/AkiOuma/atfield/example/b"
	"github.com/AkiOuma/atfield/example/c"
	"github.com/AkiOuma/atfield/example/define"
)

// var name = "yuki"
func main() {
	userA := &a.User{
		Id:           1,
		UserName:     "yuki",
		Address:      &c.Address{StreetName: "test", Location: 1},
		RegisterTime: time.Now(),
		Labels:       map[string]*c.Label{"tag": {Content: "label"}},
		PhoneNumber:  []c.PhoneNumber{},
		BirthDay:     &c.BirthDay{T: time.Now()},
		School:       "school",
	}
	userB := define.AUserToBUser(userA)
	bulkUserB := []*b.User{userB}
	bulkUserA := define.BulkBUserToAUser(bulkUserB)
	log.Printf("user A: %#v", userA)
	log.Printf("user B: %#v", userB)
	log.Printf("bulkuser A: %#v", bulkUserA)
	log.Printf("bulkuser B: %#v", bulkUserB)
}
