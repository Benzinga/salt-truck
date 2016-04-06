package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

// ListMenu is a simple TUI list select dialog.
type ListMenu struct {
	Title    string
	Options  []string
	Selected string
}

func (l *ListMenu) cursorDown(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()

	cy++
	if cy >= len(l.Options) {
		cy = len(l.Options) - 1
	}

	err := v.SetCursor(cx, cy)

	if err != nil {
		return err
	}
	return nil
}

func (l *ListMenu) cursorUp(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()

	cy--
	if cy < 0 {
		cy = 0
	}

	err := v.SetCursor(cx, cy)

	if err != nil {
		return err
	}
	return nil
}

func (l *ListMenu) choose(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err == nil {
		l.Selected = line
		return gocui.ErrQuit
	}

	return nil
}

func (l *ListMenu) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (l *ListMenu) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("header", 0, 0, maxX-1, 2); err != nil {
		fmt.Fprintln(v, l.Title)
	}
	if v, err := g.SetView("list", 0, 2, maxX-1, maxY-1); err != nil {
		v.Highlight = true
		for _, opt := range l.Options {
			fmt.Fprintln(v, opt)
		}
		g.SetCurrentView("list")
	}
	return nil
}

func (l *ListMenu) Show() string {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(l.layout)
	g.SetKeybinding("list", gocui.KeyArrowDown, gocui.ModNone, l.cursorDown)
	g.SetKeybinding("list", gocui.KeyArrowUp, gocui.ModNone, l.cursorUp)
	g.SetKeybinding("list", gocui.KeyEnter, gocui.ModNone, l.choose)
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, l.quit)
	g.SelBgColor = gocui.ColorWhite
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true
	g.Mouse = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	return l.Selected
}
