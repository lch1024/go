package main

import "testing"

type maze [][]int

type test struct {
	in, out maze
	nums int
}

func equal(s, d maze) bool {
	if len(s) != len(d) || len(s[0]) != len(d[0]) {
		return false
	}

	for i := range s {
		for j := range s[i] {
			if s[i][j] != d[i][j] {
				return false
			}
		}
	}

	return true
}

func TestWalk(t *testing.T) {
	test1_in := [][]int {
		{0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0},
		{1, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{1, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0},
	}

	test1_out := [][]int {
		{0, 1, 2, 3, 4, 0},
		{1, 0, 3, 0, 5, 6},
		{0, 5, 4, 5, 0, 7},
		{0, 6, 0, 6, 0, 8},
		{0, 7, 8, 7, 8, 0},
		{9, 8, 9, 0, 9, 10},
	}

	test2_in := [][]int {
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}

	test2_out := [][]int {
		{0, 0, 4,  5,  6},
		{1, 2, 3,  0,  7},
		{2, 0, 4,  0,  8},
		{0, 0, 0,  10, 9},
		{0, 0, 12, 11, 0},
		{0, 0, 13, 12, 13},
	}

	tests := []test{
		{test1_in, test1_out, 10},
		{test2_in, test2_out, 13},
	}

	for i := range tests {
		testmaze := tests[i].in
		outmaze := tests[i].out
		outnums := tests[i].nums
		testoutmaze, nums := walk(testmaze, point{0, 0}, point{len(testmaze) - 1, len(testmaze[0]) - 1})
		if !equal(outmaze, testoutmaze)|| outnums != nums {
			t.Errorf("Error")
		}
	}
}
