package ui

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/wiikip/go-container/client/client"
)

type Sender struct {
	SendButton *gtk.Button
	UriEntry *gtk.Entry
	NameEntry *gtk.Entry
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
		pods := make(chan *client.ResponsePods)
		go myClient.SendMsg(context.TODO(),client.GET_PODS,"", func(s string){
			response := &client.ResponsePods{}
			fmt.Println(s)
			err :=json.Unmarshal([]byte(s),response)
			if err != nil {
				fmt.Println("ERROR:", err)
			}
			pods <- response
		})
		podsInfo := <- pods
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
		textAreaUri, err := gtk.EntryNew()
		if err != nil{
			log.Println("Error creating Text Area:", err)
			return
		}
		textAreaName, err := gtk.EntryNew()
		if err != nil{
			log.Println("Error creating Text Area:", err)
			return
		}


		sender.UriEntry = textAreaUri
		sender.SendButton = btn
		sender.NameEntry = textAreaName
		
		act := sender.createSendAction(win, myClient)

		win.AddAction(act)

		grid.Add(sender.UriEntry)
		grid.Add(sender.NameEntry)

		grid.Add(sender.SendButton)

		btn.SetActionName("win.SendData")

		// add info about running container
		for ind, podInfo := range podsInfo.Payload{

			textName, err := gtk.TextViewNew()
			if err != nil {
				fmt.Println("ERROR: ",err)
			}

			textBuffer,err := textName.GetBuffer()
			if err != nil {
				fmt.Println("ERROR: ",err)
			}

			textBuffer.SetText(podInfo.Name)
			grid.Attach(textName,1,2+ind,1,1)
		}
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