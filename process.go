package process

import (
	"context"
	"os/exec"

	"github.com/NovaH00/go-process/command"
)

// NewCommand creates a new Command instance.
func NewCommand(name string, args ...string) *command.Command {
	return &command.Command{
		Name:    name,
		Args:    args,
		ExecCmd: exec.Command(name, args...),
	}
}

// NewCommandContext creates a new Command instance with a context.
func NewCommandContext(ctx context.Context, name string, args ...string) *command.Command {
	return &command.Command{
		Name:    name,
		Args:    args,
		ExecCmd: exec.CommandContext(ctx, name, args...),
	}
}
