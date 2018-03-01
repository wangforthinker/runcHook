package main

import (
	"github.com/urfave/cli"
	"github.com/wangforthinker/runcHook/hook"
	"fmt"
	"os"
	"io/ioutil"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
)

var(
	fWriteFilePath = cli.StringFlag{
		Name: "file",
		Value: "/tmp/record",
	}
)

func record(c *cli.Context)  {
	fpath := c.String("file")

	state,err := hook.GetHookState()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}

	data,err := json.Marshal(state)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}

	err = ioutil.WriteFile(fpath, data, 0x644)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}
}
