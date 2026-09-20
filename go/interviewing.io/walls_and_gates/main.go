package main

import (
	"fmt"
	"reflect"
)

const INF = 2<<30 - 1

func main() {
	testCases := []struct {
		name  string
		rooms [][]int
		want  [][]int
	}{
		{
			name: "example",
			rooms: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
	}
	for _, tc := range testCases {
		wallsAndGates(tc.rooms)
		got := tc.rooms
		if reflect.DeepEqual(tc.want, got) {
			fmt.Printf("Test %s: Passed\n", tc.name)
		} else {
			fmt.Printf("Test %s: Failed. Want: %+v, Got: %+v\n", tc.name, tc.want, got)
		}
	}
}

type Point struct {
	x     int
	y     int
	steps int
}

func wallsAndGates(rooms [][]int) {
	queue := make([]Point, 0)
	for y, row := range rooms {
		for x, room := range row {
			if room != 0 {
				continue
			}
			queue = append(queue, Point{
				x:     x,
				y:     y,
				steps: 0,
			})
		}
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]

		x := room.x
		y := room.y
		if y < 0 || y >= len(rooms) || x < 0 || x >= len(rooms[0]) {
			continue
		}

		if rooms[y][x] == INF {
			rooms[y][x] = room.steps
		} else if rooms[y][x] != 0 {
			continue
		}

		for _, direction := range directions {
			nextX := x + direction[0]
			nextY := y + direction[1]

			queue = append(queue, Point{
				x:     nextX,
				y:     nextY,
				steps: room.steps + 1,
			})
		}
	}
}

/*
func wallsAndGates(rooms [][]int) {
	for y, row := range rooms {
		for x, room := range row {
			if room != 0 {
				continue
			}

			bfs(rooms, x+1, y, 1)
			bfs(rooms, x-1, y, 1)
			bfs(rooms, x, y-1, 1)
			bfs(rooms, x, y+1, 1)
		}
	}
}

// This is dfs, not bfs
func bfs(rooms [][]int, x, y int, steps int) {
	if y < 0 || y >= len(rooms) || x < 0 || x >= len(rooms[0]) {
		return
	}
	room := rooms[y][x]
	if room == -1 || room == 0 {
		return
	}
	if room < steps {
		return
	}
	rooms[y][x] = steps
	steps++

	bfs(rooms, x+1, y, steps)
	bfs(rooms, x-1, y, steps)
	bfs(rooms, x, y+1, steps)
	bfs(rooms, x, y-1, steps)
}
*/
