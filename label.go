package termui

import (
	"strings"
	"github.com/mitchellh/go-wordwrap"
)


const (
	AlignLeft AlignH = iota
	AlignHorisontalCenter
	AlignRight
)

const (
	AlignTop AlignV = iota
	AlignVerticalCenter
	AlignBottom
)

type (
	AlignV uint8
	AlignH uint8
)

type Label struct {
	WidgetBase
	wordWrap bool

	//Вертикальное и горизонтальное выравниевание
	alignH AlignH
	alignV AlignV

	repeat int

	text     string
	style    *Style

	line []string
}

func NewLabel(text string) *Label {
	this := &Label{
		text: text,
		alignH: AlignLeft,
		alignV: AlignTop,
		line: make([]string, 0, 1),
	}
	return this
}

func (this *Label) SetText(txt string) *Label {
	this.text = txt
	return this
}

func (this *Label) SetWordWrap(enabled bool) *Label {
	this.wordWrap = enabled
	return this
}

func (this *Label) SetRepeat(num int) *Label {
	this.repeat = num
	return this
}

func (this *Label) Resize(r *Rect) {
	this.WidgetBase.Resize(r)
	if this.r.H == 0 ||this.r.W == 0 {
		return
	}
	text := this.text
	if this.repeat < 0 {
		text = ""
	}else if this.repeat > 1 {
		text = strings.Repeat(text, this.repeat)
	}

	width := stringWidth(text)
	if width > this.r.W && this.wordWrap && this.r.W > 0 {
		txt := wordwrap.WrapString(text, uint(this.r.W))
		this.line = strings.Split(txt, "\n")
	}else{
		this.line = this.line[:0]
		this.line = append(this.line, text)
	}
}


func (this *Label) Draw(p *Painter) {
	if len(this.line) > 0 {
		for i, line := range this.line {
			runes := []rune(line)
			left := true
			if this.alignH == AlignRight {
				left = false
			}
			for len(runes) > int(this.r.W) {
				if left {
					runes = runes[:len(runes)-1]
				}else {
					runes = runes[1:]
				}
				if this.alignH == AlignHorisontalCenter {
					if left {
						left = false
					}else {
						left = true
					}
				}

			}
			txt := string(runes)
			width := stringWidth(txt)

			x := this.r.X
			y := this.r.Y+uint(i)
			if this.alignH == AlignRight {
				x += this.r.W - width
			}else if this.alignH == AlignHorisontalCenter {
				x += (this.r.W - width)/2
			}
			if this.alignV == AlignBottom {
				y += this.r.H - 1
			}
			if this.alignV == AlignVerticalCenter {
				y += this.r.H/2
			}

			p.DrawText(x, y, txt, this.style)
		}
	}
}

func (this *Label) SetAlignH(align AlignH) *Label {
	this.alignH = align
	return this
}

func (this *Label) SetAlignV(align AlignV) *Label {
	this.alignV = align
	return this
}

func (this *Label) GetMinSize(s Size) (min Size)  {
	width := stringWidth(this.text)
	if width > s.W && this.wordWrap {
		txt := strings.Split(wordwrap.WrapString(this.text, uint(this.r.W)), "\n")
		min.H = uint(len(txt))
		min.W = s.W
	}else{
		min.H = 1
		min.W = width
	}
	if min.H > s.H {
		min.H = s.H
	}
	if min.W > s.W {
		min.W = s.W
	}
	return
}
