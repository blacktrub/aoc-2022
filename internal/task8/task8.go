package task8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isVisible(grid [][]int, row, col, val int) bool {
	if row == 0 || row == len(grid)-1 || col == 0 || col == len(grid[0])-1 {
		return true
	}

	var res bool
	for i := col + 1; i < len(grid[row]); i++ {
		if grid[row][i] >= val {
			res = false
			break
		}
		res = true
	}
	if res {
		return res
	}

	for i := col - 1; i >= 0; i-- {
		if grid[row][i] >= val {
			res = false
			break
		}
		res = true
	}
	if res {
		return res
	}

	for i := row + 1; i < len(grid); i++ {
		if grid[i][col] >= val {
			res = false
			break
		}
		res = true
	}

	if res {
		return res
	}

	for i := row - 1; i >= 0; i-- {
		if grid[i][col] >= val {
			res = false
			break
		}
		res = true
	}

	return res
}

func SolutionFirstPart() int {
	f, err := os.Open("internal/task8/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	grid := [][]int{}
	for s.Scan() {
		line := s.Text()
		rows := strings.Split(line, "")
		grid = append(grid, []int{})
		for i := 0; i < len(rows); i++ {
			v, _ := strconv.Atoi(rows[i])
			grid[len(grid)-1] = append(grid[len(grid)-1], v)
		}
	}

	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isVisible(grid, i, j, grid[i][j]) {
				res++
			}
		}
	}

	return res
}

type Item struct {
	i int
	v int
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task8/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	grid := [][]int{}
	for s.Scan() {
		line := s.Text()
		rows := strings.Split(line, "")
		grid = append(grid, []int{})
		for i := 0; i < len(rows); i++ {
			v, _ := strconv.Atoi(rows[i])
			grid[len(grid)-1] = append(grid[len(grid)-1], v)
		}
	}

	res := [][]int{}
	for i := 0; i < len(grid); i++ {
		res = append(res, make([]int, len(grid[i])))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			res[i][j] = 1
		}
	}

	for i := 0; i < len(grid); i++ {
		stack := []Item{}
		for j := 0; j < len(grid[i]); j++ {
			cur := grid[i][j]

			for len(stack) > 0 && stack[len(stack)-1].v < cur {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				res[i][j] *= j
			} else {
				last := stack[len(stack)-1]
				res[i][j] *= (j - last.i)

			}

			stack = append(stack, Item{i: j, v: cur})

		}
	}

	for i := 0; i < len(grid); i++ {
		stack := []Item{}
		for j := len(grid[i]) - 1; j >= 0; j-- {
			cur := grid[i][j]

			for len(stack) > 0 && stack[len(stack)-1].v < cur {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				res[i][j] *= len(grid[i]) - 1 - j
			} else {
				last := stack[len(stack)-1]
				res[i][j] *= last.i - j

			}

			stack = append(stack, Item{i: j, v: cur})
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		stack := []Item{}
		for i := 0; i < len(grid); i++ {
			cur := grid[i][j]

			for len(stack) > 0 && stack[len(stack)-1].v < cur {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				res[i][j] *= i
			} else {
				last := stack[len(stack)-1]
				res[i][j] *= (i - last.i)

			}

			stack = append(stack, Item{i: i, v: cur})

		}
	}

	for j := 0; j < len(grid[0]); j++ {
		stack := []Item{}
		for i := len(grid) - 1; i >= 0; i-- {
			cur := grid[i][j]

			for len(stack) > 0 && stack[len(stack)-1].v < cur {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				res[i][j] *= len(grid) - 1 - i
			} else {
				last := stack[len(stack)-1]
				res[i][j] *= (last.i - i)

			}

			stack = append(stack, Item{i: i, v: cur})

		}
	}

	var best int
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] > best {
				best = res[i][j]
			}
		}
	}

	return best
}
