package ui

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Notification struct {
	text        *gtk.TextView
	displayTime int
}

type Notifier interface {
	Notify(string)
}

func (notif Notification) Notify(message string) error {
	buffer, err := notif.text.GetBuffer()
	if err != nil {
		return err
	}
	buffer.SetText(message)
	notif.text.Show()
	go notif.hideAfterWait()
	return nil

}

func (notif Notification) hideAfterWait() error {
	buffer, err := notif.text.GetBuffer()
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(notif.displayTime) * time.Second)
	buffer.SetText("")
	notif.text.Hide()
	return nil
}
