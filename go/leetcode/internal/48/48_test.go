package internal48

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "happy path",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "example 2",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	copy2DInt := func(matrix [][]int) [][]int {
		result := make([][]int, len(matrix))
		for i := 0; i < len(matrix); i++ {
			result[i] = make([]int, len(matrix[0]))
			for j := 0; j < len(matrix[0]); j++ {
				result[i][j] = matrix[i][j]
			}
		}
		return result
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("rotate1", func(t *testing.T) {
				got := copy2DInt(tc.matrix)
				require.Equal(t, tc.matrix, got)

				rotate1(got)
				assert.Equal(t, tc.want, got)
			})

			t.Run("rotate2", func(t *testing.T) {
				got := copy2DInt(tc.matrix)
				require.Equal(t, tc.matrix, got)

				rotate2(got)
				assert.Equal(t, tc.want, got)
			})
		})
	}
}
