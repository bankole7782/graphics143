package main

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
)

func PickImage() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	cmd := exec.Command("fpicker", userHomeDir, "png|jpg")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return strings.TrimSpace(string(out))
}