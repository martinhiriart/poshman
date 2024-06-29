/*
Copyright © 2024 Martin Hiriart

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/martinhiriart/poshman/styling"
	"github.com/spf13/cobra"
)

// func getStdErr(stdErr io.ReadCloser) string {
// 	errStr, _ := io.ReadAll(stdErr)
// 	return fmt.Sprintf("%s", errStr)
// }

// func getStdOut(stdOut io.ReadCloser) string {
// 	outStr, _ := io.ReadAll(stdOut)
// 	return fmt.Sprintf("%s", outStr)
// }

func runCommand(cStr string) (stdOut, stdErr bytes.Buffer, errs error) {
	cmdStr := fmt.Sprintf("%s | ConvertTo-Json -Depth 100", cStr)
	cmd := exec.Command("pwsh", "-Command", cmdStr)
	var sOut bytes.Buffer
	var sErr bytes.Buffer
	cmd.Stdout = &sOut
	cmd.Stderr = &sErr
	err := cmd.Run()

	if err != nil {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error running command: %s\\n\"", err)))
	}
	if stdErr.String() != "" {
		fmt.Println(styling.StyleErrMsg(fmt.Errorf("\"[!] Error: %s\\n\"", stdErr.String())))
	}

	return sOut, sErr, err
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "poshman",
	Short: "A package manager for PowerShell modules",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		logo := `
		          __
   ___  ___  ___ / /  __ _  ___ ____
  / _ \/ _ \(_-</ _ \/  ' \/ _ ` + "`" + `/ _ \
 / .__/\___/___/_//_/_/_/_/\_,_/_//_/
/_/

		`
		fmt.Println(styling.StyleStatusMsg(logo))
		if len(args) < 1 {
			cmd.HelpFunc()(cmd, args)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.poshman.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
