package sysop

import (
	"fmt"
	"os"
	"path/filepath"
)

type SysOpProvider struct {}

func (sop *SysOpProvider) GetEnvVar(s string) string {
	return os.Getenv(s)
}

func (sop *SysOpProvider) GetPathSep() string {
	return string(filepath.Separator)
}

func (sop *SysOpProvider) CleanPath(path string) string {
	return filepath.Clean(path)
}

func (sop *SysOpProvider) FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func (sop *SysOpProvider) Log(message string) {
	fmt.Println(message)
}
