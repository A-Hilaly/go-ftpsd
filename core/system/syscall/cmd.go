package syscall

import (
    "time"
    "context"
    "os/exec"
)

// Cmd struct
type Cmd struct {
    // Wrap exec.Cmd
    Own *exec.Cmd
}

// Prepare new command
func Command(cmd string, args ...string) *Cmd {
    return &Cmd{exec.Command(cmd, args...)}
}

// shortcut
func New(cmd string, args ...string) *Cmd {
    return &Cmd{exec.Command(cmd, args...)}
}

// Command with context
func CommandContext(ctx context.Context, cmd string, args ...string) *Cmd {
    return &Cmd{exec.CommandContext(ctx, cmd, args...)}
}

// shortcut with time.Duration context
func CommandTimeout(duration time.Duration, cmd string, args ...string) (*Cmd, context.CancelFunc) {
    con, cancel := context.WithTimeout(context.Background(), duration)
    return CommandContext(con, cmd, args ...), cancel
}

// Execute a Cmd command and return output stdout
func (c *Cmd) Output() ([]byte, error) {
    return c.Own.Output()
}

// Execute a Cmd command and return stdout + stderr in one output
func (c *Cmd) CombinedOutput() ([]byte, error) {
    return c.Own.CombinedOutput()
}

// Run a command line; wait until finish
func (c *Cmd) Run() error {
    return c.Own.Run()
}

// Run start command line and doesnt wait
func (c *Cmd) Start() error {
    return c.Own.Start()
}

// Threaded version of start with an error chan
func (c *Cmd) StartThreaded(ce chan error) {
    c.Start()
    err := c.Wait()
    ce <- err
}

// Wait for command line to finish
// To use with Start()
func (c *Cmd) Wait() error {
    return c.Own.Wait()
}
