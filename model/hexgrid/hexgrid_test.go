package hexgrid

import "testing"

import "github.com/FinnStokes/gnome-tactics/model/pathfinding"

type testGrid struct {
	r int
}

func (g testGrid) Passable(a, b Hex) bool {
	return (a.Q >= -g.r && a.Q <= g.r && a.R >= -g.r && a.R <= g.r && a.Q+a.R >= -g.r && a.Q+a.R <= g.r &&
		b.Q >= -g.r && b.Q <= g.r && b.R >= -g.r && b.R <= g.r && b.Q+b.R >= -g.r && b.Q+b.R <= g.r)
}

func (g testGrid) Cost(a, b Hex) int {
	dh := a.Distance(Hex{0, 0, g}) - b.Distance(Hex{0, 0, g})
	if dh > 0 {
		return 1 + 4*dh
	} else {
		return 1
	}
}

func contains(slice []pathfinding.Node, value Hex) bool {
	for _, v := range slice {
		if v.(Hex) == value {
			return true
		}
	}
	return false
}

func TestNeighbors(t *testing.T) {
	g := testGrid{2}
	testCases := []struct {
		h          Hex
		neighbours []Hex
	}{
		{Hex{0, 0, g}, []Hex{{0, -1, g}, {1, -1, g}, {1, 0, g}, {0, 1, g}, {-1, 1, g}, {-1, 0, g}}},
		{Hex{1, 0, g}, []Hex{{1, -1, g}, {2, -1, g}, {2, 0, g}, {1, 1, g}, {0, 1, g}, {0, 0, g}}},
		{Hex{0, 1, g}, []Hex{{0, 0, g}, {1, 0, g}, {1, 1, g}, {0, 2, g}, {-1, 2, g}, {-1, 1, g}}},
		{Hex{2, -2, g}, []Hex{{1, -2, g}, {1, -1, g}, {2, -1, g}}},
		{Hex{1, 1, g}, []Hex{{0, 2, g}, {0, 1, g}, {1, 0, g}, {2, 0, g}}},
		{Hex{0, -2, g}, []Hex{{1, -2, g}, {0, -1, g}, {-1, -1, g}}},
		{Hex{-2, 1, g}, []Hex{{-2, 0, g}, {-1, 0, g}, {-1, 1, g}, {-2, 2, g}}},
		{Hex{-1, 1, g}, []Hex{{-2, 1, g}, {-1, 0, g}, {0, 0, g}, {0, 1, g}, {-1, 2, g}, {-2, 2, g}}},
	}
	for i, test := range testCases {
		neighbours := test.h.Neighbours()
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
	g := testGrid{2}
	testCases := []struct {
		h1, h2 Hex
	}{
		{Hex{0, 0, g}, Hex{1, 0, g}},
		{Hex{1, 0, g}, Hex{1, 1, g}},
		{Hex{0, 1, g}, Hex{0, 0, g}},
		{Hex{0, 1, g}, Hex{1, 1, g}},
		{Hex{1, 1, g}, Hex{2, 1, g}},
		{Hex{0, 0, g}, Hex{1, -1, g}},
		{Hex{0, -1, g}, Hex{1, -2, g}},
		{Hex{2, 0, g}, Hex{1, 1, g}},
	}

	for _, test := range testCases {
		cost := test.h1.Cost(test.h2)
		if cost != g.Cost(test.h1, test.h2) {
			t.Errorf("Incorrect cost for %v -> %v (%d != %d)", test.h1, test.h2, cost, g.Cost(test.h1, test.h2))
		}
	}
}

func TestHeuristic(t *testing.T) {
	g := testGrid{2}
	testCases := []struct {
		h1, h2    Hex
		heuristic int
	}{
		{Hex{0, 0, g}, Hex{1, 0, g}, 1},
		{Hex{1, 0, g}, Hex{2, -1, g}, 1},
		{Hex{0, 1, g}, Hex{2, -2, g}, 3},
		{Hex{-2, 2, g}, Hex{2, -2, g}, 4},
		{Hex{0, -2, g}, Hex{2, 0, g}, 4},
		{Hex{2, 0, g}, Hex{-1, -1, g}, 4},
		{Hex{1, 1, g}, Hex{1, -1, g}, 2},
		{Hex{0, 1, g}, Hex{1, 1, g}, 1},
	}

	for _, test := range testCases {
		heuristic := test.h1.Heuristic(test.h2)
		if heuristic != test.heuristic {
			t.Errorf("Incorrect heuristic for %v -> %v (%d != %d)", test.h1, test.h2, heuristic, test.heuristic)
		}
	}
}

func TestDistance(t *testing.T) {
	g := testGrid{2}
	testCases := []struct {
		h1, h2   Hex
		distance int
	}{
		{Hex{0, 0, g}, Hex{1, 0, g}, 1},
		{Hex{1, 0, g}, Hex{2, -1, g}, 1},
		{Hex{0, 1, g}, Hex{2, -2, g}, 3},
		{Hex{-2, 2, g}, Hex{2, -2, g}, 4},
		{Hex{0, -2, g}, Hex{2, 0, g}, 4},
		{Hex{2, 0, g}, Hex{-1, -1, g}, 4},
		{Hex{1, 1, g}, Hex{1, -1, g}, 2},
		{Hex{0, 1, g}, Hex{1, 1, g}, 1},
	}

	for _, test := range testCases {
		distance := test.h1.Distance(test.h2)
		if distance != test.distance {
			t.Errorf("Incorrect distance for %v -> %v (%d != %d)", test.h1, test.h2, distance, test.distance)
		}
	}
}

func TestLine(t *testing.T) {
	g := testGrid{2}
	testCases := []struct {
		h1, h2 Hex
		line   []Hex
	}{
		{Hex{0, 0, g}, Hex{3, 0, g}, []Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {3, 0, g}}},
		{Hex{-1, 0, g}, Hex{2, -2, g}, []Hex{{-1, 0, g}, {0, -1, g}, {1, -1, g}, {2, -2, g}}},
		{Hex{0, 1, g}, Hex{2, -1, g}, []Hex{{0, 1, g}, {1, 0, g}, {2, -1, g}}},
		{Hex{-1, 2, g}, Hex{2, -2, g}, []Hex{{-1, 2, g}, {0, 1, g}, {1, 0, g}, {1, -1, g}, {2, -2, g}}},
		{Hex{-2, 0, g}, Hex{0, 3, g}, []Hex{{-2, 0, g}, {-2, 1, g}, {-1, 1, g}, {-1, 2, g}, {0, 2, g}, {0, 3, g}}},
		{Hex{1, -1, g}, Hex{-6, 1, g}, []Hex{{1, -1, g}, {0, -1, g}, {-1, 0, g}, {-2, 0, g}, {-3, 0, g}, {-4, 0, g}, {-5, 1, g}, {-6, 1, g}}},
		{Hex{-1, 2, g}, Hex{-2, 14, g}, []Hex{{-1, 2, g}, {-1, 3, g}, {-1, 4, g}, {-1, 5, g}, {-1, 6, g}, {-1, 7, g}, {-2, 8, g}, {-2, 9, g}, {-2, 10, g}, {-2, 11, g}, {-2, 12, g}, {-2, 13, g}, {-2, 14, g}}},
	}

	for i, test := range testCases {
		line := test.h1.Line(test.h2)
		if len(line) != len(test.line) {
			t.Errorf("Line %d does not match expected line (%v instead of %v)", i, line, test.line)
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
	g := testGrid{2}
	testCases := []struct {
		h1, h2 Hex
		line   []Hex
	}{
		{Hex{0, 0, g}, Hex{3, 0, g}, []Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}},
		{Hex{-1, 0, g}, Hex{2, -2, g}, []Hex{{-1, 0, g}, {0, -1, g}, {1, -1, g}, {2, -2, g}}},
		{Hex{0, 1, g}, Hex{2, -1, g}, []Hex{{0, 1, g}, {1, 0, g}, {2, -1, g}}},
		{Hex{-1, 2, g}, Hex{2, -2, g}, []Hex{{-1, 2, g}, {0, 1, g}, {1, 0, g}, {1, -1, g}, {2, -2, g}}},
		{Hex{-2, 0, g}, Hex{0, 3, g}, []Hex{{-2, 0, g}, {-2, 1, g}, {-1, 1, g}, {-1, 2, g}, {0, 2, g}}},
		{Hex{1, -1, g}, Hex{-6, 1, g}, []Hex{{1, -1, g}, {0, -1, g}, {-1, 0, g}, {-2, 0, g}}},
		{Hex{-1, 2, g}, Hex{-2, 14, g}, []Hex{{-1, 2, g}}},
	}

	for i, test := range testCases {
		line := test.h1.PassableLine(test.h2)
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
	g := testGrid{2}
	testCases := [][]Hex{
		[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}},
		[]Hex{{1, 1, g}},
		[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}},
		[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}},
		[]Hex{{0, 0, g}, {2, 0, g}},
		[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}},
	}

	for i, test := range testCases {
		set := NewSet(test)
		for _, h := range test {
			if !set[h] {
				t.Errorf("Hex %v missing from set %d (%v)", h, i, test)
			}
		}
		for h, ok := range set {
			if ok {
				found := false
				for _, h2 := range test {
					if h == h2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Hex %v in set %d (%v)", h, i, test)
				}
			}
		}
	}
}

