package cassandra

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
	"github.com/wal-g/wal-g/cmd/common"
	"github.com/wal-g/wal-g/internal"
	conf "github.com/wal-g/wal-g/internal/config"
)

const WalgShortDescription = "Cassandra backup tool"

var (
	// These variables are here only to show current version. They are set in makefile during build process
	walgVersion = "devel"
	gitRevision = "devel"
	buildDate   = "devel"

	Cmd = &cobra.Command{
		Use:     "wal-g",
		Short:   WalgShortDescription,
		Version: strings.Join([]string{walgVersion, gitRevision, buildDate, "Cassandra"}, "\t"),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if _, ok := cmd.Annotations["NoStorage"]; !ok {
				err := internal.AssertRequiredSettingsSet()
				tracelog.ErrorLogger.FatalOnError(err)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() {
	configureCommand()
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetCmd() *cobra.Command {
	return Cmd
}

func configureCommand() {
	common.Init(Cmd, conf.CASSANDRA)
	conf.AddTurboFlag(Cmd)
}
