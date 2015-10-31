package main

import (
	"fmt"
    "github.com/ktorz/qmlgrid"
    "github.com/ktorz/dijkstra"
	"os"
    "time"
)

func main() {
    err := qmlgrid.New(30, 30, 20, func (grid *qmlgrid.Grid) error {
        grid.SetGoal(15, 15)
        grid.SetGoal(5, 3)
        grid.SetBlocked(4, 1)
        grid.SetBlocked(4, 2)
        grid.SetBlocked(4, 3)
        grid.SetBlocked(6, 0)
        grid.SetBlocked(6, 1)
        grid.SetBlocked(6, 2)
        grid.SetBlocked(6, 3)
        grid.SetBlocked(6, 4)
        grid.SetBlocked(6, 5)
        grid.SetBlocked(6, 6)
        grid.SetBlocked(6, 7)
        grid.SetBlocked(6, 8)
        grid.SetBlocked(6, 9)
        grid.SetBlocked(6, 10)
        grid.SetBlocked(6, 11)
        grid.SetBlocked(6, 12)
        grid.SetBlocked(6, 13)
        grid.SetBlocked(6, 14)
        grid.SetBlocked(5, 4)
        grid.SetBlocked(14, 16)
        grid.SetBlocked(14, 17)
        grid.SetBlocked(14, 18)
        grid.SetBlocked(14, 19)
        grid.SetBlocked(14, 20)
        grid.SetBlocked(14, 15)
        grid.SetBlocked(14, 14)
        grid.SetBlocked(14, 13)
        grid.SetBlocked(14, 12)
        grid.SetBlocked(14, 11)
        grid.SetBlocked(14, 10)
        grid.SetBlocked(14, 9)

        time.Sleep(time.Second)
        dijkstra.Run(grid, dijkstra.Couple{15, 15})
        return nil
    })

	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
