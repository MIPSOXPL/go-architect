package main

import (
	"github.com/MicroProcessingSolutions/go-architect/controller"
	"github.com/urfave/cli"
)

func main() {
	var application controller.Controller = new(controller.Data)
	application.SetApplication(cli.NewApp(), "go-architect", "go-architect is used to make packages from your go code and resources")
	application.PrepareApplication()
	application.Run()
}
