package atfield

type fieldParam struct {
	Name string
	Type string
}

func ifPublicField(name string) bool {
	if len(name) == 0 {
		return false
	}
	first := name[0]
	if first > 64 && first < 91 {
		return true
	}
	return false
}
