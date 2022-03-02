package main

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/utils"
	"strings"
)

func main() {
	line := "drwx------    2 ftp      ftp            24 Mar 02 05:55 abc"

	lines := strings.Split(line," ")
	var newLines []string
	for _,l := range lines{
		if utils.IsEmpty(l){
			continue
		}
		newLines = append(newLines,l)
	}

	fmt.Println(newLines,newLines[len(newLines)-1])
}