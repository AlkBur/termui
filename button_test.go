package termui

import "testing"

func TestButton_Draw(t *testing.T) {
	for _, tt := range drawButtonTests {
		surface := NewTestSurface(21, 3)
		painter := NewPainter(surface)

		b := tt.setup()
		b.Resize(&Rect{Point: Point{0, 0}, Size: painter.Size()})
		b.Draw(painter)

		if diff := surfaceEquals(surface, tt.want); diff != "" {
			t.Error(diff)
		}
	}
}

var drawButtonTests = []struct {
	test  string
	setup func() *Flexbox
	want  string
}{
	{
		test: "Simple label",
		setup: func() *Flexbox {
			flex := NewFlexbox().SetDirection(Horizontal)
			flex.AddItem(NewButton("Play").SetBorder(true), 0, 0).
				AddItem(NewButton("Stop").SetBorder(true), 0, 0).
				AddItem(NewFlexbox().SetBorder(true),0,1)

			return flex
		},
		want: `
┌──────┐┌──────┐┌───┐
│.Play.││.Stop.││...│
└──────┘└──────┘└───┘
`,
	},
}
