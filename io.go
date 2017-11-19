package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Parses a workout YAML file and returns the plan and
// an error.
func newWorkoutPlanFromFile(infile string) (workoutplan, error) {
	myplan := workoutplan{}
	dat, err := ioutil.ReadFile(infile)
	if err != nil {
		return myplan, err
	}
	err = yaml.Unmarshal([]byte(dat), &myplan)
	if err != nil {
		return myplan, err
	}
	return myplan, err
}
