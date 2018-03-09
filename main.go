package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"os/exec"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		port := ":"+c.Args().Get(0)
		out, err := exec.Command("lsof", "-i", port, "-t", "-n").Output()
		if err != nil {
			return err
		}
		err = exec.Command("kill", strings.TrimRight(string(out), "\n")).Run()
		if err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}