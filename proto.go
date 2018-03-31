package main

import (
    "fmt"
    "time"
    "context"
    "os/exec"
)

type Cmd struct {
    Own *exec.Cmd
}

func Command(cmd string, args ...string) *Cmd {
    return &Cmd{exec.Command(cmd, args...)}
}

func CommandContext(ctx context.Context, cmd string, args ...string) *Cmd {
    return &Cmd{exec.CommandContext(ctx, cmd, args...)}
}

func CommandTimeout(duration time.Duration, cmd string, args ...string) (*Cmd, context.CancelFunc) {
    con, cancel := context.WithTimeout(context.Background(), duration)
    return CommandContext(con, cmd, args ...), cancel
}

func (c *Cmd) Output() ([]byte, error) {
    return c.Own.Output()
}

func (c *Cmd) CombinedOutput() ([]byte, error) {
    return c.Own.CombinedOutput()
}

func (c *Cmd) Run() error {
    return c.Own.Run()
}

func (c *Cmd) Start() error {
    return c.Own.Start()
}

func (c *Cmd) StartThreaded(ce chan error) {
    c.Start()
    err := c.Wait()
    ce <- err
}

func (c *Cmd) Wait() error{
    return c.Own.Wait()
}

func main() {
    cmd := Command("ls", "-la", "kaka")
    ec := make(chan error, 1)
    go cmd.StartThreaded(ec)
    fmt.Println(<- ec)
    fmt.Println(a)

}
