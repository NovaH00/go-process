# Go Process

Go Process is a simple library for managing system processes in Go. It provides a wrapper around `os/exec.Cmd` with additional features for process lifecycle management.

## Features

- Start, wait for, and terminate processes.
- Terminate entire process groups.
- Access to process PID.

## Usage

The `NewCommand` and `NewCommandContext` functions accept a boolean `withNewProcessGroup` parameter. If this is set to `true`, the command will be executed in a new process group. This allows you to terminate the entire process group, including any child processes, by calling the `Terminate()` method.

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/NovaH00/go-process"
)

func main() {
	// Create a new command
	cmd := process.NewCommand(false, "sleep", "5")

	// Start the command
	stdout, stderr, err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Get the process ID
	fmt.Printf("Process started with PID: %d\n", cmd.PID())

	// Read from stdout and stderr
	go func() {
		_, _ = io.Copy(os.Stdout, stdout)
	}()
	go func() {
		_, _ = io.Copy(os.Stderr, stderr)
	}()

	// Terminate the process after 2 seconds
	time.Sleep(2 * time.Second)
	if err := cmd.Terminate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Process terminated")
}
```
