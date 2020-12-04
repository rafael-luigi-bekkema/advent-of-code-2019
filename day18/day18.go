package day18

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

type Node struct {
	x, y  int
	value byte
}

func (n *Node) isDoor() bool {
	return 'A' <= n.value && n.value <= 'Z'
}
func (n *Node) isKey() bool {
	return 'a' <= n.value && n.value <= 'z'
}

type Graph struct {
	nodes map[Node]struct{}
	edges map[Node]map[Node]float64
}

var inf = math.Inf(1)

func (g *Graph) AddEdge(n1, n2 Node, value float64) {
	if g.edges == nil {
		g.edges = make(map[Node]map[Node]float64)
	}
	if g.edges[n1] == nil {
		g.edges[n1] = make(map[Node]float64)
	}
	g.edges[n1][n2] = value
	if g.edges[n2] == nil {
		g.edges[n2] = make(map[Node]float64)
	}
	g.edges[n2][n1] = value
}

func (g *Graph) Dijkstra(source Node, keys [26]bool) (dist map[Node]float64, prev map[Node]Node) {
	Q := make(map[Node]struct{}, len(g.nodes))
	dist = make(map[Node]float64)
	prev = make(map[Node]Node)
	for node := range g.nodes {
		if node.isDoor() && !keys[node.value-'A'] {
			continue
		}
		Q[node] = struct{}{}
		dist[node] = inf
	}
	dist[source] = 0

	for len(Q) > 0 {
		var mindist float64
		var minn Node
		minfirst := true
		for n := range Q {
			if d := dist[n]; minfirst || d < mindist {
				mindist = d
				minn = n
				minfirst = false
			}
		}
		if minfirst {
			panic("oops")
		}
		delete(Q, minn)
		for edge, edgeVal := range g.edges[minn] {
			if _, ok := Q[edge]; !ok {
				continue
			}
			if edge.isDoor() && !keys[edge.value-'A'] {
				continue
			}

			alt := dist[minn] + edgeVal
			if alt < dist[edge] {
				dist[edge] = alt
				prev[edge] = minn
			}
		}
	}

	return dist, prev
}

func (graph *Graph) BFS(source Node) (map[Node]float64, map[Node]Node) {
	visited := make(map[Node]bool, len(graph.nodes))
	visited[source] = true
	queue := []Node{
		source,
	}
	prev := make(map[Node]Node)
	dist := make(map[Node]float64, len(graph.nodes))
	for node := range graph.nodes {
		dist[node] = inf
	}
	dist[source] = 0

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		for i := range graph.edges[s] {
			if !visited[i] {
				queue = append(queue, i)
				visited[i] = true
				dist[i] = dist[s] + 1
				prev[i] = s
			}
		}
	}
	return dist, prev
}

func (graph *Graph) AddNode(n Node) {
	if graph.nodes == nil {
		graph.nodes = make(map[Node]struct{})
	}
	graph.nodes[n] = struct{}{}
}

func (graph *Graph) addNode(n Node, lines []string, collected string) {
	graph.AddNode(n)

	x, y := n.x, n.y
	if y >= 1 && passable(lines[y-1][x], collected) {
		graph.AddEdge(n, Node{x, y - 1, lines[y-1][x]}, 1)
	}
	if y < len(lines)-1 && passable(lines[y+1][x], collected) {
		graph.AddEdge(n, Node{x, y + 1, lines[y+1][x]}, 1)
	}
	if x >= 1 && passable(lines[y][x-1], collected) {
		graph.AddEdge(n, Node{x - 1, y, lines[y][x-1]}, 1)
	}
	if x < len(lines[y])-1 && passable(lines[y][x+1], collected) {
		graph.AddEdge(n, Node{x + 1, y, lines[y][x+1]}, 1)
	}
}

func passable(r byte, collected string) bool {
	if 'A' <= r && r <= 'Z' {
		return true
		// for _, o := range collected {
		// 	if r == byte(o) {
		// 		return true
		// 	}
		// }
		// return false
	}
	return r == '.' || r == '@' || ('a' <= r && r <= 'z')
}

