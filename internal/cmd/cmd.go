package cmd

import (
	"os/exec"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type Cmd struct {
	cmd string
}

func New(cmd string) *Cmd {
	return &Cmd{
		cmd: cmd,
	}
}

func (c *Cmd) Run() (string, error) {
	parts := strings.Fields(c.cmd)
	head := parts[0]
	parts = parts[1:]

	out, err := exec.Command(head, parts...).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
