package main

import (
	"flag"
	"os"
)

var (
	cfgPath            string
	externalController string
)

func FromArgs() {
	flag.StringVar(
		&cfgPath,
		"cfg",
		os.Getenv("ZEROCLASH_CFG"),
		"specify configuration file",
	)

	flag.StringVar(
		&externalController,
		"e",
		os.Getenv("ZEROCLASH_EXTERNAL_CONTROLLER"),
		"specify external controller url",
	)

	flag.Parse()
}
