package heap_pq

type SmallestInfiniteSet struct {
	h       myHeap
	exclude map[int]bool
	next    int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{h: myHeap{}, exclude: make(map[int]bool), next: 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.h.size() == 0 {
		this.exclude[this.next] = true
		this.next++
		return this.next - 1
	}
	res := this.h.pop()
	this.exclude[res] = true
	return res
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.exclude[num] {
		this.h.add(num)
		this.exclude[num] = false
	}
}

// ====== myHeap =======
type myHeap struct {
	data []int
}

func (h *myHeap) pop() int {
	res := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return res
}

func (h *myHeap) add(num int) {
	h.data = append(h.data, num)
	h.up(len(h.data) - 1)
}

func (h *myHeap) size() int {
	return len(h.data)
}

func (h *myHeap) down(index int) {
	left := 2*index + 1
	if left >= len(h.data) {
		return
	}
	minIndex := left
	right := 2*index + 2
	if right < len(h.data) && h.data[right] < h.data[minIndex] {
		minIndex = right
	}
	if h.data[index] > h.data[minIndex] {
		h.data[index], h.data[minIndex] = h.data[minIndex], h.data[index]
		h.down(minIndex)
	}
}

func (h *myHeap) up(index int) {
	if index <= 0 {
		return
	}
	parent := (index - 1) / 2
	if h.data[index] < h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		h.up(parent)
	}
}
