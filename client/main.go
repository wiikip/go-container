package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/wiikip/go-container-client/client"
)

const (
	DOCKER_BUILD = "DOCKER_BUILD"
)

func main() {
	// Create a new application.
	myClient := client.Client{}
	application, err := gtk.ApplicationNew("com.test", glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)


	// Connect to server 

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})
	
	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")
		err := myClient.Connect("wiikip.viarezo.fr", "30897")
		if err != nil {
			log.Println("Error",err)
			return
		}
		win, err := gtk.ApplicationWindowNew(application)
		errorCheck(err)
		act := createSendAction(win, myClient)
		
		win.AddAction(act)

		btn, err := gtk.ButtonNewWithLabel("Envoyer")
		errorCheck(err)
	
		btn.SetActionName("win.SendData")
		win.Add(btn)
		// Show the Window and all of its components.		
		win.ShowAll()
		application.AddWindow(win)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		myClient.Disconnect()
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}


func createSendAction(win *gtk.ApplicationWindow, myClient client.Client) (*glib.SimpleAction){
	act := glib.SimpleActionNew("SendData", nil)


	act.Connect("activate", func(action *glib.SimpleAction){
		fmt.Println("Asking for Docker Creation")
			ctx := context.Background()
		err := myClient.SendMsg(ctx, DOCKER_BUILD, "10", func(s string){log.Println(s)})
		if err != nil {
			log.Println("error:",err)
			return
		}


	})

	return act
}
// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.