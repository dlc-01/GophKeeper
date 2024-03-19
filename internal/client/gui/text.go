package gui

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

func (v *View) callTextForm(signType Sign) {
	var reqNote, metadata string
	v.tui.textForm.AddInputField("note", "", 20, nil, func(note string) {
		reqNote = note
	})

	v.tui.textForm.AddInputField("metadata", "", 20, nil, func(meta string) {
		metadata = meta
	})

	v.tui.textForm.AddButton("OK", func() {
		var err error

		switch signType {
		case register:
			_, err = v.handlers.Notes.CreateText(context.Background(), v.handlers.Token, models.Text{Note: reqNote, Metadata: metadata})
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

func (v *View) createNotesPage() {
	v.tui.notesList = tview.NewList().ShowSecondaryText(false)
	v.tui.noteInfo = tview.NewTextView()

	v.tui.notesPage = tview.NewFlex().
		AddItem(v.tui.notesList, 0, 1, true).
		AddItem(v.tui.noteInfo, 0, 3, false)

	v.tui.notesPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.switchToUnitsMenu()
		}
		return event
	})

	v.tui.body.AddPage(notesPage, v.tui.notesPage, true, false)
}

func (v *View) getNotesList() {
	notes, err := v.handlers.Notes.GetText(context.Background(), v.handlers.Token)
	if err != nil {
		v.switchToUnitsMenu()
		return
	}

	v.tui.notesList.Clear()
	for _, note := range notes {
		v.tui.notesList.AddItem(strconv.Itoa(int(note.ID)), "", ' ', nil)
	}

	v.tui.notesList.SetSelectedFunc(func(index int, name string, secondName string, shortcut rune) {
		v.setNoteInfo(notes[index])
	})
}

func (v *View) setNoteInfo(note models.Text) {
	var sb strings.Builder

	v.tui.cardInfo.Clear()
	sb.WriteString(note.Note)
	sb.WriteString("\n")
	sb.WriteString(note.Metadata)
	sb.WriteString("\n")

	v.tui.noteInfo.SetText(sb.String())
}
