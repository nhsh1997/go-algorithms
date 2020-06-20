package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers (l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val: 0,
		Next: nil,
	}

	currentL1 := l1
	currentL2 := l2
	currentResult := result

	initFlag := false

	currentResult.Val = currentL1.Val + currentL2.Val
	if currentResult.Val >= 10 {
		currentResult.Val = currentResult.Val % 10
		currentResult.Next = &ListNode{
			Val: 1,
			Next: nil,
		}
		currentResult = currentResult.Next
		initFlag = true
	}

	currentL1 = currentL1.Next
	currentL2 = currentL2.Next
	flag := false

	for currentL1 != nil || currentL2 != nil  {
		sum := 0
		if currentL1 != nil {
			sum += currentL1.Val
			currentL1 = currentL1.Next
		}
		if currentL2 != nil {
			sum += currentL2.Val
			currentL2 = currentL2.Next
		}

		if flag {
			sum++
		}

		if sum >= 10 {
			sum = sum % 10
			flag = true
		} else {
			flag = false
		}

		if !initFlag {
			currentResult.Next = &ListNode{
				Val: sum,
				Next: nil,
			}
			currentResult = currentResult.Next
		} else {
			currentResult.Val += sum
			if currentResult.Val >= 10 {
				currentResult.Val = currentResult.Val % 10
				currentResult.Next = &ListNode {
					Val: 1,
					Next: nil,
				}
				currentResult = currentResult.Next
			}
		}
	}

	if flag {
		currentResult.Next = &ListNode {
			Val: 1,
			Next: nil,
		}
		currentResult = currentResult.Next
	}

	return result
}