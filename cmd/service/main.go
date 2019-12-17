package main

import (
	"fmt"

	"github.com/ingcr3at1on/letstalkaboutmonorepos/src/service"
)

func main() {
	fmt.Println(service.CallLibA("FOO"))
	fmt.Println(service.CallLibB("BAR"))
}
