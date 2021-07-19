package commontools

import (
	"os"
	"os/exec"
)

func LinuxActioncommand(input *os.File) (output *os.File) {
	cmd := exec.Command("sh")
	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Run()
	return output

}

// 获取
func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
