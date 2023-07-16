package main

import "fmt"

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	someShit := mergeTwoLists(list1, list2)

	fmt.Println(someShit.Val)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list2.Val >= list1.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)

		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)

		return list2
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
