package commands

import (
	"fmt"

	"bakku.dev/dotf"
)

// Push copies all file to the repo, commits and pushes it.
func Push(sys dotf.SysOpsProvider, message string) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("push: %v", err)
	}

	return pushDotfiles(sys, dotfilePath, message)
}

func pushDotfiles(sys dotf.SysOpsProvider, dotfilePath, message string) error {
	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("push: %v", err)
	}

	for _, tf := range cfg.TrackedFiles {
		err = sys.CopyFile(
			tf.PathOnSystem,
			sys.CleanPath(cfg.Repo+sys.GetPathSep()+tf.PathInRepo),
		)

		if err != nil {
			return fmt.Errorf("push: %v", err)
		}
	}

	err = sys.CommitRepo(cfg.Repo, message)

	if err != nil {
		return fmt.Errorf("push: %v", err)
	}

	return nil
}
