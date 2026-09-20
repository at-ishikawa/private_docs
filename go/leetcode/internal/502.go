package internal

import (
	"container/heap"
	"sort"
)

// Brute force:
// Check the current capital => Pick up one => calculate next capital
//   - Ignore the profit = 0
//
// Apply the above recursively
//   - We can store the information and apply next.
//   - From the first to the next, choose which one is picked up or not
//   - maxCapital(i, k, w) = max(maxCapital(i-1, k-1, w+capital[i]), maxCapital(i-1, k, w))

type MaxIntHeap []int

var _ heap.Interface = (*MaxIntHeap)(nil)

// Len implements heap.Interface.
func (m MaxIntHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MaxIntHeap) Less(i int, j int) bool {
	return m[i] > m[j]
}

// Pop implements heap.Interface.
func (m *MaxIntHeap) Pop() any {
	n := len(*m)
	profit := (*m)[n-1]
	*m = (*m)[:n-1]
	return profit
}

// Push implements heap.Interface.
func (m *MaxIntHeap) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m *MaxIntHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	sorts := make([][2]int, n)
	for i := 0; i < n; i++ {
		sorts[i] = [2]int{
			capital[i], profits[i],
		}
	}
	sort.Slice(sorts, func(i, j int) bool {
		return sorts[i][0] < sorts[j][0]
	})

	queue := &MaxIntHeap{}
	heap.Init(queue)
	index := 0
	for ; k > 0; k-- {
		for ; index < n && sorts[index][0] <= w; index++ {
			heap.Push(queue, sorts[index][1])
		}
		if queue.Len() == 0 {
			break
		}

		profit := heap.Pop(queue).(int)
		w += profit
	}
	return w
}

func sortByCapitals(profits []int, capital []int) {
	// should be by a quick sort for the optimal solution
	n := len(capital)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if capital[i] > capital[j] {
				capital[i], capital[j] = capital[j], capital[i]
				profits[i], profits[j] = profits[j], profits[i]
			}
		}
	}
}

var maxProfits [][][2]int

func findMaximizedCapital1(k int, w int, profits []int, capital []int) int {
	sortByCapitals(profits, capital)
	maxProfits = make([][][2]int, len(profits))
	for i := len(profits) - 1; i >= 0; i-- {
		maxProfits[i] = make([][2]int, k+1)
	}
	return w + searchMaxProfit(k, w, profits, capital, 0, 0)
}

func searchMaxProfit(k int, w int, profits []int, capital []int, currentProfit int, index int) int {
	if index >= len(profits) {
		return currentProfit
	}
	if k == 0 {
		return currentProfit
	}
	array := maxProfits[index][k]
	maxCurrentProfit, result := array[0], array[1]
	if currentProfit < maxCurrentProfit {
		return result
	}

	var maxProfit int
	if capital[index] <= (w + currentProfit) {
		maxProfit = searchMaxProfit(k, w, profits, capital, currentProfit, index+1)
		maxProfit = max(maxProfit, searchMaxProfit(k-1, w, profits, capital, currentProfit+profits[index], index+1))
	} else {
		// capital is sorted, so if it's not possible to do a project, it cannot work on following projects anymore
		maxProfit = currentProfit
	}

	maxProfits[index][k] = [2]int{currentProfit, maxProfit}
	return maxProfit
}
