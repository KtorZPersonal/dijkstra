package dijkstra

type Walkable interface {
	SetVisited(i int, j int)
	SetGoal(i int, j int)
	SetActive(i int, j int)
	IsWalkable(i, j int) bool
	IsGoal(i, j int) bool
}

type Couple [2]int

func filter(xs [][2]Couple, f func(x [2]Couple) bool) [][2]Couple {
	fxs := make([][2]Couple, 0)

	if len(xs) == 0 {
		return fxs
	}

	for _, x := range xs {
		if f(x) {
			fxs = append(fxs, x)
		}
	}

	return fxs
}

func include(xs [][2]Couple, v [2]Couple) bool {
	for _, x := range xs {
		if x[0] == v[0] {
			return true
		}
	}

	return false
}

func Run(grid Walkable, start Couple) {
	closelist := make(map[Couple]Couple)
	openlist := [][2]Couple{[2]Couple{start, start}}

	for {
		n := len(openlist)
		if n == 0 {
			break
		}
		for _, c := range openlist {
			x := c[0]
			p := c[1]
			if x != start && grid.IsGoal(x[0], x[1]) {
				grid.SetGoal(x[0], x[1])
				for {
					grid.SetGoal(p[0], p[1])
					if p == start {
						return
					}
					t := p
					p = closelist[p]
					x = t
				}
			}
			grid.SetActive(x[0], x[1])
		}

		for i := 0; i < n; i++ {
			x := openlist[i][0]
			p := openlist[i][1]

			closelist[x] = p

			openlist = append(openlist, filter([][2]Couple{
				[2]Couple{Couple{x[0] - 1, x[1]}, x},
				[2]Couple{Couple{x[0] + 1, x[1]}, x},
				[2]Couple{Couple{x[0], x[1] + 1}, x},
				[2]Couple{Couple{x[0], x[1] - 1}, x},
			}, func(x [2]Couple) bool {
				return grid.IsWalkable(x[0][0], x[0][1]) && closelist[x[0]] == Couple{0, 0} && !include(openlist, x)
			})...)

			if !(x[0] == start[0] && x[1] == start[1]) {
				grid.SetVisited(x[0], x[1])
			}

		}
		openlist = openlist[n:]
	}
}
