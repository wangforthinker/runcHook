package main

import (
	"github.com/urfave/cli"
	"os"
	"io/ioutil"
	"encoding/json"
)

var(
	fWriteFilePath = cli.StringFlag{
		Name: "file",
		Usage: "write file path",
		Value: "/tmp/record",
	}
)

func record(c *cli.Context) error {
	log.Info("in runcHook record")

	fpath := c.String("file")

	log.WithField("file",fpath).Info("print file path")

	log.WithField("bundle", state.Bundle).WithField("id",state.ID).Info("get state info")

	data,err := json.Marshal(state)
	if err != nil {
		log.Error(os.Stderr, err.Error())
		return err
	}

	err = ioutil.WriteFile(fpath, data, 0x644)
	if err != nil {
		log.Error(os.Stderr, err.Error())
		return err
	}

	return nil
}
