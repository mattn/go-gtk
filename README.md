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

    Main Loop and Events          :  24% (  6/ 25)
    GtkAccelGroup                 :   5% (  1/ 19)
    GtkAccelMap                   :   0% (  0/ 14)
    GtkClipboard                  :  23% (  7/ 30)
    Drag and Drop                 :  11% (  4/ 35)
    GtkIconTheme                  :   0% (  0/ 31)
    GtkStockItem                  :  66% (  4/  6)
    Themeable Stock Images        :   0% (  0/ 41)
    Resource Files                :   0% (  0/ 28)
    GtkSettings                   :  23% (  3/ 13)
    GtkBinding                    :   0% (  0/ 14)
    Graphics Contexts             :   0% (  0/  2)
    GtkStyle                      :   0% (  0/ 64)
    Selections                    :   8% (  4/ 47)
    Version Information           :   0% (  0/  6)
    Testing                       :   0% (  0/ 16)
    Filesystem Utilities          :   0% (  0/  7)
    GtkDialog                     :  45% (  9/ 20)
    GtkMessageDialog              :  62% (  5/  8)
    GtkWindow                     :  23% ( 24/102)
    GtkWindowGroup                :   0% (  0/  5)
    GtkAboutDialog                :  90% ( 29/ 32)
    GtkAssistant                  :  91% ( 21/ 23)
    GtkOffscreenWindow            :   0% (  0/  3)
    GtkAccelLabel                 :  83% (  5/  6)
    GtkImage                      :  29% (  9/ 31)
    GtkLabel                      :  86% ( 39/ 45)
    GtkProgressBar                :  83% ( 10/ 12)
    GtkStatusbar                  :  77% (  7/  9)
    GtkInfoBar                    : 100% ( 12/ 12)
    GtkStatusIcon                 :  68% ( 26/ 38)
    GtkSpinner                    :   0% (  0/  3)
    GtkButton                     :  51% ( 14/ 27)
    GtkCheckButton                : 100% (  3/  3)
    GtkRadioButton                : 100% (  8/  8)
    GtkToggleButton               : 100% (  9/  9)
    GtkLinkButton                 :  75% (  6/  8)
    GtkScaleButton                :   0% (  0/  9)
    GtkVolumeButton               :   0% (  0/  1)
    GtkEntry                      :  36% ( 25/ 69)
    GtkEntryBuffer                :   0% (  0/ 11)
    GtkEntryCompletion            :   0% (  0/ 25)
    GtkHScale                     : 100% (  2/  2)
    GtkVScale                     : 100% (  2/  2)
    GtkSpinButton                 :   0% (  0/ 25)
    GtkEditable                   : 100% ( 13/ 13)
    GtkTextIter                   :  20% ( 19/ 91)
    GtkTextMark                   :   0% (  0/  7)
    GtkTextBuffer                 :  67% ( 52/ 77)
    GtkTextTag                    :   0% (  0/  9)
    GtkTextTagTable               :  83% (  5/  6)
    GtkTextView                   :  28% ( 18/ 64)
    GtkTreePath                   :  84% ( 16/ 19)
    GtkTreeRowReference           :   0% (  0/ 10)
    GtkTreeIter                   : 100% (  1/  1)
    GtkTreeModel                  :  57% ( 15/ 26)
    GtkTreeSelection              :  75% ( 15/ 20)
    GtkTreeViewColumn             :  37% ( 20/ 53)
    GtkTreeView                   :  14% ( 14/ 97)
    GtkTreeView drag-and-drop     :   0% (  0/  7)
    GtkCellView                   :   0% (  0/ 11)
    GtkIconView                   :   0% (  0/ 62)
    GtkTreeSortable               :   0% (  0/  6)
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
    GtkListStore                  :  83% ( 15/ 18)
    GtkTreeStore                  :  80% ( 17/ 21)
    GtkComboBox                   :  78% ( 30/ 38)
    GtkComboBoxText               : 100% (  7/  7)
    GtkComboBoxEntry              :  80% (  4/  5)
    GtkMenu                       :  50% ( 15/ 30)
    GtkMenuBar                    : 100% (  8/  8)
    GtkMenuItem                   :  78% ( 15/ 19)
    GtkImageMenuItem              :   0% (  0/ 11)
    GtkRadioMenuItem              :   0% (  0/  9)
    GtkCheckMenuItem              : 100% ( 10/ 10)
    GtkSeparatorMenuItem          : 100% (  1/  1)
    GtkTearoffMenuItem            :   0% (  0/  1)
    GtkToolShell                  :   0% (  0/  9)
    GtkToolbar                    :   0% (  0/ 34)
    GtkToolItem                   :   0% (  0/ 29)
    GtkToolPalette                :   0% (  0/ 22)
    GtkToolItemGroup              :   0% (  0/ 17)
    GtkSeparatorToolItem          :   0% (  0/  3)
    GtkToolButton                 :   0% (  0/ 14)
    GtkMenuToolButton             :   0% (  0/  7)
    GtkToggleToolButton           :   0% (  0/  4)
    GtkRadioToolButton            :   0% (  0/  6)
    GtkUIManager                  :   0% (  0/ 17)
    GtkActionGroup                :   0% (  0/ 20)
    GtkAction                     :   0% (  0/ 46)
    GtkToggleAction               :   0% (  0/  6)
    GtkRadioAction                :   0% (  0/  5)
    GtkRecentAction               :   0% (  0/  4)
    GtkActivatable                :   0% (  0/  6)
    GtkColorButton                :   0% (  0/ 10)
    GtkColorSelectionDialog       :   0% (  0/  2)
    GtkColorSelection             :   0% (  0/ 21)
    GtkHSV                        :   0% (  0/  8)
    GtkFileChooser                :  22% ( 13/ 58)
    GtkFileChooserButton          :   0% (  0/  9)
    GtkFileChooserDialog          : 100% (  1/  1)
    GtkFileChooserWidget          :   0% (  0/  2)
    GtkFileFilter                 :  55% (  5/  9)
    GtkFontButton                 :  71% ( 10/ 14)
    GtkFontSelection              :   0% (  0/ 14)
    GtkFontSelectionDialog        :  37% (  3/  8)
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
    GtkLayout                     :   0% (  0/ 12)
    GtkNotebook                   :  90% ( 50/ 55)
    GtkTable                      :  93% ( 14/ 15)
    GtkExpander                   :  87% ( 14/ 16)
    GtkOrientable                 :   0% (  0/  2)
    GtkFrame                      : 100% (  9/  9)
    GtkHSeparator                 : 100% (  1/  1)
    GtkVSeparator                 : 100% (  1/  1)
    GtkHScrollbar                 :   0% (  0/  1)
    GtkVScrollbar                 :   0% (  0/  1)
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
    GtkAdjustment                 :  82% ( 14/ 17)
    GtkArrow                      :   0% (  0/  2)
    GtkCalendar                   :   0% (  0/ 17)
    GtkDrawingArea                : 100% (  2/  2)
    GtkEventBox                   :  20% (  1/  5)
    GtkHandleBox                  :   0% (  0/  8)
    GtkIMContextSimple            :   0% (  0/  2)
    GtkIMMulticontext             :   0% (  0/  4)
    GtkSizeGroup                  : 100% (  8/  8)
    GtkTooltip                    :   0% (  0/  9)
    GtkViewport                   :   0% (  0/  9)
    GtkAccessible                 :   0% (  0/  3)
    GtkBin                        : 100% (  1/  1)
    GtkBox                        : 100% ( 11/ 11)
    GtkButtonBox                  :   0% (  0/ 10)
    GtkContainer                  :  18% (  6/ 33)
    GtkItem                       : 100% (  3/  3)
    GtkMenuShell                  :   0% (  0/ 11)
    GtkMisc                       :   0% (  0/  4)
    GtkObject                     : 100% (  0/  0)
    GtkPaned                      :  88% (  8/  9)
    GtkRange                      :  53% ( 16/ 30)
    GtkScale                      :  90% (  9/ 10)
    GtkSeparator                  : 100% (  0/  0)
    GtkWidget                     :  50% ( 90/180)
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

    Total progress :                 31% (951/3055)
