package atfield

type FieldSet struct{}

// link diffrent types of struct instance
//
// for example:
//
// NewLink(a.User{}, b.User{})
func LinkStruct(...interface{}) *FieldSet {
	return &FieldSet{}
}

// link field
//
// for example:
//
// NewLink(a.User{}, b.User{}).LinkField(a.User{}.UserName, b.User{}.Name)
func (f *FieldSet) LinkField(...interface{}) *FieldSet { return f }
