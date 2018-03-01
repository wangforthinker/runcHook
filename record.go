package main

import (
	"github.com/urfave/cli"
	"github.com/wangforthinker/runcHook/hook"
	"os"
	"io/ioutil"
	"encoding/json"
)

var(
	fWriteFilePath = cli.StringFlag{
		Name: "file",
		Value: "/tmp/record",
	}
)

func record(c *cli.Context)  {
	fpath := c.String("file")

	log.WithField("file",fpath).Info("print file path")

	state,err := hook.GetHookState()
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.WithField("bundle", state.Bundle).WithField("id",state.ID).Info("get state info")

	data,err := json.Marshal(state)
	if err != nil {
		log.Error(os.Stderr, err.Error())
		return
	}

	err = ioutil.WriteFile(fpath, data, 0x644)
	if err != nil {
		log.Error(os.Stderr, err.Error())
		return
	}
}
