package graph

type pair struct {
	div   string
	value float64
}

func (p pair) first() string {
	return p.div
}

func (p pair) second() float64 {
	return p.value
}

type numberQueue []pair

func (q *numberQueue) pop() pair {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *numberQueue) push(x pair) {
	*q = append(*q, x)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dict := make(map[string][]pair)
	for i, equation := range equations {
		x := equation[0]
		y := equation[1]
		dict[x] = append(dict[x], pair{y, values[i]})
		dict[y] = append(dict[y], pair{x, 1 / values[i]})
	}

	ans := make([]float64, len(queries))
	for i, query := range queries {
		x := query[0]
		y := query[1]
		if _, ok := dict[x]; !ok {
			ans[i] = -1.0
			continue
		}
		if _, ok := dict[y]; !ok {
			ans[i] = -1.0
			continue
		}
		visited := make(map[string]bool)
		var q numberQueue
		q.push(pair{x, 1})
		for len(q) > 0 {
			cur := q.pop()
			if visited[cur.first()] {
				continue
			}
			visited[cur.first()] = true
			if cur.first() == y {
				ans[i] = cur.second()
				break
			}
			for _, p := range dict[cur.first()] {
				q.push(pair{p.first(), p.second() * cur.second()})
			}
		}
		if ans[i] == 0 {
			ans[i] = -1
		}
	}
	return ans
}
