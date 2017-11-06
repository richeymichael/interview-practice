package main
/*
Write a Golang program to Simulate a Gas Station:
- You have 4 Gas Pumps
- You have 10 Cars
- Each car pulls up to an available pump for a 50 millisecond fill up, once full it gets back in line an waits for an available pump
- Run the simulation for 30 seconds
- Report at the end of the simulation the number of fill ups given to each car, and provided by each pump.

Please make good use of go funcs and channels.

Assumptions: 
   1) Simulation begins when pumps start reading from the line chan, so we
      can put some cars in the line first
   2) When 30 second timer goes off, cars that are in the process of filling
      up are allowed to continue
*/

import (
    "fmt"
    "time"
)

const (
    number_of_pumps int = 4
    number_of_cars int = 10
    fill_up_time_ms int64 = 50
    simulation_time_s int64 = 30
)

type Car struct {
    id int
    fill_up_count int
    fill_up_time_ms time.Duration
}

func pump(pumping <-chan *Car, line chan <-*Car, pump *int) {
    for {
        car := <- pumping
        time.Sleep(car.fill_up_time_ms)
        car.fill_up_count++
        *pump++
        line <- car
    }
}

func run_simulation(pumping chan <-*Car, line <-chan *Car) {
    // build the timer
    timer := time.NewTimer(time.Duration(simulation_time_s) * time.Second)
    defer timer.Stop()

    // if we have time move available car from line into pumping
    LOOP:
    for {
        select {
            case <- timer.C:
                close(pumping)
                break LOOP
            default:
                car := <- line
                pumping <- car
        }
    }
}

func display_output(cars [number_of_cars]Car, counters [number_of_pumps]int) {
    for i, _ := range cars {
        fmt.Printf("Car Id: %d, number of fill ups: %d\n", cars[i].id, cars[i].fill_up_count)
    }
    for i, _ := range counters {
        fmt.Printf("Pump Id: %d, number of fill ups: %d\n", i, counters[i])
    }
}

func main () {
    // Create the array of cars
    var cars [number_of_cars]Car
    for i, _ := range cars {
        cars[i] = Car{i, 0, time.Duration(fill_up_time_ms) * time.Millisecond}
    }

    // Create communication channels 
    pumping := make(chan *Car, number_of_pumps)
    line := make(chan *Car, number_of_cars)

    // Create the pump workers
    var counters [number_of_pumps]int
    for i:=0; i<number_of_pumps; i++ {
        go pump(pumping, line, &counters[i])
    }

    // initialize the line with all the cars
    for i, _ := range cars {
        line <- &cars[i]
    }

    run_simulation(pumping, line)
    display_output(cars, counters)
}
