package dotf

type SysOpsProvider interface {
	GetEnvVar(s string) string
	GetPathSep() string
	CleanPath(path string) string
	FileExists(path string) bool
	Log(message string)
}
