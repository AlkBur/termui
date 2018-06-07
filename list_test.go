package termui

import "testing"

func TestList_Draw(t *testing.T) {
	surface := NewTestSurface(10, 8)
	painter := NewPainter(surface)

	b := NewFlexbox().SetDirection(Vertical)
	l := NewList().SetBorder(true)
	l.AddItems("foo", "bar", "foo", "bar", "foo", "bar", "foo")

	b.AddItem(l, 0, 1)

	b.Resize(&Rect{Point: Point{0, 0}, Size: painter.Size()})
	b.Draw(painter)

	want := `
┌────────┐
│foo.....▲
│bar.....▓
│foo.....░
│bar.....░
│foo.....░
│bar.....▼
└────────┘
`
	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}
