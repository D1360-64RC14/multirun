package command

import (
	"bufio"
	"bytes"
	"os/exec"

	"github.com/D1360-64RC14/multirun/interfaces"
	"github.com/D1360-64RC14/multirun/pkg"
)

type Command struct {
	name    string
	cmd     *exec.Cmd
	stop    chan bool
	running bool
	stdout  chan interfaces.NamedMessage
}

func New(name string, cmd *exec.Cmd, stdoutChannel chan interfaces.NamedMessage) *Command {
	return &Command{
		name:    name,
		cmd:     cmd,
		stop:    make(chan bool, 1),
		running: false,
		stdout:  stdoutChannel,
	}
}

func (c *Command) Run() {
	if c.running {
		return
	}

	c.cmd.Stdout = pkg.FuncWriter(c.onWrite)
	c.cmd.Stderr = c.cmd.Stdout

	c.running = true
	c.cmd.Start()

	go c.checkForChanStop()

	c.cmd.Wait()
	c.running = false

	c.stop <- true
}

func (c *Command) onWrite(p []byte) {
	pReader := bytes.NewReader(p)
	pScanner := bufio.NewScanner(pReader)

	for pScanner.Scan() {
		c.stdout <- interfaces.NamedMessage{
			Name:   c.name,
			Output: pScanner.Bytes(),
		}
	}
}

func (c *Command) checkForChanStop() {
	<-c.stop
	if c.running {
		c.cmd.Cancel()
	}
}

func (c *Command) Stop() {
	if c.running {
		c.stop <- true
	}
}
