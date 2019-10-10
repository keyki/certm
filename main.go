package main

import (
	"os"

	"github.com/ehazlett/certm/commands/bundle"
	"github.com/ehazlett/certm/commands/ca"
	"github.com/ehazlett/certm/commands/client"
	"github.com/ehazlett/certm/commands/server"
	"github.com/ehazlett/certm/version"
	"github.com/ehazlett/simplelog"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	// setup log formatter
	f := &simplelog.SimpleFormatter{}
	log.SetFormatter(f)
}

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "certificate management"
	app.Version = version.FullVersion()
	app.Author = "@ehazlett"
	app.Email = ""
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Commands = []cli.Command{
		ca.CmdCA,
		server.CmdServer,
		client.CmdClient,
		bundle.CmdBundle,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output-directory, d",
			Value: "",
			Usage: "output directory for certs",
		},
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
