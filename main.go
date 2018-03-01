package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/wangforthinker/runcHook/utils"
	"strings"
	"github.com/wangforthinker/runcHook/hook"
	"io/ioutil"
)

var(
	log = utils.GetLogger()
	state = &hook.HookState{}
)

func readStdin() error {
	//s,err := hook.GetHookState()
	//if err != nil{
	//	return err
	//}
	//
	//state = s
	//return nil

	data,err := ioutil.ReadAll(os.Stdin)
	if err != nil{
		return err
	}

	log.Info("read stdin ok")

	log.Info("data is %s",string(data))
	return nil
}

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
	err := readStdin()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.WithField("id", state.ID).WithField("bundle",state.Bundle).Info("state info")

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}