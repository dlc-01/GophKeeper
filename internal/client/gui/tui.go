package gui

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/client/config"
	"github.com/dlc-01/GophKeeper/internal/client/handlers"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	mainMenu     = "main"
	unitsMenu    = "units"
	pairsPage    = "pairs"
	cardsPage    = "cards"
	notesPage    = "notes"
	pairForm     = "pairForm"
	textForm     = "textForm"
	cardForm     = "cardForm"
	signForm     = "sign"
	registerFail = "login exist"
)

type Sign int

const (
	undefined Sign = iota
	register
	login
)

type ui struct {
	*tview.Application
	root   *tview.Flex
	header *tview.TextView
	body   *tview.Pages
	footer *tview.Flex

	mainMenu *tview.List

	unitsMenu *tview.List

	pairsPage *tview.Flex
	pairsList *tview.List
	pairsInfo *tview.TextView

	cardsPage *tview.Flex
	cardsList *tview.List
	cardInfo  *tview.TextView

	notesPage *tview.Flex
	notesList *tview.List
	noteInfo  *tview.TextView

	pairForm *tview.Form
	textForm *tview.Form
	cardForm *tview.Form

	signForm     *tview.Form
	registerFail *tview.Modal
}

type View struct {
	handlers *handlers.Handlers
	cfg      *config.Config
	tui      *ui
}

func New(ctrl *handlers.Handlers, cfg *config.Config) (v *View) {
	v = &View{
		handlers: ctrl,
		cfg:      cfg,
		tui: &ui{
			Application: tview.NewApplication(),
			body:        tview.NewPages(),
			signForm:    tview.NewForm(),
			pairForm:    tview.NewForm(),
			textForm:    tview.NewForm(),
			cardForm:    tview.NewForm(),
		},
	}

	v.createHeader()

	v.createMainMenu()
	v.createRegisterFail()
	v.createUnitsMenu()
	v.createPairsPage()
	v.createCardsPage()
	v.createNotesPage()

	v.createFooter()
	v.createRoot()

	v.tui.EnableMouse(true)
	return
}

func (v *View) Run() {
	if err := v.tui.SetRoot(v.tui.root, true).Run(); err != nil {
		panic(err)
	}
}

func (v *View) callSignForm(signType Sign) {
	var regLogin, regPassword string
	v.tui.signForm.AddInputField("login", "", 20, nil, func(login string) {
		regLogin = login
	})

	v.tui.signForm.AddPasswordField("password", "", 20, '*', func(password string) {
		regPassword = password
	})

	v.tui.signForm.AddButton("OK", func() {
		var token string
		var err error

		switch signType {
		case register:
			token, err = v.handlers.Auth.Register(context.Background(), regLogin, regPassword)
		case login:
			token, err = v.handlers.Auth.Login(context.Background(), regLogin, regPassword)
		default:
			v.switchToMainMenu()
		}

		if err != nil {
			v.tui.body.SwitchToPage(registerFail)
			return
		}

		v.handlers.Token = token
		v.switchToUnitsMenu()
	})

	v.tui.signForm.AddButton("Cancel", func() {
		v.switchToMainMenu()
	})
}

func (v *View) createMainMenu() {
	v.tui.mainMenu = tview.NewList().
		AddItem("Register", "Sign up new user", 'r', func() {
			v.tui.signForm.Clear(true)
			v.callSignForm(register)
			v.setHeader("Register")
			v.tui.body.SwitchToPage(signForm)
		}).
		AddItem("Login", "Sign in with exist user", 'l', func() {
			v.tui.signForm.Clear(true)
			v.callSignForm(login)
			v.setHeader("Login")
			v.tui.body.SwitchToPage(signForm)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			v.tui.Stop()
		})

	v.tui.body.AddPage(mainMenu, v.tui.mainMenu, true, true)
	v.tui.body.AddPage(signForm, v.tui.signForm, true, false)
}

func (v *View) createRegisterFail() {
	v.tui.registerFail = tview.NewModal().
		SetText("Login already exist!").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			v.switchToMainMenu()
		}).
		SetBackgroundColor(tcell.ColorLightCoral)

	v.tui.body.AddPage(registerFail, v.tui.registerFail, true, false)
}

func (v *View) createUnitsMenu() {
	v.tui.unitsMenu = tview.NewList().
		AddItem("Pairs", "show login/password pairs", 'r', func() {
			v.getPairsList()
			v.setHeader("Pairs (press ESC to exit)")
			v.tui.body.SwitchToPage(pairsPage)
		}).
		AddItem("CreatePair", "create pair", 'p', func() {
			v.tui.pairForm.Clear(true)
			v.callPairForm(1)
			v.setHeader("Pair (press ESC to exit)")
			v.tui.body.SwitchToPage(pairForm)
		}).
		AddItem("Notes", "show arbitrary text data", 'l', func() {
			v.getNotesList()
			v.setHeader("Notes (press ESC to exit)")
			v.tui.body.SwitchToPage(notesPage)
		}).
		AddItem("CreateNote", "create text", 'n', func() {
			v.tui.textForm.Clear(true)
			v.callTextForm(1)
			v.setHeader("Notes (press ESC to exit)")
			v.tui.body.SwitchToPage(textForm)
		}).
		AddItem("Cards", "show bank cards data", 'b', func() {
			v.getCardsList()
			v.setHeader("Cards (press ESC to exit)")
			v.tui.body.SwitchToPage(cardsPage)
		}).
		AddItem("CreateCard", "create card", 'c', func() {
			v.tui.cardForm.Clear(true)
			v.callCardForm(1)
			v.setHeader("Card (press ESC to exit)")
			v.tui.body.SwitchToPage(cardForm)
		}).
		AddItem("Binary", "show arbitrary binary data", 'b', nil).
		AddItem("Back", "... to main menu", ' ', func() {
			v.switchToMainMenu()
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			v.tui.Stop()
		})

	v.tui.body.AddPage(unitsMenu, v.tui.unitsMenu, true, false)
	v.tui.body.AddPage(pairForm, v.tui.pairForm, true, false)
	v.tui.body.AddPage(textForm, v.tui.textForm, true, false)
	v.tui.body.AddPage(cardForm, v.tui.cardForm, true, false)
}

func (v *View) createHeader() {
	v.tui.header = tview.NewTextView()
	v.tui.header.SetBorder(true)
	v.tui.header.SetText("Main menu")
}

func (v *View) setHeader(text string) {
	v.tui.header.SetText(text)
}

func (v *View) createFooter() {
	clientInfo := tview.NewTextView().
		SetText("version: " + v.cfg.App.Version)

	clientInfo.SetBorder(true).
		SetTitle("Client info")

	serverInfo := tview.NewTextView().
		SetText("version: \ntarget: ")
	serverInfo.SetBorder(true).
		SetTitle("Server info")

	v.tui.footer = tview.NewFlex().
		AddItem(clientInfo, 0, 1, false).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(serverInfo, 0, 1, false)
}

func (v *View) createRoot() {
	v.tui.root = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(v.tui.header, 0, 1, false).
		AddItem(v.tui.body, 0, 3, true).
		AddItem(v.tui.footer, 0, 1, false)
}

func (v *View) switchToMainMenu() {
	v.setHeader("Main menu")
	v.tui.body.SwitchToPage(mainMenu)
}

func (v *View) switchToUnitsMenu() {
	v.setHeader("Resources")
	v.tui.body.SwitchToPage(unitsMenu)
}
