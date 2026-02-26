package main

import (
	"flag"
	"os"
)

var (
	version bool
	cfgPath string
)

func FromArgs() {
	flag.BoolVar(
		&version,
		"v",
		false,
		"current version of zeroclash",
	)

	flag.StringVar(
		&cfgPath,
		"cfg",
		os.Getenv("ZEROCLASH_CFG"),
		"specify configuration file",
	)

	flag.Parse()
}
