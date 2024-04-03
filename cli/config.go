package cli

import (
	"time"
)

type CHealth struct {
	Interval time.Duration `yaml:"interval"`
}

type CListener struct {
	Host      string        `yaml:"host"`
	Latency   time.Duration `yaml:"latency"`
	File      string        `yaml:"file"`
	Preset    string        `yaml:"preset"`
	BitRate   string        `yaml:"bitrate"`
	FrameRate int           `yaml:"framerate"`
	KeyFrame  int           `yaml:"keyframe"`
}

type CProvider struct {
	Name    string   `yaml:"name"`
	Secret  string   `yaml:"secret"`
	Ingests []string `yaml:"ingests"`
}

type Config struct {
	Health    CHealth     `yaml:"health"`
	Listener  CListener   `yaml:"stream"`
	Providers []CProvider `yaml:"providers"`
}
