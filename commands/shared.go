package commands

import (
	"errors"
	"fmt"

	"bakku.dev/dotf"
)

const dotfileName = ".dotf"

func getDotfConfigPath(sys dotf.SysOpsProvider) (string, error) {
	home := sys.GetEnvVar("HOME")

	if home == "" {
		return "", errors.New("HOME env var is not set")
	}

	dotfilePath := sys.CleanPath(home + sys.GetPathSep() + dotfileName)

	if !sys.PathExists(dotfilePath) {
		return "", errors.New("no dotf configuration found. Please run the 'init' command first")
	}

	return dotfilePath, nil
}

func readAbsoluteFilePath(sys dotf.SysOpsProvider, prompt string) (string, error) {
	sys.Log(prompt)
	path, err := sys.ReadLine()

	if err != nil {
		return "", err
	}

	path, err = sys.ExpandPath(path)

	if err != nil {
		return "", fmt.Errorf("could not build absolute path: %v", err)
	}

	return path, nil
}

func readConfig(sys dotf.SysOpsProvider, dotfilePath string) (dotf.Config, error) {
	rawConfig, err := sys.ReadFile(dotfilePath)

	if err != nil {
		return dotf.Config{}, fmt.Errorf("could not read dotf config: %v", err)
	}

	cfg := dotf.Config{}
	err = sys.DeserializeConfig(rawConfig, &cfg)

	if err != nil {
		return dotf.Config{}, fmt.Errorf("could not deserialize dotf config: %v", err)
	}

	return cfg, nil
}

func writeConfig(sys dotf.SysOpsProvider, dotfilePath string, cfg dotf.Config) error {
	newRawConfig, err := sys.SerializeConfig(cfg)

	if err != nil {
		return fmt.Errorf("could not serialize dotf config: %v", err)
	}

	err = sys.WriteFile(dotfilePath, newRawConfig)

	if err != nil {
		return fmt.Errorf("could not write dotf config: %v", err)
	}

	return nil
}
