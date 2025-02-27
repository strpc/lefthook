package runner

import (
	"bytes"
	"io"
)

// ExecutorOptions contains the options that control the execution.
type ExecuteOptions struct {
	name, root, failText string
	args                 []string
	interactive          bool
	env                  map[string]string
}

// Executor provides an interface for command execution.
// It is used here for testing purpose mostly.
type Executor interface {
	Execute(opts ExecuteOptions, out io.Writer) error
	RawExecute(command string, args ...string) (*bytes.Buffer, error)
}
