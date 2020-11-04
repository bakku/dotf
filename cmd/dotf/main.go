package main

import (
	"fmt"
	"os"

	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/sysop"
	"github.com/urfave/cli/v2"
)

func main() {
	opProvider := &sysop.Provider{}

	app := &cli.App{
		Name:      "dotf",
		Usage:     "a simple dotfile manager",
		UsageText: "dotf command <command arguments>",
		HideHelp:  true,
		Commands: []*cli.Command{
			{
				Name:      "init",
				Aliases:   []string{"i"},
				Usage:     "initialize dotf",
				ArgsUsage: "<path to dotfile repo>",
				HideHelp:  true,
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 1 {
						return cli.ShowCommandHelp(c, "init")
					}

					return commands.Init(opProvider, c.Args().First())
				},
			},
			{
				Name:      "add",
				Aliases:   []string{"a"},
				Usage:     "track a new file",
				ArgsUsage: "<path to file> <path in repo>",
				HideHelp:  true,
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 2 {
						return cli.ShowCommandHelp(c, "add")
					}

					return commands.Add(opProvider, c.Args().First(), c.Args().Get(1))
				},
			},
			{
				Name:      "rm",
				Aliases:   []string{"r"},
				Usage:     "remove tracking of file",
				ArgsUsage: "<path to file>",
				HideHelp:  true,
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 1 {
						return cli.ShowCommandHelp(c, "rm")
					}

					return commands.Remove(opProvider, c.Args().First())
				},
			},
			{
				Name:      "pull",
				Usage:     "copy all dotfiles to the repository and push it to the remote",
				ArgsUsage: " ",
				HideHelp:  true,
				Action: func(c *cli.Context) error {
					return cli.ShowCommandHelp(c, "pull")
				},
			},
			{
				Name:      "push",
				Usage:     "pull the repository and replace all dotfiles",
				ArgsUsage: " ",
				HideHelp:  true,
				Action: func(c *cli.Context) error {
					return cli.ShowCommandHelp(c, "push")
				},
			},
			{
				Name:      "list",
				Aliases:   []string{"l"},
				Usage:     "show all tracked files",
				ArgsUsage: " ",
				HideHelp:  true,
				Action: func(c *cli.Context) error {
					return commands.List(opProvider)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(-1)
	}
}
