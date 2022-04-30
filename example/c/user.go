package c

import "time"

type Address struct {
	StreetName string
	Location   int
}

type Label struct {
	Content string
}

type Date interface {
	GetLunar() string
	GetSolar() string
}

type PhoneNumber interface {
	GetArea()
	GetNumber()
}

type BirthDay struct {
	T time.Time
}

func (d *BirthDay) GetLunar() string {
	return d.T.Local().String()
}

func (d *BirthDay) GetSolar() string {
	return d.T.Local().String()
}
