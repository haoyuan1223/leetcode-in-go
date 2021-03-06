package _086_partition_list

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	beforeHead := &ListNode{}
	before := beforeHead
	afterHead := &ListNode{}
	after := afterHead
	for head != nil {
		if head.Val >= x {
			after.Next = head
			after = after.Next
		} else {
			before.Next = head
			before = before.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
