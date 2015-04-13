package hexgrid

import "math"

import "github.com/FinnStokes/gnome-tactics/model/pathfinding"

type Grid interface {
	Passable(Hex, Hex) bool
	Cost(Hex, Hex) int
}

type Hex struct {
	Q, R int
	Grid Grid
}

type HexSet map[Hex]bool

var adjacent = [...]struct{ Q, R int }{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, -1},
	{-1, 1},
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func round(v float64) int {
	return int(math.Floor(v + .5))
}

func cube_round(q, r float64) (int, int) {
	x := round(q)
	y := round(r)
	z := round(-q - r)

	dx := math.Abs(float64(x) - q)
	dy := math.Abs(float64(y) - r)
	dz := math.Abs(float64(z) + q + r)

	if dx > dy && dx > dz {
		x = -y - z
	} else if dy > dz {
		y = -x - z
	}

	return x, y
}

func (self Hex) Neighbours() []pathfinding.Node {
	neighbours := make([]pathfinding.Node, 0, len(adjacent))
	for _, rel := range adjacent {
		n := Hex{self.Q + rel.Q, self.R + rel.R, self.Grid}
		if self.Grid.Passable(self, n) {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func (self Hex) Cost(n pathfinding.Node) int {
	return self.Grid.Cost(self, n.(Hex))
}

func (self Hex) Distance(other Hex) int {
	return (abs(self.Q-other.Q) + abs(self.R-other.R) + abs((self.Q+self.R)-(other.Q+other.R))) / 2
}

func (self Hex) Heuristic(node pathfinding.Node) int {
	return self.Distance(node.(Hex))
}

func (self Hex) Line(end Hex) (line []Hex) {
	N := self.Distance(end)

	line = make([]Hex, N+1)
	line[0] = self

	step := 1.0 / float64(N)

	for i, t := 1, step; i <= N; i, t = i+1, t+step {
		q := float64(self.Q) + float64(end.Q-self.Q)*t
		r := float64(self.R) + float64(end.R-self.R)*t

		rq, rr := cube_round(q, r)

		line[i] = Hex{rq, rr, self.Grid}
	}

	return
}

func (self Hex) PassableLine(end Hex) (line []Hex) {
	line = self.Line(end)
	for i := 1; i < len(line); i += 1 {
		if !self.Grid.Passable(line[i-1], line[i]) {
			return line[0:i]
		}
	}
	return
}

func (self Hex) Range(N int) (result HexSet) {
	result = make(HexSet)
	for dq := -N; dq < 0; dq += 1 {
		for dr := -N - dq; dr <= N; dr += 1 {
			result[Hex{self.Q + dq, self.R + dr, self.Grid}] = true
		}
	}
	for dq := 0; dq <= N; dq += 1 {
		for dr := -N; dr <= N-dq; dr += 1 {
			result[Hex{self.Q + dq, self.R + dr, self.Grid}] = true
		}
	}
	return
}

func NewSet(hexs []Hex) (set HexSet) {
	set = make(HexSet)
	for _, h := range hexs {
		set[h] = true
	}
	return
}

func (self HexSet) Intersect(other HexSet) (result HexSet) {
	result = make(HexSet)
	for hex, in := range self {
		if in && other[hex] {
			result[hex] = true
		}
	}
	return
}

func (self HexSet) Union(other HexSet) (result HexSet) {
	result = make(HexSet)
	for hex, in := range self {
		if in {
			result[hex] = true
		}
	}
	for hex, in := range other {
		if in {
			result[hex] = true
		}
	}
	return
}

func (self HexSet) Subtract(other HexSet) (result HexSet) {
	result = make(HexSet)
	for hex, in := range self {
		if in && !other[hex] {
			result[hex] = true
		}
	}
	return
}

func (self HexSet) In(other HexSet) bool {
	for hex, in := range self {
		if in && !other[hex] {
			return false
		}
	}
	return true
}

func (self HexSet) Contains(other HexSet) bool {
	return other.In(self)
}

func (self HexSet) Equals(other HexSet) bool {
	return self.In(other) && self.Contains(other)
}
