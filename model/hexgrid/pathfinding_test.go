package hexgrid

import "testing"

import "github.com/FinnStokes/gnome-tactics/model/pathfinding"

func TestAll(t *testing.T) {
	g := testGrid{3}

	m := pathfinding.All(Hex{3, -2, g})

	paths := map[Hex][][]Hex{
		Hex{3, -2, g}:  [][]Hex{[]Hex{{3, -2, g}}},                                                                                                                                                                                                                                          // 0
		Hex{3, -3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}}},                                                                                                                                                                                                                              // 1
		Hex{3, -1, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}}},                                                                                                                                                                                                                              // 1
		Hex{2, -3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}}},                                                                                                                                                                                                                  // 2
		Hex{3, 0, g}:   [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}}},                                                                                                                                                                                                                   // 2
		Hex{1, -3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}}},                                                                                                                                                                                                      // 3
		Hex{2, 1, g}:   [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}}},                                                                                                                                                                                                        // 3
		Hex{0, -3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}}},                                                                                                                                                                                          // 4
		Hex{1, 2, g}:   [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}, {1, 2, g}}},                                                                                                                                                                                             // 4
		Hex{-1, -2, g}: [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}, {-1, -2, g}}},                                                                                                                                                                             // 5
		Hex{0, 3, g}:   [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}, {1, 2, g}, {0, 3, g}}},                                                                                                                                                                                  // 5
		Hex{2, -2, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}}},                                                                                                                                                                                                                              // 5
		Hex{2, -1, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}}},                                                                                                                                                                                                                              // 5
		Hex{-2, -1, g}: [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}, {-1, -2, g}, {-2, -1, g}}},                                                                                                                                                                // 6
		Hex{-1, 3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}, {1, 2, g}, {0, 3, g}, {-1, 3, g}}},                                                                                                                                                                      // 6
		Hex{1, -2, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}}},                                                                                                                                                                                                                  // 6
		Hex{2, 0, g}:   [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {2, 0, g}}, []Hex{{3, -2, g}, {3, -1, g}, {2, 0, g}}},                                                                                                                                                                         // 6
		Hex{-3, 0, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}, {-1, -2, g}, {-2, -1, g}, {-3, 0, g}}},                                                                                                                                                    // 7
		Hex{-2, 3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}, {1, 2, g}, {0, 3, g}, {-1, 3, g}, {-2, 3, g}}},                                                                                                                                                          // 7
		Hex{0, -2, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}, {0, -2, g}}},                                                                                                                                                                                                      // 7
		Hex{1, 1, g}:   [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {2, 0, g}, {1, 1, g}}, []Hex{{3, -2, g}, {3, -1, g}, {2, 0, g}, {1, 1, g}}},                                                                                                                                                   // 7
		Hex{-3, 1, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}, {-1, -2, g}, {-2, -1, g}, {-3, 0, g}, {-3, 1, g}}},                                                                                                                                        // 8
		Hex{-3, 3, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}, {1, 2, g}, {0, 3, g}, {-1, 3, g}, {-2, 3, g}, {-3, 3, g}}},                                                                                                                                              // 8
		Hex{-1, -1, g}: [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}, {0, -2, g}, {-1, -1, g}}},                                                                                                                                                                                         // 8
		Hex{0, 2, g}:   [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {2, 0, g}, {1, 1, g}, {0, 2, g}}, []Hex{{3, -2, g}, {3, -1, g}, {2, 0, g}, {1, 1, g}, {0, 2, g}}},                                                                                                                             // 8
		Hex{-3, 2, g}:  [][]Hex{[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}, {-1, -2, g}, {-2, -1, g}, {-3, 0, g}, {-3, 1, g}, {-3, 2, g}}, []Hex{{3, -2, g}, {3, -1, g}, {3, 0, g}, {2, 1, g}, {1, 2, g}, {0, 3, g}, {-1, 3, g}, {-2, 3, g}, {-3, 3, g}, {-3, 2, g}}}, // 9
		Hex{-2, 0, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}, {0, -2, g}, {-1, -1, g}, {-2, 0, g}}},                                                                                                                                                                             // 9
		Hex{-1, 2, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {2, 0, g}, {1, 1, g}, {0, 2, g}, {-1, 2, g}}, []Hex{{3, -2, g}, {3, -1, g}, {2, 0, g}, {1, 1, g}, {0, 2, g}, {-1, 2, g}}},                                                                                                     // 9
		Hex{-2, 1, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}, {0, -2, g}, {-1, -1, g}, {-2, 0, g}, {-2, 1, g}}},                                                                                                                                                                 // 10
		Hex{-2, 2, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {2, 0, g}, {1, 1, g}, {0, 2, g}, {-1, 2, g}}, []Hex{{3, -2, g}, {3, -1, g}, {2, 0, g}, {1, 1, g}, {0, 2, g}, {-1, 2, g}, {-2, 2, g}}},                                                                                         // 10
		Hex{1, -1, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -1, g}}, []Hex{{3, -2, g}, {2, -1, g}, {1, -1, g}}},                                                                                                                                                                       // 10
		Hex{1, 0, g}:   [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {1, 0, g}}},                                                                                                                                                                                                                   // 10
		Hex{0, -1, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -1, g}, {0, -1, g}}, []Hex{{3, -2, g}, {2, -1, g}, {1, -1, g}, {0, -1, g}}, []Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}, {0, -1, g}}},                                                                                        // 11
		Hex{0, 1, g}:   [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {1, 0, g}, {0, 1, g}}},                                                                                                                                                                                                        // 11
		Hex{-1, 0, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -1, g}, {0, -1, g}, {-1, 0, g}}, []Hex{{3, -2, g}, {2, -1, g}, {1, -1, g}, {0, -1, g}, {-1, 0, g}}, []Hex{{3, -2, g}, {2, -2, g}, {1, -2, g}, {0, -1, g}, {-1, 0, g}}},                                                    // 12
		Hex{-1, 1, g}:  [][]Hex{[]Hex{{3, -2, g}, {2, -1, g}, {1, 0, g}, {0, 1, g}, {-1, 1, g}}},                                                                                                                                                                                            // 12
		Hex{0, 0, g}:   [][]Hex{[]Hex{{3, -2, g}, {2, -2, g}, {1, -1, g}, {0, 0, g}}, []Hex{{3, -2, g}, {2, -1, g}, {1, -1, g}, {0, 0, g}}, []Hex{{3, -2, g}, {2, -1, g}, {1, 0, g}, {0, 0, g}}},                                                                                            // 15
	}

	for i, p := range paths {
		found, ok := m[i]
		if !ok {
			t.Errorf("No path found from %v to %v", Hex{3, -2, g}, i)
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
			t.Errorf("Incorrect path from %v to %v", Hex{3, -2, g}, i)
			t.Errorf("%v", found)
		}
	}

	for end := range m {
		_, ok := paths[end.(Hex)]
		if !ok {
			t.Errorf("Path to disconnected node %v", end)
		}
	}
}

func TestSingle(t *testing.T) {
	g := testGrid{3}

	paths := [][]Hex{
		[]Hex{{3, -2, g}, {2, -1, g}, {1, 0, g}, {0, 1, g}, {-1, 1, g}},
		[]Hex{{3, -2, g}, {3, -3, g}, {2, -3, g}, {1, -3, g}, {0, -3, g}, {-1, -2, g}, {-2, -1, g}, {-3, 0, g}, {-3, 1, g}},
		[]Hex{{-3, 3, g}, {-3, 2, g}, {-3, 1, g}, {-3, 0, g}, {-2, -1, g}, {-1, -2, g}, {0, -3, g}, {1, -3, g}},
		[]Hex{{0, 0, g}, {1, -1, g}, {2, -2, g}, {3, -3, g}},
		[]Hex{{2, 1, g}, {1, 1, g}, {0, 1, g}},
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
