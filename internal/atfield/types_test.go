package atfield

import "testing"

func TestIfTypeConvertible(t *testing.T) {
	type1, type2 := "int", "int"
	ans, expect := ifTypeConvertible(type1, type2), true
	if ans != expect {
		t.Errorf("ifTypeConvertible failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
	type1, type2 = "int32", "int64"
	ans, expect = ifTypeConvertible(type1, type2), true
	if ans != expect {
		t.Errorf("ifTypeConvertible failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
	type1, type2 = "int32", "float64"
	ans, expect = ifTypeConvertible(type1, type2), false
	if ans != expect {
		t.Errorf("ifTypeConvertible failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
	type1, type2 = "string", "[]byte"
	ans, expect = ifTypeConvertible(type1, type2), true
	if ans != expect {
		t.Errorf("ifTypeConvertible failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
	type1, type2 = "string", "[]rune"
	ans, expect = ifTypeConvertible(type1, type2), true
	if ans != expect {
		t.Errorf("ifTypeConvertible failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
	type1, type2 = "[]byte", "[]rune"
	ans, expect = ifTypeConvertible(type1, type2), false
	if ans != expect {
		t.Errorf("ifTypeConvertible failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
}
