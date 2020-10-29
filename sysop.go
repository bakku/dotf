package dotf

type SysOpsProvider interface {
	GetEnvVar(s string) string
}
