# go-gtk

## WHATS

  Go bindings for GTK

## SCREENSHOT

![Go GTK!](https://github.com/mattn/go-gtk/raw/gh-pages/static/images/screenshot.png "Go GTK!")

## INSTALL

  You can experiment with go-gtk
  by running the various example programs:

    git clone  https://github.com/mattn/go-gtk
    cd go-gtk
    make example
    ./example/demo/demo

  Or

    go get github.com/mattn/go-gtk/gtk

  Don't forget, that you need the GTK-Development-Packages.

  If you use linux, you should install `libgtk+-2.0` and packages that depend on gtk.

  On Debian and Debian derived distributions you can run (as root):

    apt-get install libgtk2.0-dev libglib2.0-dev libgtksourceview2.0-dev

  If you use windows, find gtk binary packages from here:

  * http://www.gtk.org/download/windows.php

## LICENSE

  The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (GNU Lesser General Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.

## AUTHOR

  * Yasuhiro Matsumoto

## CONTRIBUTE AUTHORS

  * David Roundy
  * Mark Andrew Gerads
  * Tobias Kortkamp
  * Mikhail Trushnikov
  * Federico Sogaro
  * Crazy2be
  * Daniël de Kok
  * Erik Lissel
  * Jeffrey Bolle
  * Leonhard Küper
  * Matt Joiner
  * SQP
  * Steven T
  * Taru Karttunen
  * Utkan Güngördü
  * matiaslina
  * Dag Robøle
  * Denis Dyakov

## GOAL

  Hopefully support following widgets and methods enough to run general application.

(output of tools/gogtkinfo)

    Main Loop and Events          :  30% (  8/ 26)
    GtkAccelGroup                 :  15% (  3/ 19)
    GtkAccelMap                   :   0% (  0/ 14)
    GtkClipboard                  :  25% (  8/ 31)
    Drag and Drop                 :  11% (  4/ 35)
    GtkIconTheme                  :   0% (  0/ 31)
    GtkStockItem                  :  66% (  4/  6)
    Themeable Stock Images        :   0% (  0/ 41)
    Resource Files                :  14% (  4/ 28)
    GtkSettings                   :  40% (  6/ 15)
    GtkBinding                    :   0% (  0/ 14)
    Graphics Contexts             :   0% (  0/  2)
    GtkStyle                      :   9% (  6/ 64)
    Selections                    :   8% (  4/ 47)
    Version Information           :   0% (  0/  6)
    Testing                       :   0% (  0/ 16)
    Filesystem Utilities          :   0% (  0/  7)
    GtkDialog                     :  45% (  9/ 20)
    GtkMessageDialog              :  62% (  5/  8)
    GtkWindow                     :  49% ( 50/102)
    GtkWindowGroup                :   0% (  0/  5)
    GtkAboutDialog                :  90% ( 29/ 32)
    GtkAssistant                  :  91% ( 21/ 23)
    GtkOffscreenWindow            :   0% (  0/  3)
    GtkAccelLabel                 :  83% (  5/  6)
    GtkImage                      :  32% ( 10/ 31)
    GtkLabel                      :  86% ( 39/ 45)
    GtkProgressBar                : 100% ( 12/ 12)
    GtkStatusbar                  :  77% (  7/  9)
    GtkInfoBar                    : 100% ( 12/ 12)
    GtkStatusIcon                 :  68% ( 26/ 38)
    GtkSpinner                    : 100% (  3/  3)
    GtkButton                     : 100% ( 28/ 28)
    GtkCheckButton                : 100% (  3/  3)
    GtkRadioButton                : 100% (  8/  8)
    GtkToggleButton               : 100% (  9/  9)
    GtkLinkButton                 :  75% (  6/  8)
    GtkScaleButton                : 100% (  9/  9)
    GtkVolumeButton               : 100% (  1/  1)
    GtkEntry                      :  41% ( 26/ 63)
    GtkEntryBuffer                :  72% (  8/ 11)
    GtkEntryCompletion            :  96% ( 25/ 26)
    GtkHScale                     : 100% (  2/  2)
    GtkVScale                     : 100% (  2/  2)
    GtkSpinButton                 : 100% ( 30/ 30)
    GtkEditable                   : 100% ( 13/ 13)
    GtkTextIter                   :  25% ( 23/ 91)
    GtkTextMark                   :   0% (  0/  7)
    GtkTextBuffer                 :  67% ( 52/ 77)
    GtkTextTag                    :  83% (  5/  6)
    GtkTextAttributes             : 100% (  5/  5)
    GtkTextTagTable               :  83% (  5/  6)
    GtkTextView                   :  29% ( 19/ 64)
    GtkTreePath                   :  90% ( 18/ 20)
    GtkTreeRowReference           :  60% (  6/ 10)
    GtkTreeIter                   : 100% (  2/  2)
    GtkTreeModel                  :  57% ( 15/ 26)
    GtkTreeSelection              :  78% ( 18/ 23)
    GtkTreeViewColumn             :  61% ( 34/ 55)
    GtkTreeView                   :  18% ( 18/ 98)
    GtkTreeView drag-and-drop     :   0% (  0/  7)
    GtkCellView                   :   0% (  0/ 11)
    GtkIconView                   :  17% ( 11/ 62)
    GtkTreeSortable               :  62% (  5/  8)
    GtkTreeModelSort              :   0% (  0/  9)
    GtkTreeModelFilter            :   0% (  0/ 11)
    GtkCellLayout                 :   0% (  0/  9)
    GtkCellRenderer               : 100% (  2/  2)
    GtkCellEditable               :   0% (  0/  3)
    GtkCellRendererAccel          : 100% (  1/  1)
    GtkCellRendererCombo          : 100% (  1/  1)
    GtkCellRendererPixbuf         : 100% (  1/  1)
    GtkCellRendererProgress       : 100% (  1/  1)
    GtkCellRendererSpin           : 100% (  1/  1)
    GtkCellRendererText           : 100% (  2/  2)
    GtkCellRendererToggle         : 100% (  7/  7)
    GtkCellRendererSpinner        : 100% (  1/  1)
    GtkListStore                  :  84% ( 16/ 19)
    GtkTreeStore                  :  80% ( 17/ 21)
    GtkComboBox                   :  78% ( 30/ 38)
    GtkComboBoxText               : 100% (  7/  7)
    GtkComboBoxEntry              :  80% (  4/  5)
    GtkMenu                       :  50% ( 15/ 30)
    GtkMenuBar                    : 100% (  8/  8)
    GtkMenuItem                   :  90% ( 18/ 20)
    GtkImageMenuItem              :  54% (  6/ 11)
    GtkRadioMenuItem              :  44% (  4/  9)
    GtkCheckMenuItem              : 100% ( 10/ 10)
    GtkSeparatorMenuItem          : 100% (  1/  1)
    GtkTearoffMenuItem            : 100% (  1/  1)
    GtkToolShell                  :   0% (  0/  9)
    GtkToolbar                    :  63% ( 24/ 38)
    GtkToolItem                   :  76% ( 19/ 25)
    GtkToolPalette                :  59% ( 13/ 22)
    GtkToolItemGroup              :  47% (  8/ 17)
    GtkSeparatorToolItem          : 100% (  3/  3)
    GtkToolButton                 : 100% ( 15/ 15)
    GtkMenuToolButton             :  85% (  6/  7)
    GtkToggleToolButton           : 100% (  5/  5)
    GtkRadioToolButton            :  33% (  2/  6)
    GtkUIManager                  :  29% (  5/ 17)
    GtkActionGroup                :  55% ( 11/ 20)
    GtkAction                     :  93% ( 44/ 47)
    GtkToggleAction               : 100% (  6/  6)
    GtkRadioAction                : 100% (  5/  5)
    GtkRecentAction               :  75% (  3/  4)
    GtkActivatable                :  66% (  4/  6)
    GtkColorButton                : 100% ( 10/ 10)
    GtkColorSelectionDialog       :   0% (  0/  2)
    GtkColorSelection             :   0% (  0/ 21)
    GtkHSV                        :   0% (  0/  8)
    GtkFileChooser                :  22% ( 13/ 58)
    GtkFileChooserButton          :  18% (  2/ 11)
    GtkFileChooserDialog          : 100% (  1/  1)
    GtkFileChooserWidget          :  50% (  1/  2)
    GtkFileFilter                 :  55% (  5/  9)
    GtkFontButton                 : 100% ( 14/ 14)
    GtkFontSelection              :  28% (  4/ 14)
    GtkFontSelectionDialog        : 100% (  8/  8)
    GtkInputDialog                :   0% (  0/  1)
    GtkAlignment                  : 100% (  4/  4)
    GtkAspectFrame                :   0% (  0/  2)
    GtkHBox                       : 100% (  1/  1)
    GtkVBox                       : 100% (  1/  1)
    GtkHButtonBox                 :   0% (  0/  5)
    GtkVButtonBox                 :   0% (  0/  5)
    GtkFixed                      : 100% (  5/  5)
    GtkHPaned                     : 100% (  1/  1)
    GtkVPaned                     : 100% (  1/  1)
    GtkLayout                     : 100% ( 12/ 12)
    GtkNotebook                   :  90% ( 50/ 55)
    GtkTable                      :  93% ( 14/ 15)
    GtkExpander                   :  87% ( 14/ 16)
    GtkOrientable                 :   0% (  0/  2)
    GtkFrame                      : 100% (  9/  9)
    GtkHSeparator                 : 100% (  1/  1)
    GtkVSeparator                 : 100% (  1/  1)
    GtkScrollbar                  : 100% (  0/  0)
    GtkHScrollbar                 : 100% (  1/  1)
    GtkVScrollbar                 : 100% (  1/  1)
    GtkScrolledWindow             :  86% ( 13/ 15)
    GtkPrintOperation             :  13% (  5/ 36)
    GtkPrintContext               :  18% (  2/ 11)
    GtkPrintSettings              :   0% (  0/ 74)
    GtkPageSetup                  :   0% (  0/ 25)
    GtkPaperSize                  :   0% (  0/ 21)
    GtkPrinter                    :   0% (  0/ 23)
    GtkPrintJob                   :   0% (  0/ 10)
    GtkPrintUnixDialog            :   0% (  0/ 18)
    GtkPageSetupUnixDialog        :   0% (  0/  5)
    GtkAdjustment                 :  83% ( 15/ 18)
    GtkArrow                      : 100% (  2/  2)
    GtkCalendar                   :   0% (  0/ 17)
    GtkDrawingArea                : 100% (  2/  2)
    GtkEventBox                   :  20% (  1/  5)
    GtkHandleBox                  :   0% (  0/  8)
    GtkIMContextSimple            :   0% (  0/  2)
    GtkIMMulticontext             :   0% (  0/  4)
    GtkSizeGroup                  : 100% (  8/  8)
    GtkTooltip                    :  60% (  6/ 10)
    GtkViewport                   : 100% (  9/  9)
    GtkAccessible                 : 100% (  3/  3)
    GtkBin                        : 100% (  1/  1)
    GtkBox                        : 100% ( 11/ 11)
    GtkButtonBox                  :   0% (  0/ 10)
    GtkContainer                  :  48% ( 16/ 33)
    GtkItem                       : 100% (  3/  3)
    GtkMenuShell                  :   0% (  0/ 11)
    GtkMisc                       : 100% (  4/  4)
    GtkObject                     : 100% (  2/  2)
    GtkPaned                      :  88% (  8/  9)
    GtkRange                      :  53% ( 16/ 30)
    GtkScale                      :  90% (  9/ 10)
    GtkSeparator                  : 100% (  0/  0)
    GtkWidget                     :  54% (101/187)
    GtkIMContext                  :   0% (  0/ 11)
    GtkPlug                       :   0% (  0/  7)
    GtkSocket                     :   0% (  0/  5)
    GtkRecentManager              :   0% (  0/ 37)
    GtkRecentChooser              :   0% (  0/ 33)
    GtkRecentChooserDialog        :   0% (  0/  2)
    GtkRecentChooserMenu          :   0% (  0/  4)
    GtkRecentChooserWidget        :   0% (  0/  2)
    GtkRecentFilter               :   0% (  0/ 12)
    GtkBuildable                  :   0% (  0/ 10)

    Total progress :                 47% (1519/3180)

