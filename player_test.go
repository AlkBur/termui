package termui

import "testing"

func TestPlayer_Draw(t *testing.T) {
	surface := NewTestSurface(80, 14)
	painter := NewPainter(surface)

	b := NewFlexbox()

	l := NewList().SetScroll(true).SetTitle("List of songs")
	l.AddItems("Song-1", "Song-2", "Song-3", "Song-4", "Song-5", "Song-6", "Song-7")

	gr2:= NewFlexbox().SetTitle("Now play").SetBorder(true).
		AddItem(NewLabel("Song-1"),0,1)

	gr3 := NewFlexbox().SetDirection(Horizontal).
		AddItem(NewButton("Play").SetBorder(true), 0, 0).
		AddItem(NewButton("Stop").SetBorder(true), 0, 0).
		AddItem(NewProgress(100), 0, 1)


	b.AddItem(l, 0, 1).AddItem(gr2, 0, 0).AddItem(gr3, 0, 0)

	b.Resize(&Rect{Point: Point{0, 0}, Size: painter.Size()})
	b.Draw(painter)

	want := `
┌List of songs─────────────────────────────────────────────────────────────────┐
│Song-1........................................................................▲
│Song-2........................................................................▓
│Song-3........................................................................░
│Song-4........................................................................░
│Song-5........................................................................░
│Song-6........................................................................▼
└──────────────────────────────────────────────────────────────────────────────┘
┌Now play──────────────────────────────────────────────────────────────────────┐
│Song-1........................................................................│
└──────────────────────────────────────────────────────────────────────────────┘
┌──────┐┌──────┐┌──────────────────────────────────────────────────────────────┐
│.Play.││.Stop.││..............................................................│
└──────┘└──────┘└──────────────────────────────────────────────────────────────┘
`
	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}
