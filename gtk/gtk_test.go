package gtk_test

import (
	"testing"
	"github.com/mattn/go-gtk/gtk"
	"github.com/stretchr/testify/assert"
)

func TestFILE_CHOOSER(t *testing.T) {
	d := gtk.NewFileChooserWidget(gtk.FILE_CHOOSER_ACTION_OPEN)
	assert.NotNil(t, d)

	d.SetShowHidden(false)
	assert.False(t, d.GetShowHidden())
	d.SetShowHidden(true)
	assert.True(t, d.GetShowHidden())
}
