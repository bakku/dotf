package sysop

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"bakku.dev/dotf"
)

// Provider implements the dotf.SysOpProvider interface.
type Provider struct{}

// GetEnvVar returns an environment variable of the current environment.
func (sop *Provider) GetEnvVar(s string) string {
	return os.Getenv(s)
}

// GetPathSep returns the path separator of the current operating system.
func (sop *Provider) GetPathSep() string {
	return string(filepath.Separator)
}

// CleanPath cleans the given path from common error sources and returns it
func (sop *Provider) CleanPath(path string) string {
	return filepath.Clean(path)
}

// PathExists returns true if the given path exists, otherwise false.
func (sop *Provider) PathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

// Log writes some content to STDOUT.
func (sop *Provider) Log(message string) {
	fmt.Print(message)
}

// ReadLine reads a line from STDIN.
func (sop *Provider) ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error while reading: %v", err)
	}

	return strings.TrimSpace(text), nil
}

// SerializeConfig serializes an instance of dotf.Config to JSON.
func (sop *Provider) SerializeConfig(c dotf.Config) ([]byte, error) {
	return json.Marshal(c)
}

// WriteFile takes a path and content and (over)writes the content to the given path.
func (sop *Provider) WriteFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}
