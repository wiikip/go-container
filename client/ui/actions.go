package ui

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/wiikip/go-container/client/client"
)

func (sender *Sender) createSendAction(win *gtk.ApplicationWindow, myClient *client.Client ) (*glib.SimpleAction){
	act := glib.SimpleActionNew("SendData", nil)

	act.Connect("activate", func(action *glib.SimpleAction){
		fmt.Println("Asking for Docker Creation")
		uri, err := sender.UriEntry.GetText()
		if err != nil{
			fmt.Println("Error getting text", err)
		}
				
		name, err := sender.NameEntry.GetText()
		if err != nil{
			fmt.Println("Error getting text", err)
		}

		podData := client.PodsData{name,uri}

		payload, err := json.Marshal(podData)
		if err != nil{
			log.Println("Failed to marshal req: ",err)
		}

		ctx := context.Background()

		err = myClient.SendMsg(ctx, client.DOCKER_BUILD, payload, func(s string){log.Println(s)})
		if err != nil {
			log.Println("error:",err)
			return
		}


	})

	return act
}