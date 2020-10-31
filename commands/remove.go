package commands

import (
	"fmt"

	"bakku.dev/dotf"
)

// Remove removes tracked files from dotf.
func Remove(sys dotf.SysOpsProvider, systemFilePath string) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("rm: %v", err)
	}

	return removeTrackedFile(sys, dotfilePath, systemFilePath)
}

func removeTrackedFile(sys dotf.SysOpsProvider, dotfilePath, systemFilePath string) error {
	absoluteSystemFilePath, err := sys.ExpandPath(systemFilePath)

	if err != nil {
		return fmt.Errorf("could not build absolute path: %v", err)
	}

	if err != nil {
		return fmt.Errorf("rm: %v", err)
	}

	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("rm: %v", err)
	}

	var trackedFileExists bool

	for i, trackedFile := range cfg.TrackedFiles {
		if trackedFile.PathOnSystem == absoluteSystemFilePath {
			trackedFileExists = true
			cfg.TrackedFiles = append(cfg.TrackedFiles[:i], cfg.TrackedFiles[i+1:]...)
			break
		}
	}

	if !trackedFileExists {
		return fmt.Errorf("rm: given file is not a tracked file")
	}

	err = writeConfig(sys, dotfilePath, cfg)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	return nil
}
