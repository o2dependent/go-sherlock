package gosherlock

import (
	"os/exec"
)

func Test() {
	cmd := exec.Command("ls", "-l")
	cmd.Run()
}