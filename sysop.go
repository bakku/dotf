package dotf

// SysOpsProvider provides all system operation which dotf needs.
type SysOpsProvider interface {
	GetEnvVar(s string) string
	GetPathSep() string
	CleanPath(path string) string
	PathExists(path string) bool
	Log(message string)
	ReadLine() (string, error)
	SerializeConfig(c Config) ([]byte, error)
	WriteFile(path string, content []byte) error
}
