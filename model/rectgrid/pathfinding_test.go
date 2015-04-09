package rectgrid

import "testing"

import "github.com/FinnStokes/gnome-tactics/model/pathfinding"

func TestAll(t *testing.T) {
	g := testGrid{5, 5}

	m := pathfinding.All(Rect{0, 2, g})

	paths := map[Rect][][]Rect{
		Rect{0, 2, g}: [][]Rect{[]Rect{{0, 2, g}}},
		Rect{0, 1, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}}},
		Rect{0, 0, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}}},
		Rect{0, 3, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}}},
		Rect{0, 4, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}}},
		Rect{1, 2, g}: [][]Rect{[]Rect{{0, 2, g}, {1, 2, g}}},
		Rect{1, 1, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 1, g}}},
		Rect{1, 0, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}, {1, 0, g}}},
		Rect{1, 3, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {1, 3, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 3, g}}},
		Rect{1, 4, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}}},
		Rect{2, 2, g}: [][]Rect{[]Rect{{0, 2, g}, {1, 2, g}, {2, 2, g}}},
		Rect{2, 1, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 1, g}, {2, 1, g}}},
		Rect{2, 0, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}, {1, 0, g}, {2, 0, g}}},
		Rect{2, 3, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {1, 3, g}, {2, 3, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 3, g}, {2, 3, g}}},
		Rect{2, 4, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}, {2, 4, g}}},
		Rect{3, 2, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {3, 1, g}, {3, 2, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 1, g}, {2, 1, g}, {3, 1, g}, {3, 2, g}}, []Rect{{0, 2, g}, {0, 3, g}, {1, 3, g}, {2, 3, g}, {3, 3, g}, {3, 2, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 3, g}, {2, 3, g}, {3, 3, g}, {3, 2, g}}},
		Rect{3, 1, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {1, 1, g}, {2, 1, g}, {3, 1, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 1, g}, {2, 1, g}, {3, 1, g}}},
		Rect{3, 0, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}, {1, 0, g}, {2, 0, g}, {3, 0, g}}},
		Rect{3, 3, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {1, 3, g}, {2, 3, g}, {3, 3, g}}, []Rect{{0, 2, g}, {1, 2, g}, {1, 3, g}, {2, 3, g}, {3, 3, g}}},
		Rect{3, 4, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}, {2, 4, g}, {3, 4, g}}},
		Rect{4, 2, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}, {1, 0, g}, {2, 0, g}, {3, 0, g}, {4, 0, g}, {4, 1, g}, {4, 2, g}}, []Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}, {2, 4, g}, {3, 4, g}, {4, 4, g}, {4, 3, g}, {4, 2, g}}},
		Rect{4, 1, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}, {1, 0, g}, {2, 0, g}, {3, 0, g}, {4, 0, g}, {4, 1, g}}},
		Rect{4, 0, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 1, g}, {0, 0, g}, {1, 0, g}, {2, 0, g}, {3, 0, g}, {4, 0, g}}},
		Rect{4, 3, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}, {2, 4, g}, {3, 4, g}, {4, 4, g}, {4, 3, g}}},
		Rect{4, 4, g}: [][]Rect{[]Rect{{0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}, {2, 4, g}, {3, 4, g}, {4, 4, g}}},
	}

	for i, p := range paths {
		found, ok := m[i]
		if !ok {
			t.Errorf("No path found from %v to %v", Rect{0, 2, g}, i)
			continue
		}
		okay := false
		for _, path := range p {
			if len(found) == len(path) {
				okay = true
				for j, node := range path {
					if node != found[j] {
						okay = false
						break
					}
				}
				if okay == true {
					break
				}
			}
		}
		if okay == false {
			t.Errorf("Incorrect path from %v to %v", Rect{0, 2, g}, i)
			t.Errorf("%v", found)
		}
	}

	for end := range m {
		_, ok := paths[end.(Rect)]
		if !ok {
			t.Errorf("Path to disconnected node %v", end)
		}
	}
}

func TestSingle(t *testing.T) {
	g := testGrid{5, 5}

	paths := [][]Rect{
		[]Rect{{2, 1, g}, {3, 1, g}, {3, 2, g}, {3, 3, g}},
		[]Rect{{0, 0, g}, {0, 1, g}, {0, 2, g}, {0, 3, g}, {0, 4, g}, {1, 4, g}, {2, 4, g}, {3, 4, g}},
		[]Rect{{2, 4, g}, {2, 3, g}, {2, 2, g}},
		[]Rect{{2, 2, g}, {2, 1, g}, {2, 0, g}},
		[]Rect{{3, 1, g}, {2, 1, g}, {1, 1, g}, {1, 2, g}},
		[]Rect{{3, 1, g}, {2, 1, g}, {1, 1, g}, {0, 1, g}},
	}

	for _, path := range paths {
		found, ok := pathfinding.Single(path[0], path[len(path)-1])
		if !ok {
			t.Errorf("No path found from %v to %v", path[0], path[len(path)-1])
			continue
		}
		if len(found) != len(path) {
			t.Errorf("Incorrect path from %v to %v", path[0], path[len(path)-1])
			t.Errorf("%v", found)
			continue
		}
		for j, node := range path {
			if node != found[j] {
				t.Errorf("Incorrect path from %v to %v", path[0], path[len(path)-1])
				t.Errorf("%v", found)
				break
			}
		}
	}
}
