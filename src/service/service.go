package service

import (
	a "github.com/ingcr3at1on/letstalkaboutmonorepos/src/library_a"
	b "github.com/ingcr3at1on/letstalkaboutmonorepos/src/library_b"
)

// CallLibA takes a string and returns the output of a.ReturnMyString.
func CallLibA(str string) string {
	return a.ReturnMyString(str)
}

// CallLibB takes a string and returns the output of b.ReturnMyString.
func CallLibB(str string) string {
	return b.ReturnMyString(str)
}
