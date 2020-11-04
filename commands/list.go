package commands

import (
	"fmt"
	"strings"

	"bakku.dev/dotf"
	"github.com/olekukonko/tablewriter"
)

// List shows all currently tracked files
func List(sys dotf.SysOpsProvider) error {
	dotfilePath, err := getDotfConfigPath(sys)

	if err != nil {
		return fmt.Errorf("list: %v", err)
	}

	return listAllTrackedFiles(sys, dotfilePath)
}

func listAllTrackedFiles(sys dotf.SysOpsProvider, dotfilePath string) error {
	cfg, err := readConfig(sys, dotfilePath)

	if err != nil {
		return fmt.Errorf("list: %v", err)
	}

	stringBuilder := &strings.Builder{}

	table := tablewriter.NewWriter(stringBuilder)
	table.SetHeader([]string{"File", "Path in repo"})

	for _, tf := range cfg.TrackedFiles {
		table.Append([]string{tf.PathOnSystem, tf.PathInRepo})
	}

	table.Render()

	sys.Log(stringBuilder.String())

	return nil
}
