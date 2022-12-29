package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	ct = &cli.App{
		Name:    "ct",
		Version: "1.0",
		Commands: []*cli.Command{
			{
				Name:   "first",
				Usage:  "the first command",
				Action: firstAction,
				Before: func(c *cli.Context) error {
					fmt.Printf("DBG => before %q; args: %q, ctx: %v\n", c.Command.Name, c.Args(), c != nil)
					if c.NArg() < 1 {
						fmt.Printf("DBG => not enough args, show help and exit for %q\n", c.Command.Name)
						cli.ShowSubcommandHelpAndExit(c, 1)
					}
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:     "subaa",
						Category: "A",
						Usage:    "subcommand aa",
						Action:   subcommandAAction,
						Before: func(c *cli.Context) error {
							fmt.Printf("DBG => before %q; args: %q, ctx: %v\n", c.Command.Name, c.Args(), c != nil)
							if c.NArg() < 1 {
								fmt.Printf("DBG => not enough args, show help and exit for %q\n", c.Command.Name)
								cli.ShowSubcommandHelpAndExit(c, 1)
							}
							return nil
						},
					},
				},
			},
		},
	}
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s\n", c.App.Version)
	}

	if err := ct.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "💥 %v\n", err)
		os.Exit(1)
	}
}

func firstAction(c *cli.Context) error {
	fmt.Println("FIRST complete")
	return nil
}

func subcommandAAction(c *cli.Context) error {
	fmt.Println("SUBAA complete")
	return nil
}
