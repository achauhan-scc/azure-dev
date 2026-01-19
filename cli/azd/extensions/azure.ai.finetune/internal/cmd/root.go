// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"github.com/spf13/cobra"
)

type rootFlagsDefinition struct {
	Debug    bool
	NoPrompt bool
}

// Enable access to the global command flags
var rootFlags rootFlagsDefinition

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "finetuning <command> [options]",
		Short:         "Extension for Foundry Fine Tuning. (Preview)",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.PersistentFlags().BoolVar(
		&rootFlags.Debug,
		"debug",
		false,
		"Enable debug mode",
	)

<<<<<<< HEAD
	// Adds support for `--no-prompt` global flag in azd
	// Without this the extension command will error when the flag is provided
=======
	// Adds support for `--no-prompt` global flag in azd.
	// Without this the extension command will error when the flag is provided.
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
	rootCmd.PersistentFlags().BoolVar(
		&rootFlags.NoPrompt,
		"no-prompt",
		false,
<<<<<<< HEAD
		"Accepts the default value instead of prompting, or it fails if there is no default.",
	)

	rootCmd.AddCommand(newListenCommand())
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newInitCommand(rootFlags))
	rootCmd.AddCommand(newOperationCommand())
	// rootCmd.AddCommand(newOperationListCommand())
	//rootCmd.AddCommand(newOperationCheckpointsCommand())
=======
		"accepts the default value instead of prompting, or fails if there is no default",
	)

	// rootCmd.AddCommand(newListenCommand())
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newInitCommand(rootFlags))
	rootCmd.AddCommand(newOperationCommand())
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7

	return rootCmd
}
