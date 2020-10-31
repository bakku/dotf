package commands

import (
	"fmt"
	"path/filepath"

	"bakku.dev/dotf"
)

// Add adds a file to the tracked files of dotf.
func Add(sys dotf.SysOpsProvider) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	return trackNewFile(sys, dotfilePath)
}

func trackNewFile(sys dotf.SysOpsProvider, dotfilePath string) error {
	systemFilePath, err := readAbsoluteFilePath(sys, "Please insert the path of the file you want to track: ")

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	repoFilePath := filepath.Base(systemFilePath)

	sys.Log("Please insert the path of the file inside the repo (default " + repoFilePath + "): ")
	_, err = sys.ReadLine()

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	trackedFiles := cfg.TrackedFiles
	trackedFiles = append(trackedFiles, dotf.TrackedFile{repoFilePath, systemFilePath})
	cfg.TrackedFiles = trackedFiles

	err = writeConfig(sys, dotfilePath, cfg)

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	return nil
}
