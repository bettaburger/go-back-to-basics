package main

/* Singly-linked list */
type listNode struct {
	val int
	next *listNode
}

/* Given two linked list, digits are stored in reverse order.
	Add the two numbers and return the sum as a linked list. 
	*/
func twoNumbersBrute(l1 *listNode, l2 *listNode) *listNode {
	// l1 = [2, 4, 3]
	// l2 = [5, 6, 4]
	// output = [7, 0, 8]
	// 342 + 465 = 807
	// 3 + 4 = 7
	// 4 + 6 = (1) 0 
	// carry = 1
	// 2 + 5 + (1) = 8
	carry := 0
	output := new(listNode)

	for node := output; l1 != nil || l2 != nil || carry > 0; node = node.next {
		if l1 != nil {
			carry += l1.val // update carry.
			l1 = l1.next // go to next node
		}

		if l2 != nil {
			carry += l2.val
			l2 = l2.next
		}
		// this part handles the carry. 
		// 2 % 10 = 2
		// 10 % 10 = 0
		node.next = &listNode{carry%10, nil} // creates a new node 
		// updates carry for the next loop
		// 2 / 10 = 0
		// 10 / 10 = 1
		// carry = 1
		carry /= 10
	}
	return output.next
}