package main

func CountIslands(grid [][]int) int{

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

func dfs(grid [][]int, i int, j int){
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


