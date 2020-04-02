package main

/*
1. 活细胞周围八个位置的活细胞数少于两个，该位置活细胞死亡；
2. 活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
3. 活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
4. 死细胞周围正好有三个活细胞，则该位置死细胞复活；
*/
import "fmt"

type tuple struct {
	x int
	y int
}

func gameOfLife(board [][]int) {
	rows := len(board)
	if rows == 0{return}
	cols := len(board[0])
	store := make(map[tuple]int)
	for r, arr := range board {
		for c, v := range arr {
			liveCells := 0
			cellAround := []tuple{
				tuple{r - 1, c - 1},
				tuple{r - 1, c},
				tuple{r - 1, c + 1},
				tuple{r, c - 1},
				tuple{r, c + 1},
				tuple{r + 1, c - 1},
				tuple{r + 1, c},
				tuple{r + 1, c + 1},
			}
			for _, around := range cellAround {
				if around.x < 0 || around.y < 0 || around.x >= rows || around.y >= cols{
					continue
				}
				if board[around.x][around.y] == 1 {
					liveCells++
				}
			}
			curCellPos := tuple{r, c}
			if v == 1{
				if liveCells<2 || liveCells>3{
					store[curCellPos] = 0
				}
			} else {
				if liveCells == 3{
					store[curCellPos] = 1
				}
			}
		}
	}
	for i, v:=range store{
		board[i.x][i.y] = v
	}
}
func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Println("before: ", board)
	gameOfLife(board)
	fmt.Println("after: ", board)

}
