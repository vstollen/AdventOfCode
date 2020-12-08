package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	acc int = iota
	jmp
	nop
)

type operation struct {
	op  int
	val int
}

func RunHandheldCode() {
	data := ReadData("day_8.txt")
	program := parseProgram(data)

	accUnchanged, _ := runProgram(program)
	accFixed := 0

	for i, _ := range program {
		op := program[i]

		if op.op == acc {
			continue
		}

		alteredProgram := make([]operation, len(program))
		copy(alteredProgram, program)

		switch op.op {
		case jmp:
			alteredProgram[i].op = nop
		case nop:
			alteredProgram[i].op = jmp
		}

		acc, err := runProgram(alteredProgram)

		if err == nil {
			accFixed = acc
			break
		}
	}

	fmt.Printf("Accumulator value before second execution: %v\n", accUnchanged)
	fmt.Printf("Accumulator after fix: %v\n", accFixed)
}

func runProgram(program []operation) (int, error) {
	visitedLines := make([]bool, len(program))

	accumulator := 0
	for i := 0; i < len(program); {
		if visitedLines[i] {
			return accumulator, errors.New("loop detected, execution stopped")
		}

		visitedLines[i] = true

		op := program[i]

		switch op.op {
		case acc:
			accumulator += op.val
		case jmp:
			i += op.val
			continue
		case nop:
		}

		i++
	}

	return accumulator, nil
}

func parseProgram(data []string) []operation {
	program := make([]operation, len(data))

	for i, _ := range data {
		program[i] = parseOperation(data[i])
	}

	return program
}

func parseOperation(operationString string) operation {
	opcode, valueString := splitAtFirst(operationString, " ")

	op := operation{}

	switch opcode {
	case "acc":
		op.op = acc
	case "jmp":
		op.op = jmp
	case "nop":
		op.op = nop
	}

	value, _ := strconv.Atoi(valueString)

	op.val = value

	return op
}
