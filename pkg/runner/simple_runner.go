package runner

import (
	"context"
	"os/exec"

	"github.com/D1360-64RC14/multirun/config"
	"github.com/D1360-64RC14/multirun/interfaces"
	"github.com/D1360-64RC14/multirun/pkg/command"
	"github.com/D1360-64RC14/multirun/pkg/printer"
)

var _ interfaces.Runner = (*SimpleRunner)(nil)

type SimpleRunner struct {
	runningContext context.Context
	cancelContext  context.CancelFunc
	commands       map[string]*command.Command
	stdoutChannel  chan interfaces.NamedMessage
	printer        printer.Printer
	watching       bool
}

func NewSimpleRunner(commands config.Commands, printer printer.Printer) (runner *SimpleRunner) {
	ctx, cancel := context.WithCancel(context.Background())

	runner = &SimpleRunner{
		runningContext: ctx,
		cancelContext:  cancel,
		commands:       make(map[string]*command.Command, len(commands)),
		stdoutChannel:  make(chan interfaces.NamedMessage, 64),
		printer:        printer,
		watching:       false,
	}

	for name, commandLine := range commands {
		cmd := exec.CommandContext(ctx, "bash", "-c", commandLine)
		runner.commands[name] = command.New(name, cmd, runner.stdoutChannel)
	}

	return runner
}

func (r *SimpleRunner) Run() error {
	r.watching = true

	go r.stdoutWatcher()

	for name, command := range r.commands {
		r.printer.Println(name, "< Starting... >")
		go command.Run()
	}

	return nil
}

func (r *SimpleRunner) stdoutWatcher() {
	for r.watching {
		message := <-r.stdoutChannel

		if message.Output == nil {
			continue
		}

		r.printer.Println(message.Name, string(message.Output))
	}
}

func (r *SimpleRunner) Cancel() error {
	r.cancelContext()

	r.watching = false
	r.sendDummyMessage()

	return nil
}

func (r *SimpleRunner) sendDummyMessage() {
	r.stdoutChannel <- interfaces.NamedMessage{
		Name:   "dummy",
		Output: nil,
	}
}
