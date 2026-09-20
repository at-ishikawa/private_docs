package internal

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	type eachCase struct {
		nums []int
		want float64
	}
	testCases := []struct {
		name  string
		cases []eachCase
	}{
		{

			name: "example 1",
			cases: []eachCase{
				{nums: []int{1, 2}, want: 1.5},
				{nums: []int{3}, want: 2.0},
			},
		},
		{
			name: "failed 1",
			cases: []eachCase{
				{nums: []int{6}, want: 6.0},
				{nums: []int{10}, want: 8},
				{nums: []int{2}, want: 6},
				{nums: []int{6}, want: 6},
				{nums: []int{5}, want: 6},
				{nums: []int{0}, want: 5.5},
				{nums: []int{6}, want: 6},
				{nums: []int{3}, want: 5.5},
				{nums: []int{1}, want: 5},
				{nums: []int{0}, want: 4},
				{nums: []int{0}, want: 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			medianFinder := Constructor()
			for index, c := range tc.cases {
				t.Run("Case "+strconv.Itoa(index), func(t *testing.T) {
					for _, num := range c.nums {
						medianFinder.AddNum(num)
					}
					assert.Equal(t, c.want, medianFinder.FindMedian())
				})
			}
		})
	}
}
