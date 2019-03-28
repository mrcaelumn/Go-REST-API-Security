package main

import (
	"os"

	"github.com/inconshreveable/log15"
	"github.com/mrcaelumn/go-rest-api-security/command/run"
	"github.com/mrcaelumn/go-rest-api-security/version"
	"github.com/urfave/cli"
)

var logHandler log15.Handler

func main() {
	app := cli.NewApp()
	app.Name = "go-rest-api-security"
	app.Usage = "go-rest-api-security server component"
	app.Copyright = "(c) 2019 mrcaelumn"
	app.Version = version.Version + " (" + version.GitCommit + ")"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable verbose logging",
		},
	}
	app.Before = func(c *cli.Context) error {
		f := log15.JsonFormat()
		if c.Bool("debug") {
			log15.LvlFilterHandler(log15.LvlDebug, logHandler)
		} else {
			log15.Root().SetHandler(log15.LvlFilterHandler(log15.LvlError, log15.CallerFileHandler(log15.StreamHandler(os.Stdout, f))))
		}
		return nil
	}
	app.Commands = []cli.Command{
		run.Command,
	}

	if err := app.Run(os.Args); err != nil {
		log15.Crit(err.Error())
	}
}
