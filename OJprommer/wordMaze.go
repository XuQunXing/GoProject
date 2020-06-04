package main

import (
	"fmt"
)

/*
12. Word Maze
*/
func main() {
	for {
		//输入
		var n, m int
		var word string
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			break
		}
		fmt.Scanln(&word)
		var maze [][]rune = make([][]rune, n)
		for i := 0; i < n; i++ {
			var str string
			fmt.Scanln(&str)
			maze[i] = []rune(str)
		}
		//处理
		flag := false
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if search(maze, word, i, j, 0) {
					flag = true
					break
				}
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

//dfs
func search(maze [][]rune, word string, x int, y int, n int) bool {
	if x < 0 || x >= len(maze) || y < 0 || y >= len(maze[0]) {
		return false
	}
	if maze[x][y] != rune(word[n]) {
		return false
	}
	if n == len(word)-1 {
		return true
	}
	temp := maze[x][y]
	maze[x][y] = '#'
	if search(maze, word, x+1, y, n+1) || search(maze, word, x-1, y, n+1) || search(maze, word, x, y+1, n+1) || search(maze, word, x, y-1, n+1) {
		return true
	}
	maze[x][y] = temp
	return false
}
