package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesFrom(fileName string) ([]string, error) {
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

func Alphabet() string {
	b := make([]byte, 26)
	for i := range b {
		b[i] = 'a' + byte(i)
	}
	return string(b)
}

func AlphabetLowerAndUpper() string {
	b := make([]byte, 52)
	for i := range b[:len(b)/2] {
		b[i] = 'a' + byte(i)
	}
	for i := range b[len(b)/2:] {
		b[i+26] = 'A' + byte(i)
	}
	return string(b)
}
