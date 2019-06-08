package config

import (
	"path"
	"strconv"
)

type Capture struct {
	SegmentMils  int
	SegmentCount int
	Width        int
	Height       int
	BitRate      int
	FrameRate    int
	FlipVert     bool
	FlipHoriz    bool
	VideoPath    string
}

func (c *Capture) ToArgs() []string {
	args := []string{
		"--segment", strconv.Itoa(c.SegmentMils),
		"--timeout", "0",
		"--width", strconv.Itoa(c.Width),
		"--height", strconv.Itoa(c.Height),
		"--wrap", strconv.Itoa(c.SegmentCount),
		"--bitrate", strconv.Itoa(c.BitRate),
		"--output", path.Join(c.VideoPath, "segment_%c.h264"),
	}

	if c.FlipVert {
		args = append(args, "-vf")
	}

	if c.FlipHoriz {
		args = append(args, "-vh")
	}

	return args
}
