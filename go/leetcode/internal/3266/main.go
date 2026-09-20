// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-ii/description/
package main

import (
	"container/heap"
)

// use a priority queue to store the value and index
// modulo should be done by each multiplier, it's likely have the same result, especially 10^9 + 7 is a prime number

type Item struct {
	num   int
	index int
}
type PriorityQueue []Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i int, j int) bool {
	if p[i].num == p[j].num {
		return p[i].index < p[j].index
	}
	return p[i].num < p[j].num
}

func (p *PriorityQueue) Pop() any {
	n := len(*p)
	val := (*p)[n-1]
	(*p) = (*p)[:n-1]
	return val
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(Item))
}

func (p *PriorityQueue) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

var _ heap.Interface = (*PriorityQueue)(nil)

// Exponentiation by squaring
// https://www.geeksforgeeks.org/exponential-squaring-fast-modulo-multiplication/
func exponentiation(base int, exp int, modulo int) int {
	// x^n = x^(2*(n/2)) (n = even)
	// x^n = x*x^(2*(n/2)) = x*x^2*x^(n/2) = x * x^2*x^(1) (n = odd)

	// 2*4
	// exp=4: 2 * 2 => 4 (e)
	// exp=2: 4 * 4 => 16 (e)
	// exp=1: 1 * 16 => 16 (result)
	// 2^5
	// exp=6: 2 * 2 => 4(e)
	// exp=3: result=4, e=16
	// exp=1: result=64
	result := 1
	for ; exp >= 1; exp /= 2 {
		if exp%2 == 1 {
			result = (result * base) % modulo
		}
		base = (base * base) % modulo
	}
	return result
}

func getFinalState(nums []int, k int, multiplier int) []int {
	// ignore edge case where multiplier = 1
	if multiplier == 1 {
		return nums
	}

	n := len(nums)

	queue := &PriorityQueue{}
	heap.Init(queue)
	for index, num := range nums {
		heap.Push(queue, Item{
			num:   num,
			index: index,
		})
	}

	const modulo = 1e9 + 7

	// Check how many times multiplier have to apply on each num in order to multiply for all elements
	isMultiplied := make(map[int]bool, 0)
	for ; k > 0 && len(isMultiplied) < n; k-- {
		val := heap.Pop(queue).(Item)
		// val.num = (val.num * multiplier) % modulo
		val.num = val.num * multiplier
		isMultiplied[val.index] = true
		heap.Push(queue, val)
	}

	repeatCount := k / n
	remainingNumCount := k % n

	result := make([]int, n)
	for queue.Len() > 0 {
		item := heap.Pop(queue).(Item)
		exp := repeatCount
		if remainingNumCount > 0 {
			exp++
			remainingNumCount--
		}

		result[item.index] = (item.num % modulo * exponentiation(multiplier, exp, modulo)) % modulo
	}
	return result
}
