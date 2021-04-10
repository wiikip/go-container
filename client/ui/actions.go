package ui

import (
	"context"
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
		text, err := sender.MsgEntry.GetText()
		if err != nil{
			fmt.Println("Error getting text", err)
		}
		fmt.Println("Le texte est:", text)
			ctx := context.Background()
		err = myClient.SendMsg(ctx, client.GET_PODS, "10", func(s string){log.Println(s)})
		if err != nil {
			log.Println("error:",err)
			return
		}


	})

	return act
}