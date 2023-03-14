/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pScan",
	Short: "Fast TCP port scanner",
	Long: `pSacn-short for Port Scanner -executes port scan on list of hosts.
pScan allows you to add ,delere,hosts form the list.
pScan executes a port scan on specifies TCP ports,You can customize the target port using
a command line flag.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Version: "0.1",
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

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pScan.yaml)")
	rootCmd.PersistentFlags().StringP("host-file", "f", "pScan.hosts", "pScan hosts file")
	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.\
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
