package atfield

type TypeSet map[string]struct{}

var _integers = []string{
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
	"byte", "rune",
}

var _floats = []string{
	"float32", "float64",
}

var _bytes = []string{
	"string",
	"[]byte",
}

var _runes = []string{
	"string",
	"[]rune",
}

var _typesList []TypeSet

func ifTypeConvertible(type1, type2 string) bool {
	for _, typ := range typesList() {
		if _, ok := typ[type1]; ok {
			if _, ok = typ[type2]; ok {
				return true
			}
		}
	}
	return false
}

func typesList() []TypeSet {
	if _typesList == nil {
		_typesList = make([]TypeSet, 0, 4)
		for _, v := range [][]string{_integers, _floats, _bytes, _runes} {
			_typesList = append(_typesList, initTypeSet(v))
		}
	}
	return _typesList
}

func initTypeSet(types []string) TypeSet {
	result := make(TypeSet)
	for _, v := range types {
		result[v] = struct{}{}
	}
	return result
}
