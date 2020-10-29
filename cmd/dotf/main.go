package main

import (
	"fmt"
	"os"

	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/sysop"
)

func main() {
	if len(os.Args) != 2 {
		printHelp()
	} else {
		opProvider := &sysop.SysOpProvider{}

		var err error

		switch os.Args[1] {
		case "init":
			err = commands.Init(opProvider)
		case "push":
			fmt.Println("pushing")
		case "pull":
			fmt.Println("pulling")
		case "add":
			fmt.Println("adding")
		case "rm":
			fmt.Println("deleting")
		default:
			printHelp()
		}

		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(-1)
		}
	}
}

func printHelp() {
	var commandHelpList = []struct {
		command     string
		description string
	}{
		{"init", "Initialize dotf"},
		{"push", "Copy all dotfiles in repository and push it to the remote"},
		{"pull", "Pull the repository and replace all dotfiles from the repository"},
		{"add", "Add a dotfile to be sync'ed from now on"},
		{"rm", "Delete a dotfile from the repository"},
	}

	fmt.Printf("Usage: %s <command>\n", os.Args[0])
	fmt.Println("Available commands:")

	for _, c := range commandHelpList {
		fmt.Printf("\t%s\t-\t%s\n", c.command, c.description)
	}
}
