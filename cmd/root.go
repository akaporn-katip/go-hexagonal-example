/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/akaporn-katip/go-project-structure-template/cmd/serve"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() int {
	rootCmd := NewRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		return 1
	}
	return 0
}

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:              "crm",
		TraverseChildren: true,
	}

	rootCmd.AddCommand(serve.NewServeCmd())

	rootCmd.PersistentFlags().StringP("config", "c", "./config", "Config path")

	return rootCmd
}
