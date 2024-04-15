package stream

import (
	"bufio"
	"distributary/internal/stream/os"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func (listener *Listener) Init() (err error) {
	listener.Process = os.Process{
		Name: "Distributary",
		Cmd: exec.Command(
			"ffmpeg",
			"-listen", "1",
			"-i", listener.Url,
			"-c:v", listener.Video.Codec,
			"-b:v", listener.Video.BitRate,
			"-c:a", listener.Audio.Codec,
			"-b:a", listener.Audio.BitRate,
			"-y", listener.Destination,
		),
	}

	go func() {
		err = http.ListenAndServe(":1935", nil)
		if err != nil {
			log.Fatalf("Distributary: RTMP server: %v", err)
		}
	}()

	err = listener.Process.Init()

	return
}

func (listener *Listener) WaitForStream() {
	for !listener.Streaming {
		scanner := bufio.NewScanner(listener.Process.Stderr)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "Stream") {
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
