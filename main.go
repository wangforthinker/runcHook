package main

import (
	"github.com/urfave/cli"
	"github.com/Sirupsen/logrus"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "runcHook"
	app.Usage = "for runc hook test"

	app.Commands = []cli.Command{
		cli.Command{
			Name: "record",
			Flags: []cli.Flag{fWriteFilePath},
			Action: record,
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err.Error())
	}
}