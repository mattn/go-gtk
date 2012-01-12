go-gtk
======

WHATS:
------

  Go bindings for GTK 

SCREENSHOT:
-----------

![Go GTK!](https://github.com/mattn/go-gtk/raw/gh-pages/static/images/screenshot.png "Go GTK!")

INSTALL:
--------

  To experiment with go-gtk, you can just compile and run the example
  program:

    git clone  https://github.com/mattn/go-gtk
    cd go-gtk
    gomake install
    gomake example
    ./example/demo/demo

  Or

    goinstall github.com/mattn/go-gtk/gtk

  Don't forget, that you need the GTK-Development-Packages.

LICENSE:
--------

  The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.

AUTHORS:
--------

  * Yasuhiro Matsumoto
  * David Roundy
  * Mark Andrew Gerads
  * Tobias Kortkamp
  * Mikhail Trushnikov
  * Federico Sogaro

GOAL:
-----

  Hopefully support following widgets and methods enough to run general application. 

    GtkAboutDialog: 90%
    GtkAccelLabel: 90%
    GtkAdjustment: 100%
    GtkAlignment: 100%
    GtkAssistant: 90%
    GtkBox: 100%
    GtkButton: 50%
    GtkCheckButton: 100%
    GtkCheckMenuItem: 100%
    GtkColorButton: 0%
    GtkColorSelectionDialog: 0%
    GtkComboBox: 80%
    GtkComboBoxEntry: 100%
    GtkEntry: 10%
    GtkFileChooserButton: 0%
    GtkFileChooserDialog: 100%
    GtkFixed: 100%
    GtkFontButton: 100%
    GtkFontSelectionDialog: 50%
    GtkFrame: 100%
    GtkHBox: 100%
    GtkHButtonBox: 0%
    GtkIconView: 0%
    GtkImage: 30%
    GtkItem: 100%
    GtkLabel: 10%
    GtkLinkButton: 90%
    GtkMenu: 100%
    GtkMenuBar: 100%
    GtkMenuItem: 100%
    GtkMessageDialog: 10%
    GtkNotebook: 90%
    GtkPageSetupUnixDialog: 0%
    GtkPaned: 100%
    GtkPrintUnixDialog: 0%
    GtkProgressBar: 60%
    GtkRadioButton: 100%
    GtkRecentChooserDialog: 0%
    GtkScale: 90%
    GtkScrolledWindow: 50%
    GtkSeparator: 100%
    GtkSpinButton: 0%
    GtkStatusbar: 100%
    GtkTable: 100%
    GtkTextView: 10%
    GtkToggleButton: 100%
    GtkToolbar: 0%
    GtkTreePath: 90%
    GtkTreeView: 10%
    GtkTreeViewColumn: 5%
    GtkVBox: 100%
    GtkVButtonBox: 0%
    GtkVolumeButton: 0%
    GtkWidget: 5%
    GtkWindow: 5%
