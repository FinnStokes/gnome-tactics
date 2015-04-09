package rectgrid

import "testing"

import "github.com/FinnStokes/gnome-tactics/model/pathfinding"

type testGrid struct {
	w, h int
}

func (g testGrid) Passable(a, b Rect) bool {
	return (a.X >= 0 && a.X < g.w && a.Y >= 0 && a.Y < g.h &&
		b.X >= 0 && b.X < g.w && b.Y >= 0 && b.Y < g.h)
}

func (g testGrid) Cost(a, b Rect) int {
	// return 1 + g.w/2 + g.h/2 - abs(b.X - g.w/2) - abs(b.Y - g.h/2)
	dh := abs(a.X+a.Y-g.w/2-g.h/2) + abs(a.X-a.Y) - abs(b.X+b.Y-g.w/2-g.h/2) - abs(b.X-b.Y)
	if dh > 0 {
		return 1 + 2*dh
	} else {
		return 1
	}
}

func contains(slice []pathfinding.Node, value Rect) bool {
	for _, v := range slice {
		if v.(Rect) == value {
			return true
		}
	}
	return false
}

func TestNeighbors(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		r          Rect
		neighbours []Rect
	}{
		{Rect{0, 0, g}, []Rect{{1, 0, g}, {0, 1, g}}},
		{Rect{1, 0, g}, []Rect{{2, 0, g}, {1, 1, g}, {0, 0, g}}},
		{Rect{0, 1, g}, []Rect{{1, 1, g}, {0, 2, g}, {0, 0, g}}},
		{Rect{1, 1, g}, []Rect{{2, 1, g}, {1, 2, g}, {1, 0, g}, {0, 1, g}}},
		{Rect{2, 1, g}, []Rect{{2, 2, g}, {1, 1, g}, {2, 0, g}}},
	}
	for i, test := range testCases {
		neighbours := test.r.Neighbours()
		if len(neighbours) > len(test.neighbours) {
			t.Errorf("Found %d unexpected neighbours in test %d", len(neighbours)-len(test.neighbours), i)
			continue
		}
		for _, n := range test.neighbours {
			if !contains(neighbours, n) {
				t.Errorf("Missing neighbour %v in test %d (%v)", n, i, neighbours)
			}
		}
	}
}

func TestCost(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		r1, r2 Rect
		cost   int
	}{
		{Rect{0, 0, g}, Rect{1, 0, g}, 1},
		{Rect{1, 0, g}, Rect{1, 1, g}, 5},
		{Rect{0, 1, g}, Rect{0, 0, g}, 1},
		{Rect{0, 1, g}, Rect{1, 1, g}, 5},
		{Rect{1, 1, g}, Rect{2, 1, g}, 1},
		{Rect{2, 1, g}, Rect{2, 2, g}, 1},
	}

	for _, test := range testCases {
		cost := test.r1.Cost(test.r2)
		if cost != test.cost {
			t.Errorf("Incorrect cost for %v -> %v (%d != %d)", test.r1, test.r2, cost, test.cost)
		}
	}
}

func TestHeuristic(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		r1, r2    Rect
		heuristic int
	}{
		{Rect{0, 0, g}, Rect{1, 0, g}, 1},
		{Rect{1, 0, g}, Rect{2, 1, g}, 2},
		{Rect{0, 1, g}, Rect{2, 2, g}, 3},
		{Rect{0, 1, g}, Rect{1, 1, g}, 1},
		{Rect{1, 1, g}, Rect{0, 2, g}, 2},
		{Rect{0, 0, g}, Rect{2, 2, g}, 4},
	}

	for _, test := range testCases {
		heuristic := test.r1.Heuristic(test.r2)
		if heuristic != test.heuristic {
			t.Errorf("Incorrect heuristic for %v -> %v (%d != %d)", test.r1, test.r2, heuristic, test.heuristic)
		}
	}
}

