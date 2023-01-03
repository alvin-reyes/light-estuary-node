package cmd

import (
	"context"
	"github.com/urfave/cli/v2"
	"light-estuary-node/api"
	"light-estuary-node/core"
	"light-estuary-node/jobs"
	"time"
)

func DaemonCmd() []*cli.Command {
	// add a command to run API node
	var daemonCommands []*cli.Command

	daemonCmd := &cli.Command{
		Name:  "daemon",
		Usage: "A light version of Estuary that allows users to upload and download data from the Filecoin network.",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "enable-api",
			},
		},
		Action: func(c *cli.Context) error {

			ln, err := core.NewLightNode(context.Background())
			if err != nil {
				return err
			}

			// launch the API node
			api.InitializeEchoRouterConfig(ln)

			//	launch the jobs
			go runJobs()

			api.LoopForever()
			return nil
		},
	}

	// add commands.
	daemonCommands = append(daemonCommands, daemonCmd)

	return daemonCommands

}

func runJobs() {

	// run the job every 10 seconds.
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tick.C:
			// run the job.
			go jobs.NewBucketAssignProcessor().Run()
			go jobs.NewCarGeneratorProcessor().Run()
			go jobs.NewCommpProcessor().Run()
			go jobs.NewDealsProcessor().Run()
		}
	}

}
