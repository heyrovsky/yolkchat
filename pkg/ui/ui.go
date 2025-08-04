package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/heyrovsky/yolkchat/internals/config"
	"github.com/heyrovsky/yolkchat/internals/db"
	"github.com/heyrovsky/yolkchat/pkg/schema"
)

func PaintUi() error {
	canvas := app.NewWithID(config.APPID)
	canvas.Settings().SetTheme(noScrollBarTheme{Theme: theme.DefaultTheme()})
	window := canvas.NewWindow(config.APPNAME)

	storage_path := canvas.Storage().RootURI().Path()
	fmt.Println(storage_path)
	if err := db.Init(storage_path); err != nil {
		return err
	}

	if fyne.CurrentDevice().IsMobile() {
		window.SetFullScreen(true)
	} else {
		window.Resize(fyne.NewSize(400, 800))
	}

	currentView := 1 // 1 = chat list, 2 = chat window
	var selectedChat *schema.UserList

	allChats := []schema.UserList{
		{Name: "Alice", Online: true},
		{Name: "Bob"},
		{Name: "Charlie"},
	}

	messages := []schema.ChatMessage{
		{Ingress: true, Text: "Hey there!", Timestamp: time.Now().Add(-10 * time.Minute)},
		{Ingress: false, Text: "Hello Alice!", Timestamp: time.Now().Add(-8 * time.Minute)},
		{Ingress: false, Text: "Hey test", Attachments: []string{"/Users/syam/Downloads/genesis.json", "/Users/syam/Downloads/genesis.json"}, Timestamp: time.Now().Add(-5 * time.Minute)},
		{Ingress: true, Text: "Hey there!", Timestamp: time.Now().Add(-10 * time.Minute)},
		{Ingress: false, Text: "Hello Alice!", Timestamp: time.Now().Add(-8 * time.Minute)},
		{Ingress: true, Attachments: []string{"/Users/syam/Downloads/genesis.json"}, Timestamp: time.Now().Add(-5 * time.Minute)},
		{Ingress: true, Text: "Hey there!", Timestamp: time.Now().Add(-10 * time.Minute)},
		{Ingress: false, Text: "Hello Alice!", Timestamp: time.Now().Add(-8 * time.Minute)},
		{Ingress: true, Attachments: []string{"/Users/syam/Downloads/genesis.json"}, Timestamp: time.Now().Add(-5 * time.Minute)},
		{Ingress: true, Text: "Hey there!", Timestamp: time.Now().Add(-10 * time.Minute)},
		{Ingress: false, Text: "Hello Alice!", Timestamp: time.Now().Add(-8 * time.Minute)},
		{Ingress: true, Attachments: []string{"/Users/syam/Downloads/genesis.json"}, Timestamp: time.Now().Add(-5 * time.Minute)},
		{Ingress: true, Text: "Hey there!", Timestamp: time.Now().Add(-10 * time.Minute)},
		{Ingress: false, Text: "Hello Alice!", Timestamp: time.Now().Add(-8 * time.Minute)},
		{Ingress: true, Attachments: []string{"/Users/syam/Downloads/genesis.json"}, Timestamp: time.Now().Add(-5 * time.Minute)},
		{Ingress: true, Text: "Hey there!", Timestamp: time.Now().Add(-10 * time.Minute)},
		{Ingress: false, Text: "Hello Alice!", Timestamp: time.Now().Add(-8 * time.Minute)},
		{Ingress: true, Attachments: []string{"/Users/syam/Downloads/genesis.json"}, Timestamp: time.Now().Add(-5 * time.Minute)},
	}

	splitRatio := 0.3

	var updateContent func()

	updateContent = func() {
		width := window.Canvas().Size().Width
		if width > 800 {
			window.SetContent(container.NewPadded(BuildSplitView(
				BuildChatList(window, &allChats, func(chat schema.UserList) {
					selectedChat = &chat
					currentView = 2
					updateContent()
				}),
				BuildChatWindow(selectedChat, window, messages, func() {
					currentView = 1
					updateContent()
				}, func(chatID string, chatContent schema.ChatMessage) {
					fmt.Println(chatID, chatContent)
				}),
				&splitRatio,
			)))
			return
		}

		switch currentView {
		case 1:
			window.SetContent(BuildChatList(window, &allChats, func(chat schema.UserList) {
				selectedChat = &chat
				currentView = 2
				updateContent()
			}))
		case 2:
			window.SetContent(BuildChatWindow(selectedChat, window, messages, func() {
				currentView = 1
				updateContent()
			}, func(chatID string, chatContent schema.ChatMessage) {
				fmt.Println(chatID, chatContent)
			}))
		default:
			window.SetContent(BuildChatWindow(nil, window, nil, func() {
				currentView = 1
				updateContent()
			}, func(chatID string, chatContent schema.ChatMessage) {
				fmt.Println(chatID, chatContent)
			}))
		}
	}

	updateContent()

	go func() {
		prev := window.Canvas().Size()
		for {
			size := window.Canvas().Size()
			if size != prev {
				prev = size
				fyne.CurrentApp().SendNotification(&fyne.Notification{})
				fyne.Do(func() {
					updateContent()
				})
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	window.ShowAndRun()
	return nil
}
