package main

import (
	"embed"

	"wails-skeleton/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	application := app.NewApp()

	err := wails.Run(&options.App{
		Title:             app.AppName,
		Width:             1180,
		Height:            820,
		MinWidth:          980,
		MinHeight:         700,
		HideWindowOnClose: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 248, G: 245, B: 238, A: 1},
		OnStartup:        application.Startup,
		OnShutdown:       application.Shutdown,
		Bind: []interface{}{
			application,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
