package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lncapital/lnc/internal/node"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Path to config file.",
		},
		&cli.StringFlag{
			Name:  "host",
			Value: "localhost:10009",
			Usage: "Where to reach the node. Default: localhost:10009",
		},
		&cli.StringFlag{
			Name:    "tls",
			Aliases: []string{"t"},
			Usage:   "Path to your tls.cert file.",
		},
		&cli.StringFlag{
			Name:    "macaroon",
			Aliases: []string{"m"},
			Usage:   "Path to your admin.macaroon file.",
		},
	}

	app := &cli.App{
		Commands: []*cli.Command{{
			Name:  "fees",
			Usage: "Prints a table of channels, fees rates, and capacity.",
			Action: func(c *cli.Context) error {
				client, err := node.ConnectToLND(c.String("host"), c.String("tls"), c.String("macaroon"))
				if err != nil {
					fmt.Print("failed to connect to node: ", err)
				}

				node.ListFees(client)
				return nil
			},
			Flags: flags,
		},
		},
	}

	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewJSONSourceFromFlagFunc("config"))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
