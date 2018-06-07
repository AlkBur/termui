package termui

import (
	"image"
	"testing"
)

var drawBoxTests = []struct {
	test  string
	size  image.Point
	setup func() *Flexbox
	want  string
}{

	{
		test: "Empty horizontal box",
		setup: func() *Flexbox {
			b := NewFlexbox().SetBorder(true).
				SetDirection(Horizontal)
			return b
		},
		want: `
┌────────┐
│........│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Horizontal box containing one widget",
		setup: func() *Flexbox {
			b := NewFlexbox().AddItem(NewLabel("test"), 0, 1)
			b.SetBorder(true).SetDirection(Horizontal)
			return b
		},
		want: `
┌────────┐
│test....│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Horizontal box containing multiple widgets",
		setup: func() *Flexbox {
			b := NewFlexbox().SetDirection(Horizontal).
				AddItem(NewLabel("test"),0, 0).
				AddItem(NewLabel("foo"),0, 0)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│testfoo.│
│........│
│........│
└────────┘
`,
	},

	{
		test: "Empty vertical box",
		setup: func() *Flexbox {
			b := NewFlexbox().SetBorder(true).SetDirection(Vertical)
			return b
		},
		want: `
┌────────┐
│........│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Vertical box containing one widget and scroll",
		setup: func() *Flexbox {
			b := NewFlexbox().AddItem(NewLabel("test"),0, 1)
			b.SetScroll(true)
			return b
		},
		want: `
┌────────┐
│test....▲
│........▓
│........▼
└────────┘
`,
	},
	{
		test: "Horizontal box containing one widget and scroll",
		setup: func() *Flexbox {
			b := NewFlexbox().AddItem(NewLabel("test"),0, 1)
			b.SetScroll(true).SetDirection(Horizontal)
			return b
		},
		want: `
┌────────┐
│test....│
│........│
│........│
└◄▓░░░░░►┘
`,
	},
	{
		test: "Vertical box containing multiple widgets",
		size: image.Point{10, 8},
		setup: func() *Flexbox {
			b := NewFlexbox().SetBorder(true).
				AddItem(NewLabel("test"), 0, 1).
				AddItem(NewLabel("foo"),0, 1)
			return b
		},
		want: `
┌────────┐
│test....│
│........│
│........│
│foo.....│
│........│
│........│
└────────┘
`,
	},

	{
		test: "Horizontally centered box",
		size: image.Point{32, 5},
		setup: func() *Flexbox {
			b := NewFlexbox().SetDirection(Horizontal).
				AddItem(NewLabel("test"), 0, 1).
				AddItem(NewLabel("test"), 0, 1).
				AddItem(NewLabel("test"), 0, 1)
			b.SetBorder(true)

			//b := NewHBox(NewSpacer(), nested, NewSpacer())
			//b.SetBorder(true)
			return b
		},
		want: `
┌──────────────────────────────┐
│test......test......test......│
│..............................│
│..............................│
└──────────────────────────────┘
`,
	},

	{
		test: "Horizontally centered box and vertical align",
		size: image.Point{32, 5},
		setup: func() *Flexbox {
			b := NewFlexbox().SetDirection(Horizontal).
				AddItem(NewLabel("test").SetAlignV(AlignVerticalCenter), 0, 1).
				AddItem(NewLabel("test").SetAlignV(AlignTop), 0, 1).
				AddItem(NewLabel("test").SetAlignV(AlignBottom), 0, 1)
			b.SetBorder(true)

			//b := NewHBox(NewSpacer(), nested, NewSpacer())
			//b.SetBorder(true)
			return b
		},
		want: `
┌──────────────────────────────┐
│..........test................│
│test..........................│
│....................test......│
└──────────────────────────────┘
`,
	},

}

func TestBox_Draw(t *testing.T) {
	for _, tt := range drawBoxTests {
		t.Run(tt.test, func(t *testing.T) {
			var surface *TestSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = NewTestSurface(10, 5)
			} else {
				surface = NewTestSurface(tt.size.X, tt.size.Y)
			}

			painter := NewPainter(surface)
			wg := tt.setup()
			wg.Resize(&Rect{Point: Point{0, 0}, Size: painter.Size()})
			wg.Draw(painter)

			if diff := surfaceEquals(surface, tt.want); diff != "" {
				//t.Logf("%v ?? %v\n", len(surface.String()), len(tt.want))
				//arr1 := []byte(surface.String())
				//arr2 := []byte(tt.want)
				//for i:= 0; i< len(arr1); i++ {
				//	if arr1[i] != arr2[i] {
				//		t.Logf("%d ?? %d\n", arr1[i], arr2[i])
				//	}
				//}
				t.Error(diff)
			}
		})
	}
}
