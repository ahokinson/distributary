package stream

import (
	"distributary/internal/stream/os"
	"time"
)

type Video struct {
	Codec     string
	BitRate   string
	FrameRate int
	KeyFrame  int
}

type Audio struct {
	Codec   string
	BitRate string
}

type Listener struct {
	Process     os.Process
	Url         string
	Destination string
	Video       Video
	Audio       Audio
	Streaming   bool
}

type Provider struct {
	Name    string
	Process os.Process
	Dummy   bool
	Ingests []string
	Secret  string
	Source  string
	Video   Video
	Audio   Audio
	Latency time.Duration

	ingest int
}
