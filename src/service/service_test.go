package service_test

import (
	"testing"

	. "github.com/ingcr3at1on/letstalkaboutmonorepos/src/service"
)

func TestCallLibA(t *testing.T) {
	str := CallLibA("test")
	if str != "test" {
		t.Fatal("how did you possibly manage to make this test fail?!?!?!?")
	}
}

func TestCallLibB(t *testing.T) {
	str := CallLibB("test")
	if str != "test" {
		t.Fatal("how did you possibly manage to make this test fail?!?!?!?")
	}
}
