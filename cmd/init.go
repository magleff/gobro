package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new sheet for the budget",
	Long:  `Init a new sheet for the budget`,
	Run: func(cmd *cobra.Command, args []string) {
		balance := "0"
		if len(args) > 0 {
			balance = args[0]
		}
		DB := database.NewDatabase()
		controller := budget.NewController(DB)
		controller.CreateBudget(balance)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
