package cmd

import (
	"fmt"

	"github.com/hashicorp/go-version"
	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Checks for updates to installed PowerShell modules",
	Long: `Checks for updates to installed PowerShell modules. This command can
	also be used to check for updates on a specific PowerShell module.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case len(args) <= 0:
			i := listModules(false, "")
			pendingUpdates := checkForUpdates(i)
			switch {
			case len(pendingUpdates) > 0:
				fmt.Println(styling.StyleStatusMsg("Modules with available updates:"))
				printInstalledModuleInfo(pendingUpdates)
			default:
				fmt.Println(styling.StyleSuccessMsg(fmt.Sprint("[✓] All installed modules up to date!")))
			}
		default:
			for _, arg := range args {
				m := listModules(false, arg)
				pendingUpdates := checkForUpdates(m)
				switch {
				case len(pendingUpdates) > 0:
					fmt.Println(styling.StyleStatusMsg("Modules with available updates:"))
					printInstalledModuleInfo(pendingUpdates)
				default:
					fmt.Println(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module: %s up to date!", arg)))
				}
			}
		}
	},
}

func checkForUpdates(i []InstalledModules) []InstalledModules {
	var pendingUpdates []InstalledModules
	for _, mod := range i {
		s := findModule(mod.Name, false)
		sv, err := version.NewVersion(s.Version)
		if err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", err)))
		}
		mv, err := version.NewVersion(mod.Version)
		if err != nil {
			fmt.Println(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", err)))
		}
		if sv.GreaterThan(mv) {
			pendingUpdates = append(pendingUpdates, mod)
		}
	}
	return pendingUpdates
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
