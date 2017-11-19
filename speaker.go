package main

import (
	"fmt"
	"log"
	"os/exec"
)

type saySpeaker struct {
}

func (s saySpeaker) say(statement string) error {
	path, err := exec.LookPath("say")
	if err != nil {
		log.Fatal("No say command on your system")
	}
	fmt.Println(statement)
	cmd := exec.Command(path, statement)
	err = cmd.Run()
	return err
}

func (s saySpeaker) end(ex exercise) error {
	path, err := exec.LookPath("say")
	if err != nil {
		log.Fatal("No say command on your system")
	}
	fullInstructions := ex.getEndAnnouncement()
	cmd := exec.Command(path, fullInstructions)
	err = cmd.Run()
	return err
}

// Anything that has a `speak` method satisfies the `speaker` interface. In the
// future, I might want to use the Amazon Polly service for better voices.
type speaker interface {
	say(string) error
}

func speakBegin(s speaker, ex exercise) {
	ex.begin(s)
}
func speakEnd(s speaker, ex exercise) {
	ex.end(s)
}
