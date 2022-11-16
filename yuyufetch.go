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
	logo_lines := strings.Split(Yuyu_logo, "\n")
	yuyu_info := YuyuInfo.GetStyledContents()

	var max_len int
	if len(logo_lines) > len(yuyu_info) {
		max_len = len(logo_lines)
	} else {
		max_len = len(yuyu_info)
	}

	for i := 0; i < max_len; i++ {
		if i < len(logo_lines) {
			fmt.Fprintf(cli.Stdout, "%s", logo_lines[i])
		} else {
			fmt.Fprintf(cli.Stdout, "%s", strings.Repeat(" ", len(logo_lines[0])))
		}

		fmt.Fprint(cli.Stdout, "   ")

		if i < len(yuyu_info) {
			fmt.Fprintf(cli.Stdout, "%s", yuyu_info[i])
		}
		fmt.Fprintln(cli.Stdout, "")
	}

	return nil
}
