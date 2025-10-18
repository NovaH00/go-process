package command

import (
	"os/exec"
	"sync"
)

type Command struct {
	Name string
	Args []string

	ExecCmd *exec.Cmd
	pid     int

	mu sync.RWMutex
}
