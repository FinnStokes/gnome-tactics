package rectgrid

import "github.com/FinnStokes/gnome-tactics/model/pathfinding"

type Grid interface {
	Passable(Rect, Rect) bool
	Cost(Rect, Rect) int
}

type Rect struct {
	X, Y int
	Grid Grid
}

type RectSet map[Rect]bool

var adjacent = [...]struct{ X, Y int }{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func (self Rect) Neighbours() []pathfinding.Node {
	neighbours := make([]pathfinding.Node, 0, len(adjacent))
	for _, rel := range adjacent {
		n := Rect{self.X + rel.X, self.Y + rel.Y, self.Grid}
		if self.Grid.Passable(self, n) {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func (self Rect) Cost(node pathfinding.Node) int {
	return self.Grid.Cost(self, node.(Rect))
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func (self Rect) Distance(other Rect) int {
	return abs(self.X-other.X) + abs(self.Y-other.Y)
}

func (self Rect) Heuristic(node pathfinding.Node) int {
	return self.Distance(node.(Rect))
}

func (self Rect) Line(end Rect) (line []Rect) {
	dx := abs(end.X - self.X)
	dy := abs(end.Y - self.Y)
	var N int
	if dx > dy {
		N = dx
	} else {
		N = dy
	}

	line = make([]Rect, N+1)
	line[0] = self

	step := 1.0 / float64(N)

	for i, t := 1, step; i <= N; i, t = i+1, t+step {
		x := float64(self.X) + float64(end.X-self.X)*t
		y := float64(self.Y) + float64(end.Y-self.Y)*t
		line[i] = Rect{int(x + .5), int(y + .5), self.Grid}
	}

	return
}

func (self Rect) PassableLine(end Rect) (line []Rect) {
	line = self.Line(end)
	for i := 1; i < len(line); i += 1 {
		if !self.Grid.Passable(line[i-1], line[i]) {
			return line[0:i]
		}
	}
	return
}

func (self Rect) Range(N int) (result RectSet) {
	result = make(RectSet)
	for dx := -N; dx <= N; dx += 1 {
		for dy := -N + abs(dx); dy <= N-abs(dx); dy += 1 {
			result[Rect{self.X + dx, self.Y + dy, self.Grid}] = true
		}
	}
	return
}

func NewSet(rects []Rect) (set RectSet) {
	set = make(RectSet)
	for _, r := range rects {
		set[r] = true
	}
	return
}

func (self RectSet) Intersect(other RectSet) (result RectSet) {
	result = make(RectSet)
	for rect, in := range self {
		if in && other[rect] {
			result[rect] = true
		}
	}
	return
}

func (self RectSet) Union(other RectSet) (result RectSet) {
	result = make(RectSet)
	for rect, in := range self {
		if in {
			result[rect] = true
		}
	}
	for rect, in := range other {
		if in {
			result[rect] = true
		}
	}
	return
}

func (self RectSet) Subtract(other RectSet) (result RectSet) {
	result = make(RectSet)
	for rect, in := range self {
		if in && !other[rect] {
			result[rect] = true
		}
	}
	return
}

func (self RectSet) In(other RectSet) bool {
	for rect, in := range self {
		if in && !other[rect] {
			return false
		}
	}
	return true
}

func (self RectSet) Contains(other RectSet) bool {
	return other.In(self)
}

func (self RectSet) Equals(other RectSet) bool {
	return self.In(other) && self.Contains(other)
}
