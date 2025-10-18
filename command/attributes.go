package command

func (c *Command) PID() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.pid
}
