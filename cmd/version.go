package cmd

import (
	"fmt"

	"github.com/mheers/nats-seeder/models"
	"github.com/spf13/cobra"
)

// build flags
var (
	RuntimeInfo *models.RuntimeInfo
)
var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "prints the version",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
)

func printVersion() {
	fmt.Printf("\nRuntime Info:\n")
	fmt.Printf("Version: %s\n", RuntimeInfo.Version)
	fmt.Printf("BuildTime: %s\n", RuntimeInfo.BuildTime)
	fmt.Printf("CommitHash: %s\n", RuntimeInfo.CommitHash)
	fmt.Printf("CommitTime: %s\n", RuntimeInfo.CommitTime)
	fmt.Printf("DirtyBuild: %t\n", RuntimeInfo.DirtyBuild)
	fmt.Printf("GoVersion: %s\n", RuntimeInfo.GoVersion)
	fmt.Printf("GitTag: %s\n", RuntimeInfo.GitTag)
	fmt.Printf("GitBranch: %s\n", RuntimeInfo.GitBranch)
}
