package gtk_test

import (
	"testing"
	
	"github.com/mattn/go-gtk/gtk"
	"github.com/stretchr/testify/assert"
)

func TestFILE_CHOOSER(t *testing.T) {
	gtk.Init(nil)
	d := gtk.NewFileChooserWidget(gtk.FILE_CHOOSER_ACTION_OPEN)
	assert.NotNil(t, d)

	d.SetShowHidden(false)
	assert.False(t, d.GetShowHidden())
	d.SetShowHidden(true)
	assert.True(t, d.GetShowHidden())
}

func TestMisc_GetPadding(t *testing.T) {
	gtk.Init(nil)
	m := gtk.NewImage()
	m.SetPadding(1, 2)
	x, y := m.GetPadding()
	assert.Equal(t, x, 1)
	assert.Equal(t, y, 2)
}
