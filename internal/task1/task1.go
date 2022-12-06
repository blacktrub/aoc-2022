package task1

import (
	"bufio"
	"os"
	"strconv"
)

func SolutionFirstPart() int {
	f, err := os.Open("internal/task1/input.txt")
	if err != nil {
		return 0
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var res, cur int
	for s.Scan() {
		v := s.Text()
		if v == "" {
			if cur > res {
				res = cur
			}
			cur = 0
		} else {
			i, _ := strconv.Atoi(v)
			cur = cur + i
		}
	}
	return res
}

type Node struct {
	v    int
	next *Node
}

type Queue struct {
	s    int
	head *Node
}

func (q *Queue) Cut() {
	for q.s > 3 {
		q.head = q.head.next
		q.s--
	}
}

func (q *Queue) Put(v int) {
	cur := q.head
	if v < cur.v {
		return
	}

	for cur.v < v {
		n := cur.next
		if n == nil {
			break
		}

		if n.v > v {
			break
		}

		cur = n
	}

	cur.next = &Node{v: v, next: cur.next}
	q.s++
	q.Cut()
}

func (q Queue) Sum() int {
	var s int
	cur := q.head
	for cur != nil {
		s = s + cur.v
		cur = cur.next
	}
	return s
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task1/input.txt")
	if err != nil {
		return 0
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	q := Queue{s: 1, head: &Node{v: 0}}
	var cur int
	for s.Scan() {
		v := s.Text()
		if v == "" {
			q.Put(cur)
			cur = 0
		} else {
			i, _ := strconv.Atoi(v)
			cur = cur + i
		}
	}
	return q.Sum()
}
