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
	"github.com/go-git/go-git/v5"
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

// ExpandPath builds the absolute path for a given path.
func (sop *Provider) ExpandPath(path string) (string, error) {
	return filepath.Abs(path)
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

// DeserializeConfig deserializes a JSON blob into a dotf.Config struct.
func (sop *Provider) DeserializeConfig(raw []byte, c *dotf.Config) error {
	return json.Unmarshal(raw, c)
}

// WriteFile takes a path and content and (over)writes the content to the given path.
func (sop *Provider) WriteFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}

// ReadFile takes a path and reads the contents into a byte array.
func (sop *Provider) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// CopyFile copies and overwrites src to dest.
func (sop *Provider) CopyFile(src, dest string) error {
	// if src does not exist (yet) do not try to copy
	if !sop.PathExists(src) {
		return nil
	}

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("could not read file %s", src)
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		return fmt.Errorf("could not write file %s", src)
	}

	return nil
}

// UpdateRepo updates a git repository.
func (sop *Provider) UpdateRepo(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("could not open repo %s: %v", path, err)
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("could not get worktree of %s: %v", path, err)
	}

	err = workTree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		if err.Error() != "already up-to-date" {
			return fmt.Errorf("could not pull repo %s: %v", path, err)
		}
	}

	return nil
}

// CommitRepo commits and pushes a git repository.
func (sop *Provider) CommitRepo(path, message string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("could not open repo %s: %v", path, err)
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("could not get worktree of %s: %v", path, err)
	}

	err = workTree.AddGlob(".")
	if err != nil {
		return fmt.Errorf("could not add files to repo %s: %v", path, err)
	}

	_, err = workTree.Commit(message, &git.CommitOptions{})
	if err != nil {
		return fmt.Errorf("could not commit to repo %s: %v", path, err)
	}

	err = repo.Push(&git.PushOptions{})
	if err != nil {
		return fmt.Errorf("could not push repo %s: %v", path, err)
	}

	return nil
}
