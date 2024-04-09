package graph

type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func minReorder(n int, connections [][]int) int {
	var q queue
	var ans = 0
	roads := make(map[int][][]int)
	for _, connect := range connections {
		city1 := connect[0]
		city2 := connect[1]
		roads[city1] = append(roads[city1], connect)
		roads[city2] = append(roads[city2], connect)
	}
	visited := make([]bool, n)

	q.push(0)
	for len(q) != 0 {
		cur := q.pop()
		if visited[cur] {
			continue
		}
		visited[cur] = true
		for _, edges := range roads[cur] {
			// given all road is supposed to directed to root
			// the final res graph should be like a 'reversed' tree
			// so this tree has the feature that u cannot go anywhere if you started from root
			// when bfs the node, if this node has a road which directs to it but the source has
			// already been visited, the road is not correct (since each road should direct to
			// root and each root bfs should be an edge that has a child node direct to parent, in
			// which case the child is not supposed to be visited
			if edges[1] == cur && visited[edges[0]] {
				ans++
			}
			if edges[0] == cur {
				q.push(edges[1])
			} else {
				q.push(edges[0])
			}
		}
	}
	return ans
}
