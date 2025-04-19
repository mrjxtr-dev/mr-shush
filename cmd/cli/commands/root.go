package commands

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mr-shush",
	Short: "A Secure, Offline-First Password & Secrets Manager written in Go.",
	Long: `A Secure, Offline-First Password Generator & Secrets Manager written in Go.

mr-shush is a lightweight, offline-first password generator and secrets manager
built with speed and simplicity in mind — all thanks to Go. Designed for developers,
sysadmins, privacy-conscious users, and anyone who wants full control 
over their sensitive data without relying on the cloud.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("WELCOME TO MR-SHUSH CLI")
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init adds flags to rootCmd
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
