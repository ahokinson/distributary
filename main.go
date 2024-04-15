package main

import (
	"distributary/cli"
	"embed"
	"fmt"
	"gopkg.in/yaml.v3"
)

//go:embed Release.yaml
var content embed.FS

type Release struct {
	Version   string `yaml:"version"`
	Copyright string `yaml:"copyright"`
	Author    string `yaml:"author"`
}

func (r Release) String() string {
	return fmt.Sprintf("distributary version %s Copyright (c) %s %s", r.Version, r.Copyright, r.Author)
}

func main() {
	data, _ := content.ReadFile("Release.yaml")
	var release Release
	_ = yaml.Unmarshal(data, &release)
	cli.Command(release.String())
}
