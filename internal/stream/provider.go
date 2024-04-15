package stream

import (
	"distributary/internal/stream/os"
	"errors"
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"strconv"
)

func (provider *Provider) Init(ingest int) (err error) {
	provider.ingest = ingest
	switch provider.Dummy {
	case true:
		var dummyCmd []string
		seconds := strconv.Itoa(rand.Intn(61))

		switch runtime.GOOS {
		case "darwin", "linux":
			dummyCmd = []string{"sleep", seconds}
		case "windows":
			dummyCmd = []string{"timeout", "/t", seconds}
		default:
			err = errors.New("unsupported operating system")
			return
		}

		provider.Process = os.Process{
			Name: provider.Name,
			Cmd:  exec.Command(dummyCmd[0], dummyCmd[1:]...),
		}
	default:
		provider.Process = os.Process{
			Name: provider.Name,
			Cmd: exec.Command(
				"ffmpeg", "-re",
				"-sseof", strconv.Itoa(int(-provider.Latency.Seconds())),
				"-i", provider.Source,
				"-c:v", provider.Video.Codec,
				"-b:v", provider.Video.BitRate,
				"-c:a", provider.Audio.Codec,
				"-b:a", provider.Audio.BitRate,
				"-r", strconv.Itoa(provider.Video.FrameRate),
				"-g", strconv.Itoa(provider.Video.KeyFrame*provider.Video.FrameRate),
				"-f", "flv",
				fmt.Sprintf("%s/%s", provider.Ingests[provider.ingest], provider.Secret),
			)}
	}

	err = provider.Process.Init()

	return
}

func (provider *Provider) Failover() (ingest int) {
	return (provider.ingest + 1) % len(provider.Ingests)
}
