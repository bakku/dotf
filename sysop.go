package dotf

type SysOpsProvider interface {
	GetEnvVar(s string) string
	GetPathSep() string
	PathExists(path string) bool
	Log(message string)
	ReadLine() (string, error)
}
