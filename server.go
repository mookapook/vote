package main

import (
	"9mookapook/vote/core"
	"9mookapook/vote/service"
)

func main() {
	app := core.APP()
	app.HideBanner = false
	service.Init()
	app.Logger.Fatal(app.Start(":8080"))
}
