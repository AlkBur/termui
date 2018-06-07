package termui

import (
	"github.com/gdamore/tcell"
)

const TOP_RIGHT rune = '┐'
const VERTICAL_LINE rune = '│'
const HORIZONTAL_LINE rune = '─'
const TOP_LEFT rune = '┌'
const BOTTOM_RIGHT rune = '┘'
const BOTTOM_LEFT rune = '└'
const VERTICAL_LEFT rune = '┤'
const VERTICAL_RIGHT rune = '├'
const HORIZONTAL_DOWN rune = '┬'
const HORIZONTAL_UP rune = '┴'
const QUOTA_LEFT rune = '◄'
const QUOTA_RIGHT rune = '►'
const QUOTA_UP rune = '▲'   //9650
const QUOTA_DOWN rune = '▼' //9660
const BLOCK rune = '░' //176
const SCROLL_BLOCK rune = '▓' //178


const (
	//QUOTA_DOWN rune = '\u007c'
	//HORIZONTAL_LINE = '\u2500'
	//VERTICAL_LINE  = '\u2502'
	//TOP_LEFT       = '\u250c'
	//TOP_RIGHT      = '\u2510'
	//BOTTOM_LEFT    = '\u2514'
	//BOTTOM_RIGHT   = '\u2518'
	//GraphicsLeftT               = '\u251c'
	GraphicsRightT              = '\u2524'
	GraphicsTopT                = '\u252c'
	GraphicsBottomT             = '\u2534'
	GraphicsCross               = '\u253c'
	GraphicsDbVertBar           = '\u2550'
	GraphicsDbHorBar            = '\u2551'
	GraphicsDbTopLeftCorner     = '\u2554'
	GraphicsDbTopRightCorner    = '\u2557'
	GraphicsDbBottomRightCorner = '\u255d'
	GraphicsDbBottomLeftCorner  = '\u255a'
	GraphicsEllipsis            = '\u2026'

	//BLOCK = '\u176d'
)
//░▒▓
//Color
const (
	ColorDefault Color = iota
	ColorBlack
	ColorWhite
	ColorRed
	ColorGreen
	ColorBlue
	ColorCyan
	ColorMagenta
	ColorYellow
)

const (
	DecorationInherit Decoration = iota
	DecorationOn
	DecorationOff
)

// Color represents a color.
type Color uint8
type Decoration uint8

type Style struct {
	Fg Color
	Bg Color

	Reverse Decoration
	Bold Decoration
	Underline Decoration
}

type Painter struct {
	screen tcell.Screen
	style Style
}

type Rect struct {
	Size
	Point
}

type Size struct {
	W, H uint
}

type Point struct {
	X, Y uint
}

func NewRect(w,h uint) *Rect {
	return &Rect{
		Point: Point{0, 0},
		Size: Size{w, h},
	}
}

func NewPainter(screen tcell.Screen) *Painter {
	return &Painter{
		screen: screen,
		style: *NewStyle(),
	}
}

func NewStyle() *Style {
	return &Style{
		Fg: ColorWhite,
		Bg: ColorDefault,
	}
}

func (p *Painter) Size() Size {
	w,h := p.screen.Size()
	return Size{uint(w), uint(h)}
}

func (p *Painter)Clear() {
	p.screen.Clear()
}

func (p *Painter)Show() {
	p.screen.Show()
}

func (p *Painter) DrawVerticalLine(x, y, h uint, r rune, style *Style) {
	for y1 := uint(0); y1 < h; y1++ {
		p.DrawRune(x, y+y1, r, style)
	}
}

func (p *Painter) DrawHorizontalLine(x, y, h uint, r rune, style *Style) {
	for x1 := uint(0); x1 < h; x1++ {
		p.DrawRune(x+x1, y, r, style)
	}
}

func (p *Painter) DrawRect(x, y, w, h uint, style *Style) {
	for j := uint(0); j < h; j++ {
		for i := uint(0); i < w; i++ {
			m := i + x
			n := j + y

			switch {
			case i == 0 && j == 0:
				p.DrawRune(m, n, TOP_LEFT, style)
			case i == w-1 && j == 0:
				p.DrawRune(m, n, TOP_RIGHT, style)
			case i == 0 && j == h-1:
				p.DrawRune(m, n, BOTTOM_LEFT, style)
			case i == w-1 && j == h-1:
				p.DrawRune(m, n, BOTTOM_RIGHT, style)
			case i == 0 || i == w-1:
				p.DrawRune(m, n, VERTICAL_LINE, style)
			case j == 0 || j == h-1:
				p.DrawRune(m, n, HORIZONTAL_LINE, style)
			}
		}
	}
}

func (p *Painter) DrawRune(x, y uint, r rune, style *Style) {
	if style == nil {
		style = &p.style
	}

	st := tcell.StyleDefault.Normal().
		Foreground(convertColor(style.Fg, false)).
			Background(convertColor(style.Bg, false)).
			Reverse(style.Reverse == DecorationOn).
			Bold(style.Bold == DecorationOn).
			Underline(style.Underline == DecorationOn)

	p.screen.SetCell(int(x), int(y), st, r)
}

func (p *Painter) DrawText(x, y uint, text string, style *Style) {
	for _, r := range text {
		p.DrawRune(x, y, r, style)
		x += runeWidth(r)
	}
}

func convertColor(col Color, fg bool) tcell.Color {
	switch col {
	case ColorDefault:
		if fg {
			return tcell.ColorWhite
		}
		return tcell.ColorDefault
	case ColorBlack:
		return tcell.ColorBlack
	case ColorWhite:
		return tcell.ColorWhite
	case ColorRed:
		return tcell.ColorRed
	case ColorGreen:
		return tcell.ColorGreen
	case ColorBlue:
		return tcell.ColorBlue
	case ColorCyan:
		return tcell.ColorDarkCyan
	case ColorMagenta:
		return tcell.ColorDarkMagenta
	case ColorYellow:
		return tcell.ColorYellow
	default:
		if col > 0 {
			return tcell.Color(col)
		}
		return tcell.ColorDefault
	}
}

func (r *Rect) Copy() *Rect {
	return &Rect{
		Point: Point{r.X, r.Y}, Size: Size{r.W, r.H},
	}
}