package command

// PID returns the process ID of the running command.
// It is safe for concurrent use.
func (c *Command) PID() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.pid
}
