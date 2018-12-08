package task1

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (string, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return "", err
	}
	graph := NewGraph()
	graph.PopulateFromInput(lines)
	return graph.GetOrderOfWork(), nil
}

func Solution2(inputFile string, maxWorkers int) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	graph := NewGraph()
	graph.PopulateFromInput(lines)
	timeElapsed := 0
	for graph.Size() > 0 {
		currentWork := graph.FindAvailableForWork(maxWorkers)
		completedTasks := graph.CompletedTasks(currentWork)
		if len(completedTasks) > 0 {
			for _, id := range completedTasks {
				graph.RemoveNode(id)
			}
		} else {
			graph.WorkOn(currentWork)
			timeElapsed++
		}
	}
	return timeElapsed, nil
}

type Edge struct {
	from, to string
}

func NewEdge(from, to string) Edge {
	return Edge{
		from: from,
		to:   to,
	}
}

type NodeDetails struct {
	timeToComplete int
	to, from       map[string]struct{}
}

func NewNodeDetails(id string) NodeDetails {
	return NodeDetails{
		timeToComplete: int([]rune(id)[0]) - 4,
		to:             make(map[string]struct{}),
		from:           make(map[string]struct{}),
	}
}

type Graph struct {
	nodes map[string]NodeDetails
	edges []Edge
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]NodeDetails),
	}
}

func (g *Graph) PopulateFromInput(inputLines []string) {
	graphRegex := regexp.MustCompile(`^Step (.) must be finished before step (.) can begin\.$`)
	for _, line := range inputLines {
		edgeEntry := graphRegex.FindAllStringSubmatch(line, -1)
		from := edgeEntry[0][1]
		to := edgeEntry[0][2]
		g.AddEdge(NewEdge(from, to))
	}
}

func (g *Graph) AddEdge(e Edge) {
	if _, exists := g.nodes[e.from]; !exists {
		g.nodes[e.from] = NewNodeDetails(e.from)
	}
	if _, exists := g.nodes[e.to]; !exists {
		g.nodes[e.to] = NewNodeDetails(e.to)
	}

	fromNode := g.nodes[e.from]
	fromNode.to[e.to] = struct{}{}
	g.nodes[e.from] = fromNode

	toNode := g.nodes[e.to]
	toNode.from[e.from] = struct{}{}
	g.nodes[e.to] = toNode

	g.edges = append(g.edges, e)
}

func (g *Graph) ListEdges() {
	for _, edge := range g.edges {
		fmt.Printf("%s -> %s\n", edge.from, edge.to)
	}
}

func (g *Graph) ListNodes() {
	for id, node := range g.nodes {
		for f := range node.from {
			fmt.Printf("%s ", f)
		}
		fmt.Printf("-> %s (%ds) ->", id, node.timeToComplete)
		for t := range node.to {
			fmt.Printf(" %s", t)
		}
		fmt.Print("\n")
	}
}

func (g *Graph) sortedNodeNames() []string {
	var keys []string
	for key := range g.nodes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (g *Graph) FindStart() string {
	for _, key := range g.sortedNodeNames() {
		if len(g.nodes[key].from) == 0 {
			return key
		}
	}
	return ""
}

func (g *Graph) FindAvailableForWork(maxWorkers int) []string {
	var available []string
	for _, key := range g.sortedNodeNames() {
		if len(g.nodes[key].from) == 0 && maxWorkers > 0 {
			available = append(available, key)
			maxWorkers--
		}
	}
	return available
}

func (g *Graph) Size() int {
	return len(g.nodes)
}

func (g *Graph) RemoveNode(id string) {
	for node := range g.nodes {
		if node == id {
			delete(g.nodes, id)
		} else {
			delete(g.nodes[node].from, id)
			delete(g.nodes[node].to, id)
		}
	}
}

func (g *Graph) GetOrderOfWork() string {
	result := ""
	for g.Size() > 0 {
		currentStart := g.FindStart()
		result += currentStart
		g.RemoveNode(currentStart)
	}
	return result
}

func (g *Graph) CompletedTasks(ids []string) []string {
	var completeTasks []string
	for _, id := range ids {
		if g.nodes[id].timeToComplete == 0 {
			completeTasks = append(completeTasks, id)
		}
	}
	return completeTasks
}

func (g *Graph) WorkOn(ids []string) {
	for _, id := range ids {
		node := g.nodes[id]
		node.timeToComplete--
		g.nodes[id] = node
	}
}
