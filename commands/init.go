package commands

import (
	"errors"

	"bakku.dev/dotf"
)

func Init(sys dotf.SysOpsProvider) error {
	home := sys.GetEnvVar("HOME")

	if home == "" {
		return errors.New("init: HOME env var is not set")
	}

	return nil
}
