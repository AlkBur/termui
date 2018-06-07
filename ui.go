package termui

import (
	"github.com/gdamore/tcell"
	"image"
	"reflect"
)

type UI struct {
	root    Widget
	focus Widget

	// The application's screen.
	painter *Painter

	keybindings []*keybinding

	quit chan struct{}
	eventQueue chan event
}

func New(root Widget) (*UI, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	//p := NewPainter(screen)

	return &UI{
		//painter:     p,
		root:        root,
		keybindings: make([]*keybinding, 0, 5),
		quit:        make(chan struct{}, 1),
		painter:      NewPainter(screen),
		eventQueue:  make(chan event),
	}, nil

}

func (ui *UI) Quit() {
	ui.painter.screen.Fini()
	ui.quit <- struct{}{}
}

func (ui *UI) SetKeybinding(seq string, fn func()) {
	ui.keybindings = append(ui.keybindings, &keybinding{
		sequence: seq,
		handler:  fn,
	})
}

func (ui *UI) Run() error {
	if err := ui.painter.screen.Init(); err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			ui.painter.screen.Fini()
		}
	}()

	ui.painter.screen.SetStyle(tcell.StyleDefault)
	//ui.Draw()

	go func() {
		for {
			switch ev := ui.painter.screen.PollEvent().(type) {
			case *tcell.EventKey:
				ui.handleKeyEvent(ev)
			case *tcell.EventMouse:
				ui.handleMouseEvent(ev)
			case *tcell.EventResize:
				ui.handleResizeEvent(ev)
			}
		}
	}()

	for {
		select {
		case <-ui.quit:
			return nil
		case ev := <-ui.eventQueue:
			ui.handleEvent(ev)
		}
	}
}

func (ui *UI) handleEvent(ev event) {
	switch e := ev.(type) {
	case KeyEvent:
		for _, b := range ui.keybindings {
			if b.match(e) {
				b.handler()
			}
		}
		//ui.focus.OnKeyEvent(e)
		ui.root.OnKeyEvent(e)
		//ui.painter.Repaint(ui.root, !ui.bHideCursor)
	case MouseEvent:
		//ui.focus.OnMouseEvent(e, ui.root)
		ui.root.OnMouseEvent(e)
		//ui.painter.Repaint(ui.root, !ui.bHideCursor)
	case callbackEvent:
		e.cbFn()
		//ui.painter.Repaint(ui.root, !ui.bHideCursor)
	case paintEvent:
		ui.Draw()
	}
}

func (ui *UI) handleKeyEvent(tev *tcell.EventKey) {
	ui.eventQueue <- KeyEvent{
		Key:       Key(tev.Key()),
		Rune:      tev.Rune(),
		Modifiers: ModMask(tev.Modifiers()),
	}
}

func (ui *UI) handleMouseEvent(ev *tcell.EventMouse) {
	x, y := ev.Position()
	ui.eventQueue <- MouseEvent{Pos: image.Pt(x, y)}
}

func (ui *UI) handleResizeEvent(ev *tcell.EventResize) {
	ui.eventQueue <- paintEvent{}
}

func (ui *UI)Draw() {
	var root *Flexbox
	if reflect.TypeOf(ui.root) == reflect.TypeOf(&Flexbox{}) {
		root = ui.root.(*Flexbox)
	}else {
		root = NewFlexbox().AddItem(ui.root, 0, 1)
	}

	r := &Rect{Point: Point{0, 0}, Size: ui.painter.Size()}
	root.Resize(r)


	ui.painter.Clear()
	root.Draw(ui.painter)
	ui.painter.Show()
}