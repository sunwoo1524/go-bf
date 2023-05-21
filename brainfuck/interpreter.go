package interpreter

import (
	"fmt"
	"unicode/utf8"
)

func RunProgram(code string) {
	const mem_size int = 30000

	line := 0
	column := 0
	brackets := [][3]int{}
	memory := [mem_size]uint8{}
	pointer := 0

	for index := 0; index < utf8.RuneCountInString(code); index++ {
		switch string(code[index]) {
		case "<":
			pointer--

			if pointer < 0 {
				pointer = mem_size - 1
			}
		case ">":
			pointer++

			if pointer >= mem_size {
				pointer = 0
			}
		case "-":
			if memory[pointer] == 0 {
				memory[pointer] = 255
			} else {
				memory[pointer]--
			}
		case "+":
			if memory[pointer] == 255 {
				memory[pointer] = 0
			} else {
				memory[pointer]++
			}
		case ".":
			fmt.Printf("%s", string(memory[pointer]))
		case ",":
			// var input string
			// fmt.Scanf("%s", &input)

			// memory[pointer] = int(input)
		case "[":
			brackets = append(brackets, [3]int{index, line, column})
		case "]":
			if memory[pointer] != 0 {
				if len(brackets) == 0 {
					fmt.Printf("Brainf**king Error: SyntaxError %d:%d", line+1, column+1)
					return
				} else {
					start_bracket := brackets[len(brackets)-1]
					index = start_bracket[0] - 1
					line = start_bracket[1]
					column = start_bracket[2] - 1
				}
			}

			brackets = brackets[:len(brackets)-1]

			fallthrough
		case "\n":
			line++
			column = -1
		}

		column++
	}
}
