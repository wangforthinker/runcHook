package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

func record(fpath string) error {
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
