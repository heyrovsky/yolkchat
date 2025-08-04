// package ui

// import (
// 	"fmt"
// 	"image/color"
// 	"strings"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/heyrovsky/yolkchat/internals/utils"
// 	"github.com/heyrovsky/yolkchat/pkg/schema"
// )

// func BuildChatList(allChats *[]schema.UserList, onSelect func(chat schema.UserList)) fyne.CanvasObject {
// 	filteredChats := make([]schema.UserList, len(*allChats))
// 	copy(filteredChats, *allChats)

// 	title := canvas.NewText("Yolk", color.White)
// 	title.TextSize = 20
// 	title.TextStyle = fyne.TextStyle{Bold: true}

// 	addBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
// 		fmt.Println("Add new chat")
// 	})
// 	settingBtn := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
// 		fmt.Println("Settings clicked")
// 	})
// 	userBtn := widget.NewButtonWithIcon("", theme.AccountIcon(), func() {
// 		fmt.Println("User clicked")
// 	})

// 	topBar := container.NewBorder(nil, nil, nil,
// 		container.NewHBox(addBtn, settingBtn, userBtn),
// 		title)

// 	chatList := widget.NewList(
// 		func() int { return len(filteredChats) },
// 		func() fyne.CanvasObject {
// 			// avatar := widget.NewIcon(nil)
// 			// name := widget.NewLabel("Name")
// 			// lastSeen := widget.NewLabel("Last seen")
// 			// lastSeen.TextStyle = fyne.TextStyle{Italic: true}
// 			// details := container.NewHBox(avatar, container.NewVBox(name, lastSeen))

// 			// return container.NewBorder(nil, nil, details, nil)
// 			//
// 			avatar := widget.NewIcon(nil)
// 			name := widget.NewLabel("Name")
// 			name.TextStyle = fyne.TextStyle{Bold: true}

// 			placeholder := widget.NewLabel(" ")
// 			lastSeen := widget.NewLabel("Last seen")
// 			lastSeen.TextStyle = fyne.TextStyle{Italic: true}

// 			leftside := container.NewHBox(avatar, container.NewCenter(name))
// 			rightside := container.NewVBox(placeholder, lastSeen)

// 			return container.NewBorder(nil, nil, leftside, rightside)
// 		},
// 		func(i widget.ListItemID, o fyne.CanvasObject) {
// 			chat := filteredChats[i]
// 			border := o.(*fyne.Container)
// 			leftside := border.Objects[0].(*fyne.Container)
// 			rightside := border.Objects[1].(*fyne.Container)

// 			avatar := leftside.Objects[0].(*widget.Icon)
// 			name := leftside.Objects[1].(*fyne.Container).Objects[0].(*widget.Label)

// 			lastSeen := rightside.Objects[1].(*widget.Label)

// 			avatar.SetResource(utils.GenerateAvatar(chat.Name))
// 			name.SetText(chat.Name)
// 			if chat.Online {
// 				lastSeen.SetText("Online")
// 			} else {
// 				lastSeen.SetText(chat.LastSeen.String())
// 			}
// 		},
// 	)

// 	chatList.OnSelected = func(id widget.ListItemID) {
// 		onSelect(filteredChats[id])
// 	}

// 	search := widget.NewEntry()
// 	search.SetPlaceHolder("Search chats...")

// 	clearSearch := func() {
// 		search.SetText("")
// 		filteredChats = make([]schema.UserList, len(*allChats))
// 		copy(filteredChats, *allChats)
// 		chatList.Refresh()
// 	}

// 	search.OnChanged = func(text string) {
// 		lower := strings.ToLower(text)
// 		filtered := make([]schema.UserList, 0)
// 		for _, c := range *allChats {
// 			if strings.Contains(strings.ToLower(c.Name), lower) {
// 				filtered = append(filtered, c)
// 			}
// 		}
// 		filteredChats = filtered
// 		chatList.Refresh()
// 	}

// 	searchBar := container.NewBorder(nil, nil, nil,
// 		widget.NewButtonWithIcon("", theme.CancelIcon(), func() { clearSearch() }),
// 		search)

//		return container.NewPadded(
//			container.NewBorder(
//				container.NewPadded(topBar),
//				nil,
//				nil,
//				nil,
//				container.NewBorder(container.NewPadded(searchBar), nil, nil, nil, chatList),
//			),
//		)
//	}
package ui

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/heyrovsky/yolkchat/internals/utils"
	"github.com/heyrovsky/yolkchat/pkg/schema"
)

