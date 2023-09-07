package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	OpenapiKey     string
	AutoHotKeyExec string
}

func main() {

	content, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during JSON parse: ", err)
	}

	text, err := Transcribe(config)
	if err != nil {
		log.Fatal("Cannot transcribe text: ", err)
		return
	}

	log.Println("Prompt:\n  " + text)
	command, err := BuildCommand(config, text)
	if err != nil {
		log.Fatal("Cannot interpret prompt", err)
		return
	}

	fmt.Println("Running:\n  " + command)
	output, err := RunCommand(config, command)
	if err != nil {
		log.Fatal("Cannot run command", err)
	}
	log.Println("Output:\n  " + output)
}
