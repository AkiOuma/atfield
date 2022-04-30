package atfield

import (
	"fmt"
	"go/types"
	"log"
	"strings"
)

// extract items field informations
func (e *engine) extractFields() {
	structs := make(structSet)
	for _, use := range e.pkg.TypesInfo.Uses {
		use, ok := use.(*types.TypeName)
		if !ok {
			continue
		}
		var pkg string
		structName := use.Name()
		if p := use.Pkg(); p != nil {
			pkg = p.Name()
		}
		structid := structId{PkgName: pkg, StructName: structName}
		if structs[structid] == nil {
			structs[structid] = make(fieldSet)
		}
		typ, ok := use.Type().(*types.Named)
		if !ok {
			continue
		}
		fields, ok := typ.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		for i := 0; i < fields.NumFields(); i++ {
			field := fields.Field(i)
			fieldName := field.Name()
			if !ifPublicField(fieldName) {
				continue
			}
			fieldType := getFieldType(field.Type())
			// if unrecordize type then drop convert it
			if len(fieldType) == 0 {
				continue
			}
			structs[structid][fieldName] = fieldType
		}
	}
	e.structSet = structs
}

func getFieldType(fieldType types.Type, prefix ...string) string {
	typeName := strings.Join(prefix, "")
	switch typ := fieldType.(type) {
	case *types.Basic:
		typeName = typeName + typ.Name()
	case *types.Named:
		typeName = typeName + typ.Obj().Pkg().Name() + "." + typ.Obj().Name()
	case *types.Pointer:
		typeName = getFieldType(typ.Elem(), append(prefix, "*")...)
	case *types.Slice:
		typeName = getFieldType(typ.Elem(), append(prefix, "[]")...)
	case *types.Map:
		keyType := getFieldType(typ.Key(), prefix...)
		valueType := getFieldType(typ.Elem(), prefix...)
		typeName = fmt.Sprintf("map[%s]%s", keyType, valueType)
	default:
		log.Printf("unrecordize type: %v", typ)
	}
	return typeName
}