func TestDistance(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		r1, r2   Rect
		distance int
	}{
		{Rect{0, 0, g}, Rect{1, 0, g}, 1},
		{Rect{1, 0, g}, Rect{2, 1, g}, 2},
		{Rect{0, 1, g}, Rect{2, 2, g}, 3},
		{Rect{0, 1, g}, Rect{1, 1, g}, 1},
		{Rect{1, 1, g}, Rect{0, 2, g}, 2},
		{Rect{0, 0, g}, Rect{2, 2, g}, 4},
	}

	for _, test := range testCases {
		distance := test.r1.Distance(test.r2)
		if distance != test.distance {
			t.Errorf("Incorrect distance for %v -> %v (%d != %d)", test.r1, test.r2, distance, test.distance)
		}
	}
}

func TestLine(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		r1, r2 Rect
		line   []Rect
	}{
		{Rect{0, 0, g}, Rect{3, 0, g}, []Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {3, 0, g}}},
		{Rect{1, 0, g}, Rect{4, 3, g}, []Rect{{1, 0, g}, {2, 1, g}, {3, 2, g}, {4, 3, g}}},
		{Rect{0, 1, g}, Rect{2, 2, g}, []Rect{{0, 1, g}, {1, 2, g}, {2, 2, g}}},
		{Rect{0, 2, g}, Rect{0, 0, g}, []Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}}},
		{Rect{1, 1, g}, Rect{0, 2, g}, []Rect{{1, 1, g}, {0, 2, g}}},
		{Rect{0, 0, g}, Rect{2, 2, g}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}},
		{Rect{0, 2, g}, Rect{29, 3, g}, []Rect{{0, 2, g}, {1, 2, g}, {2, 2, g}, {3, 2, g}, {4, 2, g}, {5, 2, g}, {6, 2, g}, {7, 2, g}, {8, 2, g}, {9, 2, g}, {10, 2, g}, {11, 2, g}, {12, 2, g}, {13, 2, g}, {14, 2, g}, {15, 3, g}, {16, 3, g}, {17, 3, g}, {18, 3, g}, {19, 3, g}, {20, 3, g}, {21, 3, g}, {22, 3, g}, {23, 3, g}, {24, 3, g}, {25, 3, g}, {26, 3, g}, {27, 3, g}, {28, 3, g}, {29, 3, g}}},
	}

	for i, test := range testCases {
		line := test.r1.Line(test.r2)
		if len(line) != len(test.line) {
			t.Errorf("Line %d wrong length (was %d, should be %d)", i, len(line), len(test.line))
			continue
		}
		for j, n := range test.line {
			if line[j] != n {
				t.Errorf("Line %d does not match expected line (%v instead of %v)", i, line, test.line)
				break
			}
		}
	}
}

func TestPassableLine(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		r1, r2 Rect
		line   []Rect
	}{
		{Rect{0, 0, g}, Rect{3, 0, g}, []Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}},
		{Rect{1, 0, g}, Rect{4, 3, g}, []Rect{{1, 0, g}, {2, 1, g}}},
		{Rect{0, 1, g}, Rect{2, 2, g}, []Rect{{0, 1, g}, {1, 2, g}, {2, 2, g}}},
		{Rect{0, 2, g}, Rect{0, 0, g}, []Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}}},
		{Rect{1, 1, g}, Rect{0, 2, g}, []Rect{{1, 1, g}, {0, 2, g}}},
		{Rect{0, 0, g}, Rect{2, 2, g}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}},
		{Rect{0, 2, g}, Rect{29, 3, g}, []Rect{{0, 2, g}, {1, 2, g}, {2, 2, g}}},
	}

	for i, test := range testCases {
		line := test.r1.PassableLine(test.r2)
		if len(line) != len(test.line) {
			t.Errorf("Line %d wrong length (was %d, should be %d)", i, len(line), len(test.line))
			continue
		}
		for j, n := range test.line {
			if line[j] != n {
				t.Errorf("Line %d does not match expected line (%v instead of %v)", i, line, test.line)
				break
			}
		}
	}
}

func TestSet(t *testing.T) {
	g := testGrid{3, 3}
	testCases := [][]Rect{
		[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}},
		[]Rect{{1, 1, g}},
		[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}},
		[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}},
		[]Rect{{0, 0, g}, {2, 0, g}},
		[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}},
	}

	for i, test := range testCases {
		set := NewSet(test)
		for _, r := range test {
			if !set[r] {
				t.Errorf("Rect %v missing from set %d (%v)", r, i, test)
			}
		}
		for r, ok := range set {
			if ok {
				found := false
				for _, r2 := range test {
					if r == r2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Rect %v in set %d (%v)", r, i, test)
				}
			}
		}
	}
}

