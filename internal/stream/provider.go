package stream

import (
	"distributary/internal/stream/os"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

type Provider struct {
	Name      string
	Process   os.Process
	Ingests   []string
	Source    string
	Latency   time.Duration
	Preset    string
	Bitrate   string
	Framerate int
	Keyframe  int

	ingest int
}

func (provider *Provider) Init(ingest int) (err error) {
	provider.ingest = ingest
	provider.Process = os.Process{
		Name: provider.Name,
		Cmd: exec.Command(
			"ffmpeg", "-re",
			"-sseof", strconv.Itoa(int(-provider.Latency.Seconds())),
			"-i", provider.Source,
			"-c:v", "libx264",
			"-r", strconv.Itoa(provider.Framerate),
			"-g", strconv.Itoa(provider.Keyframe*provider.Framerate),
			"-preset", provider.Preset,
			"-b:v", provider.Bitrate,
			"-f", "flv", provider.Ingests[provider.ingest],
		)}

	err = provider.Process.Init()

	return
}

func (provider *Provider) Failover() (ingest int) {
	return (provider.ingest + 1) % len(provider.Ingests)
}

// The following struct is only used for testing.
// TODO: Delete when confidence in process handling increases.

type ProviderTest struct {
	Name      string
	Process   os.Process
	Url       string
	Source    string
	Latency   time.Duration
	Preset    string
	Bitrate   string
	Framerate int
	Keyframe  int
}

func (p *ProviderTest) Init() (err error) {
	p.Process = os.Process{
		Name: p.Name,
		Cmd: exec.Command(
			"sleep", strconv.Itoa(rand.Intn(61)),
		),
	}

	err = p.Process.Init()

	return
}
