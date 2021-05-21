package main

import (
	"reflect"
	"testing"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		input []int
		output int
	}{
		/*{input: []int{1, 0, 2}, output: 5},
		{input: []int{1, 2, 2}, output: 4},
		{input: []int{1, 3, 2, 2, 1}, output: 7},
		{input: []int{1,2,87,87,87,2,1}, output: 13},
		{input: []int{1,3,4,5,2}, output: 11},
		{input: []int{1,6,10,8,7,3,2}, output: 18},
		{input: []int{29,51,87,87,72,12}, output: 12},
		{input: []int{0,1,2,5,3,2,7}, output: 15},
		{input: []int{1,2,3,5,4,3,2,1}, output: 21},
		{input: []int{1,2,3,5,4,3,2,1,4,3,2,1,3,2,1,1,2,3,4}, output: 47},
		{input: []int{0, 1}, output: 3},
		{input:[]int{58,21,72,77,48,9,38,71,68,77,82,47,25,94,89,54,26,54,54,99,64,71,76,63,81,82,60,64,29,51,87,87,72,12,16,20,21,54,43,41,83,77,41,61,72,82,15,50,36,69,49,53,92,77,16,73,12,28,37,41,79,25,80,3,37,48,23,10,55,19,51,38,96,92,99,68,75,14,18,63,35,19,68,28,49,36,53,61,64,91,2,43,68,34,46,57,82,22,67,89}, output: 208},
		{input: []int{1,3,4,3,2,1}, output: 13},*/
		{input: []int{89, 54, 36, 54, 54, 99, 64}, output: 12},
	}

	for i, tc := range tests {
		got := candy(tc.input)
		if !reflect.DeepEqual(tc.output, got) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.output, got)
		}
	}
}

