package ui

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/wiikip/go-container/client/client"
)

type Sender struct {
	SendButton *gtk.Button
	MsgEntry *gtk.Entry
}

func CreateMain(application *gtk.Application, myClient *client.Client) {
		application.Connect("startup", func() {
		log.Println("application startup")
	})
	
	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")
		host := os.Getenv("HOST")

		err := myClient.Connect(host, "31384")
		if err != nil {
			log.Println("Error",err)
			return
		}
		win, err := gtk.ApplicationWindowNew(application)
		errorCheck(err)
		grid, err := gtk.GridNew()
		if err != nil {
			log.Println("Failed to create Grid", err)
			return
		}

		
		var sender Sender

		btn, err := gtk.ButtonNewWithLabel("Envoyer")
		errorCheck(err)
		textArea, err := gtk.EntryNew()
		if err != nil{
			log.Println("Error creating Text Area:", err)
			return
		}

		sender.MsgEntry = textArea
		sender.SendButton = btn
		
		act := sender.createSendAction(win, myClient)

		win.AddAction(act)

		grid.Add(sender.MsgEntry)
		grid.Add(sender.SendButton)

		btn.SetActionName("win.SendData")
		win.Add(grid)
		// Show the Window and all of its components.		
		win.ShowAll()
		application.AddWindow(win)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		myClient.Disconnect()
		log.Println("application shutdown")
	})
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Println(e)
	}
}