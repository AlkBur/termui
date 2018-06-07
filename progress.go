package termui

import "strings"

type Progress struct {
	WidgetBase
	flex *Flexbox

	current, max int
}

func NewProgress(max int) *Progress {
	return &Progress{
		max: max,
		flex: NewFlexbox().SetBorder(true).
			AddItem(NewLabel(string(SCROLL_BLOCK)).SetRepeat(-1), 0, 1),
	}
}

func (this *Progress) Draw(p *Painter) {
	this.flex.Draw(p)
}

func (this *Progress) GetMinSize(s Size) (min Size)  {
	min = this.flex.GetMinSize(s)
	return
}

func (this *Progress) Resize(r *Rect) {
	arr := this.flex.GetWidgets()
	if len(arr)==1 {
		label := arr[0].(*Label)
		repeat := int((float64(this.current) / float64(this.max)) * float64(r.W-2))
		txt := ""
		if repeat > 0 {
			txt = strings.Repeat(string(SCROLL_BLOCK), repeat)
			label.repeat = 0
		}
		label.SetText(txt)
	}
	this.flex.Resize(r)
}

func (this *Progress) SetCurrent(val int) *Progress {
	this.current = val
	return this
}