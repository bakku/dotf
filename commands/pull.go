package commands

import (
	"fmt"

	"bakku.dev/dotf"
)

// Pull updates the repository and replaces all files with newly pulled ones.
func Pull(sys dotf.SysOpsProvider) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("pull: %v", err)
	}

	return updateDotfiles(sys, dotfilePath)
}

func updateDotfiles(sys dotf.SysOpsProvider, dotfilePath string) error {
	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("pull: %v", err)
	}

	err = sys.UpdateRepo(cfg.Repo)

	if err != nil {
		return fmt.Errorf("pull: %v", err)
	}

	for _, tf := range cfg.TrackedFiles {
		if cfg.CreateBackups {
			err = sys.CopyFile(
				tf.PathOnSystem,
				tf.PathOnSystem+".bk",
			)

			if err != nil {
				return fmt.Errorf("pull: %v", err)
			}
		}

		err = sys.CopyFile(
			sys.CleanPath(cfg.Repo+sys.GetPathSep()+tf.PathInRepo),
			tf.PathOnSystem,
		)

		if err != nil {
			return fmt.Errorf("pull: %v", err)
		}
	}

	return nil
}
