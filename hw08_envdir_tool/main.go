package main

import (
	"fmt"
	"os"
)

func main() {
	envMap, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	os.Exit(RunCmd(os.Args[2:], envMap))
}
