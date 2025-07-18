// Copyright 2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
	Date    string

	rootCmd = &cobra.Command{
		SilenceUsage: true,
		Use:          "openapi-changes",
		Short:        "openapi-changes will tell you what has changed between one or more OpenAPI / Swagger specifications.",
		Long: `openapi-changes can detect every change found in an OpenAPI specification.
it can compare between two files, or a single file, over time.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			PrintBanner()

			fmt.Println("You have a few options when it comes to commands...")
			fmt.Println()

			_ = pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
				{Level: 0, Text: "console", TextStyle: pterm.NewStyle(pterm.FgLightCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgLightMagenta)},
				{Level: 0, Text: "summary", TextStyle: pterm.NewStyle(pterm.FgLightCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgLightMagenta)},
				{Level: 0, Text: "report", TextStyle: pterm.NewStyle(pterm.FgLightCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgLightMagenta)},
				{Level: 0, Text: "html-report", TextStyle: pterm.NewStyle(pterm.FgLightCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgLightMagenta)},
				{Level: 0, Text: "markdown-report", TextStyle: pterm.NewStyle(pterm.FgLightCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgLightMagenta)},
			}).Render()

			pterm.Printf("For more help, use the %s flag with any command.", pterm.LightMagenta("--help"))
			fmt.Println()
			fmt.Println()
			return nil
		},
	}
)

func Execute(version, commit, date string) {
	Version = version
	Commit = commit
	Date = date

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(GetConsoleCommand())
	rootCmd.AddCommand(GetSummaryCommand())
	rootCmd.AddCommand(GetReportCommand())
	rootCmd.AddCommand(GetHTMLReportCommand())
	rootCmd.AddCommand(GetMarkdownReportCommand())
	rootCmd.PersistentFlags().BoolP("top", "t", false, "Only show latest changes (last git revision against HEAD)")
	rootCmd.PersistentFlags().IntP("limit", "l", 5, "Limit history to number of revisions (default is 5)")
	rootCmd.PersistentFlags().BoolP("global-revisions", "R", false, "Consider all revisions in limit, not just the ones for the file")
	rootCmd.PersistentFlags().IntP("limit-time", "d", -1, "Limit history to number of days. Supersedes limit argument if present.")
	rootCmd.PersistentFlags().BoolP("no-logo", "b", false, "Don't print the big purple pb33f banner")
	rootCmd.PersistentFlags().StringP("base", "p", "", "Base URL or path to use for resolving relative or remote references")
	rootCmd.PersistentFlags().StringP("base-commit", "", "", "Base commit to compare against (will check until commit is found or limit is reached -- make sure to not shallow clone)")
	rootCmd.PersistentFlags().BoolP("remote", "r", true, "Allow remote reference (URLs and files) to be auto resolved, without a base URL or path (default is on)")
	rootCmd.PersistentFlags().BoolP("ext-refs", "", false, "Turn on $ref lookups and resolving for extensions (x-) objects")
}

func initConfig() {

}
