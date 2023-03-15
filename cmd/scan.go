/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/gopheramit/pScan/scan"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run a short scan on the hosts",

	RunE: func(cmd *cobra.Command, args []string) error {
		hostFile, err := cmd.Flags().GetString("host-file")
		if err != nil {
			return err
		}
		ports, err := cmd.Flags().GetIntSlice("ports")
		if err != nil {
			return err

		}
		return scanAction(os.Stdout, hostFile, ports)

	},
}

func scanAction(out io.Writer, hostFiles string, ports []int) error {
	hl := &scan.HostList{}
	if err := hl.Load(hostFiles); err != nil {
		return err
	}
	results := scan.Run(hl, ports)
	return printResults(out, results)
}
func printResults(out io.Writer, results []scan.Results) error {
	message := ""
	for _, r := range results {
		message += fmt.Sprintf("%s", r.Host)

		if r.NotFound {
			message += fmt.Sprintf("Host not found \n\n")
			continue
		}
		message += fmt.Sprintln()
		for _, p := range r.PortState {
			message += fmt.Sprintf("\t %d:%s\n", p.Port, p.Open)

		}
		message += fmt.Sprintln()
	}
	_, err := fmt.Fprint(out, message)
	return err
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().IntSliceP("ports", "p", []int{22, 80, 443}, "ports to scan")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
