package main

import (
	"flag"
)

type RunSettings struct {
}

func parseRunSettings() (RunSettings, error) {

	var ret RunSettings

	// flag.StringVar(&ret.thing, "t", "default", "description")
	flag.Parse()
	return ret, nil
}
