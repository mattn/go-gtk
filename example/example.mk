.PHONY: example
example:
	cd example/action && go build -x .
	cd example/alignment && go build -x .
	cd example/builder && go build -x .
	cd example/clipboard && go build -x .
	cd example/demo && go build -x .
	cd example/spinbutton && go build -x .
	cd example/toolbar && go build -x .
	cd example/dnd && go build -x .
	cd example/drawable && go build -x .
	cd example/event && go build -x .
	cd example/expander && go build -x .
	cd example/iconview && go build -x .
	cd example/listview && go build -x .
	cd example/locale && go build -x .
	cd example/notebook && go build -x .
	cd example/number && go build -x .
	cd example/sourceview && go build -x .
	cd example/statusicon && go build -x .
	cd example/table && go build -x .
	cd example/thread && go build -x .
	cd example/treeview && go build -x .
	cd example/twitter && go build -x .
	cd example/twitterstream && go get && go build -x .
