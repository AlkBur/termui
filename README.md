#Terminal UI for Go

A UI library for terminal applications.

## Installation

```
go get github.com/AlkBur/termui
```

## Usage

```go
package main

import "github.com/AlkBur/termui"

func main() {
    runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	box := termui.NewFlexbox().
	    AddItem(termui.NewLabel("ui"),0,1)

	ui, err := termui.New(box)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
```

## License

termui released under the [MIT License](LICENSE).

## Acknowledgments

The following open-source libraries were used in the creation of `termui`.
Many thanks to all these developers.

* [tui-go](https://github.com/marcusolsson/tui-go)
* [tcell](https://github.com/gdamore/tcell)
* [tview](https://github.com/rivo/tview)