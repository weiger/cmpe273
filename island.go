package main

import "fmt"

func numberOfIsland(grid *[4][5]int64) int{

    if len(grid) == 0{
	return 0
    }	

    count := 0

    for i := 0; i < len(grid); i++{
	for j := 0; j < len(grid[i]); j++{
		if grid[i][j] == 1{
			count++
			dfs(grid, i, j)
		}
	}
    }

    return count
}

func dfs(grid *[4][5]int64, i int, j int){
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]){
	return 
    }

    if grid[i][j] == 0{
	return 
    }

    grid[i][j] = 0

    dfs(grid,i+1,j)
    dfs(grid,i-1,j)
    dfs(grid,i,j+1)
    dfs(grid,i,j-1)
}

func main(){
    grid := [4][5]int64{{1,1,1,1,0},{1,1,0,1,0},{1,1,0,0,0},{0,0,0,0,0}}
    grid2 := [4][5]int64{{1,1,0,0,0},{1,1,0,0,0},{0,0,1,0,0},{0,0,0,1,1}}

    fmt.Println(numberOfIsland(&grid))
    fmt.Println(numberOfIsland(&grid2))
}