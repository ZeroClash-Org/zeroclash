package main

import (
	"flag"
	"os"
)

var cfgPath string

func FromArgs() {
	flag.StringVar(
		&cfgPath,
		"cfg",
		os.Getenv("ZEROCLASH_CFG"),
		"specify configuration file",
	)

	flag.Parse()
}
