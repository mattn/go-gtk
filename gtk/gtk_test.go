package gtk

import (
	"testing"
)

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
