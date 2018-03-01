package main

import (
	"os"
	"github.com/wangforthinker/runcHook/utils"
	"strings"
	"github.com/wangforthinker/runcHook/hook"
	"io/ioutil"
	"encoding/json"
)

var(
	log = utils.GetLogger()
	state = &hook.HookState{}
)

func readStdin() error {
	data,err := ioutil.ReadAll(os.Stdin)
	if err != nil{
		return err
	}

	log.Info("data is %s",string(data))

	return json.Unmarshal(data, &state)
}

func main() {
	log.Infof("start runc hook, cmd is %s", strings.Join(os.Args,","))

	if len(os.Args) != 1 {
		log.Fatal("args should be 1")
	}

	err := readStdin()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.WithField("id", state.ID).WithField("bundle",state.Bundle).Info("state info")
	record(os.Args[0])
}