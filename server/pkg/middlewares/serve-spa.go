package middlewares

import (
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func ServeSPA(webAppPath string) func(e *core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		subFs := echo.MustSubFS(e.Router.Filesystem, webAppPath)
		e.Router.GET("/*", apis.StaticDirectoryHandler(subFs, false), middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				return strings.HasPrefix(c.Path(), "/_") || strings.HasPrefix(c.Path(), "/api")
			},
		}))
		originalErrorHandler := e.Router.HTTPErrorHandler
		e.Router.HTTPErrorHandler = func(c echo.Context, err error) {
			if c.Path() == "/*" && err == echo.ErrNotFound {
				err = c.FileFS("index.html", subFs)
			}
			originalErrorHandler(c, err)
		}
		return nil
	}
}
