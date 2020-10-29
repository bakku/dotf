package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printHelp()
	} else {
		switch os.Args[1] {
		case "init":
			fmt.Println("initing")
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
