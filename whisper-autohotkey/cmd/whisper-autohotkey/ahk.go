package main

import (
	"os"
	"os/exec"
)

func RunCommand(config Config, script string) (string, error) {

	if err := os.WriteFile("script.ahk", []byte(script), 0666); err != nil {
		return "", err
	}

	autoHotKeyPath := config.AutoHotKeyExec
	if autoHotKeyPath == "" {
		autoHotKeyPath = ".\\bin\\autohotkey-1.1.37.01\\AutoHotkeyU64.exe"
	}
	data, err := exec.Command(autoHotKeyPath, "script.ahk").Output()
	if err != nil {
		return "", err
	}

	output := string(data)

	return output, nil
}
