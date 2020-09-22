package main

import (
	"fmt"

	"morris"
)

func main() {
	g := morris.NewGame("foo", "bar")
	for {
		// read input
		// g.Move()
		if g.IsFinished() {
			fmt.Println("Done, player %v:%v won!", g.Last().Name, g.Last().User)
			break
		}
	}
}
