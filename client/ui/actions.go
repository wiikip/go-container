package ui

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/wiikip/go-container/client/client"
)

func (sender *DockerCreationSender) createSendAction(myClient *client.Client) *glib.SimpleAction {
	act := glib.SimpleActionNew("SendData", nil)

	act.Connect("activate", func(action *glib.SimpleAction) {
		fmt.Println("Asking for Docker Creation")
		uri, err := sender.UriEntry.GetText()
		if err != nil {
			fmt.Println("Error getting text", err)
		}

		name, err := sender.NameEntry.GetText()
		if err != nil {
			fmt.Println("Error getting text", err)
		}

		podData := client.PodsData{name, uri}

		payload, err := json.Marshal(podData)
		if err != nil {
			log.Println("Failed to marshal req: ", err)
		}

		ctx := context.Background()

		err = myClient.SendMsg(ctx, client.DOCKER_BUILD, string(payload), func(s string) {
			log.Println(s)
			sender.NameEntry.SetText("")
			sender.UriEntry.SetText("")
		})
		if err != nil {
			log.Println("error:", err)
			return
		}

	})

	return act
}

func (gridPods *PodsGrid) createRefreshAction(myClient *client.Client) *glib.SimpleAction {
	act := glib.SimpleActionNew("RefreshData", nil)

	act.Connect("activate", func(action *glib.SimpleAction) {
		ctx := context.Background()

		err := myClient.SendMsg(ctx, client.GET_PODS, "{}", func(s string) {
			log.Println(s)
			response := &client.ResponsePods{}
			fmt.Println(s)
			err := json.Unmarshal([]byte(s), response)
			if err != nil {
				log.Println("ERROR: ", err)
			}
			gridPods.PodsChan <- response.Payload
			DisplayPods(gridPods)
		})
		if err != nil {
			log.Println("error:", err)
			return
		}

	})
	return act

}
