package task1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Solution1(inputFile string) (int, error) {
	lines, err := readLinesFrom(inputFile)
	if err != nil {
		return 0, nil
	}
	var result int
	for _, str := range lines {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, nil
		}
		result += num
	}
	return result, nil
}

func readLinesFrom(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