type cacheKey struct {
	source Node
	keys   [26]bool
	score  float64
}

type dCacheKey struct {
	source Node
	keys   [26]bool
}

func puzzle1(data string) int {
	var graph Graph
	allkeys := make(map[rune]Node)
	doors := make(map[rune]Node)
	lines := strings.Split(data, "\n")
	var source Node
	for y, line := range lines {
		for x, c := range line {
			if c == '#' { // wall
				continue
			}

			n := Node{x, y, byte(c)}
			if 'A' <= c && c <= 'Z' { // door
				doors[c] = n
			}
			if 'a' <= c && c <= 'z' { // key
				allkeys[c] = n
			}
			if c == '@' {
				source = n
			}

			graph.addNode(n, lines, "")
		}
	}

	// Build final graph with only doors, keys and start
	var fGraph Graph
	fGraph.AddNode(source)
	for _, door := range doors {
		fGraph.AddNode(door)
	}
	for _, key := range allkeys {
		fGraph.AddNode(key)
	}
	for node := range fGraph.nodes {
		dest, prev := graph.BFS(node)
		for n2 := range fGraph.nodes {
			if node == n2 {
				continue
			}
			if dest[n2] == inf {
				continue
			}
			{
				// Don't add edges for nodes that cannot be
				// accessed without going through a door
				p := prev[n2]
				var hasDoor bool
				for {
					isDoor := p.isDoor()
					var ok bool
					p, ok = prev[p]
					if ok && isDoor {
						hasDoor = true
						break
					}
					if !ok {
						break
					}
				}
				if hasDoor {
					continue
				}
			}
			fGraph.AddEdge(node, n2, dest[n2])
		}
	}

	var minScore float64

	found := make(map[cacheKey]float64)
	dCache := make(map[dCacheKey]map[Node]float64)

	var findPath func(source Node, score float64, graph Graph, keys [26]bool, collected string) float64
	findPath = func(source Node, score float64, graph Graph, keys [26]bool, collected string) float64 {
		cacheKey := cacheKey{source, keys, score}
		if val, ok := found[cacheKey]; ok {
			// fmt.Printf("repeat! %v %f\n", keys, val)
			return val
		}

		dKey := dCacheKey{source, keys}
		// fmt.Printf("source: %v, keys: %v, minScore: %f, score: %f\n", source, keys, minScore, score)
		var dist map[Node]float64
		if d, ok := dCache[dKey]; ok {
			dist = d
		} else {
			d, _ := graph.Dijkstra(source, keys)
			dist = d
			dCache[dKey] = d
		}

		minfirst := true
		var minr float64
		var valid bool
		for keyval, have := range keys {
			if have || keyval >= len(allkeys) {
				continue
			}
			keychar := 'a' + rune(keyval)
			keynode := allkeys[keychar]

			val := dist[keynode]
			if val == inf {
				continue
			}
			valid = true

			var r float64
			var remaining = len(allkeys)
			for _, valz := range keys {
				if valz {
					remaining--
				}
			}
			if remaining == 1 {
				r = score + val
				if r < minScore || minScore == 0 {
					fmt.Printf("new low score: %v\n", r)
					minScore = r
				}
			} else {
				if minScore > 0 && (score+val) > minScore {
					// fmt.Printf("return early, minscore: %v %s\n", minScore, collected)
					return score + val
				}
				// clone keys wit current key
				newKeys := keys
				newKeys[keyval] = true

				// Clone graph
				r = findPath(keynode, score+val, graph, newKeys, collected+string(keychar))
			}
			if minfirst || r < minr {
				minr = r
				minfirst = false
			}
		}
		if !valid {
			panic("oops not valid")
		}
		found[cacheKey] = minr
		return minr
	}

	steps := findPath(source, 0, fGraph, [26]bool{}, "")
	return int(steps)
}

func Puzzle1() int {
	data := utils.ReadAll("./input")
	return puzzle1(data[:len(data)-1])
}
