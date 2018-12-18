package task16

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

type Opcode func(registers [4]int, a, b int) int

var Opcodes = []Opcode{
	addr, addi,
	mulr, muli,
	banr, bani,
	borr, bori,
	setr, seti,
	gtir, gtri, gtrr,
	eqir, eqri, eqrr,
}

func addr(register [4]int, a, b int) int {
	return register[a] + register[b]
}

func addi(register [4]int, a, b int) int {
	return register[a] + b
}

func mulr(register [4]int, a, b int) int {
	return register[a] * register[b]
}

func muli(register [4]int, a, b int) int {
	return register[a] * b
}

func banr(register [4]int, a, b int) int {
	return register[a] & register[b]
}

func bani(register [4]int, a, b int) int {
	return register[a] & b
}

func borr(register [4]int, a, b int) int {
	return register[a] | register[b]
}

func bori(register [4]int, a, b int) int {
	return register[a] | b
}

func setr(register [4]int, a, b int) int {
	return register[a]
}

func seti(register [4]int, a, b int) int {
	return a
}

func gtir(register [4]int, a, b int) int {
	if a > register[b] {
		return 1
	} else {
		return 0
	}
}

func gtri(register [4]int, a, b int) int {
	if register[a] > b {
		return 1
	} else {
		return 0
	}
}

func gtrr(register [4]int, a, b int) int {
	if register[a] > register[b] {
		return 1
	} else {
		return 0
	}
}

func eqir(register [4]int, a, b int) int {
	if a == register[b] {
		return 1
	} else {
		return 0
	}
}

func eqri(register [4]int, a, b int) int {
	if register[a] == b {
		return 1
	} else {
		return 0
	}
}

func eqrr(register [4]int, a, b int) int {
	if register[a] == register[b] {
		return 1
	} else {
		return 0
	}
}

type InputInstruction struct {
	Before    [4]int
	Operation [4]int
	After     [4]int
}

func StoI(digitChars [4]string) [4]int {
	var result [4]int
	for i, c := range digitChars {
		j, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		result[i] = j
	}
	return result
}

func Solution1(inputFile string) int {
	var inputInstructions []InputInstruction

	lines, _ := utils.ReadLinesFrom(inputFile)

	beforeRegex := regexp.MustCompile(`^Before: \[(\d+), (\d+), (\d+), (\d+)]$`)
	operationRegex := regexp.MustCompile(`^(\d+) (\d+) (\d+) (\d+)$`)
	afterRegex := regexp.MustCompile(`^After: {2}\[(\d+), (\d+), (\d+), (\d+)]$`)
	for i := 0; i < len(lines); i += 4 {
		before := beforeRegex.FindAllStringSubmatch(lines[i], -1)[0]
		operation := operationRegex.FindAllStringSubmatch(lines[i+1], -1)[0]
		after := afterRegex.FindAllStringSubmatch(lines[i+2], -1)[0]
		inputInstructions = append(inputInstructions, InputInstruction{
			Before:    StoI([4]string{before[1], before[2], before[3], before[4]}),
			Operation: StoI([4]string{operation[1], operation[2], operation[3], operation[4]}),
			After:     StoI([4]string{after[1], after[2], after[3], after[4]}),
		})
	}

	threeMatchesOrMore := 0

	for _, inputInstruction := range inputInstructions {
		opMatchCount := 0
		for _, op := range Opcodes {
			resultIndex := inputInstruction.Operation[3]
			output := op(inputInstruction.Before, inputInstruction.Operation[1], inputInstruction.Operation[2])
			if inputInstruction.After[resultIndex] == output {
				opMatchCount += 1
			}
		}
		if opMatchCount >= 3 {
			threeMatchesOrMore++
		}
	}

	return threeMatchesOrMore
}

func Solution2(inputFile, testFile string) int {
	var inputInstructions []InputInstruction

	lines, _ := utils.ReadLinesFrom(inputFile)

	beforeRegex := regexp.MustCompile(`^Before: \[(\d+), (\d+), (\d+), (\d+)]$`)
	operationRegex := regexp.MustCompile(`^(\d+) (\d+) (\d+) (\d+)$`)
	afterRegex := regexp.MustCompile(`^After: {2}\[(\d+), (\d+), (\d+), (\d+)]$`)
	for i := 0; i < len(lines); i += 4 {
		before := beforeRegex.FindAllStringSubmatch(lines[i], -1)[0]
		operation := operationRegex.FindAllStringSubmatch(lines[i+1], -1)[0]
		after := afterRegex.FindAllStringSubmatch(lines[i+2], -1)[0]
		inputInstructions = append(inputInstructions, InputInstruction{
			Before:    StoI([4]string{before[1], before[2], before[3], before[4]}),
			Operation: StoI([4]string{operation[1], operation[2], operation[3], operation[4]}),
			After:     StoI([4]string{after[1], after[2], after[3], after[4]}),
		})
	}

	var operationFrequencies [16][16]int

	for _, inputInstruction := range inputInstructions {
		for opCode, op := range Opcodes {
			resultIndex := inputInstruction.Operation[3]
			output := op(inputInstruction.Before, inputInstruction.Operation[1], inputInstruction.Operation[2])
			if inputInstruction.After[resultIndex] == output {
				operationFrequencies[inputInstruction.Operation[0]][opCode]++
			}
		}
	}

	operations := make(map[int]Opcode)
	for len(operations) < 16 {
		for opCode, frequencies := range operationFrequencies {
			setCount := 0
			setPos := -1
			for i, count := range frequencies {
				if count > 0 {
					setCount++
					setPos = i
				}
			}
			if setCount == 1 {
				operations[opCode] = Opcodes[setPos]
				for otherOpCode := range operationFrequencies {
					if otherOpCode != opCode {
						operationFrequencies[otherOpCode][setPos] = 0
					}
				}
			}
		}
	}

	lines, _ = utils.ReadLinesFrom(testFile)

	var testOperations [][4]int

	for _, line := range lines {
		operation := operationRegex.FindAllStringSubmatch(line, -1)[0]
		testOperations = append(testOperations, StoI([4]string{operation[1], operation[2], operation[3], operation[4]}))
	}

	registers := [4]int{0, 0, 0, 0}
	for _, operation := range testOperations {
		fmt.Printf("Operation: %v\n", operation)
		resultIndex := operation[3]
		result := operations[operation[0]](registers, operation[1], operation[2])
		registers[resultIndex] = result
		fmt.Printf("Registers: %v\n", registers)
	}

	return registers[0]
}