func TestIntersect(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct{ A, B, AnB []Rect }{
		{[]Rect{}, []Rect{}, []Rect{}},
		{[]Rect{{1, 1, g}}, []Rect{}, []Rect{}},
		{[]Rect{}, []Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}}, []Rect{}},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Rect{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, []Rect{}},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}},
		{[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}},
		{[]Rect{{0, 0, g}, {2, 0, g}}, []Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, []Rect{{0, 0, g}, {2, 0, g}}},
		{[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Rect{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, []Rect{{1, 0, g}, {1, 2, g}}},
	}

	for i, test := range testCases {
		set := NewSet(test.A).Intersect(NewSet(test.B))
		for _, r := range test.AnB {
			if !set[r] {
				t.Errorf("Rect %v missing from set %d (%v)", r, i, test.AnB)
			}
		}
		for r, ok := range set {
			if ok {
				found := false
				for _, r2 := range test.AnB {
					if r == r2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Rect %v in set %d (%v)", r, i, test.AnB)
				}
			}
		}
	}
}

func TestUnion(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct{ A, B, AuB []Rect }{
		{[]Rect{}, []Rect{}, []Rect{}},
		{[]Rect{{1, 1, g}}, []Rect{}, []Rect{{1, 1, g}}},
		{[]Rect{}, []Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}}, []Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}}},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Rect{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, []Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}},
		{[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}},
		{[]Rect{{0, 0, g}, {2, 0, g}}, []Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, []Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}},
		{[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Rect{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, []Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}, {0, 0, g}, {0, 2, g}, {2, 2, g}}},
	}

	for i, test := range testCases {
		set := NewSet(test.A).Union(NewSet(test.B))
		for _, r := range test.AuB {
			if !set[r] {
				t.Errorf("Rect %v missing from set %d (%v)", r, i, test.AuB)
			}
		}
		for r, ok := range set {
			if ok {
				found := false
				for _, r2 := range test.AuB {
					if r == r2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Rect %v in set %d (%v)", r, i, test.AuB)
				}
			}
		}
	}
}

func TestSubtract(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct{ A, B, AmB []Rect }{
		{[]Rect{}, []Rect{}, []Rect{}},
		{[]Rect{{1, 1, g}}, []Rect{}, []Rect{{1, 1, g}}},
		{[]Rect{}, []Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}}, []Rect{}},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Rect{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, []Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {2, 1, g}, {1, 2, g}, {2, 2, g}}},
		{[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{}},
		{[]Rect{{0, 0, g}, {2, 0, g}}, []Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, []Rect{}},
		{[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Rect{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, []Rect{{0, 1, g}, {2, 1, g}}},
	}

	for i, test := range testCases {
		set := NewSet(test.A).Subtract(NewSet(test.B))
		for _, r := range test.AmB {
			if !set[r] {
				t.Errorf("Rect %v missing from set %d (%v)", r, i, test.AmB)
			}
		}
		for r, ok := range set {
			if ok {
				found := false
				for _, r2 := range test.AmB {
					if r == r2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Rect %v in set %d (%v)", r, i, test.AmB)
				}
			}
		}
	}
}

func TestIn(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		A, B   []Rect
		result bool
	}{
		{[]Rect{}, []Rect{}, true},
		{[]Rect{{1, 1, g}}, []Rect{}, false},
		{[]Rect{}, []Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}}, true},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Rect{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, false},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}, false},
		{[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, true},
		{[]Rect{{0, 0, g}, {2, 0, g}}, []Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, true},
		{[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Rect{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, false},
	}

	for _, test := range testCases {
		result := NewSet(test.A).In(NewSet(test.B))
		if result != test.result {
			t.Errorf("%v.In(%v) returned %v", test.A, test.B, result)
		}
	}
}

func TestContains(t *testing.T) {
	g := testGrid{3, 3}
	testCases := []struct {
		A, B   []Rect
		result bool
	}{
		{[]Rect{}, []Rect{}, true},
		{[]Rect{{1, 1, g}}, []Rect{}, true},
		{[]Rect{}, []Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}}, false},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Rect{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, false},
		{[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}, true},
		{[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}}, true},
		{[]Rect{{0, 0, g}, {2, 0, g}}, []Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, false},
		{[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Rect{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, false},
	}

	for _, test := range testCases {
		result := NewSet(test.A).Contains(NewSet(test.B))
		if result != test.result {
			t.Errorf("%v.Contains(%v) returned %v", test.A, test.B, result)
		}
	}
}

