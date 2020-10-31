package dotf

// TrackedFile represents a file that is being tracked by dotf.
type TrackedFile struct {
	PathInRepo   string `json:"pathInRepo"`
	PathOnSystem string `json:"pathOnSystem"`
}

// Config contains all attributes to parse the dotf config file.
type Config struct {
	Repo          string        `json:"repo"`
	CreateBackups bool          `json:"createBackups"`
	TrackedFiles  []TrackedFile `json:"trackedFiles"`
}
