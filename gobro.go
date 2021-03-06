package main

import (
	"fmt"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/controllers/user"
	"github.com/magleff/gobro/database"
	"os"
)

func main() {
	database.InitDatabase()

	// Init a default user if it does not exist yet
	user.UserControllerImpl{}.InitDefault()

	// Close the database's main session after running the command
	defer database.GetSession().Close()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
