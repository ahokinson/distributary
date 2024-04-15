package cli

import (
	"errors"
	"os/exec"
	"runtime"
)

func CheckRequirements() (err error) {
	switch runtime.GOOS {
	case "darwin", "linux", "windows":
	default:
		err = errors.New("unsupported operating system (I'm not sure how you got this far...)")
		return
	}

	_, err = exec.LookPath("ffmpeg")
	if err != nil {
		err = errors.New("ffmpeg is required (It's basically how this whole thing works...)")
		return
	}

	return
}
