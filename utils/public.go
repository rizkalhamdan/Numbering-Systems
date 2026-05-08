package utils

import (
	"os"
	"os/exec"
)

func Code() int {
	return 94
}

func DO(com string) bool {
	cmd := exec.Command(com)

	cmd.Stdout = os.Stdout

	cmd.Run()

	return true
} 