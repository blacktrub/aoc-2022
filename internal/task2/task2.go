package task2

import (
	"bufio"
	"os"
	"strings"
)

func SolutionFirstPart() int {
	f, err := os.Open("internal/task2/input.txt")
	if err != nil {
		return 0
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	// (1 for Rock, 2 for Paper, and 3 for Scissors)
	// X for Rock, Y for Paper, and Z for Scissors
	shapeCost := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
	// A for Rock, B for Paper, and C for Scissors
	win := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	mirror := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	var res int
	for s.Scan() {
		v := s.Text()
		values := strings.Split(v, " ")
		opponent, our := values[0], values[1]
		res = res + shapeCost[our]
		if opponent == mirror[our] {
			res = res + 3
		} else if opponent == win[our] {
			res = res + 6
		}
	}
	return res
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task2/input.txt")
	if err != nil {
		return 0
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	resultCost := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	costPerShape := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	toLose := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	toWin := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	var res int
	for s.Scan() {
		v := s.Text()
		values := strings.Split(v, " ")
		opponent, our := values[0], values[1]
		res = res + resultCost[our]

		if our == "Y" {
			res = res + costPerShape[opponent]
		} else if our == "X" {
			res = res + costPerShape[toLose[opponent]]
		} else {
			res = res + costPerShape[toWin[opponent]]
		}
	}
	return res
}
