package main

import (
	"fmt"
	"strings"
)

func main() {
	gitUrl := "http://git.homecommunity.cn"
	git_url :=  strings.Replace(gitUrl, "://", "://123:123123@", 1)

	fmt.Print(git_url)

}
