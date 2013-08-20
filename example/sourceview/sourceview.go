package main

import (
	"github.com/mattn/go-gtk/gtk"
	gsv "github.com/mattn/go-gtk/gtksourceview"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("SourceView")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	sourcebuffer := gsv.NewSourceBufferWithLanguage(gsv.SourceLanguageManagerGetDefault().GetLanguage("cpp"))
	sourceview := gsv.NewSourceViewWithBuffer(sourcebuffer)

	var start gtk.TextIter
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
