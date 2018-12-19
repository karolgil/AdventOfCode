package task16

import (
	"regexp"
	"strconv"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

type Opcode func(registers [6]int, a, b int) int

var OpcodesMatcher = map[string]Opcode{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func addr(register [6]int, a, b int) int {
	return register[a] + register[b]
}

func addi(register [6]int, a, b int) int {
	return register[a] + b
}

func mulr(register [6]int, a, b int) int {
	return register[a] * register[b]
}

func muli(register [6]int, a, b int) int {
	return register[a] * b
}

func banr(register [6]int, a, b int) int {
	return register[a] & register[b]
}

func bani(register [6]int, a, b int) int {
	return register[a] & b
}

func borr(register [6]int, a, b int) int {
	return register[a] | register[b]
}

func bori(register [6]int, a, b int) int {
	return register[a] | b
}

func setr(register [6]int, a, b int) int {
	return register[a]
}

func seti(register [6]int, a, b int) int {
	return a
}

func gtir(register [6]int, a, b int) int {
	if a > register[b] {
		return 1
	} else {
		return 0
	}
}

func gtri(register [6]int, a, b int) int {
	if register[a] > b {
		return 1
	} else {
		return 0
	}
}

func gtrr(register [6]int, a, b int) int {
	if register[a] > register[b] {
		return 1
	} else {
		return 0
	}
}

func eqir(register [6]int, a, b int) int {
	if a == register[b] {
		return 1
	} else {
		return 0
	}
}

func eqri(register [6]int, a, b int) int {
	if register[a] == b {
		return 1
	} else {
		return 0
	}
}

func eqrr(register [6]int, a, b int) int {
	if register[a] == register[b] {
		return 1
	} else {
		return 0
	}
}

func Solution1(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)
	ipRegister, _ := strconv.Atoi(string(lines[0][4]))
	lines = lines[1:]
	varsMatcher := regexp.MustCompile(`^(.+) (\d+) (\d+) (\d+)$`)
	register := [6]int{0, 0, 0, 0, 0, 0}
	ip := 0
	for ip >= 0 && ip < len(lines) {
		register[ipRegister] = ip
		vars := varsMatcher.FindAllStringSubmatch(lines[ip], -1)[0]
		op := string(vars[1])
		a, _ := strconv.Atoi(string(vars[2]))
		b, _ := strconv.Atoi(string(vars[3]))
		c, _ := strconv.Atoi(string(vars[4]))
		result := OpcodesMatcher[op](register, a, b)
		register[c] = result
		ip = register[ipRegister] + 1
	}
	return register[0]
}

func Solution2(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)
	ipRegister, _ := strconv.Atoi(string(lines[0][4]))
	lines = lines[1:]
	varsMatcher := regexp.MustCompile(`^(.+) (\d+) (\d+) (\d+)$`)
	register := [6]int{1, 0, 0, 0, 0, 0}
	ip := 0
	for ip >= 0 && ip < len(lines) {
		register[ipRegister] = ip
		vars := varsMatcher.FindAllStringSubmatch(lines[ip], -1)[0]
		op := string(vars[1])
		a, _ := strconv.Atoi(string(vars[2]))
		b, _ := strconv.Atoi(string(vars[3]))
		c, _ := strconv.Atoi(string(vars[4]))
		result := OpcodesMatcher[op](register, a, b)
		register[c] = result
		ip = register[ipRegister] + 1
		if register[0] == 0 {
			break
		}
	}
	sum := 0
	for i := 1; i <= register[1]; i++ {
		if register[1]%i == 0 {
			sum += i
		}
	}
	return sum
}
