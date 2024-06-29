/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"

	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [Module1 Module2 ...]",
	Short: "Uninstall a PowerShell module.",
	Long:  `Uninstall a PowerShell module.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.HelpFunc()(cmd, args)
		} else {
			for _, arg := range args {
				isInstalled := checkModuleInstalled(arg, true)
				if isInstalled {
					fmt.Println(styling.StyleStatusMsg(fmt.Sprintf("Getting module info for module '%s'...\n", arg)))
					results := findModule(arg, false)
					fmt.Print(styling.StyleStatusMsg(fmt.Sprintf("Removing module '%s'...\n", arg)))
					if results.Name != "" {
						switch {
						case len(results.Dependencies) > 0:
							var wg sync.WaitGroup
							wg.Add(len(results.Dependencies))
							for _, dep := range results.Dependencies {
								exists := checkModuleInstalled(dep.Name, false)
								if exists {
									go uninstallModuleParallel(dep.Name, false, &wg)
								}
							}
							uninstallModule(arg)
						default:
							uninstallModule(arg)
						}
					}
				}
			}
		}
	},
}

func checkModuleInstalled(module string, print bool) bool {
	query := fmt.Sprintf("Get-InstalledModule %s", module)
	if print {
		fmt.Print(styling.StyleStatusMsg(fmt.Sprintf("Checking if module '%s'is installed...\n", module)))
	}
	_, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Print(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	if stdErr.String() != "" {
		if print {
			fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not installed\n", module)))
		}
		return false
	} else {
		if print {
			fmt.Print(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Installed module '%s' found!\n", module)))
		}
		return true
	}
}

func uninstallModuleParallel(module string, print bool, wg *sync.WaitGroup) {
	defer wg.Done()
	query := fmt.Sprintf("Uninstall-Module %s -Force", module)
	if print {
		fmt.Print(styling.StyleStatusMsg(fmt.Sprintf("Removing module '%s'...\n", module)))
	}
	_, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Print(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	if stdErr.String() != "" {
		fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not uninstalled\n", module)))
	} else {
		if print {
			fmt.Print(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module %s successfully removed!\n", module)))
		}
	}
}

func uninstallModule(module string) {
	query := fmt.Sprintf("Uninstall-Module %s -Force", module)
	_, stdErr, err := runCommand(query)
	if err != nil {
		fmt.Print(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	if stdErr.String() != "" {
		fmt.Print(styling.StyleErrMsg(fmt.Errorf("[!] PowerShell module '%s' not uninstalled\n", module)))
	} else {
		fmt.Print(styling.StyleSuccessMsg(fmt.Sprintf("[✓] Module %s successfully removed!\n", module)))
	}
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
