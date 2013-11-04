package main

import (
	"code.google.com/p/gowut/gwu"
	"fmt"
	"github.com/biorhitm/rpn"
	"strconv"
	"strings"
)

var (
	screen gwu.TextBox
)

func newCalcButton(text string) gwu.Button {
	btn := gwu.NewButton(text)
	btn.Style().SetWidth("40px")
	return btn
}

func addSymbolToScreen(e gwu.Event) {
	btn := e.Src()
	screen.SetText(screen.Text() + btn.Attr("value"))
	e.MarkDirty(screen)
}

func addSymbolButton(p gwu.Panel, digit string) {
	btn := newCalcButton(digit)
	btn.SetAttr("value", digit)
	btn.AddEHandlerFunc(addSymbolToScreen, gwu.ETYPE_CLICK)
	p.Add(btn)
}

func calculateExpression() {
	s := strings.Split(screen.Text(), "=")
	f, n, err := rpn.Calculate(s[0])
	if err != nil {
		fmt.Printf("%s ошибка в: %d\n", err, n)
	} else {
		screen.SetText(s[0] + "=" + strconv.FormatFloat(f, 'f', -1, 64))
	}
}

func btnResultPressed(e gwu.Event) {
	calculateExpression()
	e.MarkDirty(screen)
}

func screenKeyUp(e gwu.Event) {
	if e.KeyCode() == gwu.KEY_ENTER {
		calculateExpression()
		e.MarkDirty(screen)
	}
}

func btnClearPressed(e gwu.Event) {
	screen.SetText("0")
	e.MarkDirty(screen)
}

func main() {
	win := gwu.NewWindow("main", "Калькулятор")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HA_CENTER)
	win.SetCellPadding(2)

	panel := gwu.NewHorizontalPanel()
	panel.Style().SetBorder2(1, gwu.BRD_STYLE_SOLID, gwu.CLR_BLACK)
	panel.Style().SetWidth("600")
	panel.SetCellPadding(2)
	win.Add(panel)

	screen = gwu.NewTextBox("0")
	screen.Style().SetFullWidth()
	panel.Add(screen)
	screen.AddEHandlerFunc(screenKeyUp, gwu.ETYPE_KEY_UP)

	p0 := gwu.NewHorizontalPanel()
	win.Add(p0)

	addSymbolButton(p0, "7")
	addSymbolButton(p0, "8")
	addSymbolButton(p0, "9")
	addSymbolButton(p0, "(")
	addSymbolButton(p0, ")")

	p1 := gwu.NewHorizontalPanel()
	win.Add(p1)

	addSymbolButton(p1, "4")
	addSymbolButton(p1, "5")
	addSymbolButton(p1, "6")

	addSymbolButton(p1, "*")
	addSymbolButton(p1, "/")

	p2 := gwu.NewHorizontalPanel()
	win.Add(p2)

	addSymbolButton(p2, "1")
	addSymbolButton(p2, "2")
	addSymbolButton(p2, "3")

	addSymbolButton(p2, "+")
	addSymbolButton(p2, "-")

	p3 := gwu.NewHorizontalPanel()
	win.Add(p3)

	addSymbolButton(p3, "0")
	addSymbolButton(p3, ".")

	btnResult := newCalcButton("=")
	btnResult.Style().SetWidth("80px")
	p3.Add(btnResult)
	btnResult.AddEHandlerFunc(btnResultPressed, gwu.ETYPE_CLICK)

	btnClear := newCalcButton("C")
	p3.Add(btnClear)
	btnClear.AddEHandlerFunc(btnClearPressed, gwu.ETYPE_CLICK)

	server := gwu.NewServer("guitest", "localhost:8081")
	server.SetText("Калькулятор")
	server.AddWin(win)
	server.Start("main")
}
