package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Config struct {
	OpenapiKey     string
	AutoHotKeyExec string
}

func main() {
	err := assertThatConfigFileExists()
	if err != nil {
		log.Fatal("Error when creating config file: ", err)
	}

	content, err := readConfigFile()
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during JSON parse: ", err)
	}

	if config.OpenapiKey == "" {
		log.Fatal("Please provide your OpenAI key in the file config.json")
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

func readConfigFile() ([]byte, error) {
	content, err := os.ReadFile("./config.json")
	return content, err
}

func assertThatConfigFileExists() error {
	if !exists("./config.json") {
		template, err := os.ReadFile("./config.template.json")
		if err != nil {
			return fmt.Errorf("cannot read template config file: %w", err)
		}
		err = os.WriteFile("./config.json", template, 0666)
		if err != nil {
			return fmt.Errorf("cannot write new config file: %w", err)
		}
		return nil
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
