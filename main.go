package sherlock

import (
	"fmt"
	"os/exec"
)

func Test() {
	cmd := exec.Command("ls")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
