package main

import (
	"github.com/urfave/cli"
	"fmt"
)

func test(c *cli.Context) error {
	fpath := c.String("file")
	fmt.Printf("test ok, file:%s\n", fpath)
	return nil
}
