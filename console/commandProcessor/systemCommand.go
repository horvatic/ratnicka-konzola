package commandProcessor

import (
	"io"
	"os/exec"
)

type systemCommand struct {
	cmd    *exec.Cmd
	writer *io.PipeWriter
}

func newSystemCommand(cmd *exec.Cmd) *systemCommand {
	return &systemCommand{cmd: cmd, writer: nil}
}

func newSystemCommandWithWriter(cmd *exec.Cmd, writer *io.PipeWriter) *systemCommand {
	return &systemCommand{cmd: cmd, writer: writer}
}
