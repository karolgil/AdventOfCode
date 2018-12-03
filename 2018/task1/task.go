package task1

import (
	"container/ring"
	"errors"
	"strconv"
	"time"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
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
	lines, err := utils.ReadLinesFrom(inputFile)
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
