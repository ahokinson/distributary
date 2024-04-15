package os

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

type Process struct {
	Name      string
	Cmd       *exec.Cmd
	Active    bool
	StartTime time.Time
	Stdout    io.ReadCloser
	stdoutLog *os.File
	Stderr    io.ReadCloser
	stderrLog *os.File
	restarts  int
}

func (process *Process) Init() (err error) {

	process.StartTime = time.Now()
	process.Stdout, err = process.Cmd.StdoutPipe()
	process.stdoutLog, err = os.OpenFile(fmt.Sprintf("logs/%s.%d.stdout.log", process.Name, process.StartTime.Unix()), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	go func() {
		_, err := io.Copy(process.stdoutLog, process.Stdout)
		if err != nil {
			return
		}
	}()
	process.Stderr, err = process.Cmd.StderrPipe()
	process.stderrLog, err = os.OpenFile(fmt.Sprintf("logs/%s.%d.stderr.log", process.Name, process.StartTime.Unix()), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	go func() {
		_, err := io.Copy(process.stderrLog, process.Stderr)
		if err != nil {
			return
		}
	}()
	err = process.Cmd.Start()
	if err != nil {
		return
	}

	process.Active = true
	go func() {
		err = process.Thread()
		if err != nil {
			fmt.Println(err)
		}
	}()

	return
}

func (process *Process) Thread() (err error) {
	err = process.Cmd.Wait()
	if err != nil {
		return
	}

	process.Active = false
	process.restarts += 1
	return
}

func (process *Process) Stop() (err error) {
	err = process.Cmd.Process.Kill()
	if err != nil {
		return
	}

	return
}