func TestEquals(t *testing.T) {
	g := testGrid{3, 3}
	testCases := [][]Rect{
		[]Rect{},
		[]Rect{{1, 1, g}},
		[]Rect{{2, 2, g}, {1, 0, g}, {2, 1, g}},
		[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}},
		[]Rect{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}},
		[]Rect{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}},
		[]Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}},
		[]Rect{{0, 0, g}, {1, 1, g}, {2, 2, g}},
		[]Rect{{0, 0, g}, {2, 0, g}},
		[]Rect{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}},
		[]Rect{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}},
		[]Rect{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}},
	}

	for i, A := range testCases {
		for j, B := range testCases {
			result := NewSet(A).Equals(NewSet(B))
			if result != (i == j) {
				t.Errorf("%v.Equals(%v) returned %v", A, B, result)
			}
		}
	}
}

func TestRange(t *testing.T) {
	g := testGrid{9, 9}
	testCases := []struct {
		centre Rect
		N      int
		r      []Rect
	}{
		{Rect{0, 0, g}, 0, []Rect{{0, 0, g}}},
		{Rect{2, 1, g}, 1, []Rect{{2, 1, g}, {2, 2, g}, {1, 1, g}, {3, 1, g}, {2, 0, g}}},
		{Rect{3, 3, g}, 2, []Rect{{3, 3, g}, {3, 4, g}, {4, 3, g}, {3, 2, g}, {2, 3, g}, {4, 4, g}, {4, 2, g}, {2, 2, g}, {2, 4, g}, {3, 5, g}, {5, 3, g}, {3, 1, g}, {1, 3, g}}},
		{Rect{5, 4, g}, 3, []Rect{{5, 4, g}, {4, 4, g}, {5, 5, g}, {6, 4, g}, {5, 3, g}, {3, 4, g}, {4, 5, g}, {5, 6, g}, {6, 5, g}, {7, 4, g}, {6, 3, g}, {5, 2, g}, {4, 3, g}, {2, 4, g}, {3, 5, g}, {4, 6, g}, {5, 7, g}, {6, 6, g}, {7, 5, g}, {8, 4, g}, {7, 3, g}, {6, 2, g}, {5, 1, g}, {4, 2, g}, {3, 3, g}}},
		{Rect{4, 4, g}, 4, []Rect{{4, 4, g}, {4, 5, g}, {5, 4, g}, {4, 3, g}, {3, 4, g}, {5, 5, g}, {5, 3, g}, {3, 3, g}, {3, 5, g}, {4, 6, g}, {6, 4, g}, {4, 2, g}, {2, 4, g}, {1, 4, g}, {2, 3, g}, {3, 2, g}, {4, 1, g}, {5, 2, g}, {6, 3, g}, {7, 4, g}, {6, 5, g}, {5, 6, g}, {4, 7, g}, {3, 6, g}, {2, 5, g}, {0, 4, g}, {1, 3, g}, {2, 2, g}, {3, 1, g}, {4, 0, g}, {5, 1, g}, {6, 2, g}, {7, 3, g}, {8, 4, g}, {7, 5, g}, {6, 6, g}, {5, 7, g}, {4, 8, g}, {3, 7, g}, {2, 6, g}, {1, 5, g}}},
	}

	for _, test := range testCases {
		r := test.centre.Range(test.N)
		if !r.Equals(NewSet(test.r)) {
			t.Errorf("Incorrect range for %d around %v", test.N, test.centre)
		}
	}
}
