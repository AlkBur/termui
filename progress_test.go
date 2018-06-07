package termui

import "testing"

func TestProgress_Draw(t *testing.T) {
	b := NewFlexbox()
	p := NewProgress(100).SetCurrent(50)
	b.AddItem(p, 0, 1)

	surface := NewTestSurface(12, 3)
	painter := NewPainter(surface)

	b.Resize(&Rect{Point: Point{0, 0}, Size: painter.Size()})
	b.Draw(painter)

	want := `
┌──────────┐
│▓▓▓▓▓.....│
└──────────┘
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}