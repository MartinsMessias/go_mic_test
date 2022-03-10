package main

import (
	"image/color"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const activeIcon = "mic.png"
const inactiveIcon = "mic_off.png"

var state = false

func main() {
	myApp := app.New()
	// set theme
	myApp.Settings().SetTheme(theme.LightTheme())
	// set icon
	myApp.Settings().Theme().Icon(activeIcon)

	myWindow := myApp.NewWindow("Mic Echo")
	myWindow.CenterOnScreen()

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Mic Echo, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Messias Martins"),
			), myWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)

	// Define a welcome text centered
	text := canvas.NewText("Test your microphone with pulseaudio", color.Black)
	text.Alignment = fyne.TextAlignCenter

	// get default device
	// comand := "pactl list sources short | grep RUNNING"
	// cmd := exec.Command("sh", "-c", comand)
	// stdout, _ := cmd.Output()
	// var str = string(stdout)
	// textLog := canvas.NewText(str, color.White)
	// textLog.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, _ = fyne.LoadResourceFromPath(inactiveIcon)
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 475, Height: 475}) // by default size is 0, 0

	// Define a button
	endisBtn := widget.NewButton("Enable/Disable", func() {
		state = !state

		if state {
			comand := "pactl load-module module-loopback"
			cmd := exec.Command("sh", "-c", comand)
			_, err := cmd.Output()

			if err != nil {
				fyne.LogError("Error:", err)
			}
			resource, _ = fyne.LoadResourceFromPath(activeIcon)
		} else {
			comand := "pactl unload-module module-loopback"
			cmd := exec.Command("sh", "-c", comand)
			_, err := cmd.Output()

			if err != nil {
				fyne.LogError("Error:", err)
			}
			resource, _ = fyne.LoadResourceFromPath(inactiveIcon)
		}

		gopherImg.Resource = resource
		//Redrawn the image with the new path
		gopherImg.Refresh()
	})
	endisBtn.Importance = widget.HighImportance

	// Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		gopherImg,
		// textLog,
		endisBtn,
	)

	// Display our content
	myWindow.SetContent(box)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}
