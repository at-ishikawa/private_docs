package main

type Scorer struct {
  directReports map[int][]int
  scores map[int]int
}

func (scorer *Scorer) getScore(employeeId int) int {
  if score, ok := scorer.scores[employeeId]; ok {
    return score
  }

  count := 1
  for _, directReport := range scorer.directReports[employeeId] {
    for _, reportedEmployeeId := range directReport {
      count = count + scorer.getScore(reportedEmployeeId)
    }
  }
  scorer.scores[employeeId] = count
  return count
}

func main() {
  directReports = map[int][]int{
    123: []int{234, 345},
    234: []int{456, 789},
    345: nil,
    456: nil,
    789: nil,
  }
  scorer := Scorer{
    directReports: directReports,
    scores: make(map[int]int),
  }
  testCases := []struct{
    name string
    employeeId int
    want int
  } {
    {
      name: "no reporter",
      employeeId: 345,
      want: 1,
    },
    {
      name: "Employee has a direct reporter",
      employeeId: 234,
      want: 3,
    },
    {
      name: "Employee has an indirect reporter",
      employeeId: 123,
      want: 5,
    },
  }
  for _, tc := range testCases {
    got := scorer.getScore(tc.employeeId)
    if tc.want == got {
      fmt.Printf("Test %s: passed\n", tc.name)
    } else {
      fmt.Printf("Test %s: failed. Want: %d, Got: %d\n", tc.name,. tc.want, got)
    }
  }
}
