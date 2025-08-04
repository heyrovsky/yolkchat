package ui

import (
	"fmt"
	"image/color"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/heyrovsky/yolkchat/internals/config"
	"github.com/heyrovsky/yolkchat/pkg/schema"
)

func BuildChatWindow(chat *schema.UserList, w fyne.Window, messages []schema.ChatMessage, onBack func(), sendFunc func(chatId string, chatContent schema.ChatMessage)) fyne.CanvasObject {
	showBack := w.Canvas().Size().Width <= 800

	var left fyne.CanvasObject
	if showBack {
		left = widget.NewButton("< Back", func() {
			if onBack != nil {
				onBack()
			}
		})
	} else {
		left = layout.NewSpacer()
	}

	title := config.APPNAME
	if chat != nil {
		title = chat.Name
	}

	topBar := container.NewBorder(
		nil, nil,
		left, nil,
		widget.NewLabelWithStyle(title,
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true}),
	)

	// Messages list
	messageContainer := container.NewVBox()
	for _, msg := range messages {
		bubble := buildMessageBubble(msg, w.Canvas().Size().Width*0.48) // 48% of width
		var aligned fyne.CanvasObject
		if msg.Ingress {
			aligned = container.NewHBox(bubble, layout.NewSpacer())
		} else {
			aligned = container.NewHBox(layout.NewSpacer(), bubble)
		}
		messageContainer.Add(aligned)
	}

	scroll := container.NewVScroll(messageContainer)
	scroll.SetMinSize(fyne.NewSize(400, 500))

	// Input area
	messageEntry := widget.NewMultiLineEntry()
	messageEntry.SetPlaceHolder("Type a message...")

	var attachedFiles []string
	var sendBtn *widget.Button

	updateSendBtnLabel := func() {
		if len(attachedFiles) > 0 {
			sendBtn.SetText(fmt.Sprintf("(%d)", len(attachedFiles)))
		} else {
			sendBtn.SetText("")
		}
	}

	attachBtn := widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if uc != nil && err == nil {
				attachedFiles = append(attachedFiles, uc.URI().Path())
				updateSendBtnLabel() // update label when new file is added
			}
		}, w)
		fd.Show()
	})

	sendBtn = widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		if chat != nil && (messageEntry.Text != "" || len(attachedFiles) > 0) {
			sendFunc(chat.Name, schema.ChatMessage{
				Ingress:     false,
				Text:        messageEntry.Text,
				Attachments: attachedFiles,
				Timestamp:   time.Now(),
			})

			messageEntry.SetText("")
			attachedFiles = []string{}
			updateSendBtnLabel() // reset label after sending
		}
	})

	clearBtn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		messageEntry.SetText("")
		attachedFiles = []string{}
		updateSendBtnLabel()
	})

	attachBtn.Importance = widget.LowImportance
	clearBtn.Importance = widget.LowImportance
	sendBtn.Importance = widget.HighImportance

	inputBar := container.NewBorder(nil, nil,
		container.NewVBox(attachBtn, clearBtn), sendBtn, messageEntry)

	return container.NewBorder(
		container.NewPadded(topBar),
		container.NewPadded(inputBar),
		nil,
		nil,
		scroll,
	)
}

func openFileInExplorer(path string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", path)
	case "windows":
		cmd = exec.Command("explorer", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}
	cmd.Start()
}

func buildMessageBubble(msg schema.ChatMessage, maxWidth float32) fyne.CanvasObject {
	var elements []fyne.CanvasObject

	for _, file := range msg.Attachments {
		displayName := strings.TrimSpace(file)
		btn := widget.NewButtonWithIcon(displayName, theme.DocumentIcon(), func(f string) func() {
			return func() {
				openFileInExplorer(f)
			}
		}(displayName))
		btn.Resize(fyne.NewSize(maxWidth, btn.MinSize().Height))
		elements = append(elements, btn)
	}
	if strings.TrimSpace(msg.Text) != "" {
		lbl := widget.NewLabel(msg.Text)
		lbl.Wrapping = fyne.TextWrapOff
		lbl.Alignment = fyne.TextAlignLeading
		lbl.Resize(fyne.NewSize(maxWidth, lbl.MinSize().Height))
		elements = append(elements, lbl)
	}

	content := container.NewVBox(elements...)

	timestamp := widget.NewLabel(msg.Timestamp.Format("15:04"))
	timestamp.TextStyle = fyne.TextStyle{Italic: true}
	timestamp.Alignment = fyne.TextAlignTrailing

	bubbleContent := container.NewBorder(nil, timestamp, nil, nil, content)
	paddedBubble := container.NewPadded(bubbleContent)

	bg := canvas.NewRectangle(color.NRGBA{R: 40, G: 41, B: 46, A: 255})
	bg.CornerRadius = 12

	return container.NewPadded(container.NewStack(bg, paddedBubble))
}
