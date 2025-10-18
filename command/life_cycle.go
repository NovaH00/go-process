package command

import (
	"errors"
	"fmt"
	"io"
	"syscall"
)

func (c *Command) Start() (
	stdoutPipe io.ReadCloser,
	stderrPipe io.ReadCloser,
	err error,
) {
	if c.ExecCmd.SysProcAttr == nil {
		c.ExecCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
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

func (c *Command) Wait() error {
	return c.ExecCmd.Wait()
}

func (c *Command) Terminate() (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.pid <= 0 {
		return errors.New("process not started or invalid PID")
	}

	return syscall.Kill(-c.pid, syscall.SIGTERM)
}

func (c *Command) Kill() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.pid <= 0 {
		return errors.New("process not started or invalid PID")
	}

	return syscall.Kill(-c.pid, syscall.SIGKILL)
}
