package main

import (
	a "hometask/app"
)

func main() {
	app := a.App{}
	app.Init()
	app.Run()
	<-app.Exit
}
