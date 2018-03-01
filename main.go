package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/wangforthinker/runcHook/utils"
	"strings"
)

var(
	log = utils.GetLogger()
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
		cli.Command{
			Name: "test",
			Flags: []cli.Flag{fWriteFilePath},
			Action: test,
		},
	}

	log.Infof("start runc hook, cmd is %s", strings.Join(os.Args,","))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}