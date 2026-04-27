/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 1. Create a dummy node to start the list
    dummy := &ListNode{}
    tail := dummy

    // 2. Iterate while both lists have nodes
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }

    // 3. If one list is not empty, attach the rest of it
    if list1 != nil {
        tail.Next = list1
    } else {
        tail.Next = list2
    }

    // 4. The dummy's next is the head of our merged list
    return dummy.Next
}
