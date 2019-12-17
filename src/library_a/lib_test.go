package librarya_test

import (
	"testing"

	. "github.com/ingcr3at1on/letstalkaboutmonorepos/src/library_a"
)

func TestReturnMyString(t *testing.T) {
	str := ReturnMyString("test")
	if str != "test" {
		t.Fatal("how did you possibly manage to make this test fail?!?!?!?")
	}
}
