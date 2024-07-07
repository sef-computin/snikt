package gui

import (
	"embed"
	"fmt"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sef-computin/snikt/sniff"
)

//go:embed glade/*
var fs embed.FS

var text_buf string
var packets_chan chan string

const (
	WindowName      = "main_window"
	TextBoxName     = "result_textview"
	SaveButtonName  = "save_button"
	StartButtonName = "start_button"
	UIMain          = "glade/window.glade"
)

func setupGTK() {
	gtk.Init(&os.Args)

	bldr, err := getBuilder()
	if err != nil {
		panic(err)
	}

	window, err := getWindow(bldr)
	if err != nil {
		panic(err)
	}

	result_textbox, err := getTextView(bldr, TextBoxName)
	if err != nil {
		panic(err)
	}
	result_textbox.SetEditable(false)

	buffer, err := result_textbox.GetBuffer()
	if err != nil {
		panic(err)
	}

	window.SetTitle("Snikt - simple http sniffer")
	window.SetDefaultSize(1200, 800)
	_ = window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	window.ShowAll()

	save_button, err := getButton(bldr, SaveButtonName)
	if err != nil {
		panic(err)
	}
	start_button, err := getButton(bldr, StartButtonName)
	if err != nil {
		panic(err)
	}

	_ = save_button.Connect("clicked", func() {
		text_buf += "click\n"
		buffer.SetText(text_buf)
	})
	_ = start_button.Connect("clicked", func() {
		start_button.SetVisible(false)
		packets_chan = make(chan string)

		go sniff.StartUtil("wlan0", packets_chan)

		go func() {
			for {
				msg := <-packets_chan
				text_buf = fmt.Sprintf("%s\n%s", text_buf, msg)
			}
		}()
	})

	_ = glib.TimeoutAdd(uint(1000), func() bool {
		buffer.SetText(text_buf)

		return true
	})

	gtk.Main()
}

// getBuilder returns *gtk.getBuilder loaded with glade resource (if resource is given)
func getBuilder() (*gtk.Builder, error) {
	b, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	bs, err := fs.ReadFile(UIMain)
	if err != nil {
		return nil, err
	}

	err = b.AddFromString(string(bs))
	if err != nil {
		return nil, err
	}

	return b, nil
}

// getWindow returns *gtk.Window object from the glade resource
func getWindow(b *gtk.Builder) (*gtk.Window, error) {
	obj, err := b.GetObject(WindowName)
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		return nil, err
	}

	return window, nil
}

func getButton(b *gtk.Builder, tag string) (*gtk.Button, error) {
	obj, err := b.GetObject(tag)
	if err != nil {
		return nil, err
	}

	button, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return button, nil
}

func getTextView(b *gtk.Builder, tag string) (*gtk.TextView, error) {
	obj, err := b.GetObject(tag)
	if err != nil {
		return nil, err
	}

	textview, ok := obj.(*gtk.TextView)
	if !ok {
		return nil, err
	}

	return textview, nil
}
