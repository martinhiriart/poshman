package cmd

import (
	"fmt"

	json "github.com/json-iterator/go"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

var (
	descending bool
)

type InstalledModules struct {
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
	IconURI       any      `json:"IconUri"`
	Tags          []string `json:"Tags"`
	Includes      struct {
		Workflow       []any `json:"Workflow"`
		Function       []any `json:"Function"`
		Command        []any `json:"Command"`
		DscResource    []any `json:"DscResource"`
		RoleCapability []any `json:"RoleCapability"`
		Cmdlet         []any `json:"Cmdlet"`
	} `json:"Includes"`
	PowerShellGetFormatVersion any    `json:"PowerShellGetFormatVersion"`
	ReleaseNotes               any    `json:"ReleaseNotes"`
	Dependencies               []any  `json:"Dependencies"`
	RepositorySourceLocation   string `json:"RepositorySourceLocation"`
	Repository                 string `json:"Repository"`
	PackageManagementProvider  string `json:"PackageManagementProvider"`
	AdditionalMetadata         struct {
		Summary                   string `json:"summary"`
		PackageManagementProvider string `json:"PackageManagementProvider"`
		Type                      string `json:"Type"`
		SourceName                string `json:"SourceName"`
		Tags                      string `json:"tags"`
		Copyright                 string `json:"copyright"`
		Published                 string `json:"published"`
		Description               string `json:"description"`
		Installeddate             string `json:"installeddate"`
		InstalledLocation         string `json:"InstalledLocation"`
		IsPrerelease              bool   `json:"IsPrerelease"`
	} `json:"AdditionalMetadata"`
	InstalledLocation string `json:"InstalledLocation"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed PowerShell modules",
	Long:  `List all installed PowerShell modules`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styling.StyleStatusMsg(fmt.Sprint("Getting installed modules...")))
		listModules(true, "")
	},
}

func printInstalledModuleInfo(i []InstalledModules) {
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

func listModules(print bool, module string) []InstalledModules {
	var (
		m     InstalledModules
		i     []InstalledModules
		query string
	)
	if module == "" {
		if descending {
			query = fmt.Sprintf("Get-InstalledModule | Sort-Object Name -Descending")
		} else {
			query = fmt.Sprintf("Get-InstalledModule | Sort-Object Name")
		}
		stdOut, stdErr, err := runCommand(query)
		if err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", stdErr.String())))
		}
		if stdErr.String() != "" {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error retrieving installed modules")))
		} else {
			if stdOut.Len() <= 0 {
				fmt.Println(styling.StyleWarningMsg(fmt.Sprint("[!] No modules installed")))
			} else {
				switch string(stdOut.String()[0]) {
				case "{":
					if err := json.NewDecoder(&stdOut).Decode(&m); err != nil {
						fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error parsing JSON output: %s\n", err)))
					} else {
						i = append(i, m)
						if print {
							printInstalledModuleInfo(i)
						}
					}
				default:
					if err := json.NewDecoder(&stdOut).Decode(&i); err != nil {
						fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error parsing JSON output: %s\n", err)))
					} else {
						if print {
							printInstalledModuleInfo(i)
						}
					}
				}
			}

		}
	} else {
		query = fmt.Sprintf("Get-InstalledModule -Name %s | Sort-Object Name", module)
		stdOut, stdErr, err := runCommand(query)
		if err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", stdErr.String())))
		}
		if stdErr.String() != "" {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error retrieving installed modules")))
		} else {
			switch string(stdOut.String()[0]) {
			case "{":
				if err := json.NewDecoder(&stdOut).Decode(&m); err != nil {
					fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error parsing JSON output: %s\n", err)))
				} else {
					i = append(i, m)
					if print {
						printInstalledModuleInfo(i)
					}
				}
			default:
				if err := json.NewDecoder(&stdOut).Decode(&i); err != nil {
					fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error parsing JSON output: %s\n", err)))
				} else {
					if print {
						printInstalledModuleInfo(i)
					}
				}
			}
		}
	}
	return i
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
