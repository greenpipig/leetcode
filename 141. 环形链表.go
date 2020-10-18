package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head.Next==nil||head==nil{
		return false
	}
	fastHead := head.Next       // 快指针，每次走两步
	for fastHead != nil && head != nil && fastHead.Next != nil {
		if fastHead == head {   // 快慢指针相遇，表示有环
			return true
		}
		fastHead = fastHead.Next.Next
		head = head.Next        // 慢指针，每次走一步
	}
	return false
}
