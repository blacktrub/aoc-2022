package task4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func newRange(s string) Range {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Range{start: start, end: end}
}

func SolutionFirstPart() int {
	f, err := os.Open("internal/task4/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var res int
	for s.Scan() {
		v := s.Text()
		r := strings.Split(v, ",")
		first, second := newRange(r[0]), newRange(r[1])
		if first.start <= second.start && first.end >= second.end {
			res++
		} else if second.start <= first.start && second.end >= first.end {
			res++
		}
	}
	return res
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task4/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var res int
	for s.Scan() {
		v := s.Text()
		r := strings.Split(v, ",")
		first, second := newRange(r[0]), newRange(r[1])
		if first.start > second.start {
			first, second = second, first
		}

		if second.start <= first.end {
			res++
		}

	}

	return res
}
