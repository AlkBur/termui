package termui

type Button struct {
	WidgetBase
	flex *Flexbox
	selected func()
}

func NewButton(label string) *Button {
	return &Button{
		flex: NewFlexbox().
			AddItem(NewLabel(label).SetAlignH(AlignHorisontalCenter).SetAlignV(AlignVerticalCenter),
			0, 1),
	}
}

func (this *Button)SetSelectedFunc(f func()) *Button {
	this.selected = f
	return this
}

func (this *Button)SetBorder(border bool) *Button {
	this.flex.SetBorder(border)
	return this
}

func (this *Button) GetMinSize(max Size) (min Size)  {
	min = this.flex.GetMinSize(max)
	if this.flex.IsBorder() {
		min.W += 2
	}else{
		min.H += 2
		min.W += 4
	}
	min = correctSize(max, min)
	return
}

func (this *Button) Draw(p *Painter) {
	this.flex.Draw(p)
}

func (this *Button)Resize(r *Rect) {
	this.flex.Resize(r)
}