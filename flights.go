package main
// You are going on a one-way indirect flight trip at includes an unknown very
// large number of transfers
//
// You are not stopping twice in the same airport.
// You have 1 ticket for each part of your trip.
// Each ticket contains a src and dst airport.
// All the tickets you have are randomly sorted.
// You forgot the original departure airport and your final destination.
//
// Write a function to reconstruct your trip with minimum big-O complexity.
//
// original:
// JFK -> ORD
// BOS -> MIA
// SFO -> BOS
// MIA -> LAX
// ORD -> SFO

import "fmt"


func find(possible_origin string, legs map[string]string) bool{
    found := false
    for _, dst := range legs {
        if possible_origin == dst {
            found = true
        }
    }
    return found
}

// Find the origin of trip, it'll be the one that's never a destination
func determine_origin(legs map[string]string) string{
    for possible_origin := range legs {
        if ! find(possible_origin, legs) {
            return possible_origin
        }
    }
    return ""
}

/*
    start with source = origin
    get source's destination
    print source and destination as the leg
    set source = destination and start over
*/
func display_trip(origin string, legs map[string]string) {
    src := origin
    for legs[src] != "" {
       fmt.Printf("%s -> %s\n", src, legs[src])
       src = legs[src]
    }
}

// Main, set up legs variable get origin and display trip
func main() {
   legs := map[string] string {
       "BOS": "MIA",
       "SFO": "BOS",
       "MIA": "LAX",
       "ORD": "SFO",
       "JFK": "ORD",
   }
   origin := determine_origin(legs)
   display_trip(origin, legs)
}
