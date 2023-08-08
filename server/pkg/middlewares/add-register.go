package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func AddRegister(app *pocketbase.PocketBase) func(e *core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
	e.Router.AddRoute(echo.Route{
			Method: http.MethodPut,
			Path:   "/api/register",
			Handler: func(c echo.Context) error {
				var json map[string]interface{} = map[string]interface{}{}
				if err := c.Bind(&json); err != nil {
					return err
				}
				tokenRecord, err := app.Dao().FindFirstRecordByData("tokens", "token", json["token"].(string))
				if tokenRecord.Get("used").(bool) {
					return c.String(http.StatusLocked, "")
				}
				collection, err := app.Dao().FindCollectionByNameOrId("users")
				if err != nil {
						return err
				}
				userRecord := models.NewRecord(collection)
				userRecord.Set("name", json["name"])
				userRecord.SetEmail(json["email"].(string))
				userRecord.SetPassword(json["password"].(string))
				userRecord.SetUsername(json["email"].(string))
				userRecord.Set("licenses", tokenRecord.Get("licenses"))
				if err := app.Dao().SaveRecord(userRecord); err != nil {
						return err
				}
				tokenRecord.Set("used", true)
				if err := app.Dao().SaveRecord(tokenRecord); err != nil {
					return err
				}
				return c.String(http.StatusCreated, "")
			},
		})
		return nil
	}
}