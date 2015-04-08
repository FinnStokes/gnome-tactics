package pathfinding

import "testing"

func TestEmpty(t *testing.T) {
	q := NewNodeQueue()
	if !q.Empty() {
		t.Errorf("NodeQueue.Empty() returned false when it should have returned true")
	}
	q.Put(nil, nil, 0, 0)
	if q.Empty() {
		t.Errorf("NodeQueue.Empty() returned true when it should have returned false")
	}
	q.Get()
	if !q.Empty() {
		t.Errorf("NodeQueue.Empty() returned false when it should have returned true")
	}
}

func TestLen(t *testing.T) {
	q := NewNodeQueue()

	len := q.Len()
	if len != 0 {
		t.Errorf("NodeQueue.Len() is %d (should be 0)", len)
	}
	
	for i := 1; i < 500; i += 1 {
		q.Put(testNode(2), testNode(1), 1, 1)

		len = q.Len()
		if len != i {
			t.Errorf("NodeQueue.Len() is %d (should be %d)", len, i)
		}
	}
}

func TestGet(t *testing.T) {
	q := NewNodeQueue()
	
	q.Put(testNode(2), testNode(1), 1, 1)
	q.Put(testNode(3), testNode(2), 2, 4)
	q.Put(testNode(2), testNode(4), 9, 3)
	q.Put(testNode(0), testNode(3), 1, 8)
	q.Put(testNode(1), testNode(3), 6, 2)
	q.Put(testNode(2), testNode(4), 2, 1)
	
	nodes := []testNode{2, 2, 1, 2, 3, 0}
	parents := []testNode{1, 4, 3, 4, 2, 3}
	costs := []int{1, 2, 6, 9, 2, 1}

	for i := range nodes {
		node, parent, cost := q.Get()
		if node != nodes[i] {
			t.Errorf("NodeQueue.Get() returned incorrect node (%d instead of %d on iteration %d)", node, nodes[i], i)
		}
		if parent != parents[i] {
			t.Errorf("NodeQueue.Get() returned incorrect parent (%d instead of %d on iteration %d)", parent, parents[i], i)
		}
		if cost != costs[i] {
			t.Errorf("NodeQueue.Get() returned incorrect cost (%d instead of %d on iteration %d)", cost, costs[i], i)
		}
	}

	if !q.Empty() {
                t.Errorf("NodeQueue.Empty() returned false when it should have returned true")
	}

	q.Put(testNode(4), testNode(2), 6, 3)
	q.Put(testNode(2), testNode(3), 5, 2)
	q.Put(testNode(3), testNode(1), 2, 4)
	q.Put(testNode(0), testNode(0), 9, 8)
	q.Put(testNode(1), testNode(4), 1, 7)
	
	nodes = []testNode{2, 4, 3}
	parents = []testNode{3, 2, 1}
	costs = []int{5, 6, 2}

	for i := range nodes {
		node, parent, cost := q.Get()
		if node != nodes[i] {
			t.Errorf("NodeQueue.Get() returned incorrect node (%d instead of %d on iteration %d)", node, nodes[i], i)
		}
		if parent != parents[i] {
			t.Errorf("NodeQueue.Get() returned incorrect parent (%d instead of %d on iteration %d)", parent, parents[i], i)
		}
		if cost != costs[i] {
			t.Errorf("NodeQueue.Get() returned incorrect cost (%d instead of %d on iteration %d)", cost, costs[i], i)
		}
	}

	q.Put(testNode(4), testNode(1), 3, 9)
	q.Put(testNode(2), testNode(4), 1, 5)
	
	nodes = []testNode{2, 1, 0, 4}
	parents = []testNode{4, 4, 0, 1}
	costs = []int{1, 1, 9, 3}

	for i := range nodes {
		node, parent, cost := q.Get()
		if node != nodes[i] {
			t.Errorf("NodeQueue.Get() returned incorrect node (%d instead of %d on iteration %d)", node, nodes[i], i)
		}
		if parent != parents[i] {
			t.Errorf("NodeQueue.Get() returned incorrect parent (%d instead of %d on iteration %d)", parent, parents[i], i)
		}
		if cost != costs[i] {
			t.Errorf("NodeQueue.Get() returned incorrect cost (%d instead of %d on iteration %d)", cost, costs[i], i)
		}
	}


	if !q.Empty() {
                t.Errorf("NodeQueue.Empty() returned false when it should have returned true")
	}
}
