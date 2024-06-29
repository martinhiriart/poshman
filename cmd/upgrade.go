package cmd

import (
	"fmt"

	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, arg := range args {
				upgradeModule(arg)
			}
		} else {
			upgradeModule("")
		}
	},
}

func upgradeModule(module string) {
	if module == "" {
		// Add code to update all installed modules that have a pending update
		installed := listModules(false, "")
		if len(installed) <= 0 {
			fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] No modules are installed. Update not possible\n")))
		} else {
			for _, mod := range installed {
				query := fmt.Sprintf("Update-Module %s -Force", mod.Name)

				fmt.Print(styling.StyleStatusMsg(fmt.Sprintf("Updating module '%s'...\n", mod.Name)))
				_, stdErr, err := runCommand(query)
				if err != nil {
					fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", stdErr.String())))
				}
				if stdErr.String() != "" {
					fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not updated\n", mod.Name)))
				} else {
					fmt.Print(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module %s successfully updated!\n", mod.Name)))
				}
			}
		}
	} else {
		isInstalled := checkModuleInstalled(module, true)
		if isInstalled {
			query := fmt.Sprintf("Update-Module %s -Force", module)

			fmt.Print(styling.StyleStatusMsg(fmt.Sprintf("Updating module '%s'...\n", module)))
			_, stdErr, err := runCommand(query)
			if err != nil {
				fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] Error: %s\n", stdErr.String())))
			}
			if stdErr.String() != "" {
				fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not updated\n", module)))
			} else {
				fmt.Print(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module %s successfully updated!\n", module)))
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
