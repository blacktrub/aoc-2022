package task6

import (
	"bufio"
	"os"
)

func SolutionFirstPart() int {
	f, err := os.Open("internal/task6/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	s.Scan()
	v := s.Text()

	i := 0
	for j := 3; j < len(v); j++ {
		m := map[byte]int{}
		res := true
		for k := i; k <= j; k++ {
			_, ok := m[v[k]]
			if ok {
				res = false
				break
			}
			m[v[k]]++
		}

		if res {
			return j + 1
		}

		i++
	}

	return 0
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task6/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	s.Scan()
	v := s.Text()

	i := 0
	// TODO: you can better
	for j := 13; j < len(v); j++ {
		m := map[byte]int{}
		res := true
		for k := i; k <= j; k++ {
			_, ok := m[v[k]]
			if ok {
				res = false
				break
			}
			m[v[k]]++
		}

		if res {
			return j + 1
		}

		i++
	}

	return 0
}
