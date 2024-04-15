package cli

import (
	"time"
)

type CHealth struct {
	Interval time.Duration `yaml:"interval"`
}

type CVideo struct {
	Codec     string `yaml:"codec"`
	BitRate   string `yaml:"bitrate"`
	FrameRate int    `yaml:"framerate"`
	KeyFrame  int    `yaml:"keyframe"`
}

type CAudio struct {
	Codec   string `yaml:"codec"`
	BitRate string `yaml:"bitrate"`
}

type CStream struct {
	Host    string        `yaml:"host"`
	Latency time.Duration `yaml:"latency"`
	File    string        `yaml:"file"`
	Video   CVideo        `yaml:"video"`
	Audio   CAudio        `yaml:"audio"`
}

type CProvider struct {
	Name    string   `yaml:"name"`
	Secret  string   `yaml:"secret"`
	Ingests []string `yaml:"ingests"`
}

type CExperimental struct {
	Dummy      bool `yaml:"dummy"`
	AutoDetect bool `yaml:"autoDetect"`
}

type Config struct {
	Health       CHealth       `yaml:"health"`
	Stream       CStream       `yaml:"stream"`
	Providers    []CProvider   `yaml:"providers"`
	Experimental CExperimental `yaml:"experimental"`
}
