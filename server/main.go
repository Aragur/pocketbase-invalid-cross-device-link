package main

import (
	"example.com/example/pkg/commands"
	"example.com/example/pkg/middlewares"
	"example.com/example/pkg/utils"
	"github.com/pocketbase/pocketbase"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func main() {
	pterm.Info.Println("Loading environment variables...")
	utils.LoadEnv()
	pterm.Info.Println("Starting PocketBase server...")
	app := pocketbase.New()

	app.OnBeforeServe().Add(middlewares.ServeSPA(utils.GetEnv("WEBAPP", "./webapp")))
	app.OnBeforeServe().Add(middlewares.AddRegister(app))

	app.RootCmd.AddCommand(&cobra.Command{
        Use: "generate-tokens",
        Run: commands.GenerateTokens(app),
    })


	if err := app.Start(); err != nil {
		pterm.Fatal.Println(err)
	}

}
