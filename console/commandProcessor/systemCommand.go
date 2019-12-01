package commandProcessor

import (
	"bytes"
	"io"
	"os/exec"
)

type systemCommand struct {
	cmd         *exec.Cmd
	writer      *io.PipeWriter
	errorWriter *bytes.Buffer
}

func newSystemCommand(cmd *exec.Cmd) *systemCommand {
	return &systemCommand{cmd: cmd, writer: nil, errorWriter: nil}
}
