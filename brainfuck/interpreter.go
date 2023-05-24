package interpreter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode/utf8"
)

const (
	dec_ptr = iota
	inc_ptr

	decrease
	increase

	output
	input

	loop_start
	loop_end
)

const mem_size int = 30000

var ptr int = 0
var memory [mem_size]uint8 = [mem_size]uint8{}

func Compile(code string) (compiled [][2]int, err error) {
	stack := []int{}
	i := 0

	for n := 0; n < utf8.RuneCountInString(code); n++ {
		switch string([]rune(code)[n]) {
		case "<":
			compiled = append(compiled, [2]int{dec_ptr, 0})
		case ">":
			compiled = append(compiled, [2]int{inc_ptr, 0})

		case "-":
			compiled = append(compiled, [2]int{decrease, 0})
		case "+":
			compiled = append(compiled, [2]int{increase, 0})

		case ".":
			compiled = append(compiled, [2]int{output, 0})
		case ",":
			compiled = append(compiled, [2]int{input, 0})

		case "[":
			compiled = append(compiled, [2]int{loop_start, 0})
			stack = append(stack, i)
		case "]":
			if len(stack) == 0 {
				return nil, errors.New("brainf**king syntax error")
			}

			compiled = append(compiled, [2]int{loop_end, stack[len(stack)-1]})
			compiled[stack[len(stack)-1]][1] = i
			stack = stack[:len(stack)-1]
		default:
			i--
		}

		i++
	}

	if len(stack) != 0 {
		return nil, errors.New("brainf**king syntax error")
	}

	return
}

func Execute(code string) {
	program, err := Compile(code)

	if program == nil {
		panic(err)
	}

	for i := 0; i < len(program); i++ {
		e := program[i]

		switch e[0] {
		case dec_ptr:
			if ptr <= 0 {
				ptr = mem_size - 1
				break
			}

			ptr--
		case inc_ptr:
			if ptr >= mem_size-1 {
				ptr = 0
				break
			}

			ptr++

		case decrease:
			if memory[ptr] <= 0 {
				memory[ptr] = 255
				break
			}

			memory[ptr]--
		case increase:
			if memory[ptr] >= 255 {
				memory[ptr] = 0
				break
			}

			memory[ptr]++

		case output:
			fmt.Printf("%s", string(memory[ptr]))
		case input:
			in := bufio.NewReader(os.Stdin)
			line, _ := in.ReadString('\n')
			memory[ptr] = []byte(line)[0]

		case loop_start:
			if memory[ptr] == 0 {
				i = e[1] - 1
			}
		case loop_end:
			if memory[ptr] != 0 {
				i = e[1] - 1
			}
		}

		// fmt.Printf("%d, %d, %d, %d, %d\n", memory[0], memory[1], memory[2], memory[3], memory[4])
	}
}
