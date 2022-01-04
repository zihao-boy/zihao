package main

import (
	"fmt"
	"strings"
)

func main() {

	imagesIds := "sdfsdf,sdfsf"
	if strings.HasSuffix(imagesIds,","){
		imagesIds = imagesIds[0:len(imagesIds)-1]
	}

	fmt.Print(imagesIds)

}
