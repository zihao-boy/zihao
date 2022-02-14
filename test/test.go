package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("D://1.sql")
	defer f.Close()
	if err != nil {
		fmt.Print(err.Error())
	}

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString(';')
		if err != nil || io.EOF == err {
			break
		}

		fmt.Print(line)
	}

}
