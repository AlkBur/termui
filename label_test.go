package termui

import "testing"

func TestLabel_Draw(t *testing.T) {
	for _, tt := range drawLabelTests {
		surface := NewTestSurface(10, 5)
		painter := NewPainter(surface)

		b := NewFlexbox()
		b.AddItem(tt.setup(), 0, 1)

		b.Resize(&Rect{Point: Point{0, 0}, Size: painter.Size()})
		b.Draw(painter)

		if diff := surfaceEquals(surface, tt.want); diff != "" {
			t.Error(diff)
		}
	}
}

var drawLabelTests = []struct {
	test  string
	setup func() *Label
	want  string
}{
	{
		test: "Simple label",
		setup: func() *Label {
			return NewLabel("test")
		},
		want: `
test......
..........
..........
..........
..........
`,
	},
	{
		test: "Word wrap",
		setup: func() *Label {
			l := NewLabel("this will wrap")
			l.SetWordWrap(true)
			return l
		},
		want: `
this will.
wrap......
..........
..........
..........
`,
	},
	{
		test: "Repeat",
		setup: func() *Label {
			l := NewLabel("a")
			l.SetRepeat(2)
			return l
		},
		want: `
aaa.......
..........
..........
..........
..........
`,
	},

}
