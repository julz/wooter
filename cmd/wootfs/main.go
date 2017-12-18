package main

import (
	"os"
	"strings"

	"github.com/cloudfoundry/groot"
	"github.com/julz/wooter"
)

func main() {
	privileged := false
	grootArgs := []string{}
	for _, arg := range os.Args {
		if strings.Contains(arg, "privileged") {
			privileged = true
		} else {
			grootArgs = append(grootArgs, arg)
		}
	}

	groot.Run(wooter.Cp{
		BaseDir:    "/tmp/scroot",
		Privileged: privileged,
	}, grootArgs)
}
