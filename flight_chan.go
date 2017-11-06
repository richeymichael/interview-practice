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

import (
    "fmt"
)

func find_origin(found_chan chan string, source chan string, legs map[string] string) {
    for {
        src := <- source //block
        found_bool := false
        for _, dst := range legs {
            if src == dst {
                found_bool = true
            }
        }
        if src != "" && ! found_bool {
            found_chan <- src
        }
    }
}

func display_trip(found chan string, legs map[string] string) {
    src := <-found
    for legs[src] != "" {
        fmt.Printf("%s -> %s\n", src, legs[src])
        src = legs[src]
    }
    close(found)
}

func main() {
    legs := map[string] string {
        "BOS": "MIA",
        "SFO": "BOS",
        "MIA": "LAX",
        "JFK": "ORD",
        "ORD": "SFO",
    }

    source := make(chan string, 3)
    found := make(chan string)

    for i:=0; i < 5; i++ {
        go find_origin(found, source, legs)
    }

    for src := range legs {
        source <- src
    }
    close(source)

    display_trip(found, legs) // block
}

