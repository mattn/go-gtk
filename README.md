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

(output of tools/gogtkinfo)

    Drag and Drop                 :   0% (  0/ 36)
    Filesystem Utilities          :   0% (  0/  7)
    Graphics Contexts             :   0% (  0/  2)
    GtkAboutDialog                :  90% ( 29/ 32)
    GtkAccelGroup                 :   5% (  1/ 19)
    GtkAccelLabel                 :  83% (  5/  6)
    GtkAccelMap                   :   0% (  0/ 14)
    GtkAccessible                 :   0% (  0/  3)
    GtkAction                     :   0% (  0/ 46)
    GtkActionGroup                :   0% (  0/ 20)
    GtkActivatable                :   0% (  0/  6)
    GtkAdjustment                 :  82% ( 14/ 17)
    GtkAlignment                  : 100% (  4/  4)
    GtkArrow                      :   0% (  0/  2)
    GtkAspectFrame                :   0% (  0/  2)
    GtkAssistant                  :  91% ( 21/ 23)
    GtkBin                        : 100% (  1/  1)
    GtkBinding                    :   0% (  0/ 14)
    GtkBox                        : 100% ( 11/ 11)
    GtkBuildable                  :   0% (  0/ 10)
    GtkButton                     :  37% ( 10/ 27)
    GtkButtonBox                  :   0% (  0/ 10)
    GtkCalendar                   :   0% (  0/ 17)
    GtkCellEditable               :   0% (  0/  3)
    GtkCellLayout                 :   0% (  0/  9)
    GtkCellRenderer               : 100% (  2/  2)
    GtkCellRendererAccel          :   0% (  0/  1)
    GtkCellRendererCombo          :   0% (  0/  1)
    GtkCellRendererPixbuf         : 100% (  1/  1)
    GtkCellRendererProgress       : 100% (  1/  1)
    GtkCellRendererSpin           : 100% (  1/  1)
    GtkCellRendererSpinner        : 100% (  1/  1)
    GtkCellRendererText           : 100% (  2/  2)
    GtkCellRendererToggle         : 100% (  7/  7)
    GtkCellView                   :   0% (  0/ 11)
    GtkCheckButton                : 100% (  3/  3)
    GtkCheckMenuItem              : 100% ( 10/ 10)
    GtkClipboard                  :   0% (  0/ 28)
    GtkColorButton                :   0% (  0/ 10)
    GtkColorSelection             :   0% (  0/ 21)
    GtkColorSelectionDialog       :   0% (  0/  2)
    GtkComboBox                   :  78% ( 30/ 38)
    GtkComboBoxEntry              :  80% (  4/  5)
    GtkComboBoxText               : 100% (  7/  7)
    GtkContainer                  :  18% (  6/ 33)
    GtkDialog                     :  45% (  9/ 20)
    GtkDrawingArea                : 100% (  2/  2)
    GtkEditable                   : 100% ( 13/ 13)
    GtkEntry                      :  36% ( 25/ 69)
    GtkEntryBuffer                :   0% (  0/ 11)
    GtkEntryCompletion            :   0% (  0/ 25)
    GtkEventBox                   :  20% (  1/  5)
    GtkExpander                   :  87% ( 14/ 16)
    GtkFileChooser                :  22% ( 13/ 58)
    GtkFileChooserButton          :   0% (  0/  9)
    GtkFileChooserDialog          : 100% (  1/  1)
    GtkFileChooserWidget          :   0% (  0/  2)
    GtkFileFilter                 :  55% (  5/  9)
    GtkFixed                      : 100% (  5/  5)
    GtkFontButton                 :  71% ( 10/ 14)
    GtkFontSelection              :   0% (  0/ 14)
    GtkFontSelectionDialog        :  37% (  3/  8)
    GtkFrame                      : 100% (  9/  9)
    GtkHBox                       : 100% (  1/  1)
    GtkHButtonBox                 :   0% (  0/  5)
    GtkHPaned                     : 100% (  1/  1)
    GtkHSV                        :   0% (  0/  8)
    GtkHScale                     : 100% (  1/  1)
    GtkHScrollbar                 :   0% (  0/  1)
    GtkHSeparator                 : 100% (  1/  1)
    GtkHandleBox                  :   0% (  0/  8)
    GtkIMContext                  :   0% (  0/ 11)
    GtkIMContextSimple            :   0% (  0/  2)
    GtkIMMulticontext             :   0% (  0/  4)
    GtkIconTheme                  :   0% (  0/ 31)
    GtkIconView                   :   0% (  0/ 62)
    GtkImage                      :  29% (  9/ 31)
    GtkImageMenuItem              :   0% (  0/ 11)
    GtkInfoBar                    :   0% (  0/ 12)
    GtkInputDialog                :   0% (  0/  1)
    GtkItem                       : 100% (  3/  3)
    GtkLabel                      :  82% ( 37/ 45)
    GtkLayout                     :   0% (  0/ 12)
    GtkLinkButton                 :  75% (  6/  8)
    GtkListStore                  :  83% ( 15/ 18)
    GtkMenu                       :  50% ( 15/ 30)
    GtkMenuBar                    : 100% (  8/  8)
    GtkMenuItem                   :  78% ( 15/ 19)
    GtkMenuShell                  :   0% (  0/ 11)
    GtkMenuToolButton             :   0% (  0/  7)
    GtkMessageDialog              :  12% (  1/  8)
    GtkMisc                       :   0% (  0/  4)
    GtkNotebook                   :  90% ( 50/ 55)
    GtkObject                     : 100% (  0/  0)
    GtkOffscreenWindow            :   0% (  0/  3)
    GtkOrientable                 :   0% (  0/  2)
    GtkPageSetup                  :   0% (  0/ 25)
    GtkPageSetupUnixDialog        :   0% (  0/  5)
    GtkPaned                      :  88% (  8/  9)
    GtkPaperSize                  :   0% (  0/ 21)
    GtkPlug                       :   0% (  0/  7)
    GtkPrintContext               :   0% (  0/ 11)
    GtkPrintJob                   :   0% (  0/ 10)
    GtkPrintOperation             :   0% (  0/ 35)
    GtkPrintSettings              :   0% (  0/ 74)
    GtkPrintUnixDialog            :   0% (  0/ 18)
    GtkPrinter                    :   0% (  0/ 23)
    GtkProgressBar                :  83% ( 10/ 12)
    GtkRadioAction                :   0% (  0/  5)
    GtkRadioButton                : 100% (  8/  8)
    GtkRadioMenuItem              :   0% (  0/  9)
    GtkRadioToolButton            :   0% (  0/  6)
    GtkRange                      :  53% ( 16/ 30)
    GtkRecentAction               :   0% (  0/  4)
    GtkRecentChooser              :   0% (  0/ 33)
    GtkRecentChooserDialog        :   0% (  0/  2)
    GtkRecentChooserMenu          :   0% (  0/  4)
    GtkRecentChooserWidget        :   0% (  0/  2)
    GtkRecentFilter               :   0% (  0/ 12)
    GtkRecentManager              :   0% (  0/ 37)
    GtkScale                      :  90% (  9/ 10)
    GtkScaleButton                :   0% (  0/  9)
    GtkScrolledWindow             :  86% ( 13/ 15)
    GtkSeparator                  : 100% (  0/  0)
    GtkSeparatorMenuItem          : 100% (  1/  1)
    GtkSeparatorToolItem          :   0% (  0/  3)
    GtkSettings                   :  23% (  3/ 13)
    GtkSizeGroup                  : 100% (  8/  8)
    GtkSocket                     :   0% (  0/  5)
    GtkSpinButton                 :   0% (  0/ 25)
    GtkSpinner                    :   0% (  0/  3)
    GtkStatusIcon                 :  37% ( 14/ 37)
    GtkStatusbar                  :  77% (  7/  9)
    GtkStockItem                  :  66% (  4/  6)
    GtkStyle                      :   0% (  0/ 64)
    GtkTable                      :  93% ( 14/ 15)
    GtkTearoffMenuItem            :   0% (  0/  1)
    GtkTextBuffer                 :  67% ( 52/ 77)
    GtkTextIter                   :  20% ( 19/ 91)
    GtkTextMark                   :   0% (  0/  7)
    GtkTextTag                    :   0% (  0/  9)
    GtkTextTagTable               :  83% (  5/  6)
    GtkTextView                   :  28% ( 18/ 64)
    GtkToggleAction               :   0% (  0/  6)
    GtkToggleButton               : 100% (  9/  9)
    GtkToggleToolButton           :   0% (  0/  4)
    GtkToolButton                 :   0% (  0/ 14)
    GtkToolItem                   :   0% (  0/ 29)
    GtkToolItemGroup              :   0% (  0/ 17)
    GtkToolPalette                :   0% (  0/ 22)
    GtkToolShell                  :   0% (  0/  9)
    GtkToolbar                    :   0% (  0/ 34)
    GtkTooltip                    :   0% (  0/  9)
    GtkTreeIter                   : 100% (  1/  1)
    GtkTreeModel                  :  57% ( 15/ 26)
    GtkTreeModelFilter            :   0% (  0/ 11)
    GtkTreeModelSort              :   0% (  0/  9)
    GtkTreePath                   :  84% ( 16/ 19)
    GtkTreeRowReference           :   0% (  0/ 10)
    GtkTreeSelection              :  75% ( 15/ 20)
    GtkTreeSortable               :   0% (  0/  6)
    GtkTreeStore                  :  80% ( 17/ 21)
    GtkTreeView                   :  14% ( 14/ 97)
    GtkTreeView drag-and-drop     :   0% (  0/  7)
    GtkTreeViewColumn             :  37% ( 20/ 53)
    GtkUIManager                  :   0% (  0/ 17)
    GtkVBox                       : 100% (  1/  1)
    GtkVButtonBox                 :   0% (  0/  5)
    GtkVPaned                     : 100% (  1/  1)
    GtkVScale                     : 100% (  1/  1)
    GtkVScrollbar                 :   0% (  0/  1)
    GtkVSeparator                 : 100% (  1/  1)
    GtkViewport                   :   0% (  0/  9)
    GtkVolumeButton               :   0% (  0/  1)
    GtkWidget                     :  50% ( 90/180)
    GtkWindow                     :  23% ( 24/102)
    GtkWindowGroup                :   0% (  0/  5)
    Main Loop and Events          :  24% (  6/ 25)
    Resource Files                :   0% (  0/ 28)
    Selections                    :   0% (  0/ 46)
    Testing                       :   0% (  0/ 16)
    Themeable Stock Images        :   0% (  0/ 41)
    Version Information           :   0% (  0/  6)

    Total progress :                 29% (888/3046)
