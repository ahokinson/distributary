package cli

import (
	"bufio"
	"distributary/cli/format"
	"distributary/cli/format/table"
	"distributary/cli/format/table/cell"
	"distributary/internal/stream"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"time"
)

func Command(release string) {
	fmt.Println(release)
	fmt.Println()

	err := CheckRequirements()
	if err != nil {
		log.Fatalf("Distributary: checking requirements: %v", err)
	}

	data, err := os.ReadFile("distributary.yaml")
	if err != nil {
		log.Fatalf("Distributary: reading config: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Distributary: unmarshalling config: %v", err)
	}

	if cfg.Experimental.Dummy {
		fmt.Println(fmt.Sprintf("%sDummy processes are enabled!%s", format.Colors.Yellow, format.Colors.Reset))
	}

	listener := stream.Listener{
		Url:         cfg.Stream.Host,
		Destination: cfg.Stream.File,
		Video: stream.Video{
			Codec:     cfg.Stream.Video.Codec,
			BitRate:   cfg.Stream.Video.BitRate,
			FrameRate: cfg.Stream.Video.FrameRate,
			KeyFrame:  cfg.Stream.Video.KeyFrame,
		},
		Audio: stream.Audio{
			Codec:   cfg.Stream.Audio.Codec,
			BitRate: cfg.Stream.Audio.BitRate,
		},
	}

	err = listener.Init()
	if err != nil {
		log.Fatalf("Distributary: starting listener: %v", err)
	}

	listener.Streaming = false
	fmt.Println(fmt.Sprintf("Listening for stream @ %s%s%s", format.Colors.Cyan, cfg.Stream.Host, format.Colors.Reset))

	switch cfg.Experimental.AutoDetect {
	case true:
		listener.WaitForStream()
	case false:
		fmt.Println(fmt.Sprintf("Press %sEnter%s when OBS is streaming...", format.Colors.Green, format.Colors.Reset))
		_, err := bufio.NewReader(os.Stdin).ReadByte()
		if err != nil {
			log.Fatalf("Distributary: pressing enter: %v", err)
		}
		listener.Streaming = true
	}

	for countdown := 5; countdown > 0; countdown-- {
		fmt.Print(fmt.Sprintf("\rYou're %slive%s in %d...", format.Colors.Red, format.Colors.Reset, countdown))
		time.Sleep(time.Second)
	}
	fmt.Print("\n")

	var providers []*stream.Provider

	for _, provider := range cfg.Providers {
		p := stream.Provider{
			Dummy:   cfg.Experimental.Dummy,
			Name:    provider.Name,
			Ingests: provider.Ingests,
			Secret:  provider.Secret,
			Source:  cfg.Stream.File,
			Video: stream.Video{
				Codec:     cfg.Stream.Video.Codec,
				BitRate:   cfg.Stream.Video.BitRate,
				FrameRate: cfg.Stream.Video.FrameRate,
				KeyFrame:  cfg.Stream.Video.KeyFrame,
			},
			Audio: stream.Audio{
				Codec:   cfg.Stream.Audio.Codec,
				BitRate: cfg.Stream.Audio.BitRate,
			},
			Latency: cfg.Stream.Latency,
		}

		err = p.Init(0)
		if err != nil {
			log.Fatalf("Distributary: streaming to provider: %v", err)
		}

		providers = append(providers, &p)
	}

	for listener.Streaming {
		format.Clear()

		fmt.Println(fmt.Sprintf("%s%s%s", format.Colors.Cyan, "Distributary", format.Colors.Reset))
		fmt.Println()
		fmt.Println(fmt.Sprintf("You are %s%s%s!", format.Colors.Red, "LIVE", format.Colors.Reset))
		fmt.Println()
		fmt.Println()
		display := table.New(
			[]table.Initializer{
				{Header: "Provider", Style: cell.Style{Color: format.Colors.Purple}},
				{Header: "Status", Style: cell.Style{Color: format.Colors.Green}},
				{Header: "PID", Style: cell.Style{Color: format.Colors.White}},
			},
		)

		for _, p := range providers {
			display.AddRow([]string{
				p.Name,
				fmt.Sprintf("%s (%ds)", "Active", int(time.Since(p.Process.StartTime).Seconds())),
				strconv.Itoa(p.Process.Cmd.Process.Pid)})
			if !p.Process.Active {
				err = p.Init(p.Failover())
				if err != nil {
					log.Fatalf("Distributary: restarting provider stream: %v", err)
				}
			}
		}

		display.Format()
		display.Print()
		listener.CheckForEnd()
		time.Sleep(cfg.Health.Interval)
	}

	fmt.Println("\nStream ended.")

	for _, provider := range providers {
		if provider.Process.Active {
			err = provider.Process.Stop()
			if err != nil {
				log.Fatalf("Distributary: killing provider stream stream: %v", err)
			}
		}
	}
}