func TestIntersect(t *testing.T) {
	g := testGrid{2}
	testCases := []struct{ A, B, AnB []Hex }{
		{[]Hex{}, []Hex{}, []Hex{}},
		{[]Hex{{1, 1, g}}, []Hex{}, []Hex{}},
		{[]Hex{}, []Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}}, []Hex{}},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Hex{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, []Hex{}},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}}},
		{[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}},
		{[]Hex{{0, 0, g}, {2, 0, g}}, []Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, []Hex{{0, 0, g}, {2, 0, g}}},
		{[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Hex{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, []Hex{{1, 0, g}, {1, 2, g}}},
	}

	for i, test := range testCases {
		set := NewSet(test.A).Intersect(NewSet(test.B))
		for _, h := range test.AnB {
			if !set[h] {
				t.Errorf("Hex %v missing from set %d (%v)", h, i, test.AnB)
			}
		}
		for h, ok := range set {
			if ok {
				found := false
				for _, h2 := range test.AnB {
					if h == h2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Hex %v in set %d (%v)", h, i, test.AnB)
				}
			}
		}
	}
}

func TestUnion(t *testing.T) {
	g := testGrid{2}
	testCases := []struct{ A, B, AuB []Hex }{
		{[]Hex{}, []Hex{}, []Hex{}},
		{[]Hex{{1, 1, g}}, []Hex{}, []Hex{{1, 1, g}}},
		{[]Hex{}, []Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}}, []Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}}},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Hex{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, []Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}},
		{[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}},
		{[]Hex{{0, 0, g}, {2, 0, g}}, []Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, []Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}},
		{[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Hex{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, []Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}, {0, 0, g}, {0, 2, g}, {2, 2, g}}},
	}

	for i, test := range testCases {
		set := NewSet(test.A).Union(NewSet(test.B))
		for _, h := range test.AuB {
			if !set[h] {
				t.Errorf("Hex %v missing from set %d (%v)", h, i, test.AuB)
			}
		}
		for h, ok := range set {
			if ok {
				found := false
				for _, h2 := range test.AuB {
					if h == h2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Hex %v in set %d (%v)", h, i, test.AuB)
				}
			}
		}
	}
}

