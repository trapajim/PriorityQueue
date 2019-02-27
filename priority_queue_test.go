package PriorityQueue

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPriorityQueuePush(t *testing.T) {
	item := struct{}{}
	pq := New(true)
	_, err := pq.Push(item, 1)
	if err != nil {
		t.Error(err)
	}

	quickAssert(pq.Len(), 1, t)
}

func TestPriorityQueuePull(t *testing.T) {
	item := struct{}{}
	pq := New(true)
	pq.Push(item, 1)
	pulled, err := pq.Pull()
	if err != nil {
		t.Error(err)
	}
	quickAssert(pulled, item, t)
}

func TestPriorityQueueLen(t *testing.T) {
	pq := New(true)
	quickAssert(pq.Len(), 0, t)
	pq.Push(struct{}{}, 1)
	quickAssert(pq.Len(), 1, t)
}

func quickAssert(got, want interface{}, t *testing.T) {
	if want != got {
		_, file, line, _ := runtime.Caller(1)
		s := fmt.Sprintf("\t%s:%d: %s\n", filepath.Base(file), line, fmt.Sprintf("Expected %v to be %v", got, want))
		fmt.Println(s)
		t.Fail()
	}
}
