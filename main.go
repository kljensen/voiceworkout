package main

import (
	"fmt"
	"log"
	"os"
)

const version = "v1.3"

func main() {
	fmt.Printf("voiceworkout %s\n", version)
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
