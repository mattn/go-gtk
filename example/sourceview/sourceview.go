package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	gsv "github.com/mattn/go-gtk/gtksourceview"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("SourceView")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.ScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.GTK_SHADOW_IN)
	sourcebuffer := gsv.SourceBufferWithLanguage(gsv.SourceLanguageManagerGetDefault().GetLanguage("cpp"))
	sourceview := gsv.SourceViewWithBuffer(sourcebuffer)

	var start gtk.GtkTextIter
	sourcebuffer.GetStartIter(&start)
	sourcebuffer.BeginNotUndoableAction()
	sourcebuffer.Insert(&start, `#include <iostream>
template<class T>
struct foo_base {
  T operator+(T const &rhs) const {
    T tmp(static_cast<T const &>(*this));
    tmp += rhs;
    return tmp;
  }
};

class foo : public foo_base<foo> {
private:
  int v;
public:
  foo(int v) : v(v) {}
  foo &operator+=(foo const &rhs){
    this->v += rhs.v;
    return *this;
  }
  operator int() { return v; }
};

int main(void) {
  foo a(1), b(2);
  a += b;
  std::cout << (int)a << std::endl;
}
`)
	sourcebuffer.EndNotUndoableAction()

	swin.Add(sourceview)

	window.Add(swin)
	window.SetSizeRequest(400, 300)
	window.ShowAll()

	gtk.Main()
}
