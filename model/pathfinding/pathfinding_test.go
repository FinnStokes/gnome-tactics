package pathfinding

import "testing"

type testNode int

func (t testNode) Neighbours() []Node {
	return []Node{(t-1) & 7, (t+1) & 7}
}

func (t testNode) Cost(node Node) int {
	if (t == 0 && node.(testNode) == 7) ||  (t == 7 && node.(testNode) == 0) {
		return 3
	} else {
		return 2
	}
}

func (t testNode) Heuristic(node Node) int {
	return 0
}

func TestAll(t *testing.T) {
	m := All(testNode(0))

	paths := [][]testNode{
		[]testNode{0},
		[]testNode{0,1},
		[]testNode{0,1,2},
		[]testNode{0,1,2,3},
		[]testNode{0,1,2,3,4},
		[]testNode{0,7,6,5},
		[]testNode{0,7,6},
		[]testNode{0,7},
	}

	for i, path := range paths {
		found, ok := m[testNode(i)]
		if !ok {
			t.Errorf("No path found from 0 to %d", i)
			continue
		}
		if len(found) != len(path) {
			t.Errorf("Incorrect path from 0 to %d", i)
			t.Errorf("%v", found)
			continue
		}
		for j, node := range path {
			if node != found[j] {
				t.Errorf("Incorrect path from 0 to %d", i)
				t.Errorf("%v", found)
				break
			}
		}
	}

	for end := range m {
		if int(end.(testNode)) >= len(paths) || int(end.(testNode)) < 0 {
			t.Errorf("Path to disconnected node %d", end)
		}
	}

	m = All(testNode(3))

	paths = [][]testNode{
		[]testNode{3,2,1,0},
		[]testNode{3,2,1},
		[]testNode{3,2},
		[]testNode{3},
		[]testNode{3,4},
		[]testNode{3,4,5},
		[]testNode{3,4,5,6},
		[]testNode{3,4,5,6,7},
	}

	for i, path := range paths {
		found, ok := m[testNode(i)]
		if !ok {
			t.Errorf("No path found from 3 to %d", i)
			continue
		}
		if len(found) != len(path) {
			t.Errorf("Incorrect path from 3 to %d", i)
			t.Errorf("%v", found)
			continue
		}
		for j, node := range path {
			if node != found[j] {
				t.Errorf("Incorrect path from 3 to %d", i)
				t.Errorf("%v", found)
				break
			}
		}
	}

	for end := range m {
		if int(end.(testNode)) >= len(paths) || int(end.(testNode)) < 0 {
			t.Errorf("Path found to disconnected node %d", end)
		}
	}
}

func TestSingle(t *testing.T) {
	paths := [][]testNode{
		[]testNode{0},
		[]testNode{0,1},
		[]testNode{0,1,2},
		[]testNode{0,1,2,3},
		[]testNode{0,1,2,3,4},
		[]testNode{0,7,6,5},
		[]testNode{0,7,6},
		[]testNode{0,7},
	}

	for i, path := range paths {
		found, ok := Single(testNode(0), testNode(i))
		if !ok {
			t.Errorf("No path found from 0 to %d", i)
			continue
		}
		if len(found) != len(path) {
			t.Errorf("Incorrect path from 0 to %d", i)
			t.Errorf("%v", found)
			continue
		}
		for j, node := range path {
			if node != found[j] {
				t.Errorf("Incorrect path from 0 to %d", i)
				t.Errorf("%v", found)
				break
			}
		}
	}

	found, ok := Single(testNode(3), testNode(8))
	if ok || found != nil {
			t.Errorf("Path found to disconnected node %d", 8)
	}
	
	paths = [][]testNode{
		[]testNode{3,2,1,0},
		[]testNode{3,2,1},
		[]testNode{3,2},
		[]testNode{3},
		[]testNode{3,4},
		[]testNode{3,4,5},
		[]testNode{3,4,5,6},
		[]testNode{3,4,5,6,7},
	}

	for i, path := range paths {
		found, ok := Single(testNode(3), testNode(i))
		if !ok {
			t.Errorf("No path found from 3 to %d", i)
			continue
		}
		if len(found) != len(path) {
			t.Errorf("Incorrect path from 3 to %d", i)
			t.Errorf("%v", found)
			continue
		}
		for j, node := range path {
			if node != found[j] {
				t.Errorf("Incorrect path from 3 to %d", i)
				t.Errorf("%v", found)
				break
			}
		}
	}

	found, ok = Single(testNode(3), testNode(8))
	if ok || found != nil {
			t.Errorf("Path found to disconnected node %d", 8)
	}
}
