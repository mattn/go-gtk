package gtk_test

import (
	"io/ioutil"
	"testing"

	"github.com/mattn/go-gtk/gtk"
	"github.com/stretchr/testify/assert"
	"os"
)

func gtkRun() {
	for gtk.EventsPending() {
		gtk.MainIterationDo(false)
	}
}

func TestFILE_CHOOSER(t *testing.T) {
	gtk.Init(nil)
	d := gtk.NewFileChooserWidget(gtk.FILE_CHOOSER_ACTION_SAVE)
	assert.NotNil(t, d)

	d.SetShowHidden(false)
	assert.False(t, d.GetShowHidden())
	d.SetShowHidden(true)
	assert.True(t, d.GetShowHidden())

	d.SetDoOverwriteConfirmation(false)
	assert.False(t, d.GetDoOverwriteConfirmation())
	d.SetDoOverwriteConfirmation(true)
	assert.True(t, d.GetDoOverwriteConfirmation())

	d.SetCreateFolders(false)
	assert.False(t, d.GetCreateFolders())
	d.SetCreateFolders(true)
	assert.True(t, d.GetCreateFolders())

	d.SelectFilename("foobar")
	d.UnselectFilename("foobar")

	d.SelectAll()
	d.UnselectAll()

	f1, err := ioutil.TempFile("/tmp", "go-gtk")
	assert.NoError(t, err)
	f1.Close()
	defer os.Remove(f1.Name())

	f2, err := ioutil.TempFile("/tmp", "go-gtk")
	assert.NoError(t, err)
	f2.Close()
	defer os.Remove(f2.Name())

	d.SetFilename(f1.Name())
	gtkRun()
	assert.Equal(t, f1.Name(), d.GetFilename())

	assert.Equal(t, "file://"+f1.Name(), d.GetUri())
	d.SetUri("file://" + f2.Name())
	assert.Equal(t, "file://"+f2.Name(), d.GetUri())

	assert.True(t, d.SelectUri("file://"+f1.Name()))
	d.UnselectUri("file://"+f1.Name())

	d.UnselectAll()
	gtkRun()
	assert.Equal(t, []string{f1.Name()}, d.GetFilenames())
	assert.Equal(t, []string{"file://" + f1.Name()}, d.GetUris())

	d.SelectFilename(f2.Name())
	assert.Equal(t, []string{f2.Name()}, d.GetFilenames())
	assert.Equal(t, []string{"file://" + f2.Name()}, d.GetUris())
}

func TestFileChooser_SetCurrentName(t *testing.T) {
	gtk.Init(nil)
	d := gtk.NewFileChooserWidget(gtk.FILE_CHOOSER_ACTION_SAVE)
	d.SetCurrentName("foobar")
	// no way to check this until GTK+ 3.10
}


func TestMisc_GetPadding(t *testing.T) {
	gtk.Init(nil)
	m := gtk.NewImage()
	m.SetPadding(1, 2)
	x, y := m.GetPadding()
	assert.Equal(t, x, 1)
	assert.Equal(t, y, 2)
}
