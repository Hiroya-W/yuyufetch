package yuyufetch

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	ExitOK int = 0
	ExitNG int = 1
)

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

func Main(args []string) int {
	cli := &CLI{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	if err := cli.Run(args); err != nil {
		fmt.Fprintln(cli.Stderr, "Error:", err)
		return ExitNG
	}

	return ExitOK
}

func (cli *CLI) Run(args []string) error {
    lines := strings.Split(yuyu_logo, "\n")

    for _, line := range lines {
        fmt.Fprintf(cli.Stdout, "%s\n", line)
    }
    return nil
}