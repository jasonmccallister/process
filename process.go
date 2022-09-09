package process

import (
	"io"
	"os/exec"
	"syscall"
)

// Opts for starting a process.
type Opts struct {
	Name              string
	Args              []string
	Writer            io.Writer
	ErrWriter         io.Writer
	SetProcessGroupID bool
}

// Start a process.
func Start(opts Opts) error {
	p, err := exec.LookPath(opts.Name)
	if err != nil {
		return err
	}

	c := exec.Command(p, opts.Args...)
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: opts.SetProcessGroupID}

	c.Stdout = opts.Writer
	c.Stderr = opts.ErrWriter

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
