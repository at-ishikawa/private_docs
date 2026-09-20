package internal

import (
	"math"
)

func maxPoints(points [][]int) int {
	return maxPoints2(points)
}

func maxPoints2(points [][]int) int {
	var result int
	for i := 0; i < len(points); i++ {
		slopes := make(map[float64]int)

		samePointCount := 1
		basePoint := points[i]
		for j := i + 1; j < len(points); j++ {
			x := points[j][0]
			y := points[j][1]
			if basePoint[0] == x {
				slopes[math.MaxInt]++
				continue
			} else if basePoint[0] == x && basePoint[0] == y {
				samePointCount++
				continue
			}
			slope := float64(y-basePoint[1]) / float64(x-basePoint[0])
			slopes[slope]++
		}

		var maxCount int
		for _, cnt := range slopes {
			maxCount = max(cnt, maxCount)
		}
		result = max(result, maxCount+samePointCount)
	}
	return result
}

// [x1, y1]: [x2, y2]
// The line: y=ax+b
//
//	y1=ax1+b
//	y2=ax2+b
//	y1-y2=ax1-ax2
//	a=(y1-y2)/(x1-x2)
//	b=>y1-ax1
//	a=1/2, b=1/2
//
// check i (0 <= i <= n): [xi, yi]: check axi+b <=> yi
func maxPoints1(points [][]int) int {
	maxCount := 1
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			point1 := points[i]
			point2 := points[j]

			count := 2

			// line is x=xi
			if point1[0]-point2[0] == 0 {
				for k := 0; k < len(points); k++ {
					if k == i || k == j {
						continue
					}
					if points[k][0] == point1[0] {
						count++
					}
				}
			} else {
				// y=ax+b
				a := float64(point1[1]-point2[1]) / float64(point1[0]-point2[0])
				b := float64(point1[1]) - a*float64(point1[0])

				for k := 0; k < len(points); k++ {
					if k == i || k == j {
						continue
					}
					y := a*float64(points[k][0]) + b
					if math.Abs(y-float64(points[k][1])) < (1e-9) {
						count++
					}
				}
			}

			if count > maxCount {
				maxCount = count
			}
		}
	}

	return maxCount
}
