package main

type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommand() {
	for _, c := range c.Commands {
		c.execute()
	}
}