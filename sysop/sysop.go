package sysop

import "os"

type SysOpProvider struct {}

func (sop *SysOpProvider) GetEnvVar(s string) string {
	return os.Getenv(s)
}
