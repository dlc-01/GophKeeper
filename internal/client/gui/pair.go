package gui

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

func (v *View) callPairForm(signType Sign) {
	var pair models.Pair
	v.tui.pairForm.AddInputField("login", "", 20, nil, func(login string) {
		pair.Username = login
	})

	v.tui.pairForm.AddPasswordField("password", "", 20, '*', func(password string) {
		pair.PasswordHash = password
	})

	v.tui.pairForm.AddInputField("metadata", "", 20, nil, func(meta string) {
		pair.Metadata = meta
	})

	v.tui.pairForm.AddButton("OK", func() {
		var err error

		switch signType {
		case register:
			_, err = v.handlers.Pairs.CreatePair(context.Background(), v.handlers.Token, v.handlers.SecretKey, pair)
		case login:

		default:
			v.switchToMainMenu()
		}

		if err != nil {
			v.tui.body.SwitchToPage(registerFail)
			return
		}

		v.switchToUnitsMenu()
	})

	v.tui.signForm.AddButton("Cancel", func() {
		v.switchToMainMenu()
	})
}

func (v *View) createPairsPage() {
	v.tui.pairsList = tview.NewList().ShowSecondaryText(false)
	v.tui.pairsInfo = tview.NewTextView()

	v.tui.pairsPage = tview.NewFlex().
		AddItem(v.tui.pairsList, 0, 1, true).
		AddItem(v.tui.pairsInfo, 0, 3, false)

	v.tui.pairsPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.switchToUnitsMenu()
		}
		return event
	})

	v.tui.body.AddPage(pairsPage, v.tui.pairsPage, true, false)
}

func (v *View) getPairsList() {
	pairs, err := v.handlers.Pairs.GetPair(context.Background(), v.handlers.Token, v.handlers.SecretKey)
	if err != nil {
		v.switchToUnitsMenu()
		return
	}

	v.tui.pairsList.Clear()
	for _, pair := range pairs {
		v.tui.pairsList.AddItem(strconv.Itoa(int(pair.ID)), "", ' ', nil)
	}

	v.tui.pairsList.SetSelectedFunc(func(index int, name string, secondName string, shortcut rune) {
		v.setPairInfo(pairs[index])
	})
}

func (v *View) setPairInfo(pair models.Pair) {
	var sb strings.Builder

	v.tui.pairsInfo.Clear()
	sb.WriteString("Username: " + pair.Username)
	sb.WriteString("\n")
	sb.WriteString("Password: " + pair.PasswordHash)
	sb.WriteString("\n")
	sb.WriteString("Metadata:" + pair.Metadata)
	sb.WriteString("\n")

	v.tui.pairsInfo.SetText(sb.String())
}
