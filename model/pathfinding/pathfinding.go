package pathfinding

type Node interface {
	Neighbours() []Node
	Cost(Node) int
	Heuristic(Node) int
}

type PrioritisedNode struct {
	node     Node
	priority int
}

type Path []Node

func backtrack(end Node, parent map[Node]Node) Path {
	depth := 1
	for n := end; parent[n] != nil; n = parent[n] {
		depth += 1
	}

	path := make(Path, depth)
	n := end
	for i := depth - 1; i >= 0; i -= 1 {
		path[i] = n
		n = parent[n]
	}

	return path
}

func All(start Node) map[Node]Path {
	frontier := NewNodeQueue()
	parent := make(map[Node]Node)
	visited := make(map[Node]bool)

	frontier.Put(start, nil, 0, 0)

	for !frontier.Empty() {
		current, prev, old_cost := frontier.Get()
		if visited[current] {
			continue
		}
		visited[current] = true
		parent[current] = prev

		for _, next := range current.Neighbours() {
			new_cost := old_cost + current.Cost(next)
			if !visited[next] {
				frontier.Put(next, current, new_cost, new_cost)
			}
		}
	}

	path := make(map[Node]Path)
	for end := range visited {
		path[end] = backtrack(end, parent)
	}

	return path
}

func AllWithin(start Node, N int) map[Node]Path {
	frontier := NewNodeQueue()
	parent := make(map[Node]Node)
	visited := make(map[Node]bool)

	frontier.Put(start, nil, 0, 0)

	for !frontier.Empty() {
		current, prev, old_cost := frontier.Get()
		if visited[current] {
			continue
		}
		visited[current] = true
		parent[current] = prev

		for _, next := range current.Neighbours() {
			new_cost := old_cost + current.Cost(next)
			if new_cost <= N && !visited[next] {
				frontier.Put(next, current, new_cost, new_cost)
			}
		}
	}

	path := make(map[Node]Path)
	for end := range visited {
		path[end] = backtrack(end, parent)
	}

	return path
}

func Single(start, end Node) (Path, bool) {
	frontier := NewNodeQueue()
	parent := make(map[Node]Node)
	visited := make(map[Node]bool)

	frontier.Put(start, nil, 0, start.Heuristic(end))

	for !frontier.Empty() {
		current, prev, old_cost := frontier.Get()
		if visited[current] {
			continue
		}
		visited[current] = true
		parent[current] = prev

		if current == end {
			return backtrack(end, parent), true
		}

		for _, next := range current.Neighbours() {
			new_cost := old_cost + current.Cost(next)
			if !visited[next] {
				frontier.Put(next, current, new_cost, new_cost+next.Heuristic(end))
			}
		}
	}

	return nil, false
}
