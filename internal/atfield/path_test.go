package atfield

import (
	"os"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	ans, err := absolutePath(".")
	if err != nil {
		t.Errorf("absolutePath failed to pass because of error: %s", err.Error())
	}
	expect, _ := os.Getwd()
	if ans != expect {
		t.Errorf("absolutePath failed to pass, answer is [%s], expect is [%s]", ans, expect)
	}
}
