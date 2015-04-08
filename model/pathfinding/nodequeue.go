package pathfinding

import "container/heap"

type nodeQueueElement struct {
	Priority int
	Cost int
	Parent Node
	Self Node
}

type NodeQueue []nodeQueueElement

func NewNodeQueue() NodeQueue {
	return make(NodeQueue, 0, 1000)
}

func (q NodeQueue) Empty() bool {
	return len(q) == 0
}

func (q NodeQueue) Len() int {
	return len(q)
}

func (q NodeQueue) Less(i, j int) bool {
	return q[i].Priority < q[j].Priority
}

func (q NodeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *NodeQueue) Push(n interface{}) {
	*q = append(*q, n.(nodeQueueElement))
}

func (q *NodeQueue) Pop() interface{} {
	n := len(*q)
	last := (*q)[n-1]
	*q = (*q)[0:n-1]
	return last
}

func (q *NodeQueue) Get() (Node, Node, int) {
	n := heap.Pop(q).(nodeQueueElement)
	return n.Self, n.Parent, n.Cost
}

func (q *NodeQueue) Put(self, parent Node, cost, priority int) {
	heap.Push(q, nodeQueueElement{priority, cost, parent, self})
}
