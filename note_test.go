package main

import (
	"testing"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"

	"github.com/stretchr/testify/assert"
)

func testlist() *notelist {
	a := test.NewApp()
	n := &notelist{pref: a.Preferences()}

	return n
}

func TestNoteTitle(t *testing.T) {
	str := "Some content"
	n := &note{binding.BindString(&str)}
	title, _ := n.title().Get()
	assert.Equal(t, str, title)

	n.content.Set("line1\nline2")
	title, _ = n.title().Get()
	assert.Equal(t, "line1", title)

	n.content.Set("")
	title, _ = n.title().Get()
	assert.Equal(t, "Untitled", title)
}

func TestNoteListAdd(t *testing.T) {
	notes := testlist()

	notes.add()
	assert.Equal(t, 1, len(notes.notes))
}

func TestNoteListRemove(t *testing.T) {
	str1 := "remove me"
	str2 := "remove me2"
	first := &note{content: binding.BindString(&str1)}
	second := &note{content: binding.BindString(&str2)}
	notes := testlist()
	notes.notes = []*note{first, second}

	assert.Equal(t, 2, len(notes.notes))
	notes.remove(first)
	assert.Equal(t, 1, len(notes.notes))
	notes.remove(second)
	assert.Equal(t, 0, len(notes.notes))
}

func TestNoteListLoad(t *testing.T) {
	l := testlist()
	n := l.add()
	defer l.remove(n)
	n.content.Set("Test")
	l.save()

	// get a new one
	l.load() // load fresh from preferences
	assert.Equal(t, 1, len(l.notes))
	str, _ := l.notes[0].content.Get()
	assert.Equal(t, "Test", str) // same content
}
