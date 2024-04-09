package graph

func canVisitAllRooms(rooms [][]int) bool {
	total := len(rooms)
	var toVisit [][]int
	checklist := make([]bool, len(rooms))
	count := 1
	checklist[0] = true
	toVisit = append(toVisit, rooms[0])
	for len(toVisit) > 0 {
		keys := toVisit[0]
		toVisit = toVisit[1:]
		for _, key := range keys {
			if checklist[key] {
				continue
			}
			checklist[key] = true
			count++
			toVisit = append(toVisit, rooms[key])
			if count == total {
				return true
			}
		}
	}
	return false
}
