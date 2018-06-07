package termui

import (
	"github.com/mattn/go-runewidth"
	"image"
	"github.com/gdamore/tcell"
	"fmt"
	"strings"
)

type TestSurface struct {
	cells   map[image.Point]testCell
	cursor  Point
	size    Size
	emptyCh rune
}

type testCell struct {
	Rune  []rune
	Style tcell.Style
}

func NewTestSurface(w, h int) *TestSurface {
	return &TestSurface{
		cells:   make(map[image.Point]testCell),
		size:    Size{uint(w), uint(h)},
		emptyCh: '.',
	}
}

func (this *TestSurface) Init() error {	return nil }

func (this *TestSurface) Fini(){}

func (this *TestSurface) Clear(){
	this.cells = make(map[image.Point]testCell)
}

func (this *TestSurface) Fill(rune, tcell.Style){}

func (s *TestSurface) SetCell(x int, y int, style tcell.Style, ch ...rune){
	s.cells[image.Point{x, y}] = testCell{
		Rune:  ch,
		Style: style,
	}
}

func (s *TestSurface) GetContent(x, y int) (mainc rune, combc []rune, style tcell.Style, width int) { return }

func (s *TestSurface) SetContent(x int, y int, mainc rune, combc []rune, style tcell.Style) {}

func (s *TestSurface) SetStyle(style tcell.Style){}

func (s *TestSurface) ShowCursor(x int, y int) {
	s.cursor = Point{uint(x), uint(y)}
}

func (s *TestSurface) HideCursor()  {
	s.cursor = Point{}
}

func (s *TestSurface) Size() (int, int) {
	return int(s.size.W), int(s.size.H)
}

func (s *TestSurface) PollEvent() tcell.Event { return nil }

func (s *TestSurface) PostEvent(ev tcell.Event) error {return nil}

func (s *TestSurface) PostEventWait(ev tcell.Event) {}

func (s *TestSurface) EnableMouse() {}

func (s *TestSurface) DisableMouse() {}

func (s *TestSurface) HasMouse() bool { return true }

func (s *TestSurface) Colors() int {return 0}

func (s *TestSurface) Show() {}

func (s *TestSurface) Sync() {}

func (s *TestSurface) CharacterSet() string {return ""}

func (s *TestSurface) RegisterRuneFallback(r rune, subst string) {}

func (s *TestSurface) UnregisterRuneFallback(r rune) {}

func (s *TestSurface) CanDisplay(r rune, checkFallbacks bool) bool {return true}

func (s *TestSurface) Resize(int, int, int, int) {}

func (s *TestSurface) HasKey(tcell.Key) bool {return false}

func surfaceEquals(surface *TestSurface, want string) string {
	var b [][]rune
	//b = want
	arr := strings.Split(want, "\n")
	if len(arr)-2 < 0 {
		return fmt.Sprintf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
	b = make([][]rune, len(arr)-2)
	for i:= 0; i<len(b); i++{
		line := []rune(arr[i+1])
		b[i]=line
	}
	if !surface.Compare(b) {
		return fmt.Sprintf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
	return ""
}

func (s *TestSurface) Compare(b [][]rune) bool {
	if len(b) != int(s.size.H) {
		return false
	}
	for j := uint(0); j < s.size.H; j++ {
		if len(b[j]) != int(s.size.W) {
			return false
		}
		for i := uint(0); i < s.size.W; i++ {
			if cell, ok := s.cells[image.Point{int(i), int(j)}]; ok {
				if len(cell.Rune) == 0 {
					return false
				}else if len(cell.Rune) == 1 {
					if b[j][i]!= cell.Rune[0] {
						return false
					}
				}else{
					return false
				}
			}else{
				if b[j][i]!= s.emptyCh {
					return false
				}
			}
		}
	}
	return true
}

func (s *TestSurface) String() string {
	var buf strings.Builder
	buf.WriteByte('\n')
	for j := uint(0); j < s.size.H; j++ {
		for i := uint(0); i < s.size.W; i++ {
			if cell, ok := s.cells[image.Point{int(i), int(j)}]; ok {
				for _, r := range cell.Rune {
					buf.WriteRune(r)
					if w := runeWidth(r); w > 1 {
						i += w - 1
					}
				}
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func runeWidth(r rune) uint {
	return uint(runewidth.RuneWidth(r))
}

func stringWidth(s string) uint {
	return uint(runewidth.StringWidth(s))
}

func correctSize(max, cur Size) Size {
	if cur.H > max.H {
		cur.H = max.H
	}
	if cur.W > max.W {
		cur.W = max.W
	}
	return cur
}
