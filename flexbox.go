package termui

const (
	Horizontal Alignment = iota
	Vertical
)

type Alignment uint8

type Flexbox struct {
	WidgetBase
	item []*flexboxItem
	alignment Alignment

	border bool
	title  string

	borderStyle *Style
	titleStyle *Style

	startItem uint

	scroll bool
}

type flexboxItem struct {
	fixed uint
	scale uint
	widget Widget
}

func NewFlexbox() *Flexbox {
	return &Flexbox{
		item: make([]*flexboxItem, 0, 2),
		alignment: Vertical,
		borderStyle: NewStyle(),
		titleStyle: NewStyle(),
	}
}

func (this * Flexbox)AddItem(w Widget, fixedSize, scaleSize uint) *Flexbox {
	item := &flexboxItem{
		widget: w,
		scale: scaleSize,
		fixed: fixedSize,
	}
	this.item = append(this.item, item)
	return this
}

func (this *Flexbox) GetWidgets() []Widget {
	val := make([]Widget, 0, len(this.item))
	for _, it := range this.item {
		val = append(val, it.widget)
	}
	return val
}

func (this *Flexbox) SetDirection(alignment Alignment) *Flexbox {
	this.alignment = alignment
	return this
}

func (this *Flexbox) SetBorder(border bool) *Flexbox {
	this.border = border
	return this
}

func (this *Flexbox) IsBorder() bool {
	return this.border
}

func (this *Flexbox) SetTitle(title string) *Flexbox {
	this.title = title
	return this
}

func (this *Flexbox) SetScroll(show bool) *Flexbox {
	this.scroll = show
	this.border = show
	return this
}

func (this *Flexbox)Resize(r *Rect) {
	this.WidgetBase.Resize(r)

	bodyRect := r.Copy()
	if this.title != "" || this.border {
		bodyRect.Y++
		bodyRect.H--
	}
	if this.border {
		bodyRect.H--
		bodyRect.X++
		bodyRect.W-=2
	}


	arr_rect :=  make([]*Rect, len(this.item))
	free_rect :=  make([]*Rect, 0, len(this.item))

	var free_space uint
	if this.alignment == Vertical {
		free_space = bodyRect.H
	}else {
		free_space = bodyRect.W
	}

	for i, item := range this.item {
		if item.fixed > 0 {
			free_space -= item.fixed
			if this.alignment == Vertical {
				arr_rect[i] = NewRect(bodyRect.W, item.fixed)
			}else {
				arr_rect[i] = NewRect(item.fixed, bodyRect.H)
			}
		}else {
			var r *Rect
			widgetSize := item.widget.GetMinSize(bodyRect.Size)
			if this.alignment == Vertical {
				r = NewRect(bodyRect.W, widgetSize.H)
				if item.scale == 0 {
					free_space -= widgetSize.H
				}
			}else {
				r = NewRect(widgetSize.W, bodyRect.H)
				if item.scale == 0 {
					free_space -= widgetSize.W
				}
			}
			arr_rect[i] = r
			if item.scale != 0 {
				free_rect = append(free_rect, r)
			}
		}
	}

	if len(free_rect) > 0 {
		num := len(free_rect)
		for num > 0 {
			if num == 1 {
				if this.alignment == Vertical {
					free_rect[0].H = free_space
				}else{
					free_rect[0].W = free_space
				}
			}else {
				delta := free_space / uint(num)
				if this.alignment == Vertical {
					if delta > free_rect[0].H {
						free_rect[0].H = delta
					}
				}else{
					if delta > free_rect[0].W {
						free_rect[0].W = delta
					}
				}
			}
			//end
			if this.alignment == Vertical {
				free_space -= free_rect[0].H
			}else{
				free_space -= free_rect[0].W
			}
			free_rect = free_rect[1:]
			num = len(free_rect)
		}
	}

	for i, item := range this.item {
		itemRect := arr_rect[i]
		itemRect.X = bodyRect.X
		itemRect.Y = bodyRect.Y

		item.widget.Resize(itemRect)

		if this.alignment == Vertical {
			bodyRect.Y += itemRect.H
			bodyRect.H -= itemRect.H
		} else {
			bodyRect.X += itemRect.W
			bodyRect.W -= itemRect.W
		}

	}
}

func (this *Flexbox) Draw(p *Painter) {
	if this.WidgetBase.r.W == 0 || this.WidgetBase.r.H == 0 {
		return
	}
	bodyRect := this.r.Copy()
	if this.border {
		p.DrawRect(this.r.X, this.r.Y, this.r.W, this.r.H, this.borderStyle)
		bodyRect.W -= 2
		bodyRect.H -= 2
	}
	if this.title != "" {
		p.DrawText(this.r.X+1, this.r.Y, this.title, this.titleStyle)
		if !this.border {
			bodyRect.H--
		}
	}
	if this.scroll {
		if this.alignment == Vertical {
			if this.r.H > 2 {
				p.DrawVerticalLine(this.r.X+this.r.W-1, this.r.Y+1, this.r.H-2, BLOCK, this.borderStyle)
				p.DrawRune(this.r.X+this.r.W-1, this.r.Y+2, SCROLL_BLOCK, this.borderStyle)

				p.DrawRune(this.r.X+this.r.W-1, this.r.Y+1, QUOTA_UP, this.borderStyle)
				p.DrawRune(this.r.X+this.r.W-1, this.r.Y+this.r.H-2, QUOTA_DOWN, this.borderStyle)
			}
		}else{
			p.DrawHorizontalLine(this.r.X+1, this.r.Y+this.r.H-1, this.r.W-2, BLOCK, this.borderStyle)
			p.DrawRune(this.r.X+2, this.r.Y+this.r.H-1, SCROLL_BLOCK, this.borderStyle)

			p.DrawRune(this.r.X+1, this.r.Y+this.r.H-1, QUOTA_LEFT, this.borderStyle)
			p.DrawRune(this.r.X+this.r.W-2, this.r.Y+this.r.H-1, QUOTA_RIGHT, this.borderStyle)
		}
	}

	var free_space int
	if this.alignment == Vertical {
		free_space = int(bodyRect.H)
	}else {
		free_space = int(bodyRect.W)
	}
	for i:= int(this.startItem); i < len(this.item); i++ {
		if free_space <= 0 {
			break
		}

		this.item[i].widget.Draw(p)

		size := this.item[i].widget.GetSize()
		if this.alignment == Vertical {
			free_space -= int(size.H)
		}else {
			free_space -= int(size.W)
		}
	}
}

func (this *Flexbox) GetMinSize(max Size) (min Size)  {
	for _, it := range this.item{
		tmp := it.widget.GetMinSize(max)
		if this.alignment == Vertical {
			if tmp.W > min.W {
				min.W = tmp.W
			}
			min.H += tmp.H
		}else{
			min.W += tmp.W
			if tmp.H > min.H {
				min.H = tmp.H
			}
		}
	}
	if this.scroll {
		min.H += 2
		min.W += 2
		if this.alignment == Vertical {
			if min.H < 5 {
				min.H = 5
			}
		}else{
			if min.W < 5 {
				min.W = 5
			}
		}
	}else if this.border {
		min.H += 2
		min.W += 2
	}else if this.title != "" {
		min.H++

	}
	min = correctSize(max, min)
	return
}
