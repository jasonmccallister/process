package process

import (
	"io"
	"os/exec"
	"syscall"
)

// Options for starting a process.
type Options struct {
	Name              string
	Args              []string
	Writer            io.Writer
	ErrWriter         io.Writer
	SetProcessGroupID bool
}

// Start a process.
func Start(o Options) error {
	p, err := exec.LookPath(o.Name)
	if err != nil {
		return err
	}

	c := exec.Command(p, o.Args...)
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: o.SetProcessGroupID}

	c.Stdout = o.Writer
	c.Stderr = o.ErrWriter

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
