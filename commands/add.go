package commands

import (
	"fmt"

	"bakku.dev/dotf"
)

// Add adds a file to the tracked files of dotf.
func Add(sys dotf.SysOpsProvider, systemFilePath, repoFilePath string) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	return trackNewFile(sys, dotfilePath, systemFilePath, repoFilePath)
}

func trackNewFile(sys dotf.SysOpsProvider, dotfilePath, systemFilePath, repoFilePath string) error {
	absoluteSystemFilePath, err := sys.ExpandPath(systemFilePath)

	if err != nil {
		return fmt.Errorf("add: could not build absolute path: %v", err)
	}

	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	trackedFiles := cfg.TrackedFiles
	trackedFiles = append(trackedFiles, dotf.TrackedFile{repoFilePath, absoluteSystemFilePath})
	cfg.TrackedFiles = trackedFiles

	err = writeConfig(sys, dotfilePath, cfg)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	return nil
}
