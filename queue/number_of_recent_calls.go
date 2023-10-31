package queue

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	var i int
	for i < len(this.calls) && this.calls[i] < t-3000 {
		i++
	}
	this.calls = append(this.calls, t)
	this.calls = this.calls[i:]
	return len(this.calls)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
