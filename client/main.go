package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/wiikip/go-container/client/client"
	"github.com/wiikip/go-container/client/ui"
)


func main() {
	// Create a new application.
	myClient := client.Client{}
	application, err := gtk.ApplicationNew("com.test", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Println("Error while creating the application:", err)
	}
	ui.CreateMain(application, &myClient)




	// Connect to server 

	// Connect function to application startup event, this is not required.

	// Launch the application
	os.Exit(application.Run(os.Args))
}





// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
