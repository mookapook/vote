package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	app *echo.Echo
)

func APP() *echo.Echo {
	if app == nil {
		//LoadErrorCode()
		app = echo.New()
		app.HideBanner = false

		// app.HTTPErrorHandler = HTTPError
		// app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		// 	Level: 5,
		// }))
		app.Use(middleware.Logger())
		app.Use(middleware.Recover())
		// app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// 	TokenLookup: "header:X-XSRF-TOKEN",
		// 	ContextKey:  "csrf",
		// 	CookieName:  "csrf",
		// }))
		//getPublicKey()

		// app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// 	AllowOrigins: []string{"*"},
		// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		// }))

		app.Pre(middleware.RemoveTrailingSlash())

	}
	// DEVMODE = os.Getenv("DEVMODE")

	return app
}
