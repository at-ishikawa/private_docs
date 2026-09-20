package internal

import (
	"container/heap"
	"fmt"
)

type MinIntHeap []int

// Len implements heap.Interface.
func (m MinIntHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MinIntHeap) Less(i int, j int) bool {
	return m[i] < m[j]
}

// Pop implements heap.Interface.
func (m *MinIntHeap) Pop() any {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

// Push implements heap.Interface.
func (m *MinIntHeap) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m *MinIntHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

var _ heap.Interface = (*MinIntHeap)(nil)

type MedianFinder = MedianFinder3

func Constructor() MedianFinder {
	lower := make(MinIntHeap, 0)
	heap.Init(&lower)
	upper := make(MinIntHeap, 0)
	heap.Init(&upper)
	return MedianFinder3{
		lower: &lower,
		upper: &upper,
	}
}

type MedianFinder3 struct {
	lower *MinIntHeap
	upper *MinIntHeap
}

func (this *MedianFinder3) AddNum(num int) {
	if this.lower.Len() > this.upper.Len() {
		heap.Push(this.lower, -num)
		val := heap.Pop(this.lower).(int)
		heap.Push(this.upper, -val)
		return
	}

	heap.Push(this.upper, num)
	val := heap.Pop(this.upper).(int)
	heap.Push(this.lower, -val)
}
func (this *MedianFinder) peek(h heap.Interface) int {
	if h.Len() == 0 {
		return 0
	}

	x := heap.Pop(h).(int)
	heap.Push(h, x)
	return x
}

func (this *MedianFinder) FindMedian() float64 {
	low := this.peek(this.lower)
	high := this.peek(this.upper)
	fmt.Printf("(%d, %d)\n", -low, high)
	if (this.lower.Len()+this.upper.Len())%2 == 1 {
		return float64(-low)
	}

	return float64(-low+high) / 2
}

// func (this *MedianFinder3) AddNum(num int) {
// 	lowerVal := math.MinInt
// 	if this.lower.Len() > 0 {
// 		lowerVal = heap.Pop(this.lower).(int)
// 	}
// 	upperVal := math.MaxInt
// 	if this.upper.Len() > 0 {
// 		upperVal = heap.Pop(this.upper).(int)
// 	}

// 	// fmt.Printf("%d: (%d, %d)\n", num, lowerVal, upperVal)
// 	if num < lowerVal {
// 		// go to a lowerVal
// 		heap.Push(this.lower, num)

// 		heap.Push(this.upper, upperVal)
// 		if this.upper.Len() >= this.lower.Len() {
// 			heap.Push(this.lower, lowerVal)
// 		} else {
// 			heap.Push(this.upper, lowerVal)
// 		}
// 		return
// 	}
// 	if num > upperVal {
// 		heap.Push(this.upper, num)
// 		heap.Push(this.lower, lowerVal)

// 		if this.upper.Len() >= this.lower.Len() {
// 			heap.Push(this.lower, upperVal)
// 		} else {
// 			heap.Push(this.upper, upperVal)
// 		}
// 		return
// 	}

// 	// the num is equal to either lowerVal or upperVal
// 	heap.Push(this.lower, lowerVal)
// 	heap.Push(this.upper, upperVal)
// 	if this.lower.Len() > this.upper.Len() {
// 		heap.Push(this.upper, num)
// 	} else {
// 		heap.Push(this.lower, num)
// 	}
// }

// type MedianFinder = MedianFinder2

// func Constructor() MedianFinder {
// 	return MedianFinder2{}
// }

// // What if all numbers are from 0-100
// type MedianFinder2 struct {
// 	histogram  [101]int // index => count
// 	totalCount int
// }

// func (this *MedianFinder2) AddNum(num int) {
// 	this.histogram[num]++
// 	this.totalCount++
// }

// func (this *MedianFinder2) FindMedian() float64 {
// 	medianCount := (this.totalCount + 1) / 2
// 	count := 0
// 	var num int
// 	var numCount int
// 	for num, numCount = range this.histogram {
// 		count += numCount
// 		if count >= medianCount {
// 			break
// 		}
// 	}
// 	if count == medianCount && this.totalCount%2 == 0 {
// 		return float64(num+num+1) / 2
// 	}

// 	return float64(num)
// }

// type MedianFinder = MedianFinder1

// func Constructor() MedianFinder {
// 	return MedianFinder1{
// 		nums: make([]int, 0),
// 	}
// }

// type MedianFinder1 struct {
// 	// Any numbers
// 	nums []int // sorted
// }

// func (this *MedianFinder1) AddNum(num int) {
// 	// can replace this with a binary search
// 	var i int
// 	for i = 0; i < len(this.nums); i++ {
// 		if i == 0 {
// 			if num < this.nums[i] {
// 				break
// 			}
// 			continue
// 		}

// 		if this.nums[i-1] <= num && num <= this.nums[i] {
// 			break
// 		}
// 	}

// 	var former []int
// 	if i > 0 {
// 		former = append([]int{}, this.nums[:i]...)
// 	} else {
// 		former = make([]int, 0)
// 	}
// 	former = append(former, num)

// 	var later []int
// 	if len(this.nums)-i > 0 {
// 		later = this.nums[i:]
// 	}
// 	this.nums = append(former, later...)
// }

// func (this *MedianFinder1) FindMedian() float64 {
// 	middle := len(this.nums) / 2
// 	if len(this.nums)%2 == 1 {
// 		// 1, 2, 3: len 3 => index 1
// 		return float64(this.nums[middle])
// 	}
// 	// 1, 2, 3, 4: len 4 => index 2
// 	return float64(this.nums[middle-1]+this.nums[middle]) / 2.0
// }
