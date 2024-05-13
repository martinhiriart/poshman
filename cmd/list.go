/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

var (
	descending bool
)

type InstalledModules []struct {
	Name          string   `json:"Name"`
	Version       string   `json:"Version"`
	Type          string   `json:"Type"`
	Description   string   `json:"Description"`
	Author        string   `json:"Author"`
	CompanyName   string   `json:"CompanyName"`
	Copyright     string   `json:"Copyright"`
	PublishedDate string   `json:"PublishedDate"`
	InstalledDate string   `json:"InstalledDate"`
	UpdatedDate   any      `json:"UpdatedDate"`
	LicenseURI    string   `json:"LicenseUri"`
	ProjectURI    string   `json:"ProjectUri"`
	IconURI       string   `json:"IconUri"`
	Tags          []string `json:"Tags"`
	Includes      struct {
		RoleCapability []any `json:"RoleCapability"`
		Workflow       []any `json:"Workflow"`
		Cmdlet         []any `json:"Cmdlet"`
		Function       []any `json:"Function"`
		DscResource    []any `json:"DscResource"`
		Command        []any `json:"Command"`
	} `json:"Includes"`
	PowerShellGetFormatVersion any    `json:"PowerShellGetFormatVersion"`
	ReleaseNotes               string `json:"ReleaseNotes"`
	Dependencies               []any  `json:"Dependencies"`
	RepositorySourceLocation   string `json:"RepositorySourceLocation"`
	Repository                 string `json:"Repository"`
	PackageManagementProvider  string `json:"PackageManagementProvider"`
	AdditionalMetadata         struct {
		Summary                   string `json:"summary"`
		Description               string `json:"description"`
		InstalledLocation         string `json:"InstalledLocation"`
		Copyright                 string `json:"copyright"`
		Tags                      string `json:"tags"`
		Installeddate             string `json:"installeddate"`
		PackageManagementProvider string `json:"PackageManagementProvider"`
		ReleaseNotes              string `json:"releaseNotes"`
		Published                 string `json:"published"`
		SourceName                string `json:"SourceName"`
		Type                      string `json:"Type"`
		IsPrerelease              bool   `json:"IsPrerelease"`
	} `json:"AdditionalMetadata"`
	InstalledLocation string `json:"InstalledLocation"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listModules()
	},
}

func printInstalledModuleInfo(i InstalledModules) {
	tableHeaderStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}).
		Bold(true)
	tableBaseStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}).
		Align(lipgloss.Center)
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#366EBA"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return tableHeaderStyle
			}
			return tableBaseStyle
		}).
		Width(100).
		Headers("NAME", "VERSION", "REPOSITORY")
	for _, mod := range i {
		t.Row(mod.Name, mod.Version, mod.Repository)
	}
	fmt.Println(t)
	fmt.Println()
}

func listModules() {
	var (
		i     InstalledModules
		query string
	)
	if descending {
		query = fmt.Sprintf("Get-InstalledModule | Sort-Object Name -Descending")
	} else {
		query = fmt.Sprintf("Get-InstalledModule | Sort-Object Name")
	}

	fmt.Println(styling.StyleStatusMsg(fmt.Sprintln("Getting installed modules...")))
	stdOut, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}
	if stdErr.String() != "" {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error retrieving installed modules")))
	} else {
		if err := json.NewDecoder(&stdOut).Decode(&i); err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error parsing JSON output: %s\\n\"", err)))
		}
		printInstalledModuleInfo(i)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&descending, "desc", "d", false, "Sort installed modules by name in descending order")
}
