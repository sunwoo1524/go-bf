package main

import (
	"fmt"

	"io/ioutil"

	"os"

	interpreter "github.com/sunwoo1524/go-bf/brainfuck"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args

	if len(args) == 0 {
		fmt.Println("Please input brainfuck file name.")
		return
	}

	data, err := ioutil.ReadFile(args[1])

	errCheck(err)

	interpreter.RunProgram(string(data))
}
