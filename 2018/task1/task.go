package task1

import (
	"bufio"
	"container/ring"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
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

func Solution2(inputFile string) (int, error) {
	lines, err := readLinesFrom(inputFile)
	if err != nil {
		return 0, nil
	}

	r := ring.New(len(lines))
	for _, str := range lines {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, nil
		}
		r.Value = num
		r = r.Next()
	}

	result, err := findFirstCycleSum(r)
	if err != nil {
		return 0, nil
	}

	return result, nil
}

func findFirstCycleSum(r *ring.Ring) (int, error) {
	timeout := time.After(1 * time.Second)
	resultsSoFar := make(map[int]bool)
	var result int
	for {
		select {
		case <-timeout:
			return 0, errors.New("timeout reached")
		default:
			result += r.Value.(int)
			if resultsSoFar[result] {
				return result, nil
			} else {
				r = r.Next()
				resultsSoFar[result] = true
			}
		}
	}
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
