package main

import (
	"os"
	"os/exec"
)

func RunCommand(config Config, script string) (string, error) {

	if err := os.WriteFile("script.ahk", []byte(script), 0666); err != nil {
		return "", err
	}

	data, err := exec.Command(config.AutoHotKeyExec, "script.ahk").Output()
	if err != nil {
		return "", err
	}

	output := string(data)

	return output, nil
}
