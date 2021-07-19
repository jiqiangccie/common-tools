package commontools

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// 返回命令到终端
func LinuxActioncommand(input *os.File) bytes.Buffer {
	cmd := exec.Command("sh")
	cmd.Stdin = input

	var ret bytes.Buffer
	cmd.Stdout = &ret
	cmd.Run()
	return ret

}

// 获取
func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func main() {
	input := os.Stdin
	var output bytes.Buffer
	output = LinuxActioncommand(input)
	fmt.Println(output.String())
}
