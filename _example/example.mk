EXAMPLES := \
	_example/action/action \
	_example/alignment/alignment \
	_example/builder/builder \
	_example/clipboard/clipboard \
	_example/demo/demo \
	_example/arrow/arrow \
	_example/dnd/dnd \
	_example/drawable/drawable \
	_example/event/event \
	_example/expander/expander \
	_example/iconview/iconview \
	_example/listview/listview \
	_example/locale/locale \
	_example/notebook/notebook \
	_example/number/number \
	_example/sourceview/sourceview \
	_example/spinbutton/spinbutton \
	_example/statusicon/statusicon \
	_example/table/table \
	_example/thread/thread \
	_example/toolbar/toolbar \
	_example/treeview/treeview \
	_example/twitterstream/twitterstream

.PHONY: example
example: $(EXAMPLES)
	@true

.PHONY: clean-example
clean-example:
	rm -f $(EXAMPLES)

%: %.go
	cd $(@D) && go build -x

example/action/action: example/action/action.go
example/alignment/alignment: example/alignment/alignment.go
example/builder/builder: example/builder/builder.go
example/clipboard/clipboard: example/clipboard/clipboard.go
example/demo/demo: example/demo/demo.go
example/arrow/arrow: example/arrow/arrow.go
example/dnd/dnd: example/dnd/dnd.go
example/drawable/drawable: example/drawable/drawable.go
example/event/event: example/event/event.go
example/expander/expander: example/expander/expander.go
example/iconview/iconview: example/iconview/iconview.go
example/listview/listview: example/listview/listview.go
example/locale/locale: example/locale/locale.go
example/notebook/notebook: example/notebook/notebook.go
example/number/number: example/number/number.go
example/sourceview/sourceview: example/sourceview/sourceview.go
example/spinbutton/spinbutton: example/spinbutton/spinbutton.go
example/statusicon/statusicon: example/statusicon/statusicon.go
example/table/table: example/table/table.go
example/thread/thread: example/thread/thread.go
example/toolbar/toolbar: example/toolbar/toolbar.go
example/treeview/treeview: example/treeview/treeview.go
example/twitterstream/twitterstream: example/twitterstream/twitterstream.go
