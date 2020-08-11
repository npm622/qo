package commands

import (
	"github.com/urfave/cli/v2"
)

// set of global args
const (
	ArgURL     = "url"
	ArgVerbose = "verbose"
	ArgFerbose = "ferbose"
)

type globalArgs struct {
	url     string
	verbose bool
	ferbose bool
}

func parseGlobalArgs(c *cli.Context) globalArgs {
	return globalArgs{
		url:     c.String(ArgURL),
		verbose: c.Bool(ArgVerbose),
		ferbose: c.Bool(ArgFerbose),
	}
}
