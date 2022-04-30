package atfield

// const buildTag = "-tags=atfield"

func Unfold(in, out string) error {
	atf := NewATField(in, out)
	atf.ReadPackages()
	atf.AnalyseSets()
	atf.ExtractFields()
	atf.GenerateConverts()
	return atf.Error()
}

type ATField interface {
	ReadPackages()
	AnalyseSets()
	ExtractFields()
	GenerateConverts()
	Error() error
}
