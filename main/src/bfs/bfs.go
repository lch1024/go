package main

import (
	"fmt"
	"os"
)

// 深度优先遍历找最短路径
// 原理 树的深度优先遍历 保证每一层到达的点位都是步数都是最小的 如果走过再走一遍 那么这条路径一定不是最小的

// 迷宫形状
//
// 0  0  0  0  0  1
// 0  1  0  1  0  0
// 1  0  0  0  1  0
// 1  0  1  0  1  0
// 1  0  0  0  0  1
// 0  0  0  1  0  0

// 走到迷宫各个位置最小路径图记录如下
// *为墙壁
//
// 0  1  2  3  4  *
// 1  *  3  *  5  6
// *  5  4  5  *  7
// *  6  *  6  *  8
// *  7  8  7  8  *
// 9  8  9  *  9  10

// 分析
//
// 需求的变量 ：
// 1. 原始迷宫数组  2. 开始坐标i, j  3. 目标坐标i, j
// 4. 标记走过最短路径的迷宫数组(此数组 从终点 依次找小于他1步地 位置依次输出i, j就是一种最短路径走法吗, 而终点就是最短路径值)
//

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

func (p point) move(s point) point {
	r := point{
		i: p.i + s.i,
		j: p.j + s.j,
	}
	return r
}

var dirs = []point{
	{0, 1},    // 上
	{0, -1},   // 下
	{-1, 0},   // 左
	{1, 0},    // 右
}

func (p point) stepLegal(steps [][]int) bool {
	if p.i < 0 || p.i >= len(steps) {
		return false
	}

	if p.j < 0 || p.j >= len(steps[0]) {
		return false
	}

	if steps[p.i][p.j] != 0 {
		return false
	}
	return true
}

func walk(maze [][]int, start, end point) ([][]int, int) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
		for j := range steps[i] {
			steps[i][j] = 0
		}
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for i := range dirs {
			next := cur.move(dirs[i])

			if start == next {
				continue
			}

			if !next.stepLegal(steps) {
				continue
			}

			if maze[next.i][next.j] != 0 {
				continue
			}

			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			Q = append(Q, next)
		}
	}

	return steps, steps[end.i][end.j]
}

func main() {
	maze := readMaze("src/bfs/maze.in")

	fmt.Println("读取的迷宫为")
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	steps, num := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println("最短步数为：", num)
	fmt.Println("最短路径步数迷宫为：")
	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d ", steps[i][j])
		}
		fmt.Println()
	}
}