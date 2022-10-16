package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	myMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		myMenu.Append(menu.AppMenu())
		myMenu.Append(menu.EditMenu())
	}

	testMenu := myMenu.AddSubmenu("Test")
	testMenu.AddText("Hello world!", keys.CmdOrCtrl("h"), func(_ *menu.CallbackData) {
		wailsRuntime.MessageDialog(app.ctx, wailsRuntime.MessageDialogOptions{
			Type:    wailsRuntime.InfoDialog,
			Title:   "Hello world!",
			Message: "From the menubar",
		})
	})
	testMenu.AddSeparator()
	testMenu.AddText("Do other things", keys.CmdOrCtrl("d"), func(_ *menu.CallbackData) {
		wailsRuntime.MessageDialog(app.ctx, wailsRuntime.MessageDialogOptions{
			Type:    wailsRuntime.InfoDialog,
			Title:   "Do other things...",
			Message: "You just clicked a menubar item.",
		})
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "GraphGuard Presentation",
		Width:            1024,
		Height:           768,
		MinWidth:         400,
		MinHeight:        400,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Menu:             myMenu,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "GraphGuard Presentation",
				Message: "© 2022 GraphGuard",
				Icon:    icon,
			},
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			BackdropType:         windows.Auto,
			Theme:                windows.SystemDefault,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
