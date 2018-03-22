package system

import (
	"context"
	"io"
	osexec "os/exec"
	"syscall"
	"time"
)

var ErrExecutableNotFound = osexec.ErrNotFound


type Interface interface {
	Command(cmd string, args ...string) Cmd
	CommandContext(ctx context.Context, cmd string, args ...string) Cmd
	LookPath(file string) (string, error)
}

type Cmd interface {
	// Run runs the command to the completion.
	Run() error
	// CombinedOutput runs the command and returns its combined standard output
	// and standard error. This follows the pattern of package os/exec
	CombinedOutput() ([]byte, error)
	Output() ([]byte, error)
	SetDir(dir string)
	SetStdin(in io.Reader)
	SetStdout(out io.Writer)
	SetStderr(out io.Writer)
	Stop()
}

type ExitError interface {
	String() string
	Error() string
	Exited() bool
	ExitStatus() int
}

type executor struct{}

func New() Interface {
	return &executor{}
}

func (executor *executor) Command(cmd string, args ...string) Cmd {
	return (*cmdWrapper)(osexec.Command(cmd, args...))
}

func (executor *executor) CommandContext(ctx context.Context, cmd string, args ...string) Cmd {
	return (*cmdWrapper)(osexec.CommandContext(ctx, cmd, args...))
}

// LookPath is part of the Interface interface
func (executor *executor) LookPath(file string) (string, error) {
	return osexec.LookPath(file)
}

// Wraps exec.Cmd so we can capture errors
type cmdWrapper osexec.Cmd

var _ Cmd = &cmdWrapper{}

func (cmd *cmdWrapper) SetDir(dir string) {
	cmd.Dir = dir
}

func (cmd *cmdWrapper) SetStdin(in io.Reader) {
	cmd.Stdin = in
}

func (cmd *cmdWrapper) SetStdout(out io.Writer) {
	cmd.Stdout = out
}

func (cmd *cmdWrapper) SetStderr(out io.Writer) {
	cmd.Stderr = out
}

// Run is part of the Cmd interface
func (cmd *cmdWrapper) Run() error {
	err := (*osexec.Cmd)(cmd).Run()
	return handleError(err)
}

// CombinedOutput is part of the Cmd interface
func (cmd *cmdWrapper) CombinedOutput() ([]byte, error) {
	out, err := (*osexec.Cmd)(cmd).CombinedOutput()
	return out, handleError(err)
}

func (cmd *cmdWrapper) Output() ([]byte, error) {
	out, err := (*osexec.Cmd)(cmd).Output()
	return out, handleError(err)
}

// Stop is part of the Cmd interface.
func (cmd *cmdWrapper) Stop() {
	c := (*osexec.Cmd)(cmd)

	if c.Process == nil {
		return
	}

	c.Process.Signal(syscall.SIGTERM)

	time.AfterFunc(10*time.Second, func() {
		if !c.ProcessState.Exited() {
			c.Process.Signal(syscall.SIGKILL)
		}
	})
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case *osexec.ExitError:
		return &ExitErrorWrapper{e}
	case *osexec.Error:
		if e.Err == osexec.ErrNotFound {
			return ErrExecutableNotFound
		}
	}

	return err
}

type ExitErrorWrapper struct {
	*osexec.ExitError
}

var _ ExitError = &ExitErrorWrapper{}

// ExitStatus is part of the ExitError interface.
func (eew ExitErrorWrapper) ExitStatus() int {
	ws, ok := eew.Sys().(syscall.WaitStatus)
	if !ok {
		panic("can't call ExitStatus() on a non-WaitStatus exitErrorWrapper")
	}
	return ws.ExitStatus()
}

type CodeExitError struct {
	Err  error
	Code int
}

var _ ExitError = CodeExitError{}

func (e CodeExitError) Error() string {
	return e.Err.Error()
}

func (e CodeExitError) String() string {
	return e.Err.Error()
}

func (e CodeExitError) Exited() bool {
	return true
}

func (e CodeExitError) ExitStatus() int {
	return e.Code
}
