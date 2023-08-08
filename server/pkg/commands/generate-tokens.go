package commands

import (
	"strconv"
	"strings"

	"example.com/example/pkg/utils"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func GenerateTokens(app *pocketbase.PocketBase) func(_ *cobra.Command, args []string) {
	return func(_ *cobra.Command, args []string) {
		if len(args) != 2 {
			pterm.Error.Println("Please provide a number of tokens to generate and license ids to assign to them. Usage: generate-tokens 10 example")
			return
		}
		var amount, err = strconv.Atoi(args[0])
		if err != nil {
			pterm.Error.Println("Please provide a valid number of tokens to generate")
			return
		}
		for i := 0; i < amount; i++ {
			collection, err := app.Dao().FindCollectionByNameOrId("tokens")
				if err != nil {
					pterm.Error.Println(err)
					return
				}
			tokenRecord := models.NewRecord(collection)
			tokenRecord.Set("token", utils.RandStringRunes(6))
			tokenRecord.Set("licenses", strings.Split(args[1], ","))
			if err := app.Dao().SaveRecord(tokenRecord); err != nil {
					pterm.Error.Println(err)
					return
				}
		}
	}
}