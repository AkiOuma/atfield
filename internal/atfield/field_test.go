package atfield

import "testing"

func TestIfPublicField(t *testing.T) {
	name := "User"
	expect := true
	ans := ifPublicField(name)
	if expect != ans {
		t.Errorf("ifPublicField failed to pass, answer is [%v], expect is [%v]", ans, expect)
	}
}
