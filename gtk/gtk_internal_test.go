package gtk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAboutDialog_GetAuthors(t *testing.T) {
	Init(nil)
	dialog := NewAboutDialog()
	assert.Equal(t, len(dialog.GetAuthors()), 0)
	dialog.SetAuthors([]string{"a", ""})
	assert.Equal(t, len(dialog.GetAuthors()), 2)
	dialog.SetAuthors(nil)
	assert.Equal(t, len(dialog.GetAuthors()), 0)
}

func TestUpdateTreeViewColumns(t *testing.T) {
	Init(nil)

	columnsDesc := []struct {
		Title   string
		Type    string
		GTKType ICellRenderer
	}{
		{"Col_1", "text", NewCellRendererText()},
		{"Col_2", "text", NewCellRendererText()},
	}

	tw := NewTreeView()
	for i, c := range columnsDesc {
		tw.AppendColumn(NewTreeViewColumnWithAttributes(c.Title, c.GTKType, c.Type, i))
	}

	columns := tw.GetColumns()

	if len(columns) != len(columnsDesc) {
		t.Error("Wrong number of the columns:", len(columns), len(columnsDesc))
	} else {

		for i, _ := range columnsDesc {
			if columns[i].GetTitle() != columnsDesc[i].Title {
				t.Error("Wrong column title:", columns[i].GetTitle(), columnsDesc[i].Title)
			}
		}
	}

	for lastIndex := len(columns) - 1; lastIndex >= 0; {
		c := columns[lastIndex]

		after := tw.RemoveColumn(c)

		numbersColumns := lastIndex
		if numbersColumns != after {
			t.Error("Failed remove column:", numbersColumns, after)
		}

		lastIndex -= 1
		if lastIndex >= 0 {
			if title := tw.GetColumns()[lastIndex].GetTitle(); title != columnsDesc[lastIndex].Title {
				t.Error("Wrong column title:", title, columnsDesc[lastIndex].Title)
			}
		}
	}

	if count := len(tw.GetColumns()); count != 0 {
		t.Error("Wrong number of the columns:", count)
	}
}

func TestDialog_GetWidgetForResponse(t *testing.T) {
	Init(nil)
	dialog := NewDialog()
	dialog.AddButton("A", RESPONSE_ACCEPT)
	dialog.AddButton("B", RESPONSE_CANCEL)
	accept := dialog.GetWidgetForResponse(RESPONSE_ACCEPT)
	cancel := dialog.GetWidgetForResponse(RESPONSE_CANCEL)
	assert.Equal(t, newButtonInternal(accept.GWidget).GetLabel(), "A")
	assert.Equal(t, newButtonInternal(cancel.GWidget).GetLabel(), "B")
}

func TestEntry_SetInnerBorder(t *testing.T) {
	Init(nil)
	e := NewEntry()
	assert.Nil(t, e.GetInnerBorder())
	e.SetInnerBorder(&Border{1,2,3,4,})
	border := e.GetInnerBorder()
	assert.NotNil(t, border)
	assert.Equal(t, *border, Border{1,2,3,4,})
	e.SetInnerBorder(nil)
	assert.Nil(t, e.GetInnerBorder())
}
