package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/spf13/cobra"
)

func parseArguments(args []string) (string, string) {
	var amount string
	var description string

	amount = args[1]

	if len(args) > 2 {
		description = args[2]
	}

	return amount, description
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add something",
	Long:  `Add something`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, description := parseArguments(args)

		DB := database.NewDatabase()
		controller := expensefixed.NewController(DB)
		controller.CreateExpenseFixed(amount, description)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
