
package main

import (
    "fmt"
)

const (
    N = 10
    steps = 1000
    alpha = 0.01
    dt = 0.01
    dx = 1.0
)


func main (){
    //initialize the temperature grid
    grid := [N][N]float64{}
    nextGrid := [N][N]float64{}

    //set initial conditions (hot spot in the center)
    grid[N/2][N/2] = 100.0


    //fdm loop
    for t:=0; t<steps; t++ {
        for i:=1; i<N-1; i++ {
            for j:=1; j<N-1; j++ {
                nextGrid[i][j] = grid[i][j] + alpha*dt*(
                    (grid[i+1][j] - 2*grid[i][j] + grid[i-1][j])/(dx*dx) +
                    (grid[i][j+1] - 2*grid[i][j] + grid[i][j-1])/(dx*dx))
            }
        }
        //swap grids
        grid, nextGrid = nextGrid, grid
    }

    //print the final temperature distribution
    for i:=0; i<N; i++ {
        for j:=0; j<N; j++ { 
            fmt.Printf("%6.2f ", grid[i][j])
        }
        fmt.Println()
    }
}
