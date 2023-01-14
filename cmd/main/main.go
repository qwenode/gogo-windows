package main

import (
	"github.com/qwenode/gogo-windows/process"
	"log"
	"os/exec"
)

func main() {
	command := exec.Command("cmd.exe", "/C", "rundll32 user32.dll LockWorkStation")
	byProcess, err := process.RunCommandByProcess("explorer", command)
	log.Println(byProcess, err)
}
