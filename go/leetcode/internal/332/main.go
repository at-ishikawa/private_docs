package main

import (
	"slices"
	"sort"
)

var destinations map[string][]string
var route []string

func visit(airport string) {
	for {
		destinationList, ok := destinations[airport]
		if !ok || len(destinationList) == 0 {
			break
		}
		destination := destinationList[len(destinationList)-1]
		destinations[airport] = destinationList[:len(destinationList)-1]
		visit(destination)
	}
	route = append(route, airport)
}

func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] > tickets[j][1]
		}
		return tickets[i][0] > tickets[j][0]
	})

	destinations = make(map[string][]string, 0)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if _, ok := destinations[from]; !ok {
			destinations[from] = make([]string, 0)
		}
		destinations[from] = append(destinations[from], to)
	}

	route = make([]string, 0)
	visit("JFK")
	slices.Reverse(route)
	return route
}

/*
type Itinerary struct {
	airports []string
	count    map[string]int
}

// Create a graph
// Store the count from a destination to a to while traversing a graph to make sure it doesn't use the same ticket
// Pass an argument as a result while traversing to store the result
// and sort lexiographically and choose the one after all possible itineraries were returned
// sort didn't work for time limit
func searchItinerary(airport string, currentItinerary []string, allItinerary map[string]Itinerary, index int, tickets [][]string) []string {
	if index > len(tickets) {
		return currentItinerary
	}

	for _, to := range allItinerary[airport].airports {
		count := allItinerary[airport].count[to]
		if count == 0 {
			continue
		}

		currentItinerary[index] = to
		allItinerary[airport].count[to]--
		result := searchItinerary(to, currentItinerary, allItinerary, index+1, tickets)
		if len(result) > 0 {
			return result
		}
		allItinerary[airport].count[to]++
		currentItinerary[index] = ""
	}
	return nil
}

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string, 0)
	counts := make(map[string]map[string]int, 0)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if _, ok := graph[from]; !ok {
			graph[from] = make([]string, 0)
			counts[from] = make(map[string]int, 0)
		}
		if !slices.Contains(graph[from], to) {
			graph[from] = append(graph[from], to)
		}
		counts[from][to]++
	}
	for _, ticket := range tickets {
		from, _ := ticket[0], ticket[1]
		sort.Slice(graph[from], func(i, j int) bool {
			return graph[from][i] < graph[from][j]
		})
	}

	allItinerary := make(map[string]Itinerary)
	for _, ticket := range tickets {
		from, _ := ticket[0], ticket[1]
		if _, ok := allItinerary[from]; ok {
			continue
		}

		allItinerary[from] = Itinerary{
			airports: graph[from],
			count:    counts[from],
		}
	}

	// fmt.Printf("%+v\n", allItinerary)

	itinerary := make([]string, len(tickets)+1)
	startAirport := "JFK"
	itinerary[0] = startAirport

	return searchItinerary(
		startAirport,
		itinerary,
		allItinerary,
		1,
		tickets,
	)
}
*/
