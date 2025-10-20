//go:build windows

package command

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"syscall"

	"golang.org/x/sys/windows"
)

// Start starts the command. It initializes stdout and stderr pipes and
// ensures the command is started in a new process group.
func (c *Command) Start() (
	stdoutPipe io.ReadCloser,
	stderrPipe io.ReadCloser,
	err error,
) {
	if c.ExecCmd.SysProcAttr == nil && c.WithNewProcessGroup {
		c.ExecCmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: windows.CREATE_NEW_PROCESS_GROUP,
		}
	}

	stdoutPipe, err = c.ExecCmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("stdout init: %w", err)
	}

	stderrPipe, err = c.ExecCmd.StderrPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("stderr init: %w", err)
	}

	if err := c.ExecCmd.Start(); err != nil {
		return nil, nil, err
	}

	c.mu.Lock()
	c.pid = c.ExecCmd.Process.Pid
	c.mu.Unlock()

	return stdoutPipe, stderrPipe, nil
}

// Wait waits for the command to exit and waits for any copying to stdin or
// copying from stdout or stderr to complete.
func (c *Command) Wait() error {
	return c.ExecCmd.Wait()
}

// Terminate sends a SIGTERM signal to the entire process group of the command.
func (c *Command) Terminate() (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.pid <= 0 {
		return errors.New("process not started or invalid PID")
	}

	err = windows.GenerateConsoleCtrlEvent(
		windows.CTRL_BREAK_EVENT,
		uint32(c.pid),
	)

	return err
}

// Kill sends a SIGKILL signal to the entire process group of the command.
func (c *Command) Kill() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.pid <= 0 {
		return errors.New("process not started or invalid PID")
	}

	cmd := exec.Command("taskkill", "/PID", strconv.Itoa(c.pid), "/T", "/F")
	err := cmd.Run()

	return err
}
