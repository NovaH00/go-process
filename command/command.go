package command

import (
	"os/exec"
	"sync"
)

// Command represents a command to be executed, wrapping the standard os/exec.Cmd.
type Command struct {
	Name                string
	Args                []string
	ExecCmd             *exec.Cmd
	WithNewProcessGroup bool
	pid                 int
	mu                  sync.RWMutex
}