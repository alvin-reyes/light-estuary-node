// It creates a new Echo instance, adds some middleware, creates a new WhyPFS node, creates a new GatewayHandler, and then
// adds a route to the Echo instance
package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"light-estuary-node/cmd"
	_ "net/http"
	"os"
)

var (
	log = logging.Logger("api")
)

func main() {

	// get all the commands
	var commands []*cli.Command

	// commands
	commands = append(commands, cmd.DaemonCmd()...)
	commands = append(commands, cmd.JobCmd()...)
	commands = append(commands, cmd.PinCmd()...)
	commands = append(commands, cmd.RetrieveCmds()...)
	commands = append(commands, cmd.MinerCmd()...)
	commands = append(commands, cmd.DealsCmd()...)
	commands = append(commands, cmd.BucketCmds()...)

	app := &cli.App{
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
