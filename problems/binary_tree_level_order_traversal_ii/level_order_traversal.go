package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	mapResult := make(map[int][]int)
	depth := 0
	explore(root, 1, &depth, mapResult)
	if depth == 0 {
		return [][]int{}
	} else {
		result := make([][]int, depth)
		for key, arr := range mapResult{
			result[depth - key] = arr
		}
		return result
	}
}

func explore(node *TreeNode, level int, depth *int, mapResult map[int][]int)  {
	if node == nil {
		return
	} else {
		if level > *depth {
			*depth = level
		}
		arr, ok := mapResult[level]
		if !ok {
			mapResult[level] = []int{}
		}
		mapResult[level] = append(arr, node.Val)
	}

	if node.Left != nil {
		explore(node.Left, level + 1, depth, mapResult)
	}
	if node.Right != nil {
		explore(node.Right, level + 1, depth, mapResult)
	}
}



func main()  {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	result := levelOrderBottom(root)
	fmt.Println(result)
}



