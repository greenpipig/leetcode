package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//不是很理解。。。不是操作得p吗为什么head变了是因为地址的变化吗
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p!=nil {
		val := p.Val
		for p.Next!=nil && p.Next.Val ==  val{
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}
