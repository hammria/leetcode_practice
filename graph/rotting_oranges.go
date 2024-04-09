package graph

func orangesRotting(grid [][]int) int {
	q := posQ{}
	var good int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q.push(opos{x: i, y: j})
			} else if grid[i][j] == 1 {
				good++
			}
		}
	}
	if good == 0 {
		return 0
	}
	res := -1
	for !q.isEmpty() {
		res++
		var nextMinute posQ
		for _, p := range q.data {
			spread(p.x, p.y, grid, &nextMinute)
		}
		q = nextMinute
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return res
}

func spread(x, y int, grid [][]int, nextMinute *posQ) {
	m := len(grid)
	n := len(grid[0])
	if x+1 < m && grid[x+1][y] == 1 {
		grid[x+1][y] = 2
		nextMinute.push(opos{x: x + 1, y: y})
	}
	if x > 0 && grid[x-1][y] == 1 {
		grid[x-1][y] = 2
		nextMinute.push(opos{x: x - 1, y: y})
	}
	if y+1 < n && grid[x][y+1] == 1 {
		grid[x][y+1] = 2
		nextMinute.push(opos{x: x, y: y + 1})
	}
	if y > 0 && grid[x][y-1] == 1 {
		grid[x][y-1] = 2
		nextMinute.push(opos{x: x, y: y - 1})
	}
}

type opos struct {
	x int
	y int
}

type posQ struct {
	data []opos
}

func (q *posQ) push(pos opos) {
	(*q).data = append((*q).data, pos)
}

func (q *posQ) isEmpty() bool {
	return len((*q).data) == 0
}
