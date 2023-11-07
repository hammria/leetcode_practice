package linkedlist

func pairSum(head *ListNode) int {
	var index int
	cur := head
	var list []int
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
		index++
	}
	var ans int
	for i := 0; i < index/2; i++ {
		if list[i]+list[index-1-i] > ans {
			ans = list[i] + list[index-1-i]
		}
	}
	return ans

}
