// Quick Random Number Generator intended to be imported to other programs

package main

// First, we import the packages necessary for creating random numbers
// fmt to return values, math/rand to generate numbers, time to facilitate randomness
import (
	"fmt"
	"math/rand"
	"time"
)

func RandNum(start, end int) {
	// Creates a random number between the input numbers using the current time in nanoseconds as a seed

	/*
		The Go language needs to seed (choose a starting number) to generate random numbers
		However, each seed provides the same set of numbers, so we need a unique number each time
		Our solution for this is to use the current time, in nanoseconds, as our seed
	*/

	rand.Seed(time.Now().UnixNano())
	rNum := rand.Intn(end)
	fmt.Sprintln(rNum)
	// return rNum
}

func main() {
	fmt.Println(RandNum(10))
}
