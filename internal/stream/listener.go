package stream

import (
	"bufio"
	"distributary/internal/stream/os"
	"fmt"
	"os/exec"
	"strings"
)

type Listener struct {
	Process     os.Process
	Url         string
	Destination string
	Streaming   bool
}

func (listener *Listener) Init() (err error) {
	listener.Process = os.Process{
		Name: "OBS Listener",
		Cmd: exec.Command(
			"ffmpeg", "-listen", "1",
			"-f", "flv",
			"-i", listener.Url,
			"-y", listener.Destination,
		),
	}

	err = listener.Process.Init()

	return
}

func (listener *Listener) WaitForStream() {
	fmt.Print("Waiting for stream... ")
	// I have no idea why ffmpeg writes this stuff to stderr...
	scanner := bufio.NewScanner(listener.Process.Stderr)
	for !listener.Streaming {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "Stream #0") {
				fmt.Println("Stream detected!")
				listener.Streaming = true
				break
			}
		}
	}
}

func (listener *Listener) CheckForEnd() {
	if !listener.Process.Active {
		listener.Streaming = false
	}
}
