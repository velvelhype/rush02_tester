package main

import (
	// "fmt"
	"ft"
	// "os"
)

const INF = 1 << 63 - 1

func printBoard(info *Info) {
	for _, row := range info.AnsBoard.Field {
		for _, b := range row {
			ft.PrintRune(rune(b))
		}
		ft.PrintRune('\n')
	}
	ft.PrintRune('\n')
}

func (board *Board) score() int {
	dist := INF
	for i:=0; i< len(board.Field); i++ {
		for j:=0; j< len(board.Field); j++ {
			if board.Field[i][j] == SPACE && i+j < dist {
				dist = i+j;
			}
		}
	}
	return dist
}

func all(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}

func dfs(info *Info, y, x int) {
	// fmt.Println(y, x, info)
	if y == info.Length - 1 && x == info.Length - 1 {
		if all(info.Used) && (info.AnsBoard.Field == nil || info.AnsBoard.score() < info.Board.score()) {
			info.AnsBoard = info.Board.Dup()
			printBoard(info)
		}
		return 
	}
	if x == info.Length - 1 {
		dfs(info, y+1, 0)
		return
	}
	cand := []int{}
	for i, p := range info.Pieces {
		if info.Used[i] {
			continue
		}
		if info.Board.IsFit(p, y, x) {
			cand = append(cand, i)
		}
	}
	if len(cand) == 0 {
		dfs(info, y, x+1)
	}
	for _, idx := range cand {
		info.Board.Set(info.Pieces[idx], byte('A'+idx), y, x)
		info.Used[idx] = true
		dfs(info, y, x+1)
		info.Board.Set(info.Pieces[idx], byte('.'), y, x)
		info.Used[idx] = false
	}
}

func Solve(info *Info) {
	dfs(info, 0, 0)
	printBoard(info)
}
