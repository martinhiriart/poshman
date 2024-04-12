/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package search

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

type SearchModuleInfo struct {
	Name          string   `json:"Name"`
	Version       string   `json:"Version"`
	Type          string   `json:"Type"`
	Description   string   `json:"Description"`
	Author        string   `json:"Author"`
	CompanyName   string   `json:"CompanyName"`
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
	PowerShellGetFormatVersion any    `json:"PowerShellGetFormatVersion"`
	ReleaseNotes               any    `json:"ReleaseNotes"`
	Dependencies               []any  `json:"Dependencies"`
	RepositorySourceLocation   string `json:"RepositorySourceLocation"`
	Repository                 string `json:"Repository"`
	PackageManagementProvider  string `json:"PackageManagementProvider"`
	AdditionalMetadata         struct {
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

// SearchCmd represents the search command
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a specific PowerShell module",
	Long:  `Search for a specific PowerShell module`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.HelpFunc()(cmd, args)
		}
		//fmt.Println("search called")
		////if len(args) > 0 {
		////	fmt.Println(args[0])
		////}
		for _, arg := range args {
			findModule(arg)
		}

	},
}

func getStdErr(stdErr io.ReadCloser) string {
	errStr, _ := io.ReadAll(stdErr)
	return fmt.Sprintf("%s", errStr)
}

func findModule(module string) {
	queryText := fmt.Sprintf("pwsh -Command \"Find-Module %s | ConvertTo-Json -Depth 100\"", module)
	query := exec.Command("bash", "-c", queryText)
	//output, err := query.StdoutPipe()
	stdErr, err := query.StderrPipe()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	if err := query.Start(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Printf("Searching for module '%s'...\n", module)

	errStr := getStdErr(stdErr)
	if errStr != "" {
		fmt.Printf("PowerShell module '%s' not found\n", module)
	}
	if err := query.Wait(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	//switch {
	//case strings.Contains(string(query), fmt.Sprintf("Find-Package: No match was found for the specified search criteria and module name '%s'.", module)):
	//	fmt.Printf("ERROR: %s\n", query)
	//default:
	//	var modInfo SearchModuleInfo
	//	if err := json.Unmarshal(query, &modInfo); err != nil {
	//		fmt.Printf("ERROR: %v\n", err)
	//	}
	//	fmt.Println(modInfo)
	//}

}

func init() {
	//var mod  SearchModuleInfo{}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
