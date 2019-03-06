# PriorityQueue [![Go Report Card](https://goreportcard.com/badge/github.com/trapajim/PriorityQueue)](https://goreportcard.com/report/github.com/trapajim/PriorityQueue) [![GoDoc](https://godoc.org/github.com/trapajim/PriorityQueue?status.svg)](https://godoc.org/github.com/trapajim/PriorityQueue)

Thread safe priority queue implemented on top of container/heap

## Usage

```go
// create a new queue
// pass true for mutex autolock
q := PriorityQueue.New(true)

// push a new element
q.Push("a", 1)
q.Push("b", 2)

// pull the element with the highest priority
first, _ := q.Pull()
// output: first is b
fmt.Println("first is ", first)

second, _ := q.Pull()
// output: second is a
fmt.Println("second is ", second)
```
