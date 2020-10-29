package commands

import (
	"errors"

	"bakku.dev/dotf"
)

const dotfileName = ".dotf"

func Init(sys dotf.SysOpsProvider) error {
	home := sys.GetEnvVar("HOME")

	if home == "" {
		return errors.New("init: HOME env var is not set")
	}

	dotfilePath := sys.CleanPath(home + sys.GetPathSep() + dotfileName)

	if !sys.FileExists(dotfilePath) {
		return nil
	}

	sys.Log(dotfilePath + " already exists")
	
	return nil
}
