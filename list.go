package termui

type List struct {
	WidgetBase
	selected int

	items    []string
	flex *Flexbox

	onItemActivated    func(*List)
	onSelectionChanged func(*List)
}

func NewList() *List {
	return &List{
		items: make([]string, 0, 2),
		selected: -1,
		flex: NewFlexbox().SetDirection(Vertical).SetScroll(true),
	}
}

func (this *List) AddItems(items ...string) *List {
	this.items = append(this.items, items...)
	for _, item :=range items {
		this.flex.AddItem(NewLabel(item), 0, 0)
	}
	return this
}

func (this *List) SetBorder(border bool) *List {
	this.flex.SetBorder(border)
	return this
}

func (this *List) SetScroll(show bool) *List {
	this.flex.SetScroll(show)
	return this
}

func (this *List) SetTitle(title string) *List {
	this.flex.SetTitle(title)
	return this
}

func (l *List) SetSelected(i int) {
	l.selected = i
}

func (l *List) Selected() int {
	return l.selected
}

func (this *List)Resize(r *Rect) {
	this.flex.Resize(r)
}

func (this *List) Draw(p *Painter) {
	this.flex.Draw(p)
}

func (this *List) GetMinSize(s Size) (min Size)  {
	min = this.flex.GetMinSize(s)
	return
}

