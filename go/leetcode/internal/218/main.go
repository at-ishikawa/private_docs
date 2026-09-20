package main

import (
	"container/heap"
	"slices"
	"sort"
)

type Queue []int

var _ heap.Interface = (*Queue)(nil)

// Len implements heap.Interface.
func (q Queue) Len() int {
	return len(q)
}

// Less implements heap.Interface.
func (q Queue) Less(i int, j int) bool {
	return q[i] > q[j]
}

// Pop implements heap.Interface.
func (q *Queue) Pop() any {
	n := len(*q)
	val := (*q)[n-1]
	*q = (*q)[:n-1]
	return val
}

// Push implements heap.Interface.
func (q *Queue) Push(x any) {
	*q = append(*q, x.(int))
}

// Swap implements heap.Interface.
func (q *Queue) Swap(i int, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func getSkyline(buildings [][]int) [][]int {
	// divide buildings the positions of each left and right.
	// On the right side is set as a negative value for a height
	points := make([][]int, 0)
	for _, building := range buildings {
		points = append(points,
			[]int{building[0], building[2]},
			[]int{building[1], -building[2]},
		)
	}
	// order by position X asc, and height desc
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] > points[j][1]
	})

	var pq Queue
	heap.Init(&pq)
	heap.Push(&pq, 0)

	currentHeight := 0
	skyline := make([][]int, 0)
	for _, point := range points {
		x, height := point[0], point[1]
		if height < 0 {
			heap.Remove(&pq, slices.Index(pq, -height))
		} else {
			heap.Push(&pq, height)
		}

		// In order to get tne tallest building at any point, use a priority queue
		// highestHeight := heap.Pop(&pq).(int)
		// heap.Push(&pq, highestHeight)
		highestHeight := pq[0]

		// Push an updated height to the result. This assumes that points slice is sorted by the height
		// such that shorter building in the same position isn't pushed
		if highestHeight != currentHeight {
			currentHeight = highestHeight
			skyline = append(skyline, []int{
				x, highestHeight,
			})
		}
	}
	return skyline
}

/*
import "sort"


// buildings patterns
//   1. One building is a inner one of another. This should be ignored
//   1. There can be a gap between buildings
//   1. Multiple adjacent buildings can be same heights
//   1. Shorter building's height will be ignored if there is a taller one in the same X

// Follow from the first building: (left, height)
// ignore as long as height is the same
// heigher building: (left, height)
// shorter building: (previosu left, current height)
// not adjacent: (right, 0)
func getSkyline(buildings [][]int) [][]int {
	// buildings are sorted by left
	var previousBuilding []int
	currentHeight := 0
	result := make([][]int, 0)

	isAdjacent := func(previousBuilding []int, currentBuilding []int) bool {
		// currentBuilding.left < previousBuilding.right
		return currentBuilding[0] <= previousBuilding[1]
	}

	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i][0] != buildings[j][0] {
			return buildings[i][0] < buildings[j][0]
		}
		if buildings[i][2] == buildings[j][2] {
			// order by right desc if the height is the same
			return buildings[i][1] > buildings[j][1]
		}
		// order by height desc
		return buildings[i][2] > buildings[j][2]
	})

	for _, building := range buildings {
		left, _, height := building[0], building[1], building[2]

		if previousBuilding != nil && !isAdjacent(previousBuilding, building) {
			// if there is a gap between buildings, the previous building's right position on the earth should be the point
			result = append(result, []int{
				previousBuilding[1], 0,
			})
			currentHeight = 0
		}
		if currentHeight == 0 {
			// from the surface, always add the point on the height of a building
			result = append(result, []int{
				left, height,
			})
			currentHeight = height
			previousBuilding = building
			continue
		}
		if currentHeight == height {
			// no skyline point if the heights of buildings are the same
			// TODO: is the right value always bigger on the later element?
			previousBuilding = building
			continue
		}
		if currentHeight < height {
			result = append(result, []int{
				left, height,
			})
			currentHeight = height
			previousBuilding = building
			continue
		}
		if building[1] <= previousBuilding[1] {
			// if a building is included in the previous building
			// simply, skip this building
			continue
		}

		result = append(result, []int{
			previousBuilding[1], height,
		})
		currentHeight = height
		previousBuilding = building
	}

	result = append(result, []int{
		previousBuilding[1], 0,
	})
	return result
}
*/
