package process

import (
	"io"
	"os/exec"
	"syscall"
)

type Options struct {
	Name      string
	Args      []string
	Writer    io.Writer
	ErrWriter io.Writer
}

func Start(o Options) error {
	p, err := exec.LookPath(o.Name)
	if err != nil {
		return err
	}

	c := exec.Command(p, o.Args...)
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	c.Stdout = o.Writer
	c.Stderr = o.ErrWriter

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
