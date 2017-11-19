package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var exampleWorkout = `
title: Reddit body weight fitness
exercises:
  - instruction: band dislocations
    number: 30
    isTimed: true
  - instruction: pushups
    number: 50
    isTimed: false
`

type exercise struct {
	Instruction string
	Number      int
	IsTimed     bool
}

type workoutplan struct {
	Title     string
	Exercises []exercise
}

// Parses a workout YAML file and returns the plan and
// an error.
func readInputFile(infile string) (workoutplan, error) {
	myplan := workoutplan{}
	dat, err := ioutil.ReadFile(infile)
	if err != nil {
		return myplan, err
	}
	err = yaml.Unmarshal([]byte(dat), &myplan)
	return myplan, err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Input file is missing.")
	}

	myplan, err := readInputFile(os.Args[1])
	if err != nil {
		log.Fatalf("Bad YAML input: error: %v", err)
	}
	fmt.Printf("--- myplan:\n%v\n\n", myplan)

}
