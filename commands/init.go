package commands

import (
	"errors"
	"fmt"

	"bakku.dev/dotf"
)

const dotfileName = ".dotf"

func Init(sys dotf.SysOpsProvider) error {
	home := sys.GetEnvVar("HOME")

	if home == "" {
		return errors.New("init: HOME env var is not set")
	}

	dotfilePath := home + sys.GetPathSep() + dotfileName

	if !sys.PathExists(dotfilePath) {
		return createDotfile(sys, dotfilePath)
	}

	sys.Log(dotfilePath + " already exists\n")

	return nil
}

func createDotfile(sys dotf.SysOpsProvider, dotfilePath string) error {
	sys.Log("Insert path to dotfile repo: ")
	repoPath, err := sys.ReadLine()

	if err != nil {
		return fmt.Errorf("init: %v", err)
	}

	if !sys.PathExists(repoPath) {
		return fmt.Errorf("init: path %v does not exist", repoPath)
	}

	return nil
}
