package cmd

import (
	"github.com/openag/vault-casc/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes new local vault configuration folder",
	Long: `Initializes local configuration folder by pulling configs from the specified 
vault instance.`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := internal.GetFlags()
		vault := internal.NewVaultClient(flags)
		resources := internal.InitResources(flags.GetDir())

		// dump policies
		vault.Policies(resources.WritePolicy)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
