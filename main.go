package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var (
	content *widget.Entry
	list *widget.Box
	current *note
)

func setNote(n *note) {
	current = n
	if n == nil {
		content.SetText("")
		return
	}

	content.SetText(n.content)
}

func refreshList(n *notelist) {
	var items []fyne.CanvasObject
	for _, n := range n.notes {
		theNote := n
		b := widget.NewButton(n.title(), func() {
			setNote(theNote)
		})

		if theNote == current {
			b.Style = widget.PrimaryButton
		}
		items = append(items, b)
	}
	list.Children = items
	list.Refresh()
}

func loadUI(n *notelist) fyne.CanvasObject {
	list = widget.NewVBox()
	refreshList(n)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			setNote(n.add())
			refreshList(n)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			n.remove(current)
			refreshList(n)
			if len(n.notes) == 0 {
				setNote(nil)
			}
			setNote(n.notes[0])
		}))

	content = widget.NewMultiLineEntry()
	if len(n.notes) > 0 {
		setNote(n.notes[0])
	}
	content.OnChanged = func(text string) {
		if current == nil {
			return
		}

		current.content = text
		n.save()
		refreshList(n)
	}

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)
	split := widget.NewHSplitContainer(side, content)
	split.Offset = 0.25
	return split
}

func main() {
	a := app.NewWithID("xyz.andy.notes")
	w := a.NewWindow("Notes")

	notes := &notelist{pref: a.Preferences()}
	notes.load()

	w.SetContent(loadUI(notes))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}