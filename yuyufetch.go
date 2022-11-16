package yuyufetch

import (
	"flag"
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
    var showVersion bool
    fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
    fs.BoolVar(&showVersion, "v", false, "show version")
    fs.BoolVar(&showVersion, "version", false, "show version")
    fs.SetOutput(cli.Stderr)
	fs.Usage = func() {
        fmt.Printf("yuyufetch %s\n\n", version)
        fmt.Printf("Usage:\n  %s [Options]\n\n", "yuyufetch")
        fmt.Println("Options:")
		fs.PrintDefaults()
        fmt.Println("")
	}

    if err := fs.Parse(args[1:]); err != nil {
		return err
	}

    if showVersion {
        fmt.Fprintln(cli.Stdout, "yuyufetch", version)
        return nil
    }

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
