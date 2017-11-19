package main

import (
	"fmt"
	"time"
)

type exercise struct {
	Instruction string
	Number      int
	IsTimed     bool
}

type workoutplan struct {
	Title     string
	Exercises []exercise
}

func (ex *exercise) getBeginAnnouncement() string {
	units := "repps and hit any key when finished"
	if ex.IsTimed {
		units = "seconds"
	}
	template := "Now doing %s. Do %s for %d %s"
	return fmt.Sprintf(template, ex.Instruction, ex.Instruction, ex.Number, units)
}

func (ex *exercise) getEndAnnouncement() string {
	return fmt.Sprintf("Done with %s", ex.Instruction)
}

func (ex *exercise) begin(s speaker) error {
	return s.say(ex.getBeginAnnouncement())
}

func (ex *exercise) end(s speaker) error {
	return s.say(ex.getEndAnnouncement())
}

func (ex *exercise) waitForCompletion() {
	if ex.IsTimed == true {
		time.Sleep(time.Duration(ex.Number) * time.Second)
	} else {
		getChar()
	}
}

func (w *workoutplan) begin(s speaker) {
	for _, ex := range w.Exercises {
		ex.begin(s)
		ex.waitForCompletion()
		ex.end(s)
	}
}
