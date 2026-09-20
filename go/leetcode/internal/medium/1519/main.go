// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/description/
// 1519. Number of Nodes in the Sub-Tree With the Same Label
package main

func countSubTrees(n int, edges [][]int, labels string) []int {
	return countSubTrees2(n, edges, labels)
}

func countSubTrees2(n int, edges [][]int, labels string) []int {
	edgeMap := make(map[int][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		edgeMap[a] = append(edgeMap[a], b)
		edgeMap[b] = append(edgeMap[b], a)
	}

	result := make([]int, n)
	searcher := Searcher{
		EdgeMap: edgeMap,
		Visited: make([]bool, n),
	}
	searcher.dfs(0, labels, result)
	return result
}

type Searcher struct {
	EdgeMap map[int][]int
	Visited []bool
}

func (searcher Searcher) dfs(number int, labels string, result []int) map[byte]int {
	labelCounts := make(map[byte]int, 0)
	visited := searcher.Visited
	visited[number] = true

	for _, child := range searcher.EdgeMap[number] {
		if visited[child] {
			continue
		}

		count := searcher.dfs(child, labels, result)
		for childLabel, count := range count {
			labelCounts[childLabel] += count
		}
	}
	label := labels[number]
	labelCounts[label]++
	result[number] = labelCounts[label]
	return labelCounts
}

type Node struct {
	Number      int
	Label       rune
	Children    []*Node
	LabelCounts map[rune]int
}

func countSubTrees1(n int, edges [][]int, labels string) []int {
	nodeMap := buildNodeMap(edges, labels)

	result := make([]int, len(labels))
	for index, label := range labels {
		result[index] = countLabels(label, nodeMap[index])
	}
	return result
}

func buildNodeMap(edges [][]int, labels string) map[int]*Node {
	edgeMaps := make(map[int][]int, 0)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		edgeMaps[from] = append(edgeMaps[from], to)
		edgeMaps[to] = append(edgeMaps[to], from)
		// nodeMap[from].Children = append(nodeMap[from].Children, nodeMap[to])
	}

	visited := make(map[int]bool, len(labels))
	nodeMap := make(map[int]*Node, 0)
	nodeMap[0] = &Node{
		Number:      0,
		Label:       rune(labels[0]),
		Children:    make([]*Node, 0),
		LabelCounts: make(map[rune]int, 0),
	}
	queue := []int{0}
	for len(queue) > 0 {
		number := queue[0]
		queue = queue[1:]

		parent := nodeMap[number]
		visited[number] = true

		for _, childNumber := range edgeMaps[parent.Number] {
			if visited[childNumber] {
				continue
			}
			queue = append(queue, childNumber)

			child, ok := nodeMap[childNumber]
			if !ok {
				child = &Node{
					Number:      childNumber,
					Label:       rune(labels[childNumber]),
					Children:    make([]*Node, 0),
					LabelCounts: make(map[rune]int, 0),
				}
				nodeMap[childNumber] = child
			}
			parent.Children = append(parent.Children, child)
		}
	}
	return nodeMap

}

func countLabels(label rune, node *Node) int {
	if node == nil {
		return 0
	}
	if count, ok := node.LabelCounts[label]; ok {
		return count
	}

	sum := 0
	for _, child := range node.Children {
		sum += countLabels(label, child)
	}
	if node.Label == label {
		sum++
	}
	node.LabelCounts[label] = sum
	return sum
}
