package main

import (
	"github.com/{{cookiecutter.githubname}}/{{cookiecutter.botname}}/commands"
	"github.com/{{cookiecutter.githubname}}/{{cookiecutter.botname}}/bot"
	"github.com/{{cookiecutter.githubname}}/{{cookiecutter.botname}}/icon"

	"os"

	"github.com/getlantern/systray"
)

var (
	botState bot.State
)

func main() {
	botState = bot.GetDefaultState()
	botState.Selfbot = true

	// Add functionality like so
	botState.AddCommand(commands.PingCommand{})
	// making sure the struct you supply
	// inherits everything from the command interface.
	// See _examples/blank_command.go for a barebones command.

	go botState.Start(os.Getenv("{{cookiecutter.botname}}_token"))
	go systray.Run(onReady, onExit)

	<-make(chan struct{}) // Simple blocking.
}

// Required as of the latest systray updated, nothing to do in it yet though.
func onExit() {}

// Runs the system tray icon.
func onReady() {
	systray.SetIcon(icon.DataOn)
	systray.SetTitle("Bot")
	systray.SetTooltip("{{cookiecutter.botname}}")

	mToggle := systray.AddMenuItem("Toggle", "Toggle the bot off.")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		// When clicked exit, close everything
		case <-mQuit.ClickedCh:
			systray.Quit()
			os.Exit(0)
			return
		// Otherwise toggle the bot.
		case <-mToggle.ClickedCh:
			if botState.Running {
				botState.Running = false
				systray.SetIcon(icon.DataOff)
			} else {
				botState.Running = true
				systray.SetIcon(icon.DataOn)
			}
		}
	}
}
