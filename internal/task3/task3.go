package task3

import (
	"bufio"
	"os"
)

func findDuplicate(strs []string, t int) rune {
	mem := make(map[rune]int)
	for _, s := range strs {
		uniq := make(map[rune]int)
		for _, ch := range s {
			_, ok := uniq[ch]
			if ok {
				continue
			}
			uniq[ch]++
			mem[ch]++
		}
	}
	var res rune
	for key := range mem {
		if mem[key] == t {
			return key
		}
	}
	return res
}

func SolutionFirstPart() int {
	f, err := os.Open("internal/task3/input.txt")
	if err != nil {
		return 0
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var res int
	for s.Scan() {
		s := s.Text()
		head, tail := s[:len(s)/2], s[len(s)/2:]
		d := findDuplicate([]string{head, tail}, 2)
		if int(d) >= 97 {
			// int('a') == 97
			res += int(d) - 96
		} else {
			// int('A') == 65
			res += int(d) - 38
		}
	}

	return res
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task3/input.txt")
	if err != nil {
		return 0
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var res int
	var stack []string
	for s.Scan() {
		s := s.Text()
		stack = append(stack, s)
		if len(stack) < 3 {
			continue
		}

		d := findDuplicate(stack, 3)
		if int(d) >= 97 {
			// int('a') == 97
			res += int(d) - 96
		} else {
			// int('A') == 65
			res += int(d) - 38
		}
		stack = []string{}
	}

	return res
}
