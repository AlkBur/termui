package termui

type Widget interface {
	Draw(p *Painter)
	OnKeyEvent(ev KeyEvent)
	OnMouseEvent(ev MouseEvent)
	Resize(r *Rect)
	GetSize() Size
	GetMinSize(s Size) Size
}

type WidgetBase struct {
	r Rect
}

func (this *WidgetBase)Draw(p *Painter) {}

func (this *WidgetBase)Resize(r *Rect) {
	this.r = *r.Copy()
}

func (this *WidgetBase)OnKeyEvent(ev KeyEvent) {}

func (this *WidgetBase)OnMouseEvent(ev MouseEvent) {}

func (this *WidgetBase)GetSize() Size {
	return this.r.Size
}

func (this *WidgetBase)GetMinSize(max Size) Size {
	return Size{0, 0}
}