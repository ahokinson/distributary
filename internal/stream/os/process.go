package os

import (
	"fmt"
	"io"
	"os/exec"
	"time"
)

var colors = map[string]string{
	"red":    "\033[91m",
	"green":  "\033[92m",
	"orange": "\033[93m",
	"blue":   "\033[94m",
}

var reset = "\033[0m"

type Process struct {
	Name      string
	Cmd       *exec.Cmd
	Active    bool
	stdout    io.ReadCloser
	Stderr    io.ReadCloser
	startTime time.Time
}

func (process *Process) Init() (err error) {
	process.stdout, err = process.Cmd.StdoutPipe()
	process.Stderr, err = process.Cmd.StderrPipe()
	process.startTime = time.Now()
	err = process.Cmd.Start()
	process.Active = true
	go func() {
		process.Thread()
	}()
	return
}

func (process *Process) Thread() {
	_ = process.Cmd.Wait()
	process.Active = false
	return
}

func (process *Process) Stop() (err error) {
	err = process.Cmd.Process.Kill()
	return
}

func (process *Process) HealthCheck() (message string) {
	if !process.Active {
		message = fmt.Sprintf(
			"%s%s (%d):%s",
			colors["blue"],
			process.Name,
			process.Cmd.Process.Pid,
			reset) + fmt.Sprintf(
			"\t%sInactive%s\t",
			colors["orange"],
			reset)
		return
	}
	message = fmt.Sprintf(
		"%s%s (%d):%s",
		colors["blue"],
		process.Name,
		process.Cmd.Process.Pid,
		reset) + fmt.Sprintf(
		"\t%sActive (%ds)%s\t",
		colors["green"],
		int(time.Since(process.startTime).Seconds()),
		reset)
	return
}
