package main

import "fmt"

func findItinerary(tickets [][]string) []string {
	return nil
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))

	println("UseCase 2......")
	fmt.Printf("%+v\n", findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
