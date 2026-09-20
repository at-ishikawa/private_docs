package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		name     string
		meetings [][]int
		want     int
	}{
		{
			name: "example 1",
			meetings: [][]int{
				{5, 10},
				{2, 3},
			},
			want: 1,
		},
		{
			name: "example 2",
			meetings: [][]int{
				{1, 3},
				{5, 7},
				{4, 6},
				{7, 9},
				{9, 10},
			},
			want: 2,
		},
		{
			name: "example 3",
			meetings: [][]int{
				{7, 9},
				{1, 3},
				{7, 10},
				{4, 6},
				{8, 10},
			},
			want: 3,
		},
		{
			name: "example 4",
			meetings: [][]int{
				{1, 2},
				{1, 3},
				{1, 6},
				{2, 6},
				{3, 6},
				{4, 6},
				{6, 7},
				{6, 10},
				{7, 8},
				{8, 9},
				{9, 10},
			},
			want: 4,
		},
	}

	for _, tc := range testCases {
		got := getMinimumMeetingRooms(tc.meetings)
		if tc.want != got {
			fmt.Printf("Test %s: Failed. Want %d, Got %d\n", tc.name, tc.want, got)
		} else {
			fmt.Printf("Test %s: Passed\n", tc.name)
		}
	}
}

type TimeRange struct {
	startTime int
	endTime   int
}

type PriorityQueue []TimeRange

// Len implements heap.Interface.
func (m PriorityQueue) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m PriorityQueue) Less(i int, j int) bool {
	return m[i].endTime < m[j].endTime
}

// Pop implements heap.Interface.
func (m *PriorityQueue) Pop() any {
	endIndex := len(*m) - 1
	result := (*m)[endIndex]
	*m = (*m)[:endIndex]
	return result
}

// Push implements heap.Interface.
func (m *PriorityQueue) Push(x any) {
	*m = append(*m, x.(TimeRange))
}

// Swap implements heap.Interface.
func (m PriorityQueue) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*PriorityQueue)(nil)

// https://interviewing.io/questions/meeting-rooms
func getMinimumMeetingRooms(meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	rooms := make(PriorityQueue, 0)
	heap.Init(&rooms)
	heap.Push(&rooms, TimeRange{
		startTime: meetings[0][0],
		endTime:   meetings[0][1],
	})
	var answer PriorityQueue

	for _, meeting := range meetings[1:] {
		for rooms.Len() > 0 {
			top := heap.Pop(&rooms).(TimeRange)
			if !(top.endTime <= meeting[0]) {
				heap.Push(&rooms, top)
				break
			}
		}
		heap.Push(&rooms, TimeRange{
			startTime: meeting[0],
			endTime:   meeting[1],
		})
		if answer.Len() < rooms.Len() {
			answer = make(PriorityQueue, rooms.Len())
			for i, room := range rooms {
				answer[i] = room
			}
		}
		// fmt.Printf("Meeting: %+v, Top: %#v, Heap: %#v\n", meeting, top, rooms)
	}

	fmt.Printf("answer: %+v\n", answer)
	return len(answer)
}

/*
func getMinimumMeetingRooms(meetings [][]int) int {
	startTimes := make([]int, len(meetings))
	endTimes := make([]int, len(meetings))

	for i, meeting := range meetings {
		startTimes[i] = meeting[0]
		endTimes[i] = meeting[1]
	}

	sort.Slice(startTimes, func(i, j int) bool {
		return startTimes[i] < startTimes[j]
	})
	sort.Slice(endTimes, func(i, j int) bool {
		return endTimes[i] < endTimes[j]
	})

	roomCount := 0
	// endIndex := 0

	//	for startIndex := 0; startIndex < len(startTimes); startIndex++ {
	//		// fmt.Printf("(%d, %d): (%d, %d), %d\n", startIndex, endIndex, startTimes[startIndex], endTimes[endIndex], roomCount)
	//		if startTimes[startIndex] < endTimes[endIndex] {
	//			roomCount++
	//		} else {
	//			// Why does when a startIndex is incremented at the same time?
	//			endIndex++
	//		}
	//	}

	startIndex := 0
	for endIndex := 0; endIndex < len(endTimes); endIndex++ {
		count := 0
		for ; startIndex < len(startTimes) && startTimes[startIndex] < endTimes[endIndex]; startIndex++ {
			count++
		}
		if count > roomCount {
			roomCount = count
		}
	}
	return roomCount
}
*/

// Time Complexity: O(n * log(n)) by Sort
// Space complexity: O(2n) => O(n)

/*
func getMinimumMeetingRooms(meetings [][]int) int {
	maxCount := 0
	for hour := 0; hour < 24; hour++ {
		count := 0
		for _, meeting := range meetings {
			if len(meeting) != 2 {
				log.Fatalf("Failed: Meeting should have: %v\n", meeting)
				continue
			}
			if meeting[0] <= hour && hour < meeting[1] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
*/
