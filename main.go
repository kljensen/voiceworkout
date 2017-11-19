package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Input file is missing.")
	}

	myplan, err := newWorkoutPlanFromFile(os.Args[1])
	if err != nil {
		log.Fatalf("Bad YAML input: error: %v", err)
	}

	s := saySpeaker{}
	myplan.begin(s)
}
