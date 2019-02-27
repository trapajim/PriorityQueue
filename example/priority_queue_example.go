package main

import (
	"fmt"

	"github.com/trapajim/PriorityQueue"
)

func main() {
	l := PriorityQueue.New(false)
	l.Push("a", 1)
	l.Push("b", 2)
	first, _ := l.Pull()
	// output: first is b
	fmt.Println("first is ", first)
	second, _ := l.Pull()
	// output: second is a
	fmt.Println("second is ", second)
}