func TestSubtract(t *testing.T) {
	g := testGrid{2}
	testCases := []struct{ A, B, AmB []Hex }{
		{[]Hex{}, []Hex{}, []Hex{}},
		{[]Hex{{1, 1, g}}, []Hex{}, []Hex{{1, 1, g}}},
		{[]Hex{}, []Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}}, []Hex{}},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Hex{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, []Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {2, 1, g}, {1, 2, g}, {2, 2, g}}},
		{[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{}},
		{[]Hex{{0, 0, g}, {2, 0, g}}, []Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, []Hex{}},
		{[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Hex{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, []Hex{{0, 1, g}, {2, 1, g}}},
	}

	for i, test := range testCases {
		set := NewSet(test.A).Subtract(NewSet(test.B))
		for _, h := range test.AmB {
			if !set[h] {
				t.Errorf("Hex %v missing from set %d (%v)", h, i, test.AmB)
			}
		}
		for h, ok := range set {
			if ok {
				found := false
				for _, h2 := range test.AmB {
					if h == h2 {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Unexpected Hex %v in set %d (%v)", h, i, test.AmB)
				}
			}
		}
	}
}

func TestIn(t *testing.T) {
	g := testGrid{2}
	testCases := []struct {
		A, B   []Hex
		result bool
	}{
		{[]Hex{}, []Hex{}, true},
		{[]Hex{{1, 1, g}}, []Hex{}, false},
		{[]Hex{}, []Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}}, true},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Hex{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, false},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}}, false},
		{[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, true},
		{[]Hex{{0, 0, g}, {2, 0, g}}, []Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, true},
		{[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Hex{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, false},
	}

	for _, test := range testCases {
		result := NewSet(test.A).In(NewSet(test.B))
		if result != test.result {
			t.Errorf("%v.In(%v) returned %v", test.A, test.B, result)
		}
	}
}

func TestContains(t *testing.T) {
	g := testGrid{2}
	testCases := []struct {
		A, B   []Hex
		result bool
	}{
		{[]Hex{}, []Hex{}, true},
		{[]Hex{{1, 1, g}}, []Hex{}, true},
		{[]Hex{}, []Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}}, false},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}}, []Hex{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}}, false},
		{[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}}, []Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}}, true},
		{[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, []Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}}, true},
		{[]Hex{{0, 0, g}, {2, 0, g}}, []Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}}, false},
		{[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}}, []Hex{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}}, false},
	}

	for _, test := range testCases {
		result := NewSet(test.A).Contains(NewSet(test.B))
		if result != test.result {
			t.Errorf("%v.Contains(%v) returned %v", test.A, test.B, result)
		}
	}
}

