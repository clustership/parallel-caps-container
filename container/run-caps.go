package main

import (
	"log"
  "os"
	"os/exec"
)

func main() {
	// cmd := exec.Command("nice", "-20", "sleep", "5")
	cmd := exec.Command("nice", "-20", "cat", "/etc/passwd")
  cmd.Stdout = os.Stdout

	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
  if err != nil {
	  log.Printf("Command finished with error: %v", err)
  }
}
