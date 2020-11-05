package dotf

// SysOpsProvider provides all system operation which dotf needs.
type SysOpsProvider interface {
	GetEnvVar(s string) string
	GetPathSep() string
	CleanPath(path string) string
	PathExists(path string) bool
	ExpandPath(path string) (string, error)
	Log(message string)
	ReadLine() (string, error)
	SerializeConfig(c Config) ([]byte, error)
	DeserializeConfig(raw []byte, c *Config) error
	WriteFile(path string, content []byte) error
	ReadFile(path string) ([]byte, error)
	CopyFile(src, dest string) error
	UpdateRepo(path string) error
	CommitRepo(path, message string) error
}
