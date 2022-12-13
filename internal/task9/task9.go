package task9

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x int
	y int
}

func (h *Node) step(dir string) {

	switch dir {
	case "L":
		h.y--
	case "R":
		h.y++
	case "D":
		h.x--
	case "U":
		h.x++
	}
}
func (t *Node) follow(h Node) {
	if h.x == t.x && h.y == t.y {
		return
	}

	if h.x == t.x || h.y == t.y {
		switch h.x - t.x {
		case 2:
			t.x++
		case -2:
			t.x--
		}

		switch h.y - t.y {
		case 2:
			t.y++
		case -2:
			t.y--
		}
	} else {
		var dx, dy int
		if h.x > t.x {
			dx = h.x - t.x
		} else {
			dx = t.x - h.x
		}

		if h.y > t.y {
			dy = h.y - t.y
		} else {
			dy = t.y - h.y
		}

		if dx > 1 || dy > 1 {
			if h.x > t.x {
				t.x++
			} else {
				t.x--
			}

			if h.y > t.y {
				t.y++
			} else {
				t.y--
			}
		}

	}

}

func SolutionFirstPart() int {
	f, err := os.Open("internal/task9/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	head := Node{}
	tail := Node{}

	seen := map[[2]int]bool{}
	for s.Scan() {
		v := strings.Split(s.Text(), " ")
		dir := v[0]
		steps, _ := strconv.Atoi(v[1])

		for i := 0; i < steps; i++ {
			head.step(dir)
			tail.follow(head)
			seen[[2]int{tail.x, tail.y}] = true
		}
	}

	return len(seen)
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task9/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	snake := [10]Node{}

	seen := map[[2]int]bool{}
	for s.Scan() {
		v := strings.Split(s.Text(), " ")
		dir := v[0]
		steps, _ := strconv.Atoi(v[1])

		for i := 0; i < steps; i++ {
			snake[9].step(dir)
			for j := len(snake) - 2; j >= 0; j-- {
				snake[j].follow(snake[j+1])
			}
			t := snake[0]
			seen[[2]int{t.x, t.y}] = true
		}

	}

	return len(seen)
}
