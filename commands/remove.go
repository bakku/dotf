package commands

import (
	"fmt"

	"bakku.dev/dotf"
)

// Remove removes tracked files from dotf.
func Remove(sys dotf.SysOpsProvider) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("rm: %v", err)
	}

	return removeTrackedFile(sys, dotfilePath)
}

func removeTrackedFile(sys dotf.SysOpsProvider, dotfilePath string) error {
	systemFilePath, err := readAbsoluteFilePath(sys, "Please insert the path of the file you want dotf to stop tracking: ")

	if err != nil {
		return fmt.Errorf("rm: %v", err)
	}

	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("rm: %v", err)
	}

	var trackedFileExists bool

	for i, trackedFile := range cfg.TrackedFiles {
		if trackedFile.PathOnSystem == systemFilePath {
			trackedFileExists = true
			cfg.TrackedFiles = append(cfg.TrackedFiles[:i], cfg.TrackedFiles[i+1:]...)
			break
		}
	}

	if !trackedFileExists {
		return fmt.Errorf("rm: given file is not tracked")
	}

	err = writeConfig(sys, dotfilePath, cfg)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	return nil
}