type ChatItem struct {
	widget.BaseWidget
	Content      *fyne.Container
	OnRightClick func()
}

func NewChatItem(content *fyne.Container, onRightClick func()) *ChatItem {
	item := &ChatItem{
		Content:      content,
		OnRightClick: onRightClick,
	}
	item.ExtendBaseWidget(item)
	return item
}

func (c *ChatItem) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.Content)
}

func (c *ChatItem) TappedSecondary(*fyne.PointEvent) {
	if c.OnRightClick != nil {
		c.OnRightClick()
	}
}

func BuildChatList(win fyne.Window, allChats *[]schema.UserList, onSelect func(chat schema.UserList)) fyne.CanvasObject {
	filteredChats := make([]schema.UserList, len(*allChats))
	copy(filteredChats, *allChats)

	title := canvas.NewText("Yolk", color.White)
	title.TextSize = 20
	title.TextStyle = fyne.TextStyle{Bold: true}

	addBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		fmt.Println("Add new chat")
	})
	settingBtn := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		fmt.Println("Settings clicked")
	})
	userBtn := widget.NewButtonWithIcon("", theme.AccountIcon(), func() {
		fmt.Println("User clicked")
	})

	topBar := container.NewBorder(nil, nil, nil,
		container.NewHBox(addBtn, settingBtn, userBtn),
		title)

	var chatList *widget.List
	chatList = widget.NewList(
		func() int { return len(filteredChats) },
		func() fyne.CanvasObject {
			avatar := widget.NewIcon(nil)
			name := widget.NewLabel("Name")
			name.TextStyle = fyne.TextStyle{Bold: true}

			placeholder := widget.NewLabel(" ")
			lastSeen := widget.NewLabel("Last seen")
			lastSeen.TextStyle = fyne.TextStyle{Italic: true}

			leftside := container.NewHBox(avatar, container.NewCenter(name))
			rightside := container.NewVBox(placeholder, lastSeen)

			item := container.NewBorder(nil, nil, leftside, rightside)
			return NewChatItem(item, nil)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			chat := filteredChats[i]
			chatItem := o.(*ChatItem)
			border := chatItem.Content
			leftside := border.Objects[0].(*fyne.Container)
			rightside := border.Objects[1].(*fyne.Container)

			avatar := leftside.Objects[0].(*widget.Icon)
			name := leftside.Objects[1].(*fyne.Container).Objects[0].(*widget.Label)
			lastSeen := rightside.Objects[1].(*widget.Label)

			avatar.SetResource(utils.GenerateAvatar(chat.Name))
			name.SetText(chat.Name)
			if chat.Online {
				lastSeen.SetText("Online")
			} else {
				lastSeen.SetText(chat.LastSeen.String())
			}

			chatItem.OnRightClick = func() {
				dialog.NewConfirm(
					"Delete Chat",
					fmt.Sprintf("Are you sure you want to delete %s?", chat.Name),
					func(confirm bool) {
						if confirm {
							*allChats = removeUser(*allChats, chat)
							filteredChats = make([]schema.UserList, len(*allChats))
							copy(filteredChats, *allChats)
							chatList.Refresh()
						}
					},
					win,
				).Show()
			}
		},
	)

	chatList.OnSelected = func(id widget.ListItemID) {
		onSelect(filteredChats[id])
	}

	search := widget.NewEntry()
	search.SetPlaceHolder("Search chats...")

	clearSearch := func() {
		search.SetText("")
		filteredChats = make([]schema.UserList, len(*allChats))
		copy(filteredChats, *allChats)
		chatList.Refresh()
	}

	search.OnChanged = func(text string) {
		lower := strings.ToLower(text)
		filtered := make([]schema.UserList, 0)
		for _, c := range *allChats {
			if strings.Contains(strings.ToLower(c.Name), lower) {
				filtered = append(filtered, c)
			}
		}
		filteredChats = filtered
		chatList.Refresh()
	}

	searchBar := container.NewBorder(nil, nil, nil,
		widget.NewButtonWithIcon("", theme.CancelIcon(), func() { clearSearch() }),
		search)

	return container.NewPadded(
		container.NewBorder(
			container.NewPadded(topBar),
			nil,
			nil,
			nil,
			container.NewBorder(container.NewPadded(searchBar), nil, nil, nil, chatList),
		),
	)
}

// Helper function to remove a user from the list
func removeUser(users []schema.UserList, target schema.UserList) []schema.UserList {
	for i, u := range users {
		if u.Name == target.Name {
			return append(users[:i], users[i+1:]...)
		}
	}
	return users
}