func TestEquals(t *testing.T) {
	g := testGrid{2}
	testCases := [][]Hex{
		[]Hex{},
		[]Hex{{1, 1, g}},
		[]Hex{{2, 2, g}, {1, 0, g}, {2, 1, g}},
		[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}},
		[]Hex{{1, 1, g}, {1, 2, g}, {2, 1, g}, {0, 2, g}},
		[]Hex{{0, 0, g}, {1, 0, g}, {2, 0, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {0, 2, g}, {1, 2, g}, {2, 2, g}},
		[]Hex{{0, 2, g}, {0, 1, g}, {1, 1, g}},
		[]Hex{{0, 0, g}, {1, 1, g}, {2, 2, g}},
		[]Hex{{0, 0, g}, {2, 0, g}},
		[]Hex{{1, 0, g}, {2, 0, g}, {2, 1, g}, {0, 0, g}, {1, 1, g}},
		[]Hex{{0, 1, g}, {1, 0, g}, {2, 1, g}, {1, 2, g}},
		[]Hex{{0, 0, g}, {1, 2, g}, {0, 2, g}, {1, 0, g}, {2, 2, g}},
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
	g := testGrid{12}
	testCases := []struct {
		centre Hex
		N      int
		h      []Hex
	}{
		{Hex{0, 0, g}, 0, []Hex{{0, 0, g}}},
		{Hex{2, 1, g}, 1, []Hex{{2, 1, g}, {2, 2, g}, {1, 2, g}, {1, 1, g}, {2, 0, g}, {3, 0, g}, {3, 1, g}}},
		{Hex{3, 3, g}, 2, []Hex{{3, 3, g}, {3, 4, g}, {2, 4, g}, {2, 3, g}, {3, 2, g}, {4, 2, g}, {4, 3, g}, {4, 4, g}, {3, 5, g}, {2, 5, g}, {1, 5, g}, {1, 4, g}, {1, 3, g}, {2, 2, g}, {3, 1, g}, {4, 1, g}, {5, 1, g}, {5, 2, g}, {5, 3, g}}},
		{Hex{5, 4, g}, 3, []Hex{{5, 4, g}, {5, 5, g}, {4, 5, g}, {4, 4, g}, {5, 3, g}, {6, 3, g}, {6, 4, g}, {6, 5, g}, {5, 6, g}, {4, 6, g}, {3, 6, g}, {3, 5, g}, {3, 4, g}, {4, 3, g}, {5, 2, g}, {6, 2, g}, {7, 2, g}, {7, 3, g}, {7, 4, g}, {7, 5, g}, {6, 6, g}, {5, 7, g}, {4, 7, g}, {3, 7, g}, {2, 7, g}, {2, 6, g}, {2, 5, g}, {2, 4, g}, {3, 3, g}, {4, 2, g}, {5, 1, g}, {6, 1, g}, {7, 1, g}, {8, 1, g}, {8, 2, g}, {8, 3, g}, {8, 4, g}}},
		{Hex{2, -2, g}, 4, []Hex{{2, -2, g}, {2, -1, g}, {1, -1, g}, {1, -2, g}, {2, -3, g}, {3, -3, g}, {3, -2, g}, {3, -1, g}, {2, 0, g}, {1, 0, g}, {0, 0, g}, {0, -1, g}, {0, -2, g}, {1, -3, g}, {2, -4, g}, {3, -4, g}, {4, -4, g}, {4, -3, g}, {4, -2, g}, {4, -1, g}, {3, 0, g}, {2, 1, g}, {1, 1, g}, {0, 1, g}, {-1, 1, g}, {-1, 0, g}, {-1, -1, g}, {-1, -2, g}, {0, -3, g}, {1, -4, g}, {2, -5, g}, {3, -5, g}, {4, -5, g}, {5, -5, g}, {5, -4, g}, {5, -3, g}, {5, -2, g}, {5, -1, g}, {4, 0, g}, {3, 1, g}, {2, 2, g}, {1, 2, g}, {0, 2, g}, {-1, 2, g}, {-2, 2, g}, {-2, 1, g}, {-2, 0, g}, {-2, -1, g}, {-2, -2, g}, {-1, -3, g}, {0, -4, g}, {1, -5, g}, {2, -6, g}, {3, -6, g}, {4, -6, g}, {5, -6, g}, {6, -6, g}, {6, -5, g}, {6, -4, g}, {6, -3, g}, {6, -2, g}}},
	}

	for _, test := range testCases {
		h := test.centre.Range(test.N)
		if !h.Equals(NewSet(test.h)) {
			t.Errorf("Incorrect range for %d around %v", test.N, test.centre)
		}
	}
}
