package graph

type pos struct {
	x    int
	y    int
	step int
}

type posQueue []pos

func (p *posQueue) poll() pos {
	cur := (*p)[0]
	*p = (*p)[1:]
	return cur
}

func (p *posQueue) add(cur pos) {
	*p = append(*p, cur)
}

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])
	start := pos{entrance[0], entrance[1], 0}
	maze[entrance[0]][entrance[1]] = '*'
	var q posQueue
	q.add(start)
	for len(q) > 0 {
		cur := q.poll()
		x := cur.x
		y := cur.y
		curStep := cur.step
		if maze[x][y] == '.' && (x == 0 || y == 0 || x == m-1 || y == n-1) {
			return curStep
		}
		if maze[x][y] == '+' {
			continue
		}
		maze[x][y] = '+'
		if x > 0 && maze[x-1][y] == '.' {
			q.add(pos{x - 1, y, curStep + 1})
			maze[x-1][y] = '+'
		}
		if y > 0 && maze[x][y-1] == '.' {
			q.add(pos{x, y - 1, curStep + 1})
			maze[x][y-1] = '+'
		}
		if x < m-1 && maze[x+1][y] == '.' {
			q.add(pos{x + 1, y, curStep + 1})
			maze[x+1][y] = '+'
		}
		if y < n-1 && maze[x][y+1] == '.' {
			q.add(pos{x, y + 1, curStep + 1})
			maze[x][y+1] = '+'
		}
	}
	return -1
}
