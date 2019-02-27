package PriorityQueue

type Item struct {
	priority uint8
	index    int
	value    interface{}
}

type Queue []*Item

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].priority > q[j].priority
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

func (q *Queue) update(item *Item, value string, priority uint8) {
	item.value = value
	item.priority = priority
}
