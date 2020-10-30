package sysop

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SysOpProvider struct{}

func (sop *SysOpProvider) GetEnvVar(s string) string {
	return os.Getenv(s)
}

func (sop *SysOpProvider) GetPathSep() string {
	return string(filepath.Separator)
}

func (sop *SysOpProvider) PathExists(path string) bool {
	if _, err := os.Stat(filepath.Clean(path)); err == nil {
		return true
	}

	return false
}

func (sop *SysOpProvider) Log(message string) {
	fmt.Print(message)
}

func (sop *SysOpProvider) ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error while reading: %v", err)
	}

	return strings.TrimSpace(text), nil
}
