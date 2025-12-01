package util

import (
	"bufio"
	"os"
)

func ReadInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines = make([]string, 0)

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	return lines
}
