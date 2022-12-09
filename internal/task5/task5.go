package task5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//         [H]     [W] [B]
//     [D] [B]     [L] [G] [N]
// [P] [J] [T]     [M] [R] [D]
// [V] [F] [V]     [F] [Z] [B]     [C]
// [Z] [V] [S]     [G] [H] [C] [Q] [R]
// [W] [W] [L] [J] [B] [V] [P] [B] [Z]
// [D] [S] [M] [S] [Z] [W] [J] [T] [G]
// [T] [L] [Z] [R] [C] [Q] [V] [P] [H]
//  1   2   3   4   5   6   7   8   9

func getStack() [][]string {
	return [][]string{
		[]string{"T", "D", "W", "Z", "V", "P"},
		[]string{"L", "S", "W", "V", "F", "J", "D"},
		[]string{"Z", "M", "L", "S", "V", "T", "B", "H"},
		[]string{"R", "S", "J"},
		[]string{"C", "Z", "B", "G", "F", "M", "L", "W"},
		[]string{"Q", "W", "V", "H", "Z", "R", "G", "B"},
		[]string{"V", "J", "P", "C", "B", "D", "N"},
		[]string{"P", "T", "B", "Q"},
		[]string{"H", "G", "Z", "R", "C"},
	}

}

func SolutionFirstPart() string {
	f, err := os.Open("internal/task5/input.txt")
	if err != nil {
		return ""
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	stack := getStack()
	for s.Scan() {
		// move 3 from 2 to 9
		v := s.Text()
		parts := strings.Split(v, " ")
		times, _ := strconv.Atoi(parts[1])
		frm, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		frm--
		to--

		for i := 0; i < times; i++ {
			item := stack[frm][len(stack[frm])-1]
			stack[frm] = stack[frm][:len(stack[frm])-1]
			stack[to] = append(stack[to], item)
		}

	}

	var res string
	for i := 0; i < len(stack); i++ {
		res += stack[i][len(stack[i])-1]
	}

	return res
}

func SolutionSecondPart() string {
	f, err := os.Open("internal/task5/input.txt")
	if err != nil {
		return ""
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	stack := getStack()
	for s.Scan() {
		// move 3 from 2 to 9
		v := s.Text()
		parts := strings.Split(v, " ")
		times, _ := strconv.Atoi(parts[1])
		frm, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		frm--
		to--

		item := stack[frm][len(stack[frm])-times:]
		stack[frm] = stack[frm][:len(stack[frm])-times]
		for i := 0; i < len(item); i++ {
			stack[to] = append(stack[to], item[i])
		}

	}

	var res string
	for i := 0; i < len(stack); i++ {
		res += stack[i][len(stack[i])-1]
	}

	return res
}
