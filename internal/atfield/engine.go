package atfield

import (
	"bytes"

	"golang.org/x/tools/go/packages"
)

type engine struct {
	pkg          *packages.Package
	infileDir    string
	outfileDir   string
	buf          *bytes.Buffer
	structSet    structSet
	linkStructs  [][2]structId
	linkFieldSet linkFieldSet
	linkPackages []*pkg
	err          error
}

type linkFieldSet map[fieldDir]beLinkedFields

type fieldDir struct {
	StructId  structId
	FieldName string
}

type beLinkedFields map[structId]string

type pkg struct {
	Alias, Dir string
}

type structId struct {
	PkgName, StructName string
}

type structSet map[structId]fieldSet

type fieldSet map[string]string

func NewATField(infileDir, outfileDir string) ATField {
	atf := &engine{}
	atf.SetInfileDir(infileDir)
	atf.SetoutfileDir(outfileDir)
	atf.buf = bytes.NewBuffer([]byte{})
	atf.structSet = make(structSet)
	atf.linkStructs = make([][2]structId, 0)
	atf.linkFieldSet = make(linkFieldSet)
	atf.linkPackages = make([]*pkg, 0)
	return atf
}

var _ ATField = (*engine)(nil)

func (e *engine) ReadPackages() {
	e.errorIntercept(e.readPackage)
}

func (e *engine) AnalyseSets() {
	e.errorIntercept(e.analyseSets)
}

func (e *engine) ExtractFields() {
	e.errorIntercept(e.extractFields)
}

func (e *engine) GenerateConverts() {
	e.errorIntercept(e.generateConverts)
}

func (e *engine) Error() error {
	return e.err
}

func (e *engine) errorIntercept(f func()) {
	if e.Error() != nil {
		return
	}
	f()
}

// read package information of infile directory
func (e *engine) readPackage() {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedFiles,
		BuildFlags: []string{"-tags=atfield"},
	}
	pkgs, err := packages.Load(cfg, e.infileDir)
	if err != nil {
		e.err = errReadPackage
		return
	}
	if len(pkgs) > 0 {
		e.pkg = pkgs[0]
	} else {
		e.err = errReadPackage
	}
}

// analyse items convert definition
func (e *engine) analyseSets() {
	e.analyseImports()
	e.analyseSyntaxes()
}
