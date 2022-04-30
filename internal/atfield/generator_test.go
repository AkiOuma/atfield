package atfield

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/tools/go/packages"
)

var e = &engine{
	infileDir:  ".",
	outfileDir: ".",
	pkg:        &packages.Package{Name: "test"},
	buf:        bytes.NewBuffer([]byte{}),
	linkPackages: []*pkg{{
		Alias: "a",
		Dir:   "github.com/test/a",
	}, {
		Alias: "b",
		Dir:   "github.com/test/c",
	}},
	linkStructs: [][2]structId{
		{{PkgName: "a", StructName: "User"}, {PkgName: "b", StructName: "User"}},
	},
	linkFieldSet: linkFieldSet{
		{StructId: structId{PkgName: "a", StructName: "User"}, FieldName: "Name"}:     beLinkedFields{{PkgName: "b", StructName: "User"}: "UserName"},
		{StructId: structId{PkgName: "b", StructName: "User"}, FieldName: "UserName"}: beLinkedFields{{PkgName: "a", StructName: "User"}: "Name"},
	},
	structSet: structSet{
		{PkgName: "a", StructName: "User"}: {"Name": "string", "Age": "int32"},
		{PkgName: "b", StructName: "User"}: {"UserName": "[]rune", "Age": "int"},
	},
}

func TestFunctionName(t *testing.T) {
	struct1 := structId{PkgName: "a", StructName: "User"}
	struct2 := structId{PkgName: "b", StructName: "User"}
	expect := "AUserToBUser"
	ans := functionName(struct1, struct2)
	if ans != expect {
		t.Errorf("functionName failed to pass, answer is [%s], expect is [%s]", ans, expect)
	}
}

func TestDefinePackage(t *testing.T) {
	e.buf.Reset()
	e.definePackage()
	expect := fmt.Sprintf("package %s\n\n", e.pkg.Name)
	ans := e.buf.String()
	if ans != expect {
		t.Errorf("definePackage failed to pass, answer is [%s], expect is [%s]", ans, expect)
	}
}

func TestGenerateImportPackage(t *testing.T) {
	e.buf.Reset()
	pkg := e.linkPackages[0]
	e.generateImportPackage(pkg)
	expect := fmt.Sprintf("\t\"%s\"\n", pkg.Dir)
	ans := e.buf.String()
	if ans != expect {
		t.Errorf("generateImportPackage failed to pass, answer is [%s], expect is [%s]", ans, expect)
	}
	e.buf.Reset()
	pkg = e.linkPackages[1]
	e.generateImportPackage(pkg)
	expect = fmt.Sprintf("\t%s \"%s\"\n", pkg.Alias, pkg.Dir)
	ans = e.buf.String()
	if ans != expect {
		t.Errorf("generateImportPackage failed to pass, answer is [%s], expect is [%s]", ans, expect)
	}
}

func TestGenerateConvertFunction(t *testing.T) {
	e.buf.Reset()
	e.generateConvertFunction(e.linkStructs[0][0], e.linkStructs[0][1])
	expect1 := `func AUserToBUser(source *a.User) *b.User {
	result := &b.User{}
	result.UserName = []rune(source.Name)
	result.Age = int(source.Age)
	return result
}

`
	expect2 := `func AUserToBUser(source *a.User) *b.User {
	result := &b.User{}
	result.Age = int(source.Age)
	result.UserName = []rune(source.Name)
	return result
}
	
`
	ans := e.buf.String()
	if ans != expect1 && ans != expect2 {
		t.Errorf("generateConvertFunction failed to pass, answer is [%s], expect is [%s] or [%s]", ans, expect1, expect2)
	}
}

func TestGenerateBulkConvertFunction(t *testing.T) {
	e.buf.Reset()
	e.generateBulkConvertFunction(e.linkStructs[0][0], e.linkStructs[0][1])
	expect := `func BulkAUserToBUser(source []*a.User) []*b.User {
		result := make([]*b.User, 0, len(source))
		for _, v := range source {
			result = append(result, AUserToBUser(v))
		}
		return result
	}

`
	ans := e.buf.String()
	if ans != expect {
		t.Errorf("generateConvertFunction failed to pass, answer is [%s], expect is [%s]", ans, expect)
	}
}
