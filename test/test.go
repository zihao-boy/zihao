package main

import (
	"fmt"
	"path"
)

func main() {
	a := "/abd/asdf/dfdf/sdfsd.jar"
	fmt.Print(path.Dir(a))

}
