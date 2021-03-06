package commands

import (
	"errors"
	"fmt"

	"bakku.dev/dotf"
)

// Init tries to create the dotfile of dotf under $HOME/.dotf.
func Init(sys dotf.SysOpsProvider, repoPath string) error {
	home := sys.GetEnvVar("HOME")

	if home == "" {
		return errors.New("init: HOME env var is not set")
	}

	dotfilePath := sys.CleanPath(home + sys.GetPathSep() + dotfileName)

	if !sys.PathExists(dotfilePath) {
		return createDotfile(sys, dotfilePath, repoPath)
	}

	sys.Log(dotfilePath + " already exists\n")

	return nil
}

func createDotfile(sys dotf.SysOpsProvider, dotfilePath string, repoPath string) error {
	repoPath, err := sys.ExpandPath(repoPath)

	if err != nil {
		return fmt.Errorf("init: could not get absolute path: %v", err)
	}

	if !sys.PathExists(repoPath) {
		return fmt.Errorf("init: path %v does not exist", repoPath)
	}

	var resp string

	for resp != "y" && resp != "n" {
		sys.Log("Do you want to create backups of your dotfiles when pulling? (y/n): ")

		resp, err = sys.ReadLine()
		if err != nil {
			return fmt.Errorf("init: %v", err)
		}
	}

	var createBackups bool = false

	if resp == "y" {
		createBackups = true
	}

	conf := dotf.Config{repoPath, createBackups, []dotf.TrackedFile{}}
	bytes, err := sys.SerializeConfig(conf)

	if err != nil {
		return fmt.Errorf("init: could not serialize config: %v", err)
	}

	err = sys.WriteFile(dotfilePath, bytes)
	if err != nil {
		return fmt.Errorf("init: count not write to file %s: %v", dotfilePath, err)
	}

	sys.Log("Successfully created file at " + dotfilePath + "\n")

	return nil
}
