package task7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	children map[string]*Node
	size     int
	total    int
	name     string
}

func (n *Node) cd(name string) *Node {
	if name == ".." {
		return n.parent
	}

	node, ok := n.children[name]
	if ok {
		return node
	}

	node = &Node{name: name, children: map[string]*Node{}, parent: n}
	n.children[name] = node
	return node
}

func (n *Node) totalSize() int {
	if n.total != 0 {
		return n.total
	}

	size := n.size
	for key := range n.children {
		size += n.children[key].totalSize()
	}
	n.total = size
	return size
}

func findDirs(n *Node) int {
	if n == nil {
		return 0
	}

	total := n.totalSize()
	if total > 100_000 {
		total = 0
	}

	for key := range n.children {
		total += findDirs(n.children[key])
	}
	return total
}

func findWhichDelete(n *Node, m int, req int) int {
	if n == nil {
		return m
	}

	total := n.totalSize()
	if total >= req && total < m {
		m = total
	}

	for key := range n.children {
		newM := findWhichDelete(n.children[key], m, req)
		if newM < m {
			m = newM
		}
	}
	return m
}

func SolutionFirstPart() int {
	f, err := os.Open("internal/task7/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	root := &Node{name: "/", children: map[string]*Node{}}
	root.parent = root
	head := root
	for s.Scan() {
		v := s.Text()
		parts := strings.Split(v, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					head = root
				} else {
					head = head.cd(parts[2])
				}
			}
		} else {
			line := strings.Split(s.Text(), " ")
			if line[0] == "dir" {
				_, ok := head.children[line[1]]
				if !ok {
					head.children[line[1]] = &Node{name: line[1], children: map[string]*Node{}, parent: head}
				}
			} else {
				size, _ := strconv.Atoi(line[0])
				head.size += size
			}
		}
	}

	return findDirs(root)
}

func SolutionSecondPart() int {
	f, err := os.Open("internal/task7/input.txt")
	if err != nil {
		return 0
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	root := &Node{name: "/", children: map[string]*Node{}}
	root.parent = root
	head := root
	for s.Scan() {
		v := s.Text()
		parts := strings.Split(v, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					head = root
				} else {
					head = head.cd(parts[2])
				}
			}
		} else {
			line := strings.Split(s.Text(), " ")
			if line[0] == "dir" {
				_, ok := head.children[line[1]]
				if !ok {
					head.children[line[1]] = &Node{name: line[1], children: map[string]*Node{}, parent: head}
				}
			} else {
				size, _ := strconv.Atoi(line[0])
				head.size += size
			}
		}
	}
	totalMem := 70000000
	updateSize := 30000000
	return findWhichDelete(root, root.totalSize(), updateSize-(totalMem-root.totalSize()))
}
