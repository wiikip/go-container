package ui

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/wiikip/go-container/client/client"
)

type PodsGrid struct {
	Pods     []client.PodsData
	PodsChan chan []client.PodsData
	Grid     *gtk.Grid
	Win      *gtk.ApplicationWindow
}
type DockerCreationSender struct {
	SendButton *gtk.Button
	UriEntry   *gtk.Entry
	NameEntry  *gtk.Entry
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
			log.Println("Error", err)
			return
		}

		win, err := gtk.ApplicationWindowNew(application)
		if err != nil {
			log.Println("Failed to create Application Windows: ", err)
		}
		grid, err := gtk.GridNew()
		if err != nil {
			log.Println("Failed to create Grid", err)
			return
		}
		grid.SetColumnSpacing(20)
		grid.SetRowSpacing(20)
		grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

		formGrid, err := createCreationForm(win, myClient)
		if err != nil {
			log.Println("Error while creating docker creation form: ", err)
		}

		gridPods, err := createPodsList(win)
		if err != nil {
			log.Println("Error while creating pods list: ", err)
		}

		go activateLiveUpdatePods(myClient, gridPods)

		// grid.Attach(formGrid, 0, 0, 5, 5)
		// grid.Attach(gridPods.Grid, 0, 1, 1, 1)
		grid.Add(formGrid)
		grid.Add(gridPods.Grid)
		win.Add(grid)
		win.ShowAll()
		// Show the Window and all of its components.
		application.AddWindow(win)

	})
	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		myClient.Disconnect()
		log.Println("application shutdown")
	})
}

func createCreationForm(win *gtk.ApplicationWindow, myClient *client.Client) (*gtk.Grid, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}
	var sender DockerCreationSender

	uriEntry, err := gtk.EntryNew()
	if err != nil {
		return nil, err
	}
	uriLabel, err := gtk.LabelNew("Docker's image URI")
	if err != nil {
		return nil, err
	}
	uriLabel.SetMarginEnd(30)
	uriLabel.SetMarginStart(30)
	nameEntry, err := gtk.EntryNew()
	if err != nil {
		return nil, err
	}
	nameLabel, err := gtk.LabelNew("Name for your pod")
	if err != nil {
		return nil, err
	}
	nameLabel.SetMarginEnd(30)
	nameLabel.SetMarginStart(30)

	btnCreation, err := gtk.ButtonNewWithLabel("Demander la création")
	if err != nil {
		return nil, err
	}

	btnCreation.SetActionName("win.SendData")
	// Field to chose Docker URI
	grid.Attach(uriLabel, 0, 0, 2, 1)
	grid.Attach(uriEntry, 3, 0, 2, 1)

	grid.Attach(nameLabel, 0, 1, 2, 1)
	grid.Attach(nameEntry, 3, 1, 2, 1)

	fmt.Println("Attached button")
	grid.Attach(btnCreation, 0, 2, 2, 1)

	grid.SetRowSpacing(5)
	sender.NameEntry = nameEntry
	sender.UriEntry = uriEntry
	sender.SendButton = btnCreation

	// Actions
	act := sender.createSendAction(myClient)
	win.AddAction(act)

	return grid, nil
}
func createPodsList(win *gtk.ApplicationWindow) (*PodsGrid, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}
	chanPods := make(chan []client.PodsData)
	gridPods := PodsGrid{[]client.PodsData{}, chanPods, grid, win}

	return &gridPods, nil

}

func activateLiveUpdatePods(myClient *client.Client, gridPods *PodsGrid) {
	for {
		time.Sleep(5 * time.Second)
		log.Println("Displaying")
		go myClient.FetchPods(gridPods.PodsChan)
		glib.IdleAdd(func() { DisplayPods(gridPods) })
	}
}
func DisplayPods(gridsPod *PodsGrid) {
	pods := <-gridsPod.PodsChan
	gridsPod.Pods = pods
	fmt.Println("Displaying")
	for ind, podInfo := range gridsPod.Pods {
		textName, err := gtk.TextViewNew()
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		textName.SetEditable(false)
		textName.SetCanFocus(false)
		textBuffer, err := textName.GetBuffer()
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		textBuffer.SetText(podInfo.Name)
		gridsPod.Grid.Attach(textName, 1, ind, 2, 1)
		gridsPod.Grid.ShowAll()
	}
}
