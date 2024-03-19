package gui

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

func (v *View) callCardForm(signType Sign) {
	var newCardData models.BankAccountString
	var reqN string

	v.tui.cardForm.AddInputField("CardHolder", "", 20, nil, func(holder string) {
		newCardData.CardHolder = holder
	})

	v.tui.cardForm.AddInputField("number", "", 20, nil, func(number string) {
		reqN = number
	})

	v.tui.cardForm.AddPasswordField("security code", "", 4, '*', func(sec string) {
		newCardData.SecurityCode = sec
	})

	v.tui.cardForm.AddInputField("explanation date in format 2022-02-24", "", 10, nil, func(data string) {
		newCardData.ExpirationDate = data
	})

	v.tui.cardForm.AddInputField("metadata", "", 20, nil, func(meta string) {
		newCardData.Metadata = meta
	})

	v.tui.cardForm.AddButton("OK", func() {
		var err error

		switch signType {
		case register:
			newCardData.Number, err = strconv.ParseUint(reqN, 10, 64)
			if err != nil {
				v.tui.body.SwitchToPage(registerFail)
				return
			}
			_, err = v.handlers.Cards.CreateBank(context.Background(), v.handlers.Token, newCardData)
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

func (v *View) createCardsPage() {
	v.tui.cardsList = tview.NewList().ShowSecondaryText(false)
	v.tui.cardInfo = tview.NewTextView()

	v.tui.cardsPage = tview.NewFlex().
		AddItem(v.tui.cardsList, 0, 1, true).
		AddItem(v.tui.cardInfo, 0, 3, false)

	v.tui.cardsPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.switchToUnitsMenu()
		}
		return event
	})

	v.tui.body.AddPage(cardsPage, v.tui.cardsPage, true, false)
}

func (v *View) getCardsList() {
	cards, err := v.handlers.Cards.GetBank(context.Background(), v.handlers.Token)
	if err != nil {
		v.switchToUnitsMenu()
		return
	}

	v.tui.cardsList.Clear()
	for _, card := range cards {
		v.tui.cardsList.AddItem(strconv.Itoa(int(card.ID)), "", ' ', nil)
	}

	v.tui.cardsList.SetSelectedFunc(func(index int, name string, secondName string, shortcut rune) {
		v.setCardInfo(cards[index])
	})
}

func (v *View) setCardInfo(card models.BankAccount) {
	var sb strings.Builder

	v.tui.cardInfo.Clear()
	sb.WriteString("CardHolder: " + card.CardHolder)
	sb.WriteString("\n")
	sb.WriteString("Number: " + strconv.FormatUint(card.Number, 10))
	sb.WriteString("\n")
	sb.WriteString("ExplanationDate: " + card.ExpirationDate.String())
	sb.WriteString("\n")
	sb.WriteString("Security code: " + card.SecurityCode)
	sb.WriteString("\n")
	sb.WriteString("Metadata: " + card.Metadata)
	sb.WriteString("\n")

	v.tui.cardInfo.SetText(sb.String())
}
