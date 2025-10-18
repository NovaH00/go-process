package process

import (
	"context"
	"os/exec"

	"github.com/NovaH00/go-process/command"
)

func NewCommand(name string, args []string) *command.Command {
	return &command.Command{
		Name:    name,
		Args:    args,
		ExecCmd: exec.Command(name, args...),
	}
}

func NewCommandContext(ctx context.Context, name string, args []string) *command.Command {
	return &command.Command{
		Name:    name,
		Args:    args,
		ExecCmd: exec.CommandContext(ctx, name, args...),
	}
}
