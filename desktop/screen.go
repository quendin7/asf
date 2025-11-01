package desktop

import (
	"os/exec"
	"strings"
)

func GetScreenResolution() string {
	cmd := exec.Command("xrandr", "--current")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "*") {
			fields := strings.Fields(line)
			return fields[0]
		}
	}

	return "unknown"
}
