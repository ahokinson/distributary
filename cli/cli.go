package cli

import (
	stream "distributary/internal/stream"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

func Command() {
	data, err := os.ReadFile("distributary.yaml")
	if err != nil {
		log.Fatalf("Distributary: reading config: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Distributary: unmarshalling config: %v", err)
	}

	listener := stream.Listener{
		Url:         cfg.Listener.Host,
		Destination: cfg.Listener.File,
	}

	err = listener.Init()
	if err != nil {
		log.Fatalf("Distributary: starting listener: %v", err)
	}

	listener.Streaming = false
	listener.WaitForStream()

	fmt.Println("Get ready, you're live in 5...")
	time.Sleep(5 * time.Second)
	var activeProviders []*stream.Provider

	for _, provider := range cfg.Providers {
		p := stream.Provider{
			Name:      provider.Name,
			Ingests:   provider.Ingests,
			Source:    cfg.Listener.File,
			Latency:   cfg.Listener.Latency,
			Preset:    cfg.Listener.Preset,
			Bitrate:   cfg.Listener.BitRate,
			Framerate: cfg.Listener.FrameRate,
			Keyframe:  cfg.Listener.KeyFrame,
		}

		err = p.Init(0)
		if err != nil {
			log.Fatalf("Distributary: streaming to provider: %v", err)
		}

		activeProviders = append(activeProviders, &p)
	}
	for listener.Streaming {
		update := "\033[K"
		update += listener.Process.HealthCheck()

		for _, provider := range activeProviders {
			update += provider.Process.HealthCheck()
		}
		update += "\r"
		fmt.Print(update)

		for _, p := range activeProviders {
			if !p.Process.Active {
				err = p.Init(p.Failover())
				if err != nil {
					log.Fatalf("Distributary: restarting provider stream: %v", err)
				}
			}
		}

		listener.CheckForEnd()
		time.Sleep(cfg.Health.Interval)
	}

	fmt.Println("\nStream ended.")

	for _, provider := range activeProviders {
		if provider.Process.Active {
			err = provider.Process.Stop()
			if err != nil {
				log.Fatalf("Distributary: killing provider stream stream: %v", err)
			}
		}
	}
}
