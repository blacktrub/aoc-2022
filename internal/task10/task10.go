package task10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Register struct {
	c         int
	v         int
	strength  int
	nextCycle int
	maxCycle  int
}

func (r *Register) add(x int) {
	r.v += x
}

func (r *Register) cycle() {
	r.c++

	if r.c > r.maxCycle {
		return
	}

	if r.c == r.nextCycle {
		r.strength = r.strength + (r.c * r.v)
		r.nextCycle += 40
	}
}

type SecondRegister struct {
	c   int
	v   int
	cur string
	pic []string
}

func (r *SecondRegister) add(x int) {
	r.v += x
}

func (r *SecondRegister) cycle() {
	r.c++

	symbol := "."
	if r.c >= r.v && r.c <= r.v+2 {
		symbol = "#"
	}
	r.cur += symbol

	if r.c%40 == 0 {
		r.c = 0
		r.pic = append(r.pic, r.cur)
		r.cur = ""
	}
}

func SolutionFirstPart() int {
	f, err := os.Open("internal/task10/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	reg := Register{v: 1, c: 0, nextCycle: 20, maxCycle: 220}
	for s.Scan() {
		v := strings.Split(s.Text(), " ")
		switch v[0] {
		case "noop":
			reg.cycle()
		case "addx":
			val, _ := strconv.Atoi(v[1])
			reg.cycle()
			reg.cycle()
			reg.add(val)
		}
	}
	return reg.strength
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task10/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	reg := SecondRegister{c: 0, v: 1, pic: []string{}}
	for s.Scan() {
		v := strings.Split(s.Text(), " ")
		switch v[0] {
		case "noop":
			reg.cycle()
		case "addx":
			val, _ := strconv.Atoi(v[1])
			reg.cycle()
			reg.cycle()
			reg.add(val)
		}
	}

	for i := 0; i < len(reg.pic); i++ {
		fmt.Println(reg.pic[i])
	}

	return 0
}
