package commands

import (
	"errors"
	"fmt"
	"path/filepath"

	"bakku.dev/dotf"
)

// Add adds a file to the tracked files of dotf.
func Add(sys dotf.SysOpsProvider) error {
	home := sys.GetEnvVar("HOME")

	if home == "" {
		return errors.New("add: HOME env var is not set")
	}

	dotfilePath := sys.CleanPath(home + sys.GetPathSep() + dotfileName)

	if !sys.PathExists(dotfilePath) {
		sys.Log("No dotf configuration found. Please run the 'init' command first\n")
		return nil
	}

	return trackNewFile(sys, dotfilePath)
}

func trackNewFile(sys dotf.SysOpsProvider, dotfilePath string) error {
	sys.Log("Please insert the path of the file you want to track: ")
	systemFilePath, err := sys.ReadLine()

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	systemFilePath, err = sys.ExpandPath(systemFilePath)

	if err != nil {
		return fmt.Errorf("add: could not build absolute path: %v", err)
	}

	repoFilePath := filepath.Base(systemFilePath)

	sys.Log("Please insert the path of the file inside the repo (default " + repoFilePath + "): ")
	_, err = sys.ReadLine()

	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	rawConfig, err := sys.ReadFile(dotfilePath)

	if err != nil {
		return fmt.Errorf("add: could not read dotf config: %v", err)
	}

	cfg := dotf.Config{}
	err = sys.DeserializeConfig(rawConfig, &cfg)

	if err != nil {
		return fmt.Errorf("add: could not deserialize dotf config: %v", err)
	}

	trackedFiles := cfg.TrackedFiles
	trackedFiles = append(trackedFiles, dotf.TrackedFile{repoFilePath, systemFilePath})
	cfg.TrackedFiles = trackedFiles

	newRawConfig, err := sys.SerializeConfig(cfg)

	if err != nil {
		return fmt.Errorf("add: could not serialize dotf config: %v", err)
	}

	err = sys.WriteFile(dotfilePath, newRawConfig)

	if err != nil {
		return fmt.Errorf("add: could not write dotf config: %v", err)
	}

	return nil
}
