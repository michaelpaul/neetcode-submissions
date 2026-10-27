/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    // iterate the list
	// note next value
	// swap pointers between current node and next
	// return last pointer
	// 0 -> 1
	// 0 -> nil
	if head == nil {
		return nil
	}
	
	var previous *ListNode
	node := head

	for node.Next != nil {
		temp := node.Next 
		node.Next = previous
		previous = node
		node = temp 
	}

	node.Next = previous
	return node
}
