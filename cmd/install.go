package cmd

import (
	"fmt"
	"strings"

	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [Module1 Module2 ...]",
	Short: "Install specific PowerShell module(s)",
	Long:  `Install specific PowerShell module(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.HelpFunc()(cmd, args)
		} else {
			for _, arg := range args {
				var (
					v string
					m string
				)
				switch {
				case strings.Contains(arg, "@"):
					m = strings.SplitN(arg, "@", 2)[0]
					v = strings.SplitN(arg, "@", 2)[1]
					fmt.Printf("Module %s, version %s requested\n", m, v)
					fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Searching for module '%s'...\n", m)))
					results := findModule(m, false)
					if results.Name != "" {
						toggleTrustedRepo("Trusted", "PSGallery", false)
						fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Installing module '%s'...", m)))
						installModule(m, v)
						toggleTrustedRepo("Untrusted", "PSGallery", false)
					}
				default:
					fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Searching for module '%s'...\n", arg)))
					results := findModule(arg, false)
					if results.Name != "" {
						toggleTrustedRepo("Trusted", "PSGallery", false)
						fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Installing module '%s'...", arg)))
						installModule(arg, "")
						toggleTrustedRepo("Untrusted", "PSGallery", false)
					}
				}
			}
		}
	},
}

func toggleTrustedRepo(status, repo string, print bool) {
	query := fmt.Sprintf("Set-PSRepository -Name '%s' -InstallationPolicy %s", repo, status)
	if print {
		fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Setting '%s' status to %s...\n", repo, status)))
	}
	_, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	if stdErr.String() != "" {
		if print {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] '%s' trust status not updated correctly\n", repo)))
		}
	} else {
		if print {
			fmt.Println(styling.StyleSuccessMsg(fmt.Sprintf("[✓] '%s' trust status updated correctly!\n", repo)))
		}
	}
}

func installModule(module, ver string) {
	var query string
	switch {
	case ver != "":
		query = fmt.Sprintf("Install-Module %s -RequiredVersion %s", module, ver)
	default:
		query = fmt.Sprintf("Install-Module %s", module)
	}

	_, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	if stdErr.String() != "" {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not installed\n", module)))
	} else {
		fmt.Println(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module %s successfully installed!\n", module)))
	}
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
