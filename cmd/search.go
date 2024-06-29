package cmd

import (
	"fmt"
	"sync"
	"time"

	json "github.com/json-iterator/go"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

type SearchResults struct {
	Name          string   `json:"Name"`
	Version       string   `json:"Version"`
	Type          string   `json:"Type"`
	Description   string   `json:"Description"`
	Author        string   `json:"Author"`
	CompanyName   any      `json:"CompanyName"`
	Copyright     string   `json:"Copyright"`
	PublishedDate string   `json:"PublishedDate"`
	InstalledDate any      `json:"InstalledDate"`
	UpdatedDate   any      `json:"UpdatedDate"`
	LicenseURI    string   `json:"LicenseUri"`
	ProjectURI    string   `json:"ProjectUri"`
	IconURI       any      `json:"IconUri"`
	Tags          []string `json:"Tags"`
	Includes      struct {
		Command        []string `json:"Command"`
		DscResource    []any    `json:"DscResource"`
		RoleCapability []any    `json:"RoleCapability"`
		Function       []string `json:"Function"`
		Workflow       []any    `json:"Workflow"`
		Cmdlet         []any    `json:"Cmdlet"`
	} `json:"Includes"`
	PowerShellGetFormatVersion any `json:"PowerShellGetFormatVersion"`
	ReleaseNotes               any `json:"ReleaseNotes"`
	Dependencies               []struct {
		Name            string `json:"Name"`
		MinimumVersion  string `json:"MinimumVersion,omitempty"`
		CanonicalID     string `json:"CanonicalId"`
		RequiredVersion string `json:"RequiredVersion,omitempty"`
	} `json:"Dependencies"`
	RepositorySourceLocation  string `json:"RepositorySourceLocation"`
	Repository                string `json:"Repository"`
	PackageManagementProvider string `json:"PackageManagementProvider"`
	AdditionalMetadata        struct {
		Summary                   string    `json:"summary"`
		Copyright                 string    `json:"copyright"`
		GUID                      string    `json:"GUID"`
		IsAbsoluteLatestVersion   string    `json:"isAbsoluteLatestVersion"`
		Created                   string    `json:"created"`
		DownloadCount             string    `json:"downloadCount"`
		IsLatestVersion           string    `json:"isLatestVersion"`
		Description               string    `json:"description"`
		Updated                   time.Time `json:"updated"`
		PackageSize               string    `json:"packageSize"`
		RequireLicenseAcceptance  string    `json:"requireLicenseAcceptance"`
		PackageManagementProvider string    `json:"PackageManagementProvider"`
		FileList                  string    `json:"FileList"`
		Authors                   string    `json:"Authors"`
		Published                 string    `json:"published"`
		CompanyName               string    `json:"CompanyName"`
		Tags                      string    `json:"tags"`
		NormalizedVersion         string    `json:"NormalizedVersion"`
		IsPrerelease              string    `json:"IsPrerelease"`
		LastUpdated               string    `json:"lastUpdated"`
		Functions                 string    `json:"Functions"`
		VersionDownloadCount      string    `json:"versionDownloadCount"`
		SourceName                string    `json:"SourceName"`
		ItemType                  string    `json:"ItemType"`
		DevelopmentDependency     string    `json:"developmentDependency"`
	} `json:"AdditionalMetadata"`
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [Module1 Module2 ...]",
	Short: "Search for a specific PowerShell module(s)",
	Long:  `Search for a specific PowerShell module(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.HelpFunc()(cmd, args)
		}
		if len(args) > 1 {
			var wg sync.WaitGroup
			wg.Add(len(args))
			for _, arg := range args {
				// fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Searching for module '%s'...\n", arg)))
				go findModuleParallel(arg, true, &wg)
			}
			wg.Wait()
		} else {
			// fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Searching for module '%s'...\n", args)))
			for _, arg := range args {
				findModule(arg, true)
			}
		}

	},
}

func printModuleSearchInfo(module string, s SearchResults) {
	fmt.Println(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module found: %s", module)))
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
		Headers("NAME", "VERSION", "REPOSITORY").
		Row(s.Name, s.Version, s.Repository).
		String()
	fmt.Println(t)
	fmt.Println()
}

func findModule(module string, print bool) SearchResults {
	var s SearchResults
	query := fmt.Sprintf("Find-Module -Name %s", module)

	stdOut, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", stdErr.String())))
	}

	if stdErr.String() != "" {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not found\n", module)))
	} else {
		if err := json.NewDecoder(&stdOut).Decode(&s); err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error parsing JSON output: %s\n", err)))
		}
		if print {
			printModuleSearchInfo(module, s)
		}
	}
	return s
}

func returnSearchResults(s SearchResults) SearchResults {
	return s
}

func findModuleParallel(module string, print bool, wg *sync.WaitGroup) SearchResults {
	defer wg.Done()
	var s SearchResults
	query := fmt.Sprintf("Find-Module %s", module)

	stdOut, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	if stdErr.String() != "" {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not found\n", module)))
	} else {
		if err := json.NewDecoder(&stdOut).Decode(&s); err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error parsing JSON output: %s\\n\"", err)))
		}
		if print {
			printModuleSearchInfo(module, s)
		}
	}
	return s
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
