package main

import (
	"fmt"
	"math"
)

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	i := "Welcome to my party, " + name + "!\n"
	h := fmt.Sprintf("%X", table)
	k := "You have been assigned to table " + h + "."
	n := " Your table is " + direction + ", "
	dis := float64(int(distance*100)) / 100
	r := Round(dis, 0.05)
	d := fmt.Sprint(r)
	m := "exactly " + d + " meters from here.\n"
	b := "You will be sitting next to " + neighbour + "."
	return i + k + n + m + b
	// panic("Please implement the AssignTable function")
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	a := !prisonerIsAwake && !archerIsAwake && !petDogIsPresent
	b := petDogIsPresent || !archerIsAwake
	if a || b {
		return true
	}
	return false
}

func main() {
	// e := AssignTable("alex", 27, "zieme", "on the left", 23.7834298)
	// fmt.Println(e)
	e := CanFreePrisoner(true, true, true, false)
	fmt.Println(e)
}
