// +build !cgocheck

package gdk

// #include "gdk.go.h"
// #cgo pkg-config: gdk-2.0 gthread-2.0
import "C"
import (
	"unsafe"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/pango"
)

func guint16(v uint16) C.guint16 { return C.guint16(v) }
func gint(v int) C.gint          { return C.gint(v) }
func gstring(s *C.char) *C.gchar { return C.toGstr(s) }

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func cfree(s *C.char) { C.freeCstr(s) }

//-----------------------------------------------------------------------
// Types
//-----------------------------------------------------------------------
type Point struct {
	X int
	Y int
}

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

func Beep() {
	C.gdk_beep()
}

func Flush() {
	C.gdk_flush()
}

//-----------------------------------------------------------------------
// Threads
//-----------------------------------------------------------------------
func ThreadsInit() {
	C.gdk_threads_init()
}

func ThreadsEnter() {
	C.gdk_threads_enter()
}

func ThreadsLeave() {
	C.gdk_threads_leave()
}

//-----------------------------------------------------------------------
// GdkSelection
//-----------------------------------------------------------------------
const (
	SELECTION_PRIMARY   Atom = Atom(uintptr(1))
	SELECTION_SECONDARY Atom = Atom(uintptr(2))
	SELECTION_CLIPBOARD Atom = Atom(uintptr(69))
)

//-----------------------------------------------------------------------
// GdkCursor
//-----------------------------------------------------------------------
type CursorType int

const (
	X_CURSOR            CursorType = 0
	ARROW               CursorType = 2
	BASED_ARROW_DOWN    CursorType = 4
	BASED_ARROW_UP      CursorType = 6
	BOAT                CursorType = 8
	BOGOSITY            CursorType = 10
	BOTTOM_LEFT_CORNER  CursorType = 12
	BOTTOM_RIGHT_CORNER CursorType = 14
	BOTTOM_SIDE         CursorType = 16
	BOTTOM_TEE          CursorType = 18
	BOX_SPIRAL          CursorType = 20
	CENTER_PTR          CursorType = 22
	CIRCLE              CursorType = 24
	CLOCK               CursorType = 26
	COFFEE_MUG          CursorType = 28
	CROSS               CursorType = 30
	CROSS_REVERSE       CursorType = 32
	CROSSHAIR           CursorType = 34
	DIAMOND_CROSS       CursorType = 36
	DOT                 CursorType = 38
	DOTBOX              CursorType = 40
	DOUBLE_ARROW        CursorType = 42
	DRAFT_LARGE         CursorType = 44
	DRAFT_SMALL         CursorType = 46
	DRAPED_BOX          CursorType = 48
	EXCHANGE            CursorType = 50
	FLEUR               CursorType = 52
	GOBBLER             CursorType = 54
	GUMBY               CursorType = 56
	HAND1               CursorType = 58
	HAND2               CursorType = 60
	HEART               CursorType = 62
	ICON                CursorType = 64
	IRON_CROSS          CursorType = 66
	LEFT_PTR            CursorType = 68
	LEFT_SIDE           CursorType = 70
	LEFT_TEE            CursorType = 72
	LEFTBUTTON          CursorType = 74
	LL_ANGLE            CursorType = 76
	LR_ANGLE            CursorType = 78
	MAN                 CursorType = 80
	MIDDLEBUTTON        CursorType = 82
	MOUSE               CursorType = 84
	PENCIL              CursorType = 86
	PIRATE              CursorType = 88
	PLUS                CursorType = 90
	QUESTION_ARROW      CursorType = 92
	RIGHT_PTR           CursorType = 94
	RIGHT_SIDE          CursorType = 96
	RIGHT_TEE           CursorType = 98
	RIGHTBUTTON         CursorType = 100
	RTL_LOGO            CursorType = 102
	SAILBOAT            CursorType = 104
	SB_DOWN_ARROW       CursorType = 106
	SB_H_DOUBLE_ARROW   CursorType = 108
	SB_LEFT_ARROW       CursorType = 110
	SB_RIGHT_ARROW      CursorType = 112
	SB_UP_ARROW         CursorType = 114
	SB_V_DOUBLE_ARROW   CursorType = 116
	SHUTTLE             CursorType = 118
	SIZING              CursorType = 120
	SPIDER              CursorType = 122
	SPRAYCAN            CursorType = 124
	STAR                CursorType = 126
	TARGET              CursorType = 128
	TCROSS              CursorType = 130
	TOP_LEFT_ARROW      CursorType = 132
	TOP_LEFT_CORNER     CursorType = 134
	TOP_RIGHT_CORNER    CursorType = 136
	TOP_SIDE            CursorType = 138
	TOP_TEE             CursorType = 140
	TREK                CursorType = 142
	UL_ANGLE            CursorType = 144
	UMBRELLA            CursorType = 146
	UR_ANGLE            CursorType = 148
	WATCH               CursorType = 150
	XTERM               CursorType = 152
	LAST_CURSOR         CursorType = 153
	BLANK_CURSOR        CursorType = -2
	CURSOR_IS_PIXMAP    CursorType = -1
)

type Cursor struct {
	GCursor *C.GdkCursor
}

func NewCursor(cursor_type CursorType) *Cursor {
	return &Cursor{C.gdk_cursor_new(C.GdkCursorType(cursor_type))}
}

//-----------------------------------------------------------------------
// GdkColor
//-----------------------------------------------------------------------
type Color struct {
	GColor C.GdkColor
}

func NewColor(name string) *Color {
	var color C.GdkColor
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gdk_color_parse(gstring(ptr), &color)
	return &Color{color}
}

func NewColorRGB(r, g, b uint16) *Color {
	color := C.GdkColor{red: C.guint16(r), green: C.guint16(g), blue: C.guint16(b)}
	return &Color{color}
}

func (v Color) Red() uint16   { return uint16(v.GColor.red) }
func (v Color) Green() uint16 { return uint16(v.GColor.green) }
func (v Color) Blue() uint16  { return uint16(v.GColor.blue) }

//-----------------------------------------------------------------------
// GdkColormap
// ----------------------------------------------------------------------
type Colormap struct {
	GColormap *C.GdkColormap
}

func ColormapFromUnsafe(colormap unsafe.Pointer) *Colormap {
	return &Colormap{(*C.GdkColormap)(colormap)}
}

func (v *Colormap) AllocColorRGB(r, g, b uint16) *Color {
	req := &Color{GColor: C.GdkColor{pixel: 0, red: guint16(r), green: guint16(g), blue: guint16(b)}}
	C.gdk_colormap_alloc_color(v.GColormap, &req.GColor, gbool(false), gbool(true))
	// XXX fixme -- should check for failure
	return req
}

//-----------------------------------------------------------------------
// GdkFont
//-----------------------------------------------------------------------
type Font struct {
	GFont *C.GdkFont
}

func FontFromUnsafe(window unsafe.Pointer) *Font {
	return &Font{C.toGdkFont(window)}
}

func FontLoad(name string) *Font {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &Font{C.gdk_font_load(gstring(ptr))}
}

func FontsetLoad(name string) *Font {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &Font{C.gdk_fontset_load(gstring(ptr))}
}

// GdkFont* gdk_font_ref (GdkFont *font);
// void gdk_font_unref (GdkFont *font);
// gint gdk_font_id (const GdkFont *font);
// gboolean gdk_font_equal (const GdkFont *fonta, const GdkFont *fontb);
// GdkFont *gdk_font_load_for_display (GdkDisplay *display, const gchar *font_name);
// GdkFont *gdk_fontset_load_for_display (GdkDisplay *display, const gchar *fontset_name);
// GdkFont *gdk_font_from_description_for_display (GdkDisplay *display, PangoFontDescription *font_desc);
// GdkFont* gdk_font_load (const gchar *font_name);
// GdkFont* gdk_fontset_load (const gchar *fontset_name);
// GdkFont* gdk_font_from_description (PangoFontDescription *font_desc);
// gint gdk_string_width (GdkFont *font, const gchar *string);
// gint gdk_text_width (GdkFont *font, const gchar *text, gint text_length);
// gint gdk_text_width_wc (GdkFont *font, const GdkWChar *text, gint text_length);
// gint gdk_char_width (GdkFont *font, gchar character);
// gint gdk_char_width_wc (GdkFont *font, GdkWChar character);
// gint gdk_string_measure (GdkFont *font, const gchar *string);
// gint gdk_text_measure (GdkFont *font, const gchar *text, gint text_length);
// gint gdk_char_measure (GdkFont *font, gchar character);
// gint gdk_string_height (GdkFont *font, const gchar *string);
// gint gdk_text_height (GdkFont *font, const gchar *text, gint text_length);
// gint gdk_char_height (GdkFont *font, gchar character);
// void gdk_text_extents (GdkFont *font, const gchar *text, gint text_length, gint *lbearing, gint *rbearing, gint *width, gint *ascent, gint *descent);
// void gdk_text_extents_wc (GdkFont *font, const GdkWChar *text, gint text_length, gint *lbearing, gint *rbearing, gint *width, gint *ascent, gint *descent);
// void gdk_string_extents (GdkFont *font, const gchar *string, gint *lbearing, gint *rbearing, gint *width, gint *ascent, gint *descent);
// GdkDisplay * gdk_font_get_display (GdkFont *font);

//-----------------------------------------------------------------------
// GdkGC
//-----------------------------------------------------------------------
type GC struct {
	GGC *C.GdkGC
}

func NewGC(drawable *Drawable) *GC {
	return &GC{C.gdk_gc_new(drawable.GDrawable)}
}

// GdkGC *gdk_gc_new_with_values (GdkDrawable *drawable, GdkGCValues *values, GdkGCValuesMask values_mask);
// GdkGC *gdk_gc_ref (GdkGC *gc);
// void gdk_gc_unref (GdkGC *gc);
// void gdk_gc_get_values (GdkGC *gc, GdkGCValues *values);
// void gdk_gc_set_values (GdkGC *gc, GdkGCValues *values, GdkGCValuesMask values_mask);
func (v *GC) SetForeground(color *Color) {
	C.gdk_gc_set_foreground(v.GGC, &color.GColor)
}

func (v *GC) SetBackground(color *Color) {
	C.gdk_gc_set_background(v.GGC, &color.GColor)
}

// void gdk_gc_set_font (GdkGC *gc, GdkFont *font);
// void gdk_gc_set_function (GdkGC *gc, GdkFunction function);
// void gdk_gc_set_fill (GdkGC *gc, GdkFill fill);
// void gdk_gc_set_tile (GdkGC *gc, GdkPixmap *tile);
// void gdk_gc_set_stipple (GdkGC *gc, GdkPixmap *stipple);
// void gdk_gc_set_ts_origin (GdkGC *gc, gint x, gint y);
// void gdk_gc_set_clip_origin (GdkGC *gc, gint x, gint y);
// void gdk_gc_set_clip_mask (GdkGC *gc, GdkBitmap *mask);
func (v *GC) SetClipRectangle(x, y, width, height int) {
	var r C.GdkRectangle
	r.x = C.gint(x)
	r.y = C.gint(y)
	r.width = C.gint(width)
	r.height = C.gint(height)
	C.gdk_gc_set_clip_rectangle(v.GGC, &r)
}

// void gdk_gc_set_clip_region (GdkGC *gc, const GdkRegion *region);
// void gdk_gc_set_subwindow (GdkGC *gc, GdkSubwindowMode mode);
// void gdk_gc_set_exposures (GdkGC *gc, gboolean exposures);
// void gdk_gc_set_line_attributes (GdkGC *gc, gint line_width, GdkLineStyle line_style, GdkCapStyle cap_style, GdkJoinStyle join_style);
// void gdk_gc_set_dashes (GdkGC *gc, gint dash_offset, gint8 dash_list[], gint n);
// void gdk_gc_offset (GdkGC *gc, gint x_offset, gint y_offset);
// void gdk_gc_copy (GdkGC *dst_gc, GdkGC *src_gc);
// void gdk_gc_set_colormap (GdkGC *gc, GdkColormap *colormap);
// GdkColormap *gdk_gc_get_colormap (GdkGC *gc);
func (v *GC) GetColormap() *Colormap {
	return &Colormap{C.gdk_gc_get_colormap(v.GGC)}
}

func (v *GC) SetRgbFgColor(color *Color) {
	C.gdk_gc_set_rgb_fg_color(v.GGC, &color.GColor)
}

func (v *GC) SetRgbBgColor(color *Color) {
	C.gdk_gc_set_rgb_bg_color(v.GGC, &color.GColor)
}

//-----------------------------------------------------------------------
// GdkScreen
//-----------------------------------------------------------------------
type Screen struct {
	GScreen *C.GdkScreen
}

// GdkScreen * gdk_gc_get_screen (GdkGC *gc);

func ScreenFromUnsafe(screen unsafe.Pointer) *Screen {
	return &Screen{(*C.GdkScreen)(screen)}
}

func GetDefaultScreen() (screen *Screen) {
	return &Screen{GScreen: C.gdk_screen_get_default()}
}

func ScreenWidth() (width int) {
	return int(C.gdk_screen_width())
}

func ScreenHeight() (height int) {
	return int(C.gdk_screen_height())
}

//-----------------------------------------------------------------------
// GdkDrawable
//-----------------------------------------------------------------------
type Drawable struct {
	GDrawable *C.GdkDrawable
}

func (v *Drawable) DrawPoint(gc *GC, x int, y int) {
	C.gdk_draw_point(v.GDrawable, gc.GGC, gint(x), gint(y))
}

func (v *Drawable) DrawLine(gc *GC, x1 int, y1 int, x2 int, y2 int) {
	C.gdk_draw_line(v.GDrawable, gc.GGC, gint(x1), gint(y1), gint(x2), gint(y2))
}

func (v *Drawable) DrawRectangle(gc *GC, filled bool, x int, y int, width int, height int) {
	C.gdk_draw_rectangle(v.GDrawable, gc.GGC, gbool(filled), gint(x), gint(y), gint(width), gint(height))
}

func (v *Drawable) DrawArc(gc *GC, filled bool, x int, y int, width int, height int, angle1 int, angle2 int) {
	C.gdk_draw_arc(v.GDrawable, gc.GGC, gbool(filled), gint(x), gint(y), gint(width), gint(height), gint(angle1), gint(angle2))
}

// void gdk_draw_polygon (GdkDrawable *drawable, GdkGC *gc, gboolean filled, const GdkPoint *points, gint n_points);
// void gdk_draw_string (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *string);
func (v *Drawable) DrawString(font *Font, gc *GC, x int, y int, str string) {
	ptr := C.CString(str)
	defer cfree(ptr)
	C.gdk_draw_string(v.GDrawable, font.GFont, gc.GGC, gint(x), gint(y), gstring(ptr))
}

// void gdk_draw_text (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *text, gint text_length);
// void gdk_draw_text_wc (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const GdkWChar *text, gint text_length);
func (v *Drawable) DrawDrawable(gc *GC, src *Drawable, xsrc int, ysrc int, xdest int, ydest int, width int, height int) {
	C.gdk_draw_drawable(v.GDrawable, gc.GGC, src.GDrawable, gint(xsrc), gint(ysrc), gint(xdest), gint(ydest), gint(width), gint(height))
}

// void gdk_draw_image (GdkDrawable *drawable, GdkGC *gc, GdkImage *image, gint xsrc, gint ysrc, gint xdest, gint ydest, gint width, gint height);
// void gdk_draw_points (GdkDrawable *drawable, GdkGC *gc, const GdkPoint *points, gint n_points);
// void gdk_draw_segments (GdkDrawable *drawable, GdkGC *gc, const GdkSegment *segs, gint n_segs);
// void gdk_draw_lines (GdkDrawable *drawable, GdkGC *gc, const GdkPoint *points, gint n_points);

type GdkRgbDither int

const (
	//Never use dithering.
	RGB_DITHER_NONE GdkRgbDither = C.GDK_RGB_DITHER_NONE

	// Use dithering in 8 bits per pixel (and below) only.
	RGB_DITHER_NORMAL GdkRgbDither = C.GDK_RGB_DITHER_NORMAL

	// Use dithering in 16 bits per pixel and below.
	RGB_DITHER_MAX GdkRgbDither = C.GDK_RGB_DITHER_MAX
)

// void gdk_draw_pixbuf (GdkDrawable *drawable, GdkGC *gc, const GdkPixbuf *pixbuf,
//                       gint src_x, gint src_y, gint dest_x, gint dest_y, gint width, gint height,
//                       GdkRgbDither dither, gint x_dither, gint y_dither);
func (v *Drawable) DrawPixbuf(gc *GC, pixbuf *gdkpixbuf.Pixbuf,
	src_x, src_y, dest_x, dest_y, width, height int,
	dither GdkRgbDither, x_dither, y_dither int,
) {
	C.gdk_draw_pixbuf(v.GDrawable, gc.GGC, (*C.struct__GdkPixbuf)(pixbuf.GdkPixbuf.GPixbuf),
		gint(src_x), gint(src_y), gint(dest_x), gint(dest_y), gint(width), gint(height),
		C.GdkRgbDither(dither), gint(x_dither), gint(y_dither))
}

// void gdk_draw_glyphs (GdkDrawable *drawable, GdkGC *gc, PangoFont *font, gint x, gint y, PangoGlyphString *glyphs);
// void gdk_draw_layout_line (GdkDrawable *drawable, GdkGC *gc, gint x, gint y, PangoLayoutLine *line);
// void gdk_draw_layout (GdkDrawable *drawable, GdkGC *gc, gint x, gint y, PangoLayout *layout);
func (v *Drawable) DrawLayout(gc *GC, x, y int, layout *pango.Layout) {
	C.gdk_draw_layout(v.GDrawable, gc.GGC, C.gint(x), C.gint(y), (*C.PangoLayout)(layout.GLayout))
}

// void gdk_draw_layout_line_with_colors (GdkDrawable *drawable, GdkGC *gc, gint x, gint y, PangoLayoutLine *line, const GdkColor *foreground, const GdkColor *background);
// void gdk_draw_layout_with_colors (GdkDrawable *drawable, GdkGC *gc, gint x, gint y, PangoLayout *layout, const GdkColor *foreground, const GdkColor *background);
// void gdk_draw_glyphs_transformed (GdkDrawable *drawable, GdkGC	 *gc, const PangoMatrix *matrix, PangoFont *font, gint x, gint y, PangoGlyphString *glyphs);
// void gdk_draw_trapezoids (GdkDrawable *drawable, GdkGC	 *gc, const GdkTrapezoid *trapezoids, gint n_trapezoids);
// #define gdk_draw_pixmap gdk_draw_drawable
// #define gdk_draw_bitmap gdk_draw_drawable
// GdkImage* gdk_drawable_get_image (GdkDrawable *drawable, gint x, gint y, gint width, gint height);
// GdkImage *gdk_drawable_copy_to_image (GdkDrawable *drawable, GdkImage *image, gint src_x, gint src_y, gint dest_x, gint dest_y, gint width, gint height);
// GdkRegion *gdk_drawable_get_clip_region (GdkDrawable *drawable);
// GdkRegion *gdk_drawable_get_visible_region (GdkDrawable *drawable);

//-----------------------------------------------------------------------
// GdkWindow
//-----------------------------------------------------------------------
type ModifierType int

const (
	SHIFT_MASK   ModifierType = 1 << 0
	LOCK_MASK    ModifierType = 1 << 1
	CONTROL_MASK ModifierType = 1 << 2
	MOD1_MASK    ModifierType = 1 << 3
	MOD2_MASK    ModifierType = 1 << 4
	MOD3_MASK    ModifierType = 1 << 5
	MOD4_MASK    ModifierType = 1 << 6
	MOD5_MASK    ModifierType = 1 << 7
	BUTTON1_MASK ModifierType = 1 << 8
	BUTTON2_MASK ModifierType = 1 << 9
	BUTTON3_MASK ModifierType = 1 << 10
	BUTTON4_MASK ModifierType = 1 << 11
	BUTTON5_MASK ModifierType = 1 << 12

	/* The next few modifiers are used by XKB, so we skip to the end.
	 * Bits 15 - 25 are currently unused. Bit 29 is used internally.
	 */

	SUPER_MASK ModifierType = 1 << 26
	HYPER_MASK ModifierType = 1 << 27
	META_MASK  ModifierType = 1 << 28

	RELEASE_MASK ModifierType = 1 << 30

	MODIFIER_MASK ModifierType = 0x5c001fff
)

type EventType int

const (
	NOTHING           EventType = -1
	DELETE            EventType = 0
	DESTROY           EventType = 1
	EXPOSE            EventType = 2
	MOTION_NOTIFY     EventType = 3
	BUTTON_PRESS      EventType = 4
	BUTTON2_PRESS     EventType = 5
	BUTTON3_PRESS     EventType = 6
	BUTTON_RELEASE    EventType = 7
	KEY_PRESS         EventType = 8
	KEY_RELEASE       EventType = 9
	ENTER_NOTIFY      EventType = 10
	LEAVE_NOTIFY      EventType = 11
	FOCUS_CHANGE      EventType = 12
	CONFIGURE         EventType = 13
	MAP               EventType = 14
	UNMAP             EventType = 15
	PROPERTY_NOTIFY   EventType = 16
	SELECTION_CLEAR   EventType = 17
	SELECTION_REQUEST EventType = 18
	SELECTION_NOTIFY  EventType = 19
	PROXIMITY_IN      EventType = 20
	PROXIMITY_OUT     EventType = 21
	DRAG_ENTER        EventType = 22
	DRAG_LEAVE        EventType = 23
	DRAG_MOTION       EventType = 24
	DRAG_STATUS       EventType = 25
	DROP_START        EventType = 26
	DROP_FINISHED     EventType = 27
	CLIENT_EVENT      EventType = 28
	VISIBILITY_NOTIFY EventType = 29
	NO_EXPOSE         EventType = 30
	SCROLL            EventType = 31
	WINDOW_STATE      EventType = 32
	SETTING           EventType = 33
	OWNER_CHANGE      EventType = 34
	GRAB_BROKEN       EventType = 35
	DAMAGE            EventType = 36
	EVENT_LAST        EventType = 37 /* helper variable for decls */
)

type EventMask int

const (
	EXPOSURE_MASK            EventMask = 1 << 1
	POINTER_MOTION_MASK      EventMask = 1 << 2
	POINTER_MOTION_HINT_MASK EventMask = 1 << 3
	BUTTON_MOTION_MASK       EventMask = 1 << 4
	BUTTON1_MOTION_MASK      EventMask = 1 << 5
	BUTTON2_MOTION_MASK      EventMask = 1 << 6
	BUTTON3_MOTION_MASK      EventMask = 1 << 7
	BUTTON_PRESS_MASK        EventMask = 1 << 8
	BUTTON_RELEASE_MASK      EventMask = 1 << 9
	KEY_PRESS_MASK           EventMask = 1 << 10
	KEY_RELEASE_MASK         EventMask = 1 << 11
	ENTER_NOTIFY_MASK        EventMask = 1 << 12
	LEAVE_NOTIFY_MASK        EventMask = 1 << 13
	FOCUS_CHANGE_MASK        EventMask = 1 << 14
	STRUCTURE_MASK           EventMask = 1 << 15
	PROPERTY_CHANGE_MASK     EventMask = 1 << 16
	VISIBILITY_NOTIFY_MASK   EventMask = 1 << 17
	PROXIMITY_IN_MASK        EventMask = 1 << 18
	PROXIMITY_OUT_MASK       EventMask = 1 << 19
	SUBSTRUCTURE_MASK        EventMask = 1 << 20
	SCROLL_MASK              EventMask = 1 << 21
	ALL_EVENTS_MASK          EventMask = 0x3FFFFE
)

type DragAction int

const (
	ACTION_DEFAULT DragAction = 1 << 0
	ACTION_COPY    DragAction = 1 << 1
	ACTION_MOVE    DragAction = 1 << 2
	ACTION_LINK    DragAction = 1 << 3
	ACTION_PRIVATE DragAction = 1 << 4
	ACTION_ASK     DragAction = 1 << 5
)

type WindowTypeHint int

const (
	WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = 0
	WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = 1
	WINDOW_TYPE_HINT_MENU          WindowTypeHint = 2
	WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = 3
	WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = 4
	WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = 5
	WINDOW_TYPE_HINT_DOCK          WindowTypeHint = 6
	WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = 7
	WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = 8
	WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = 9
	WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = 10
	WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = 11
	WINDOW_TYPE_HINT_COMBO         WindowTypeHint = 12
	WINDOW_TYPE_HINT_DND           WindowTypeHint = 13
)

type Gravity int

const (
	GRAVITY_NORTH_WEST Gravity = 1
	GRAVITY_NORTH      Gravity = 2
	GRAVITY_NORTH_EAST Gravity = 3
	GRAVITY_WEST       Gravity = 4
	GRAVITY_CENTER     Gravity = 5
	GRAVITY_EAST       Gravity = 6
	GRAVITY_SOUTH_WEST Gravity = 7
	GRAVITY_SOUTH      Gravity = 8
	GRAVITY_SOUTH_EAST Gravity = 9
	GRAVITY_STATIC     Gravity = 10
)

type ScrollDirection int

const (
	SCROLL_UP    ScrollDirection = 0
	SCROLL_Down  ScrollDirection = 1
	SCROLL_LEFT  ScrollDirection = 2
	SCROLL_RIGHT ScrollDirection = 3
)

type Window struct {
	GWindow *C.GdkWindow
}

func WindowFromUnsafe(window unsafe.Pointer) *Window {
	return &Window{C.toGdkWindow(window)}
}

func (v *Window) GetPointer(x *int, y *int, mask *ModifierType) *Window {
	var cx, cy C.gint
	var mt C.GdkModifierType
	ret := &Window{C.gdk_window_get_pointer(v.GWindow, &cx, &cy, &mt)}
	*x = int(cx)
	*y = int(cy)
	*mask = ModifierType(mt)
	return ret
}

func (v *Window) GetDrawable() *Drawable {
	return &Drawable{
		(*C.GdkDrawable)(v.GWindow)}
}

func (v *Window) Invalidate(rect *Rectangle, invalidate_children bool) {
	if rect != nil {
		var _rect C.GdkRectangle
		_rect.x = gint(rect.X)
		_rect.y = gint(rect.Y)
		_rect.width = gint(rect.Width)
		_rect.height = gint(rect.Height)
		C.gdk_window_invalidate_rect(v.GWindow, &_rect, gbool(invalidate_children))
	} else {
		C.gdk_window_invalidate_rect(v.GWindow, nil, gbool(invalidate_children))
	}
}

func (v *Window) Show() {
	C.gdk_window_show(v.GWindow)
}

func (v *Window) Raise() {
	C.gdk_window_raise(v.GWindow)
}

//-----------------------------------------------------------------------
// GdkPixmap
//-----------------------------------------------------------------------
type Pixmap struct {
	GPixmap *C.GdkPixmap
}

func NewPixmap(drawable *Drawable, width int, height int, depth int) *Pixmap {
	return &Pixmap{C.gdk_pixmap_new(drawable.GDrawable, gint(width), gint(height), gint(depth))}
}

// GdkBitmap* gdk_bitmap_create_from_data (GdkDrawable *drawable, const gchar *data, gint width, gint height);
// GdkPixmap* gdk_pixmap_create_from_data (GdkDrawable *drawable, const gchar *data, gint width, gint height, gint depth, const GdkColor *fg, const GdkColor *bg);
// GdkPixmap* gdk_pixmap_create_from_xpm (GdkDrawable *drawable, GdkBitmap **mask, const GdkColor *transparent_color, const gchar *filename);
// GdkPixmap* gdk_pixmap_colormap_create_from_xpm (GdkDrawable *drawable, GdkColormap *colormap, GdkBitmap **mask, const GdkColor *transparent_color, const gchar *filename);
// GdkPixmap* gdk_pixmap_create_from_xpm_d (GdkDrawable *drawable, GdkBitmap **mask, const GdkColor *transparent_color, gchar **data);
// GdkPixmap* gdk_pixmap_colormap_create_from_xpm_d (GdkDrawable *drawable, GdkColormap *colormap, GdkBitmap **mask, const GdkColor *transparent_color, gchar **data);
// GdkPixmap* gdk_pixmap_foreign_new (GdkNativeWindow anid);
// GdkPixmap* gdk_pixmap_lookup (GdkNativeWindow anid);
// GdkPixmap* gdk_pixmap_foreign_new_for_display (GdkDisplay *display, GdkNativeWindow anid);
// GdkPixmap* gdk_pixmap_lookup_for_display (GdkDisplay *display, GdkNativeWindow anid);
// GdkPixmap* gdk_pixmap_foreign_new_for_screen (GdkScreen *screen, GdkNativeWindow anid, gint width, gint height, gint depth);
func (v *Pixmap) Ref() {
	C.g_object_ref(C.gpointer(v.GPixmap))
}

func (v *Pixmap) Unref() {
	C.g_object_unref(C.gpointer(v.GPixmap))
}

func (v *Pixmap) GetDrawable() *Drawable {
	return &Drawable{(*C.GdkDrawable)(v.GPixmap)}
}

// Subset of gdkkeysyms.h
const (
	KEY_VoidSymbol                  = 0xffffff
	KEY_BackSpace                   = 0xff08
	KEY_Tab                         = 0xff09
	KEY_Linefeed                    = 0xff0a
	KEY_Clear                       = 0xff0b
	KEY_Return                      = 0xff0d
	KEY_Pause                       = 0xff13
	KEY_Scroll_Lock                 = 0xff14
	KEY_Sys_Req                     = 0xff15
	KEY_Escape                      = 0xff1b
	KEY_Delete                      = 0xffff
	KEY_Multi_key                   = 0xff20
	KEY_Codeinput                   = 0xff37
	KEY_SingleCandidate             = 0xff3c
	KEY_MultipleCandidate           = 0xff3d
	KEY_PreviousCandidate           = 0xff3e
	KEY_Kanji                       = 0xff21
	KEY_Muhenkan                    = 0xff22
	KEY_Henkan_Mode                 = 0xff23
	KEY_Henkan                      = 0xff23
	KEY_Romaji                      = 0xff24
	KEY_Hiragana                    = 0xff25
	KEY_Katakana                    = 0xff26
	KEY_Hiragana_Katakana           = 0xff27
	KEY_Zenkaku                     = 0xff28
	KEY_Hankaku                     = 0xff29
	KEY_Zenkaku_Hankaku             = 0xff2a
	KEY_Touroku                     = 0xff2b
	KEY_Massyo                      = 0xff2c
	KEY_Kana_Lock                   = 0xff2d
	KEY_Kana_Shift                  = 0xff2e
	KEY_Eisu_Shift                  = 0xff2f
	KEY_Eisu_toggle                 = 0xff30
	KEY_Kanji_Bangou                = 0xff37
	KEY_Zen_Koho                    = 0xff3d
	KEY_Mae_Koho                    = 0xff3e
	KEY_Home                        = 0xff50
	KEY_Left                        = 0xff51
	KEY_Up                          = 0xff52
	KEY_Right                       = 0xff53
	KEY_Down                        = 0xff54
	KEY_Prior                       = 0xff55
	KEY_Page_Up                     = 0xff55
	KEY_Next                        = 0xff56
	KEY_Page_Down                   = 0xff56
	KEY_End                         = 0xff57
	KEY_Begin                       = 0xff58
	KEY_Select                      = 0xff60
	KEY_Print                       = 0xff61
	KEY_Execute                     = 0xff62
	KEY_Insert                      = 0xff63
	KEY_Undo                        = 0xff65
	KEY_Redo                        = 0xff66
	KEY_Menu                        = 0xff67
	KEY_Find                        = 0xff68
	KEY_Cancel                      = 0xff69
	KEY_Help                        = 0xff6a
	KEY_Break                       = 0xff6b
	KEY_Mode_switch                 = 0xff7e
	KEY_script_switch               = 0xff7e
	KEY_Num_Lock                    = 0xff7f
	KEY_KP_Space                    = 0xff80
	KEY_KP_Tab                      = 0xff89
	KEY_KP_Enter                    = 0xff8d
	KEY_KP_F1                       = 0xff91
	KEY_KP_F2                       = 0xff92
	KEY_KP_F3                       = 0xff93
	KEY_KP_F4                       = 0xff94
	KEY_KP_Home                     = 0xff95
	KEY_KP_Left                     = 0xff96
	KEY_KP_Up                       = 0xff97
	KEY_KP_Right                    = 0xff98
	KEY_KP_Down                     = 0xff99
	KEY_KP_Prior                    = 0xff9a
	KEY_KP_Page_Up                  = 0xff9a
	KEY_KP_Next                     = 0xff9b
	KEY_KP_Page_Down                = 0xff9b
	KEY_KP_End                      = 0xff9c
	KEY_KP_Begin                    = 0xff9d
	KEY_KP_Insert                   = 0xff9e
	KEY_KP_Delete                   = 0xff9f
	KEY_KP_Equal                    = 0xffbd
	KEY_KP_Multiply                 = 0xffaa
	KEY_KP_Add                      = 0xffab
	KEY_KP_Separator                = 0xffac
	KEY_KP_Subtract                 = 0xffad
	KEY_KP_Decimal                  = 0xffae
	KEY_KP_Divide                   = 0xffaf
	KEY_KP_0                        = 0xffb0
	KEY_KP_1                        = 0xffb1
	KEY_KP_2                        = 0xffb2
	KEY_KP_3                        = 0xffb3
	KEY_KP_4                        = 0xffb4
	KEY_KP_5                        = 0xffb5
	KEY_KP_6                        = 0xffb6
	KEY_KP_7                        = 0xffb7
	KEY_KP_8                        = 0xffb8
	KEY_KP_9                        = 0xffb9
	KEY_F1                          = 0xffbe
	KEY_F2                          = 0xffbf
	KEY_F3                          = 0xffc0
	KEY_F4                          = 0xffc1
	KEY_F5                          = 0xffc2
	KEY_F6                          = 0xffc3
	KEY_F7                          = 0xffc4
	KEY_F8                          = 0xffc5
	KEY_F9                          = 0xffc6
	KEY_F10                         = 0xffc7
	KEY_F11                         = 0xffc8
	KEY_L1                          = 0xffc8
	KEY_F12                         = 0xffc9
	KEY_L2                          = 0xffc9
	KEY_F13                         = 0xffca
	KEY_L3                          = 0xffca
	KEY_F14                         = 0xffcb
	KEY_L4                          = 0xffcb
	KEY_F15                         = 0xffcc
	KEY_L5                          = 0xffcc
	KEY_F16                         = 0xffcd
	KEY_L6                          = 0xffcd
	KEY_F17                         = 0xffce
	KEY_L7                          = 0xffce
	KEY_F18                         = 0xffcf
	KEY_L8                          = 0xffcf
	KEY_F19                         = 0xffd0
	KEY_L9                          = 0xffd0
	KEY_F20                         = 0xffd1
	KEY_L10                         = 0xffd1
	KEY_F21                         = 0xffd2
	KEY_R1                          = 0xffd2
	KEY_F22                         = 0xffd3
	KEY_R2                          = 0xffd3
	KEY_F23                         = 0xffd4
	KEY_R3                          = 0xffd4
	KEY_F24                         = 0xffd5
	KEY_R4                          = 0xffd5
	KEY_F25                         = 0xffd6
	KEY_R5                          = 0xffd6
	KEY_F26                         = 0xffd7
	KEY_R6                          = 0xffd7
	KEY_F27                         = 0xffd8
	KEY_R7                          = 0xffd8
	KEY_F28                         = 0xffd9
	KEY_R8                          = 0xffd9
	KEY_F29                         = 0xffda
	KEY_R9                          = 0xffda
	KEY_F30                         = 0xffdb
	KEY_R10                         = 0xffdb
	KEY_F31                         = 0xffdc
	KEY_R11                         = 0xffdc
	KEY_F32                         = 0xffdd
	KEY_R12                         = 0xffdd
	KEY_F33                         = 0xffde
	KEY_R13                         = 0xffde
	KEY_F34                         = 0xffdf
	KEY_R14                         = 0xffdf
	KEY_F35                         = 0xffe0
	KEY_R15                         = 0xffe0
	KEY_Shift_L                     = 0xffe1
	KEY_Shift_R                     = 0xffe2
	KEY_Control_L                   = 0xffe3
	KEY_Control_R                   = 0xffe4
	KEY_Caps_Lock                   = 0xffe5
	KEY_Shift_Lock                  = 0xffe6
	KEY_Meta_L                      = 0xffe7
	KEY_Meta_R                      = 0xffe8
	KEY_Alt_L                       = 0xffe9
	KEY_Alt_R                       = 0xffea
	KEY_Super_L                     = 0xffeb
	KEY_Super_R                     = 0xffec
	KEY_Hyper_L                     = 0xffed
	KEY_Hyper_R                     = 0xffee
	KEY_ISO_Lock                    = 0xfe01
	KEY_ISO_Level2_Latch            = 0xfe02
	KEY_ISO_Level3_Shift            = 0xfe03
	KEY_ISO_Level3_Latch            = 0xfe04
	KEY_ISO_Level3_Lock             = 0xfe05
	KEY_ISO_Level5_Shift            = 0xfe11
	KEY_ISO_Level5_Latch            = 0xfe12
	KEY_ISO_Level5_Lock             = 0xfe13
	KEY_ISO_Group_Shift             = 0xff7e
	KEY_ISO_Group_Latch             = 0xfe06
	KEY_ISO_Group_Lock              = 0xfe07
	KEY_ISO_Next_Group              = 0xfe08
	KEY_ISO_Next_Group_Lock         = 0xfe09
	KEY_ISO_Prev_Group              = 0xfe0a
	KEY_ISO_Prev_Group_Lock         = 0xfe0b
	KEY_ISO_First_Group             = 0xfe0c
	KEY_ISO_First_Group_Lock        = 0xfe0d
	KEY_ISO_Last_Group              = 0xfe0e
	KEY_ISO_Last_Group_Lock         = 0xfe0f
	KEY_ISO_Left_Tab                = 0xfe20
	KEY_ISO_Move_Line_Up            = 0xfe21
	KEY_ISO_Move_Line_Down          = 0xfe22
	KEY_ISO_Partial_Line_Up         = 0xfe23
	KEY_ISO_Partial_Line_Down       = 0xfe24
	KEY_ISO_Partial_Space_Left      = 0xfe25
	KEY_ISO_Partial_Space_Right     = 0xfe26
	KEY_ISO_Set_Margin_Left         = 0xfe27
	KEY_ISO_Set_Margin_Right        = 0xfe28
	KEY_ISO_Release_Margin_Left     = 0xfe29
	KEY_ISO_Release_Margin_Right    = 0xfe2a
	KEY_ISO_Release_Both_Margins    = 0xfe2b
	KEY_ISO_Fast_Cursor_Left        = 0xfe2c
	KEY_ISO_Fast_Cursor_Right       = 0xfe2d
	KEY_ISO_Fast_Cursor_Up          = 0xfe2e
	KEY_ISO_Fast_Cursor_Down        = 0xfe2f
	KEY_ISO_Continuous_Underline    = 0xfe30
	KEY_ISO_Discontinuous_Underline = 0xfe31
	KEY_ISO_Emphasize               = 0xfe32
	KEY_ISO_Center_Object           = 0xfe33
	KEY_ISO_Enter                   = 0xfe34
	KEY_dead_grave                  = 0xfe50
	KEY_dead_acute                  = 0xfe51
	KEY_dead_circumflex             = 0xfe52
	KEY_dead_tilde                  = 0xfe53
	KEY_dead_perispomeni            = 0xfe53
	KEY_dead_macron                 = 0xfe54
	KEY_dead_breve                  = 0xfe55
	KEY_dead_abovedot               = 0xfe56
	KEY_dead_diaeresis              = 0xfe57
	KEY_dead_abovering              = 0xfe58
	KEY_dead_doubleacute            = 0xfe59
	KEY_dead_caron                  = 0xfe5a
	KEY_dead_cedilla                = 0xfe5b
	KEY_dead_ogonek                 = 0xfe5c
	KEY_dead_iota                   = 0xfe5d
	KEY_dead_voiced_sound           = 0xfe5e
	KEY_dead_semivoiced_sound       = 0xfe5f
	KEY_dead_belowdot               = 0xfe60
	KEY_dead_hook                   = 0xfe61
	KEY_dead_horn                   = 0xfe62
	KEY_dead_stroke                 = 0xfe63
	KEY_dead_abovecomma             = 0xfe64
	KEY_dead_psili                  = 0xfe64
	KEY_dead_abovereversedcomma     = 0xfe65
	KEY_dead_dasia                  = 0xfe65
	KEY_dead_doublegrave            = 0xfe66
	KEY_dead_belowring              = 0xfe67
	KEY_dead_belowmacron            = 0xfe68
	KEY_dead_belowcircumflex        = 0xfe69
	KEY_dead_belowtilde             = 0xfe6a
	KEY_dead_belowbreve             = 0xfe6b
	KEY_dead_belowdiaeresis         = 0xfe6c
	KEY_dead_invertedbreve          = 0xfe6d
	KEY_dead_belowcomma             = 0xfe6e
	KEY_dead_currency               = 0xfe6f
	KEY_dead_a                      = 0xfe80
	KEY_dead_A                      = 0xfe81
	KEY_dead_e                      = 0xfe82
	KEY_dead_E                      = 0xfe83
	KEY_dead_i                      = 0xfe84
	KEY_dead_I                      = 0xfe85
	KEY_dead_o                      = 0xfe86
	KEY_dead_O                      = 0xfe87
	KEY_dead_u                      = 0xfe88
	KEY_dead_U                      = 0xfe89
	KEY_dead_small_schwa            = 0xfe8a
	KEY_dead_capital_schwa          = 0xfe8b
	KEY_First_Virtual_Screen        = 0xfed0
	KEY_Prev_Virtual_Screen         = 0xfed1
	KEY_Next_Virtual_Screen         = 0xfed2
	KEY_Last_Virtual_Screen         = 0xfed4
	KEY_Terminate_Server            = 0xfed5
	KEY_AccessX_Enable              = 0xfe70
	KEY_AccessX_Feedback_Enable     = 0xfe71
	KEY_RepeatKeys_Enable           = 0xfe72
	KEY_SlowKeys_Enable             = 0xfe73
	KEY_BounceKeys_Enable           = 0xfe74
	KEY_StickyKeys_Enable           = 0xfe75
	KEY_MouseKeys_Enable            = 0xfe76
	KEY_MouseKeys_Accel_Enable      = 0xfe77
	KEY_Overlay1_Enable             = 0xfe78
	KEY_Overlay2_Enable             = 0xfe79
	KEY_AudibleBell_Enable          = 0xfe7a
	KEY_Pointer_Left                = 0xfee0
	KEY_Pointer_Right               = 0xfee1
	KEY_Pointer_Up                  = 0xfee2
	KEY_Pointer_Down                = 0xfee3
	KEY_Pointer_UpLeft              = 0xfee4
	KEY_Pointer_UpRight             = 0xfee5
	KEY_Pointer_DownLeft            = 0xfee6
	KEY_Pointer_DownRight           = 0xfee7
	KEY_Pointer_Button_Dflt         = 0xfee8
	KEY_Pointer_Button1             = 0xfee9
	KEY_Pointer_Button2             = 0xfeea
	KEY_Pointer_Button3             = 0xfeeb
	KEY_Pointer_Button4             = 0xfeec
	KEY_Pointer_Button5             = 0xfeed
	KEY_Pointer_DblClick_Dflt       = 0xfeee
	KEY_Pointer_DblClick1           = 0xfeef
	KEY_Pointer_DblClick2           = 0xfef0
	KEY_Pointer_DblClick3           = 0xfef1
	KEY_Pointer_DblClick4           = 0xfef2
	KEY_Pointer_DblClick5           = 0xfef3
	KEY_Pointer_Drag_Dflt           = 0xfef4
	KEY_Pointer_Drag1               = 0xfef5
	KEY_Pointer_Drag2               = 0xfef6
	KEY_Pointer_Drag3               = 0xfef7
	KEY_Pointer_Drag4               = 0xfef8
	KEY_Pointer_Drag5               = 0xfefd
	KEY_Pointer_EnableKeys          = 0xfef9
	KEY_Pointer_Accelerate          = 0xfefa
	KEY_Pointer_DfltBtnNext         = 0xfefb
	KEY_Pointer_DfltBtnPrev         = 0xfefc
	KEY_3270_Duplicate              = 0xfd01
	KEY_3270_FieldMark              = 0xfd02
	KEY_3270_Right2                 = 0xfd03
	KEY_3270_Left2                  = 0xfd04
	KEY_3270_BackTab                = 0xfd05
	KEY_3270_EraseEOF               = 0xfd06
	KEY_3270_EraseInput             = 0xfd07
	KEY_3270_Reset                  = 0xfd08
	KEY_3270_Quit                   = 0xfd09
	KEY_3270_PA1                    = 0xfd0a
	KEY_3270_PA2                    = 0xfd0b
	KEY_3270_PA3                    = 0xfd0c
	KEY_3270_Test                   = 0xfd0d
	KEY_3270_Attn                   = 0xfd0e
	KEY_3270_CursorBlink            = 0xfd0f
	KEY_3270_AltCursor              = 0xfd10
	KEY_3270_KeyClick               = 0xfd11
	KEY_3270_Jump                   = 0xfd12
	KEY_3270_Ident                  = 0xfd13
	KEY_3270_Rule                   = 0xfd14
	KEY_3270_Copy                   = 0xfd15
	KEY_3270_Play                   = 0xfd16
	KEY_3270_Setup                  = 0xfd17
	KEY_3270_Record                 = 0xfd18
	KEY_3270_ChangeScreen           = 0xfd19
	KEY_3270_DeleteWord             = 0xfd1a
	KEY_3270_ExSelect               = 0xfd1b
	KEY_3270_CursorSelect           = 0xfd1c
	KEY_3270_PrintScreen            = 0xfd1d
	KEY_3270_Enter                  = 0xfd1e
	KEY_space                       = 0x020
	KEY_exclam                      = 0x021
	KEY_quotedbl                    = 0x022
	KEY_numbersign                  = 0x023
	KEY_dollar                      = 0x024
	KEY_percent                     = 0x025
	KEY_ampersand                   = 0x026
	KEY_apostrophe                  = 0x027
	KEY_quoteright                  = 0x027
	KEY_parenleft                   = 0x028
	KEY_parenright                  = 0x029
	KEY_asterisk                    = 0x02a
	KEY_plus                        = 0x02b
	KEY_comma                       = 0x02c
	KEY_minus                       = 0x02d
	KEY_period                      = 0x02e
	KEY_slash                       = 0x02f
	KEY_0                           = 0x030
	KEY_1                           = 0x031
	KEY_2                           = 0x032
	KEY_3                           = 0x033
	KEY_4                           = 0x034
	KEY_5                           = 0x035
	KEY_6                           = 0x036
	KEY_7                           = 0x037
	KEY_8                           = 0x038
	KEY_9                           = 0x039
	KEY_colon                       = 0x03a
	KEY_semicolon                   = 0x03b
	KEY_less                        = 0x03c
	KEY_equal                       = 0x03d
	KEY_greater                     = 0x03e
	KEY_question                    = 0x03f
	KEY_at                          = 0x040
	KEY_A                           = 0x041
	KEY_B                           = 0x042
	KEY_C                           = 0x043
	KEY_D                           = 0x044
	KEY_E                           = 0x045
	KEY_F                           = 0x046
	KEY_G                           = 0x047
	KEY_H                           = 0x048
	KEY_I                           = 0x049
	KEY_J                           = 0x04a
	KEY_K                           = 0x04b
	KEY_L                           = 0x04c
	KEY_M                           = 0x04d
	KEY_N                           = 0x04e
	KEY_O                           = 0x04f
	KEY_P                           = 0x050
	KEY_Q                           = 0x051
	KEY_R                           = 0x052
	KEY_S                           = 0x053
	KEY_T                           = 0x054
	KEY_U                           = 0x055
	KEY_V                           = 0x056
	KEY_W                           = 0x057
	KEY_X                           = 0x058
	KEY_Y                           = 0x059
	KEY_Z                           = 0x05a
	KEY_bracketleft                 = 0x05b
	KEY_backslash                   = 0x05c
	KEY_bracketright                = 0x05d
	KEY_asciicircum                 = 0x05e
	KEY_underscore                  = 0x05f
	KEY_grave                       = 0x060
	KEY_quoteleft                   = 0x060
	KEY_a                           = 0x061
	KEY_b                           = 0x062
	KEY_c                           = 0x063
	KEY_d                           = 0x064
	KEY_e                           = 0x065
	KEY_f                           = 0x066
	KEY_g                           = 0x067
	KEY_h                           = 0x068
	KEY_i                           = 0x069
	KEY_j                           = 0x06a
	KEY_k                           = 0x06b
	KEY_l                           = 0x06c
	KEY_m                           = 0x06d
	KEY_n                           = 0x06e
	KEY_o                           = 0x06f
	KEY_p                           = 0x070
	KEY_q                           = 0x071
	KEY_r                           = 0x072
	KEY_s                           = 0x073
	KEY_t                           = 0x074
	KEY_u                           = 0x075
	KEY_v                           = 0x076
	KEY_w                           = 0x077
	KEY_x                           = 0x078
	KEY_y                           = 0x079
	KEY_z                           = 0x07a
	KEY_braceleft                   = 0x07b
	KEY_bar                         = 0x07c
	KEY_braceright                  = 0x07d
	KEY_asciitilde                  = 0x07e
	KEY_nobreakspace                = 0x0a0
	KEY_exclamdown                  = 0x0a1
	KEY_cent                        = 0x0a2
	KEY_sterling                    = 0x0a3
	KEY_currency                    = 0x0a4
	KEY_yen                         = 0x0a5
	KEY_brokenbar                   = 0x0a6
	KEY_section                     = 0x0a7
	KEY_diaeresis                   = 0x0a8
	KEY_copyright                   = 0x0a9
	KEY_ordfeminine                 = 0x0aa
	KEY_guillemotleft               = 0x0ab
	KEY_notsign                     = 0x0ac
	KEY_hyphen                      = 0x0ad
	KEY_registered                  = 0x0ae
	KEY_macron                      = 0x0af
	KEY_degree                      = 0x0b0
	KEY_plusminus                   = 0x0b1
	KEY_twosuperior                 = 0x0b2
	KEY_threesuperior               = 0x0b3
	KEY_acute                       = 0x0b4
	KEY_mu                          = 0x0b5
	KEY_paragraph                   = 0x0b6
	KEY_periodcentered              = 0x0b7
	KEY_cedilla                     = 0x0b8
	KEY_onesuperior                 = 0x0b9
	KEY_masculine                   = 0x0ba
	KEY_guillemotright              = 0x0bb
	KEY_onequarter                  = 0x0bc
	KEY_onehalf                     = 0x0bd
	KEY_threequarters               = 0x0be
	KEY_questiondown                = 0x0bf
	KEY_Agrave                      = 0x0c0
	KEY_Aacute                      = 0x0c1
	KEY_Acircumflex                 = 0x0c2
	KEY_Atilde                      = 0x0c3
	KEY_Adiaeresis                  = 0x0c4
	KEY_Aring                       = 0x0c5
	KEY_AE                          = 0x0c6
	KEY_Ccedilla                    = 0x0c7
	KEY_Egrave                      = 0x0c8
	KEY_Eacute                      = 0x0c9
	KEY_Ecircumflex                 = 0x0ca
	KEY_Ediaeresis                  = 0x0cb
	KEY_Igrave                      = 0x0cc
	KEY_Iacute                      = 0x0cd
	KEY_Icircumflex                 = 0x0ce
	KEY_Idiaeresis                  = 0x0cf
	KEY_ETH                         = 0x0d0
	KEY_Eth                         = 0x0d0
	KEY_Ntilde                      = 0x0d1
	KEY_Ograve                      = 0x0d2
	KEY_Oacute                      = 0x0d3
	KEY_Ocircumflex                 = 0x0d4
	KEY_Otilde                      = 0x0d5
	KEY_Odiaeresis                  = 0x0d6
	KEY_multiply                    = 0x0d7
	KEY_Oslash                      = 0x0d8
	KEY_Ooblique                    = 0x0d8
	KEY_Ugrave                      = 0x0d9
	KEY_Uacute                      = 0x0da
	KEY_Ucircumflex                 = 0x0db
	KEY_Udiaeresis                  = 0x0dc
	KEY_Yacute                      = 0x0dd
	KEY_THORN                       = 0x0de
	KEY_Thorn                       = 0x0de
	KEY_ssharp                      = 0x0df
	KEY_agrave                      = 0x0e0
	KEY_aacute                      = 0x0e1
	KEY_acircumflex                 = 0x0e2
	KEY_atilde                      = 0x0e3
	KEY_adiaeresis                  = 0x0e4
	KEY_aring                       = 0x0e5
	KEY_ae                          = 0x0e6
	KEY_ccedilla                    = 0x0e7
	KEY_egrave                      = 0x0e8
	KEY_eacute                      = 0x0e9
	KEY_ecircumflex                 = 0x0ea
	KEY_ediaeresis                  = 0x0eb
	KEY_igrave                      = 0x0ec
	KEY_iacute                      = 0x0ed
	KEY_icircumflex                 = 0x0ee
	KEY_idiaeresis                  = 0x0ef
	KEY_eth                         = 0x0f0
	KEY_ntilde                      = 0x0f1
	KEY_ograve                      = 0x0f2
	KEY_oacute                      = 0x0f3
	KEY_ocircumflex                 = 0x0f4
	KEY_otilde                      = 0x0f5
	KEY_odiaeresis                  = 0x0f6
	KEY_division                    = 0x0f7
	KEY_oslash                      = 0x0f8
	KEY_ooblique                    = 0x0f8
	KEY_ugrave                      = 0x0f9
	KEY_uacute                      = 0x0fa
	KEY_ucircumflex                 = 0x0fb
	KEY_udiaeresis                  = 0x0fc
	KEY_yacute                      = 0x0fd
	KEY_thorn                       = 0x0fe
	KEY_ydiaeresis                  = 0x0ff
	KEY_Aogonek                     = 0x1a1
	KEY_breve                       = 0x1a2
	KEY_Lstroke                     = 0x1a3
	KEY_Lcaron                      = 0x1a5
	KEY_Sacute                      = 0x1a6
	KEY_Scaron                      = 0x1a9
	KEY_Scedilla                    = 0x1aa
	KEY_Tcaron                      = 0x1ab
	KEY_Zacute                      = 0x1ac
	KEY_Zcaron                      = 0x1ae
	KEY_Zabovedot                   = 0x1af
	KEY_aogonek                     = 0x1b1
	KEY_ogonek                      = 0x1b2
	KEY_lstroke                     = 0x1b3
	KEY_lcaron                      = 0x1b5
	KEY_sacute                      = 0x1b6
	KEY_caron                       = 0x1b7
	KEY_scaron                      = 0x1b9
	KEY_scedilla                    = 0x1ba
	KEY_tcaron                      = 0x1bb
	KEY_zacute                      = 0x1bc
	KEY_doubleacute                 = 0x1bd
	KEY_zcaron                      = 0x1be
	KEY_zabovedot                   = 0x1bf
	KEY_Racute                      = 0x1c0
	KEY_Abreve                      = 0x1c3
	KEY_Lacute                      = 0x1c5
	KEY_Cacute                      = 0x1c6
	KEY_Ccaron                      = 0x1c8
	KEY_Eogonek                     = 0x1ca
	KEY_Ecaron                      = 0x1cc
	KEY_Dcaron                      = 0x1cf
	KEY_Dstroke                     = 0x1d0
	KEY_Nacute                      = 0x1d1
	KEY_Ncaron                      = 0x1d2
	KEY_Odoubleacute                = 0x1d5
	KEY_Rcaron                      = 0x1d8
	KEY_Uring                       = 0x1d9
	KEY_Udoubleacute                = 0x1db
	KEY_Tcedilla                    = 0x1de
	KEY_racute                      = 0x1e0
	KEY_abreve                      = 0x1e3
	KEY_lacute                      = 0x1e5
	KEY_cacute                      = 0x1e6
	KEY_ccaron                      = 0x1e8
	KEY_eogonek                     = 0x1ea
	KEY_ecaron                      = 0x1ec
	KEY_dcaron                      = 0x1ef
	KEY_dstroke                     = 0x1f0
	KEY_nacute                      = 0x1f1
	KEY_ncaron                      = 0x1f2
	KEY_odoubleacute                = 0x1f5
	KEY_udoubleacute                = 0x1fb
	KEY_rcaron                      = 0x1f8
	KEY_uring                       = 0x1f9
	KEY_tcedilla                    = 0x1fe
	KEY_abovedot                    = 0x1ff
	KEY_Hstroke                     = 0x2a1
	KEY_Hcircumflex                 = 0x2a6
	KEY_Iabovedot                   = 0x2a9
	KEY_Gbreve                      = 0x2ab
	KEY_Jcircumflex                 = 0x2ac
	KEY_hstroke                     = 0x2b1
	KEY_hcircumflex                 = 0x2b6
	KEY_idotless                    = 0x2b9
	KEY_gbreve                      = 0x2bb
	KEY_jcircumflex                 = 0x2bc
	KEY_Cabovedot                   = 0x2c5
	KEY_Ccircumflex                 = 0x2c6
	KEY_Gabovedot                   = 0x2d5
	KEY_Gcircumflex                 = 0x2d8
	KEY_Ubreve                      = 0x2dd
	KEY_Scircumflex                 = 0x2de
	KEY_cabovedot                   = 0x2e5
	KEY_ccircumflex                 = 0x2e6
	KEY_gabovedot                   = 0x2f5
	KEY_gcircumflex                 = 0x2f8
	KEY_ubreve                      = 0x2fd
	KEY_scircumflex                 = 0x2fe
	KEY_kra                         = 0x3a2
	KEY_kappa                       = 0x3a2
	KEY_Rcedilla                    = 0x3a3
	KEY_Itilde                      = 0x3a5
	KEY_Lcedilla                    = 0x3a6
	KEY_Emacron                     = 0x3aa
	KEY_Gcedilla                    = 0x3ab
	KEY_Tslash                      = 0x3ac
	KEY_rcedilla                    = 0x3b3
	KEY_itilde                      = 0x3b5
	KEY_lcedilla                    = 0x3b6
	KEY_emacron                     = 0x3ba
	KEY_gcedilla                    = 0x3bb
	KEY_tslash                      = 0x3bc
	KEY_ENG                         = 0x3bd
	KEY_eng                         = 0x3bf
	KEY_Amacron                     = 0x3c0
	KEY_Iogonek                     = 0x3c7
	KEY_Eabovedot                   = 0x3cc
	KEY_Imacron                     = 0x3cf
	KEY_Ncedilla                    = 0x3d1
	KEY_Omacron                     = 0x3d2
	KEY_Kcedilla                    = 0x3d3
	KEY_Uogonek                     = 0x3d9
	KEY_Utilde                      = 0x3dd
	KEY_Umacron                     = 0x3de
	KEY_amacron                     = 0x3e0
	KEY_iogonek                     = 0x3e7
	KEY_eabovedot                   = 0x3ec
	KEY_imacron                     = 0x3ef
	KEY_ncedilla                    = 0x3f1
	KEY_omacron                     = 0x3f2
	KEY_kcedilla                    = 0x3f3
	KEY_uogonek                     = 0x3f9
	KEY_utilde                      = 0x3fd
	KEY_umacron                     = 0x3fe
	KEY_Babovedot                   = 0x1001e02
	KEY_babovedot                   = 0x1001e03
	KEY_Dabovedot                   = 0x1001e0a
	KEY_Wgrave                      = 0x1001e80
	KEY_Wacute                      = 0x1001e82
	KEY_dabovedot                   = 0x1001e0b
	KEY_Ygrave                      = 0x1001ef2
	KEY_Fabovedot                   = 0x1001e1e
	KEY_fabovedot                   = 0x1001e1f
	KEY_Mabovedot                   = 0x1001e40
	KEY_mabovedot                   = 0x1001e41
	KEY_Pabovedot                   = 0x1001e56
	KEY_wgrave                      = 0x1001e81
	KEY_pabovedot                   = 0x1001e57
	KEY_wacute                      = 0x1001e83
	KEY_Sabovedot                   = 0x1001e60
	KEY_ygrave                      = 0x1001ef3
	KEY_Wdiaeresis                  = 0x1001e84
	KEY_wdiaeresis                  = 0x1001e85
	KEY_sabovedot                   = 0x1001e61
	KEY_Wcircumflex                 = 0x1000174
	KEY_Tabovedot                   = 0x1001e6a
	KEY_Ycircumflex                 = 0x1000176
	KEY_wcircumflex                 = 0x1000175
	KEY_tabovedot                   = 0x1001e6b
	KEY_ycircumflex                 = 0x1000177
	KEY_OE                          = 0x13bc
	KEY_oe                          = 0x13bd
	KEY_Ydiaeresis                  = 0x13be
	KEY_overline                    = 0x47e
	KEY_kana_fullstop               = 0x4a1
	KEY_kana_openingbracket         = 0x4a2
	KEY_kana_closingbracket         = 0x4a3
	KEY_kana_comma                  = 0x4a4
	KEY_kana_conjunctive            = 0x4a5
	KEY_kana_middledot              = 0x4a5
	KEY_kana_WO                     = 0x4a6
	KEY_kana_a                      = 0x4a7
	KEY_kana_i                      = 0x4a8
	KEY_kana_u                      = 0x4a9
	KEY_kana_e                      = 0x4aa
	KEY_kana_o                      = 0x4ab
	KEY_kana_ya                     = 0x4ac
	KEY_kana_yu                     = 0x4ad
	KEY_kana_yo                     = 0x4ae
	KEY_kana_tsu                    = 0x4af
	KEY_kana_tu                     = 0x4af
	KEY_prolongedsound              = 0x4b0
	KEY_kana_A                      = 0x4b1
	KEY_kana_I                      = 0x4b2
	KEY_kana_U                      = 0x4b3
	KEY_kana_E                      = 0x4b4
	KEY_kana_O                      = 0x4b5
	KEY_kana_KA                     = 0x4b6
	KEY_kana_KI                     = 0x4b7
	KEY_kana_KU                     = 0x4b8
	KEY_kana_KE                     = 0x4b9
	KEY_kana_KO                     = 0x4ba
	KEY_kana_SA                     = 0x4bb
	KEY_kana_SHI                    = 0x4bc
	KEY_kana_SU                     = 0x4bd
	KEY_kana_SE                     = 0x4be
	KEY_kana_SO                     = 0x4bf
	KEY_kana_TA                     = 0x4c0
	KEY_kana_CHI                    = 0x4c1
	KEY_kana_TI                     = 0x4c1
	KEY_kana_TSU                    = 0x4c2
	KEY_kana_TU                     = 0x4c2
	KEY_kana_TE                     = 0x4c3
	KEY_kana_TO                     = 0x4c4
	KEY_kana_NA                     = 0x4c5
	KEY_kana_NI                     = 0x4c6
	KEY_kana_NU                     = 0x4c7
	KEY_kana_NE                     = 0x4c8
	KEY_kana_NO                     = 0x4c9
	KEY_kana_HA                     = 0x4ca
	KEY_kana_HI                     = 0x4cb
	KEY_kana_FU                     = 0x4cc
	KEY_kana_HU                     = 0x4cc
	KEY_kana_HE                     = 0x4cd
	KEY_kana_HO                     = 0x4ce
	KEY_kana_MA                     = 0x4cf
	KEY_kana_MI                     = 0x4d0
	KEY_kana_MU                     = 0x4d1
	KEY_kana_ME                     = 0x4d2
	KEY_kana_MO                     = 0x4d3
	KEY_kana_YA                     = 0x4d4
	KEY_kana_YU                     = 0x4d5
	KEY_kana_YO                     = 0x4d6
	KEY_kana_RA                     = 0x4d7
	KEY_kana_RI                     = 0x4d8
	KEY_kana_RU                     = 0x4d9
	KEY_kana_RE                     = 0x4da
	KEY_kana_RO                     = 0x4db
	KEY_kana_WA                     = 0x4dc
	KEY_kana_N                      = 0x4dd
	KEY_voicedsound                 = 0x4de
	KEY_semivoicedsound             = 0x4df
	KEY_kana_switch                 = 0xff7e
	KEY_Farsi_0                     = 0x10006f0
	KEY_Farsi_1                     = 0x10006f1
	KEY_Farsi_2                     = 0x10006f2
	KEY_Farsi_3                     = 0x10006f3
	KEY_Farsi_4                     = 0x10006f4
	KEY_Farsi_5                     = 0x10006f5
	KEY_Farsi_6                     = 0x10006f6
	KEY_Farsi_7                     = 0x10006f7
	KEY_Farsi_8                     = 0x10006f8
	KEY_Farsi_9                     = 0x10006f9
	KEY_Arabic_percent              = 0x100066a
	KEY_Arabic_superscript_alef     = 0x1000670
	KEY_Arabic_tteh                 = 0x1000679
	KEY_Arabic_peh                  = 0x100067e
	KEY_Arabic_tcheh                = 0x1000686
	KEY_Arabic_ddal                 = 0x1000688
	KEY_Arabic_rreh                 = 0x1000691
	KEY_Arabic_comma                = 0x5ac
	KEY_Arabic_fullstop             = 0x10006d4
	KEY_Arabic_0                    = 0x1000660
	KEY_Arabic_1                    = 0x1000661
	KEY_Arabic_2                    = 0x1000662
	KEY_Arabic_3                    = 0x1000663
	KEY_Arabic_4                    = 0x1000664
	KEY_Arabic_5                    = 0x1000665
	KEY_Arabic_6                    = 0x1000666
	KEY_Arabic_7                    = 0x1000667
	KEY_Arabic_8                    = 0x1000668
	KEY_Arabic_9                    = 0x1000669
	KEY_Arabic_semicolon            = 0x5bb
	KEY_Arabic_question_mark        = 0x5bf
	KEY_Arabic_hamza                = 0x5c1
	KEY_Arabic_maddaonalef          = 0x5c2
	KEY_Arabic_hamzaonalef          = 0x5c3
	KEY_Arabic_hamzaonwaw           = 0x5c4
	KEY_Arabic_hamzaunderalef       = 0x5c5
	KEY_Arabic_hamzaonyeh           = 0x5c6
	KEY_Arabic_alef                 = 0x5c7
	KEY_Arabic_beh                  = 0x5c8
	KEY_Arabic_tehmarbuta           = 0x5c9
	KEY_Arabic_teh                  = 0x5ca
	KEY_Arabic_theh                 = 0x5cb
	KEY_Arabic_jeem                 = 0x5cc
	KEY_Arabic_hah                  = 0x5cd
	KEY_Arabic_khah                 = 0x5ce
	KEY_Arabic_dal                  = 0x5cf
	KEY_Arabic_thal                 = 0x5d0
	KEY_Arabic_ra                   = 0x5d1
	KEY_Arabic_zain                 = 0x5d2
	KEY_Arabic_seen                 = 0x5d3
	KEY_Arabic_sheen                = 0x5d4
	KEY_Arabic_sad                  = 0x5d5
	KEY_Arabic_dad                  = 0x5d6
	KEY_Arabic_tah                  = 0x5d7
	KEY_Arabic_zah                  = 0x5d8
	KEY_Arabic_ain                  = 0x5d9
	KEY_Arabic_ghain                = 0x5da
	KEY_Arabic_tatweel              = 0x5e0
	KEY_Arabic_feh                  = 0x5e1
	KEY_Arabic_qaf                  = 0x5e2
	KEY_Arabic_kaf                  = 0x5e3
	KEY_Arabic_lam                  = 0x5e4
	KEY_Arabic_meem                 = 0x5e5
	KEY_Arabic_noon                 = 0x5e6
	KEY_Arabic_ha                   = 0x5e7
	KEY_Arabic_heh                  = 0x5e7
	KEY_Arabic_waw                  = 0x5e8
	KEY_Arabic_alefmaksura          = 0x5e9
	KEY_Arabic_yeh                  = 0x5ea
	KEY_Arabic_fathatan             = 0x5eb
	KEY_Arabic_dammatan             = 0x5ec
	KEY_Arabic_kasratan             = 0x5ed
	KEY_Arabic_fatha                = 0x5ee
	KEY_Arabic_damma                = 0x5ef
	KEY_Arabic_kasra                = 0x5f0
	KEY_Arabic_shadda               = 0x5f1
	KEY_Arabic_sukun                = 0x5f2
	KEY_Arabic_madda_above          = 0x1000653
	KEY_Arabic_hamza_above          = 0x1000654
	KEY_Arabic_hamza_below          = 0x1000655
	KEY_Arabic_jeh                  = 0x1000698
	KEY_Arabic_veh                  = 0x10006a4
	KEY_Arabic_keheh                = 0x10006a9
	KEY_Arabic_gaf                  = 0x10006af
	KEY_Arabic_noon_ghunna          = 0x10006ba
	KEY_Arabic_heh_doachashmee      = 0x10006be
	KEY_Farsi_yeh                   = 0x10006cc
	KEY_Arabic_farsi_yeh            = 0x10006cc
	KEY_Arabic_yeh_baree            = 0x10006d2
	KEY_Arabic_heh_goal             = 0x10006c1
	KEY_Arabic_switch               = 0xff7e
	KEY_Cyrillic_GHE_bar            = 0x1000492
	KEY_Cyrillic_ghe_bar            = 0x1000493
	KEY_Cyrillic_ZHE_descender      = 0x1000496
	KEY_Cyrillic_zhe_descender      = 0x1000497
	KEY_Cyrillic_KA_descender       = 0x100049a
	KEY_Cyrillic_ka_descender       = 0x100049b
	KEY_Cyrillic_KA_vertstroke      = 0x100049c
	KEY_Cyrillic_ka_vertstroke      = 0x100049d
	KEY_Cyrillic_EN_descender       = 0x10004a2
	KEY_Cyrillic_en_descender       = 0x10004a3
	KEY_Cyrillic_U_straight         = 0x10004ae
	KEY_Cyrillic_u_straight         = 0x10004af
	KEY_Cyrillic_U_straight_bar     = 0x10004b0
	KEY_Cyrillic_u_straight_bar     = 0x10004b1
	KEY_Cyrillic_HA_descender       = 0x10004b2
	KEY_Cyrillic_ha_descender       = 0x10004b3
	KEY_Cyrillic_CHE_descender      = 0x10004b6
	KEY_Cyrillic_che_descender      = 0x10004b7
	KEY_Cyrillic_CHE_vertstroke     = 0x10004b8
	KEY_Cyrillic_che_vertstroke     = 0x10004b9
	KEY_Cyrillic_SHHA               = 0x10004ba
	KEY_Cyrillic_shha               = 0x10004bb
	KEY_Cyrillic_SCHWA              = 0x10004d8
	KEY_Cyrillic_schwa              = 0x10004d9
	KEY_Cyrillic_I_macron           = 0x10004e2
	KEY_Cyrillic_i_macron           = 0x10004e3
	KEY_Cyrillic_O_bar              = 0x10004e8
	KEY_Cyrillic_o_bar              = 0x10004e9
	KEY_Cyrillic_U_macron           = 0x10004ee
	KEY_Cyrillic_u_macron           = 0x10004ef
	KEY_Serbian_dje                 = 0x6a1
	KEY_Macedonia_gje               = 0x6a2
	KEY_Cyrillic_io                 = 0x6a3
	KEY_Ukrainian_ie                = 0x6a4
	KEY_Ukranian_je                 = 0x6a4
	KEY_Macedonia_dse               = 0x6a5
	KEY_Ukrainian_i                 = 0x6a6
	KEY_Ukranian_i                  = 0x6a6
	KEY_Ukrainian_yi                = 0x6a7
	KEY_Ukranian_yi                 = 0x6a7
	KEY_Cyrillic_je                 = 0x6a8
	KEY_Serbian_je                  = 0x6a8
	KEY_Cyrillic_lje                = 0x6a9
	KEY_Serbian_lje                 = 0x6a9
	KEY_Cyrillic_nje                = 0x6aa
	KEY_Serbian_nje                 = 0x6aa
	KEY_Serbian_tshe                = 0x6ab
	KEY_Macedonia_kje               = 0x6ac
	KEY_Ukrainian_ghe_with_upturn   = 0x6ad
	KEY_Byelorussian_shortu         = 0x6ae
	KEY_Cyrillic_dzhe               = 0x6af
	KEY_Serbian_dze                 = 0x6af
	KEY_numerosign                  = 0x6b0
	KEY_Serbian_DJE                 = 0x6b1
	KEY_Macedonia_GJE               = 0x6b2
	KEY_Cyrillic_IO                 = 0x6b3
	KEY_Ukrainian_IE                = 0x6b4
	KEY_Ukranian_JE                 = 0x6b4
	KEY_Macedonia_DSE               = 0x6b5
	KEY_Ukrainian_I                 = 0x6b6
	KEY_Ukranian_I                  = 0x6b6
	KEY_Ukrainian_YI                = 0x6b7
	KEY_Ukranian_YI                 = 0x6b7
	KEY_Cyrillic_JE                 = 0x6b8
	KEY_Serbian_JE                  = 0x6b8
	KEY_Cyrillic_LJE                = 0x6b9
	KEY_Serbian_LJE                 = 0x6b9
	KEY_Cyrillic_NJE                = 0x6ba
	KEY_Serbian_NJE                 = 0x6ba
	KEY_Serbian_TSHE                = 0x6bb
	KEY_Macedonia_KJE               = 0x6bc
	KEY_Ukrainian_GHE_WITH_UPTURN   = 0x6bd
	KEY_Byelorussian_SHORTU         = 0x6be
	KEY_Cyrillic_DZHE               = 0x6bf
	KEY_Serbian_DZE                 = 0x6bf
	KEY_Cyrillic_yu                 = 0x6c0
	KEY_Cyrillic_a                  = 0x6c1
	KEY_Cyrillic_be                 = 0x6c2
	KEY_Cyrillic_tse                = 0x6c3
	KEY_Cyrillic_de                 = 0x6c4
	KEY_Cyrillic_ie                 = 0x6c5
	KEY_Cyrillic_ef                 = 0x6c6
	KEY_Cyrillic_ghe                = 0x6c7
	KEY_Cyrillic_ha                 = 0x6c8
	KEY_Cyrillic_i                  = 0x6c9
	KEY_Cyrillic_shorti             = 0x6ca
	KEY_Cyrillic_ka                 = 0x6cb
	KEY_Cyrillic_el                 = 0x6cc
	KEY_Cyrillic_em                 = 0x6cd
	KEY_Cyrillic_en                 = 0x6ce
	KEY_Cyrillic_o                  = 0x6cf
	KEY_Cyrillic_pe                 = 0x6d0
	KEY_Cyrillic_ya                 = 0x6d1
	KEY_Cyrillic_er                 = 0x6d2
	KEY_Cyrillic_es                 = 0x6d3
	KEY_Cyrillic_te                 = 0x6d4
	KEY_Cyrillic_u                  = 0x6d5
	KEY_Cyrillic_zhe                = 0x6d6
	KEY_Cyrillic_ve                 = 0x6d7
	KEY_Cyrillic_softsign           = 0x6d8
	KEY_Cyrillic_yeru               = 0x6d9
	KEY_Cyrillic_ze                 = 0x6da
	KEY_Cyrillic_sha                = 0x6db
	KEY_Cyrillic_e                  = 0x6dc
	KEY_Cyrillic_shcha              = 0x6dd
	KEY_Cyrillic_che                = 0x6de
	KEY_Cyrillic_hardsign           = 0x6df
	KEY_Cyrillic_YU                 = 0x6e0
	KEY_Cyrillic_A                  = 0x6e1
	KEY_Cyrillic_BE                 = 0x6e2
	KEY_Cyrillic_TSE                = 0x6e3
	KEY_Cyrillic_DE                 = 0x6e4
	KEY_Cyrillic_IE                 = 0x6e5
	KEY_Cyrillic_EF                 = 0x6e6
	KEY_Cyrillic_GHE                = 0x6e7
	KEY_Cyrillic_HA                 = 0x6e8
	KEY_Cyrillic_I                  = 0x6e9
	KEY_Cyrillic_SHORTI             = 0x6ea
	KEY_Cyrillic_KA                 = 0x6eb
	KEY_Cyrillic_EL                 = 0x6ec
	KEY_Cyrillic_EM                 = 0x6ed
	KEY_Cyrillic_EN                 = 0x6ee
	KEY_Cyrillic_O                  = 0x6ef
	KEY_Cyrillic_PE                 = 0x6f0
	KEY_Cyrillic_YA                 = 0x6f1
	KEY_Cyrillic_ER                 = 0x6f2
	KEY_Cyrillic_ES                 = 0x6f3
	KEY_Cyrillic_TE                 = 0x6f4
	KEY_Cyrillic_U                  = 0x6f5
	KEY_Cyrillic_ZHE                = 0x6f6
	KEY_Cyrillic_VE                 = 0x6f7
	KEY_Cyrillic_SOFTSIGN           = 0x6f8
	KEY_Cyrillic_YERU               = 0x6f9
	KEY_Cyrillic_ZE                 = 0x6fa
	KEY_Cyrillic_SHA                = 0x6fb
	KEY_Cyrillic_E                  = 0x6fc
	KEY_Cyrillic_SHCHA              = 0x6fd
	KEY_Cyrillic_CHE                = 0x6fe
	KEY_Cyrillic_HARDSIGN           = 0x6ff
	KEY_Greek_ALPHAaccent           = 0x7a1
	KEY_Greek_EPSILONaccent         = 0x7a2
	KEY_Greek_ETAaccent             = 0x7a3
	KEY_Greek_IOTAaccent            = 0x7a4
	KEY_Greek_IOTAdieresis          = 0x7a5
	KEY_Greek_IOTAdiaeresis         = 0x7a5
	KEY_Greek_OMICRONaccent         = 0x7a7
	KEY_Greek_UPSILONaccent         = 0x7a8
	KEY_Greek_UPSILONdieresis       = 0x7a9
	KEY_Greek_OMEGAaccent           = 0x7ab
	KEY_Greek_accentdieresis        = 0x7ae
	KEY_Greek_horizbar              = 0x7af
	KEY_Greek_alphaaccent           = 0x7b1
	KEY_Greek_epsilonaccent         = 0x7b2
	KEY_Greek_etaaccent             = 0x7b3
	KEY_Greek_iotaaccent            = 0x7b4
	KEY_Greek_iotadieresis          = 0x7b5
	KEY_Greek_iotaaccentdieresis    = 0x7b6
	KEY_Greek_omicronaccent         = 0x7b7
	KEY_Greek_upsilonaccent         = 0x7b8
	KEY_Greek_upsilondieresis       = 0x7b9
	KEY_Greek_upsilonaccentdieresis = 0x7ba
	KEY_Greek_omegaaccent           = 0x7bb
	KEY_Greek_ALPHA                 = 0x7c1
	KEY_Greek_BETA                  = 0x7c2
	KEY_Greek_GAMMA                 = 0x7c3
	KEY_Greek_DELTA                 = 0x7c4
	KEY_Greek_EPSILON               = 0x7c5
	KEY_Greek_ZETA                  = 0x7c6
	KEY_Greek_ETA                   = 0x7c7
	KEY_Greek_THETA                 = 0x7c8
	KEY_Greek_IOTA                  = 0x7c9
	KEY_Greek_KAPPA                 = 0x7ca
	KEY_Greek_LAMDA                 = 0x7cb
	KEY_Greek_LAMBDA                = 0x7cb
	KEY_Greek_MU                    = 0x7cc
	KEY_Greek_NU                    = 0x7cd
	KEY_Greek_XI                    = 0x7ce
	KEY_Greek_OMICRON               = 0x7cf
	KEY_Greek_PI                    = 0x7d0
	KEY_Greek_RHO                   = 0x7d1
	KEY_Greek_SIGMA                 = 0x7d2
	KEY_Greek_TAU                   = 0x7d4
	KEY_Greek_UPSILON               = 0x7d5
	KEY_Greek_PHI                   = 0x7d6
	KEY_Greek_CHI                   = 0x7d7
	KEY_Greek_PSI                   = 0x7d8
	KEY_Greek_OMEGA                 = 0x7d9
	KEY_Greek_alpha                 = 0x7e1
	KEY_Greek_beta                  = 0x7e2
	KEY_Greek_gamma                 = 0x7e3
	KEY_Greek_delta                 = 0x7e4
	KEY_Greek_epsilon               = 0x7e5
	KEY_Greek_zeta                  = 0x7e6
	KEY_Greek_eta                   = 0x7e7
	KEY_Greek_theta                 = 0x7e8
	KEY_Greek_iota                  = 0x7e9
	KEY_Greek_kappa                 = 0x7ea
	KEY_Greek_lamda                 = 0x7eb
	KEY_Greek_lambda                = 0x7eb
	KEY_Greek_mu                    = 0x7ec
	KEY_Greek_nu                    = 0x7ed
	KEY_Greek_xi                    = 0x7ee
	KEY_Greek_omicron               = 0x7ef
	KEY_Greek_pi                    = 0x7f0
	KEY_Greek_rho                   = 0x7f1
	KEY_Greek_sigma                 = 0x7f2
	KEY_Greek_finalsmallsigma       = 0x7f3
	KEY_Greek_tau                   = 0x7f4
	KEY_Greek_upsilon               = 0x7f5
	KEY_Greek_phi                   = 0x7f6
	KEY_Greek_chi                   = 0x7f7
	KEY_Greek_psi                   = 0x7f8
	KEY_Greek_omega                 = 0x7f9
	KEY_Greek_switch                = 0xff7e
	KEY_leftradical                 = 0x8a1
	KEY_topleftradical              = 0x8a2
	KEY_horizconnector              = 0x8a3
	KEY_topintegral                 = 0x8a4
	KEY_botintegral                 = 0x8a5
	KEY_vertconnector               = 0x8a6
	KEY_topleftsqbracket            = 0x8a7
	KEY_botleftsqbracket            = 0x8a8
	KEY_toprightsqbracket           = 0x8a9
	KEY_botrightsqbracket           = 0x8aa
	KEY_topleftparens               = 0x8ab
	KEY_botleftparens               = 0x8ac
	KEY_toprightparens              = 0x8ad
	KEY_botrightparens              = 0x8ae
	KEY_leftmiddlecurlybrace        = 0x8af
	KEY_rightmiddlecurlybrace       = 0x8b0
	KEY_topleftsummation            = 0x8b1
	KEY_botleftsummation            = 0x8b2
	KEY_topvertsummationconnector   = 0x8b3
	KEY_botvertsummationconnector   = 0x8b4
	KEY_toprightsummation           = 0x8b5
	KEY_botrightsummation           = 0x8b6
	KEY_rightmiddlesummation        = 0x8b7
	KEY_lessthanequal               = 0x8bc
	KEY_notequal                    = 0x8bd
	KEY_greaterthanequal            = 0x8be
	KEY_integral                    = 0x8bf
	KEY_therefore                   = 0x8c0
	KEY_variation                   = 0x8c1
	KEY_infinity                    = 0x8c2
	KEY_nabla                       = 0x8c5
	KEY_approximate                 = 0x8c8
	KEY_similarequal                = 0x8c9
	KEY_ifonlyif                    = 0x8cd
	KEY_implies                     = 0x8ce
	KEY_identical                   = 0x8cf
	KEY_radical                     = 0x8d6
	KEY_includedin                  = 0x8da
	KEY_includes                    = 0x8db
	KEY_intersection                = 0x8dc
	KEY_union                       = 0x8dd
	KEY_logicaland                  = 0x8de
	KEY_logicalor                   = 0x8df
	KEY_partialderivative           = 0x8ef
	KEY_function                    = 0x8f6
	KEY_leftarrow                   = 0x8fb
	KEY_uparrow                     = 0x8fc
	KEY_rightarrow                  = 0x8fd
	KEY_downarrow                   = 0x8fe
	KEY_blank                       = 0x9df
	KEY_soliddiamond                = 0x9e0
	KEY_checkerboard                = 0x9e1
	KEY_ht                          = 0x9e2
	KEY_ff                          = 0x9e3
	KEY_cr                          = 0x9e4
	KEY_lf                          = 0x9e5
	KEY_nl                          = 0x9e8
	KEY_vt                          = 0x9e9
	KEY_lowrightcorner              = 0x9ea
	KEY_uprightcorner               = 0x9eb
	KEY_upleftcorner                = 0x9ec
	KEY_lowleftcorner               = 0x9ed
	KEY_crossinglines               = 0x9ee
	KEY_horizlinescan1              = 0x9ef
	KEY_horizlinescan3              = 0x9f0
	KEY_horizlinescan5              = 0x9f1
	KEY_horizlinescan7              = 0x9f2
	KEY_horizlinescan9              = 0x9f3
	KEY_leftt                       = 0x9f4
	KEY_rightt                      = 0x9f5
	KEY_bott                        = 0x9f6
	KEY_topt                        = 0x9f7
	KEY_vertbar                     = 0x9f8
	KEY_emspace                     = 0xaa1
	KEY_enspace                     = 0xaa2
	KEY_em3space                    = 0xaa3
	KEY_em4space                    = 0xaa4
	KEY_digitspace                  = 0xaa5
	KEY_punctspace                  = 0xaa6
	KEY_thinspace                   = 0xaa7
	KEY_hairspace                   = 0xaa8
	KEY_emdash                      = 0xaa9
	KEY_endash                      = 0xaaa
	KEY_signifblank                 = 0xaac
	KEY_ellipsis                    = 0xaae
	KEY_doubbaselinedot             = 0xaaf
	KEY_onethird                    = 0xab0
	KEY_twothirds                   = 0xab1
	KEY_onefifth                    = 0xab2
	KEY_twofifths                   = 0xab3
	KEY_threefifths                 = 0xab4
	KEY_fourfifths                  = 0xab5
	KEY_onesixth                    = 0xab6
	KEY_fivesixths                  = 0xab7
	KEY_careof                      = 0xab8
	KEY_figdash                     = 0xabb
	KEY_leftanglebracket            = 0xabc
	KEY_decimalpoint                = 0xabd
	KEY_rightanglebracket           = 0xabe
	KEY_marker                      = 0xabf
	KEY_oneeighth                   = 0xac3
	KEY_threeeighths                = 0xac4
	KEY_fiveeighths                 = 0xac5
	KEY_seveneighths                = 0xac6
	KEY_trademark                   = 0xac9
	KEY_signaturemark               = 0xaca
	KEY_trademarkincircle           = 0xacb
	KEY_leftopentriangle            = 0xacc
	KEY_rightopentriangle           = 0xacd
	KEY_emopencircle                = 0xace
	KEY_emopenrectangle             = 0xacf
	KEY_leftsinglequotemark         = 0xad0
	KEY_rightsinglequotemark        = 0xad1
	KEY_leftdoublequotemark         = 0xad2
	KEY_rightdoublequotemark        = 0xad3
	KEY_prescription                = 0xad4
	KEY_minutes                     = 0xad6
	KEY_seconds                     = 0xad7
	KEY_latincross                  = 0xad9
	KEY_hexagram                    = 0xada
	KEY_filledrectbullet            = 0xadb
	KEY_filledlefttribullet         = 0xadc
	KEY_filledrighttribullet        = 0xadd
	KEY_emfilledcircle              = 0xade
	KEY_emfilledrect                = 0xadf
	KEY_enopencircbullet            = 0xae0
	KEY_enopensquarebullet          = 0xae1
	KEY_openrectbullet              = 0xae2
	KEY_opentribulletup             = 0xae3
	KEY_opentribulletdown           = 0xae4
	KEY_openstar                    = 0xae5
	KEY_enfilledcircbullet          = 0xae6
	KEY_enfilledsqbullet            = 0xae7
	KEY_filledtribulletup           = 0xae8
	KEY_filledtribulletdown         = 0xae9
	KEY_leftpointer                 = 0xaea
	KEY_rightpointer                = 0xaeb
	KEY_club                        = 0xaec
	KEY_diamond                     = 0xaed
	KEY_heart                       = 0xaee
	KEY_maltesecross                = 0xaf0
	KEY_dagger                      = 0xaf1
	KEY_doubledagger                = 0xaf2
	KEY_checkmark                   = 0xaf3
	KEY_ballotcross                 = 0xaf4
	KEY_musicalsharp                = 0xaf5
	KEY_musicalflat                 = 0xaf6
	KEY_malesymbol                  = 0xaf7
	KEY_femalesymbol                = 0xaf8
	KEY_telephone                   = 0xaf9
	KEY_telephonerecorder           = 0xafa
	KEY_phonographcopyright         = 0xafb
	KEY_caret                       = 0xafc
	KEY_singlelowquotemark          = 0xafd
	KEY_doublelowquotemark          = 0xafe
	KEY_cursor                      = 0xaff
	KEY_leftcaret                   = 0xba3
	KEY_rightcaret                  = 0xba6
	KEY_downcaret                   = 0xba8
	KEY_upcaret                     = 0xba9
	KEY_overbar                     = 0xbc0
	KEY_downtack                    = 0xbc2
	KEY_upshoe                      = 0xbc3
	KEY_downstile                   = 0xbc4
	KEY_underbar                    = 0xbc6
	KEY_jot                         = 0xbca
	KEY_quad                        = 0xbcc
	KEY_uptack                      = 0xbce
	KEY_circle                      = 0xbcf
	KEY_upstile                     = 0xbd3
	KEY_downshoe                    = 0xbd6
	KEY_rightshoe                   = 0xbd8
	KEY_leftshoe                    = 0xbda
	KEY_lefttack                    = 0xbdc
	KEY_righttack                   = 0xbfc
	KEY_hebrew_doublelowline        = 0xcdf
	KEY_hebrew_aleph                = 0xce0
	KEY_hebrew_bet                  = 0xce1
	KEY_hebrew_beth                 = 0xce1
	KEY_hebrew_gimel                = 0xce2
	KEY_hebrew_gimmel               = 0xce2
	KEY_hebrew_dalet                = 0xce3
	KEY_hebrew_daleth               = 0xce3
	KEY_hebrew_he                   = 0xce4
	KEY_hebrew_waw                  = 0xce5
	KEY_hebrew_zain                 = 0xce6
	KEY_hebrew_zayin                = 0xce6
	KEY_hebrew_chet                 = 0xce7
	KEY_hebrew_het                  = 0xce7
	KEY_hebrew_tet                  = 0xce8
	KEY_hebrew_teth                 = 0xce8
	KEY_hebrew_yod                  = 0xce9
	KEY_hebrew_finalkaph            = 0xcea
	KEY_hebrew_kaph                 = 0xceb
	KEY_hebrew_lamed                = 0xcec
	KEY_hebrew_finalmem             = 0xced
	KEY_hebrew_mem                  = 0xcee
	KEY_hebrew_finalnun             = 0xcef
	KEY_hebrew_nun                  = 0xcf0
	KEY_hebrew_samech               = 0xcf1
	KEY_hebrew_samekh               = 0xcf1
	KEY_hebrew_ayin                 = 0xcf2
	KEY_hebrew_finalpe              = 0xcf3
	KEY_hebrew_pe                   = 0xcf4
	KEY_hebrew_finalzade            = 0xcf5
	KEY_hebrew_finalzadi            = 0xcf5
	KEY_hebrew_zade                 = 0xcf6
	KEY_hebrew_zadi                 = 0xcf6
	KEY_hebrew_qoph                 = 0xcf7
	KEY_hebrew_kuf                  = 0xcf7
	KEY_hebrew_resh                 = 0xcf8
	KEY_hebrew_shin                 = 0xcf9
	KEY_hebrew_taw                  = 0xcfa
	KEY_hebrew_taf                  = 0xcfa
	KEY_Hebrew_switch               = 0xff7e
	KEY_Thai_kokai                  = 0xda1
	KEY_Thai_khokhai                = 0xda2
	KEY_Thai_khokhuat               = 0xda3
	KEY_Thai_khokhwai               = 0xda4
	KEY_Thai_khokhon                = 0xda5
	KEY_Thai_khorakhang             = 0xda6
	KEY_Thai_ngongu                 = 0xda7
	KEY_Thai_chochan                = 0xda8
	KEY_Thai_choching               = 0xda9
	KEY_Thai_chochang               = 0xdaa
	KEY_Thai_soso                   = 0xdab
	KEY_Thai_chochoe                = 0xdac
	KEY_Thai_yoying                 = 0xdad
	KEY_Thai_dochada                = 0xdae
	KEY_Thai_topatak                = 0xdaf
	KEY_Thai_thothan                = 0xdb0
	KEY_Thai_thonangmontho          = 0xdb1
	KEY_Thai_thophuthao             = 0xdb2
	KEY_Thai_nonen                  = 0xdb3
	KEY_Thai_dodek                  = 0xdb4
	KEY_Thai_totao                  = 0xdb5
	KEY_Thai_thothung               = 0xdb6
	KEY_Thai_thothahan              = 0xdb7
	KEY_Thai_thothong               = 0xdb8
	KEY_Thai_nonu                   = 0xdb9
	KEY_Thai_bobaimai               = 0xdba
	KEY_Thai_popla                  = 0xdbb
	KEY_Thai_phophung               = 0xdbc
	KEY_Thai_fofa                   = 0xdbd
	KEY_Thai_phophan                = 0xdbe
	KEY_Thai_fofan                  = 0xdbf
	KEY_Thai_phosamphao             = 0xdc0
	KEY_Thai_moma                   = 0xdc1
	KEY_Thai_yoyak                  = 0xdc2
	KEY_Thai_rorua                  = 0xdc3
	KEY_Thai_ru                     = 0xdc4
	KEY_Thai_loling                 = 0xdc5
	KEY_Thai_lu                     = 0xdc6
	KEY_Thai_wowaen                 = 0xdc7
	KEY_Thai_sosala                 = 0xdc8
	KEY_Thai_sorusi                 = 0xdc9
	KEY_Thai_sosua                  = 0xdca
	KEY_Thai_hohip                  = 0xdcb
	KEY_Thai_lochula                = 0xdcc
	KEY_Thai_oang                   = 0xdcd
	KEY_Thai_honokhuk               = 0xdce
	KEY_Thai_paiyannoi              = 0xdcf
	KEY_Thai_saraa                  = 0xdd0
	KEY_Thai_maihanakat             = 0xdd1
	KEY_Thai_saraaa                 = 0xdd2
	KEY_Thai_saraam                 = 0xdd3
	KEY_Thai_sarai                  = 0xdd4
	KEY_Thai_saraii                 = 0xdd5
	KEY_Thai_saraue                 = 0xdd6
	KEY_Thai_sarauee                = 0xdd7
	KEY_Thai_sarau                  = 0xdd8
	KEY_Thai_sarauu                 = 0xdd9
	KEY_Thai_phinthu                = 0xdda
	KEY_Thai_maihanakat_maitho      = 0xdde
	KEY_Thai_baht                   = 0xddf
	KEY_Thai_sarae                  = 0xde0
	KEY_Thai_saraae                 = 0xde1
	KEY_Thai_sarao                  = 0xde2
	KEY_Thai_saraaimaimuan          = 0xde3
	KEY_Thai_saraaimaimalai         = 0xde4
	KEY_Thai_lakkhangyao            = 0xde5
	KEY_Thai_maiyamok               = 0xde6
	KEY_Thai_maitaikhu              = 0xde7
	KEY_Thai_maiek                  = 0xde8
	KEY_Thai_maitho                 = 0xde9
	KEY_Thai_maitri                 = 0xdea
	KEY_Thai_maichattawa            = 0xdeb
	KEY_Thai_thanthakhat            = 0xdec
	KEY_Thai_nikhahit               = 0xded
	KEY_Thai_leksun                 = 0xdf0
	KEY_Thai_leknung                = 0xdf1
	KEY_Thai_leksong                = 0xdf2
	KEY_Thai_leksam                 = 0xdf3
	KEY_Thai_leksi                  = 0xdf4
	KEY_Thai_lekha                  = 0xdf5
	KEY_Thai_lekhok                 = 0xdf6
	KEY_Thai_lekchet                = 0xdf7
	KEY_Thai_lekpaet                = 0xdf8
	KEY_Thai_lekkao                 = 0xdf9
	KEY_Hangul                      = 0xff31
	KEY_Hangul_Start                = 0xff32
	KEY_Hangul_End                  = 0xff33
	KEY_Hangul_Hanja                = 0xff34
	KEY_Hangul_Jamo                 = 0xff35
	KEY_Hangul_Romaja               = 0xff36
	KEY_Hangul_Codeinput            = 0xff37
	KEY_Hangul_Jeonja               = 0xff38
	KEY_Hangul_Banja                = 0xff39
	KEY_Hangul_PreHanja             = 0xff3a
	KEY_Hangul_PostHanja            = 0xff3b
	KEY_Hangul_SingleCandidate      = 0xff3c
	KEY_Hangul_MultipleCandidate    = 0xff3d
	KEY_Hangul_PreviousCandidate    = 0xff3e
	KEY_Hangul_Special              = 0xff3f
	KEY_Hangul_switch               = 0xff7e
	KEY_Hangul_Kiyeog               = 0xea1
	KEY_Hangul_SsangKiyeog          = 0xea2
	KEY_Hangul_KiyeogSios           = 0xea3
	KEY_Hangul_Nieun                = 0xea4
	KEY_Hangul_NieunJieuj           = 0xea5
	KEY_Hangul_NieunHieuh           = 0xea6
	KEY_Hangul_Dikeud               = 0xea7
	KEY_Hangul_SsangDikeud          = 0xea8
	KEY_Hangul_Rieul                = 0xea9
	KEY_Hangul_RieulKiyeog          = 0xeaa
	KEY_Hangul_RieulMieum           = 0xeab
	KEY_Hangul_RieulPieub           = 0xeac
	KEY_Hangul_RieulSios            = 0xead
	KEY_Hangul_RieulTieut           = 0xeae
	KEY_Hangul_RieulPhieuf          = 0xeaf
	KEY_Hangul_RieulHieuh           = 0xeb0
	KEY_Hangul_Mieum                = 0xeb1
	KEY_Hangul_Pieub                = 0xeb2
	KEY_Hangul_SsangPieub           = 0xeb3
	KEY_Hangul_PieubSios            = 0xeb4
	KEY_Hangul_Sios                 = 0xeb5
	KEY_Hangul_SsangSios            = 0xeb6
	KEY_Hangul_Ieung                = 0xeb7
	KEY_Hangul_Jieuj                = 0xeb8
	KEY_Hangul_SsangJieuj           = 0xeb9
	KEY_Hangul_Cieuc                = 0xeba
	KEY_Hangul_Khieuq               = 0xebb
	KEY_Hangul_Tieut                = 0xebc
	KEY_Hangul_Phieuf               = 0xebd
	KEY_Hangul_Hieuh                = 0xebe
	KEY_Hangul_A                    = 0xebf
	KEY_Hangul_AE                   = 0xec0
	KEY_Hangul_YA                   = 0xec1
	KEY_Hangul_YAE                  = 0xec2
	KEY_Hangul_EO                   = 0xec3
	KEY_Hangul_E                    = 0xec4
	KEY_Hangul_YEO                  = 0xec5
	KEY_Hangul_YE                   = 0xec6
	KEY_Hangul_O                    = 0xec7
	KEY_Hangul_WA                   = 0xec8
	KEY_Hangul_WAE                  = 0xec9
	KEY_Hangul_OE                   = 0xeca
	KEY_Hangul_YO                   = 0xecb
	KEY_Hangul_U                    = 0xecc
	KEY_Hangul_WEO                  = 0xecd
	KEY_Hangul_WE                   = 0xece
	KEY_Hangul_WI                   = 0xecf
	KEY_Hangul_YU                   = 0xed0
	KEY_Hangul_EU                   = 0xed1
	KEY_Hangul_YI                   = 0xed2
	KEY_Hangul_I                    = 0xed3
	KEY_Hangul_J_Kiyeog             = 0xed4
	KEY_Hangul_J_SsangKiyeog        = 0xed5
	KEY_Hangul_J_KiyeogSios         = 0xed6
	KEY_Hangul_J_Nieun              = 0xed7
	KEY_Hangul_J_NieunJieuj         = 0xed8
	KEY_Hangul_J_NieunHieuh         = 0xed9
	KEY_Hangul_J_Dikeud             = 0xeda
	KEY_Hangul_J_Rieul              = 0xedb
	KEY_Hangul_J_RieulKiyeog        = 0xedc
	KEY_Hangul_J_RieulMieum         = 0xedd
	KEY_Hangul_J_RieulPieub         = 0xede
	KEY_Hangul_J_RieulSios          = 0xedf
	KEY_Hangul_J_RieulTieut         = 0xee0
	KEY_Hangul_J_RieulPhieuf        = 0xee1
	KEY_Hangul_J_RieulHieuh         = 0xee2
	KEY_Hangul_J_Mieum              = 0xee3
	KEY_Hangul_J_Pieub              = 0xee4
	KEY_Hangul_J_PieubSios          = 0xee5
	KEY_Hangul_J_Sios               = 0xee6
	KEY_Hangul_J_SsangSios          = 0xee7
	KEY_Hangul_J_Ieung              = 0xee8
	KEY_Hangul_J_Jieuj              = 0xee9
	KEY_Hangul_J_Cieuc              = 0xeea
	KEY_Hangul_J_Khieuq             = 0xeeb
	KEY_Hangul_J_Tieut              = 0xeec
	KEY_Hangul_J_Phieuf             = 0xeed
	KEY_Hangul_J_Hieuh              = 0xeee
	KEY_Hangul_RieulYeorinHieuh     = 0xeef
	KEY_Hangul_SunkyeongeumMieum    = 0xef0
	KEY_Hangul_SunkyeongeumPieub    = 0xef1
	KEY_Hangul_PanSios              = 0xef2
	KEY_Hangul_KkogjiDalrinIeung    = 0xef3
	KEY_Hangul_SunkyeongeumPhieuf   = 0xef4
	KEY_Hangul_YeorinHieuh          = 0xef5
	KEY_Hangul_AraeA                = 0xef6
	KEY_Hangul_AraeAE               = 0xef7
	KEY_Hangul_J_PanSios            = 0xef8
	KEY_Hangul_J_KkogjiDalrinIeung  = 0xef9
	KEY_Hangul_J_YeorinHieuh        = 0xefa
	KEY_Korean_Won                  = 0xeff
	KEY_Armenian_ligature_ew        = 0x1000587
	KEY_Armenian_full_stop          = 0x1000589
	KEY_Armenian_verjaket           = 0x1000589
	KEY_Armenian_separation_mark    = 0x100055d
	KEY_Armenian_but                = 0x100055d
	KEY_Armenian_hyphen             = 0x100058a
	KEY_Armenian_yentamna           = 0x100058a
	KEY_Armenian_exclam             = 0x100055c
	KEY_Armenian_amanak             = 0x100055c
	KEY_Armenian_accent             = 0x100055b
	KEY_Armenian_shesht             = 0x100055b
	KEY_Armenian_question           = 0x100055e
	KEY_Armenian_paruyk             = 0x100055e
	KEY_Armenian_AYB                = 0x1000531
	KEY_Armenian_ayb                = 0x1000561
	KEY_Armenian_BEN                = 0x1000532
	KEY_Armenian_ben                = 0x1000562
	KEY_Armenian_GIM                = 0x1000533
	KEY_Armenian_gim                = 0x1000563
	KEY_Armenian_DA                 = 0x1000534
	KEY_Armenian_da                 = 0x1000564
	KEY_Armenian_YECH               = 0x1000535
	KEY_Armenian_yech               = 0x1000565
	KEY_Armenian_ZA                 = 0x1000536
	KEY_Armenian_za                 = 0x1000566
	KEY_Armenian_E                  = 0x1000537
	KEY_Armenian_e                  = 0x1000567
	KEY_Armenian_AT                 = 0x1000538
	KEY_Armenian_at                 = 0x1000568
	KEY_Armenian_TO                 = 0x1000539
	KEY_Armenian_to                 = 0x1000569
	KEY_Armenian_ZHE                = 0x100053a
	KEY_Armenian_zhe                = 0x100056a
	KEY_Armenian_INI                = 0x100053b
	KEY_Armenian_ini                = 0x100056b
	KEY_Armenian_LYUN               = 0x100053c
	KEY_Armenian_lyun               = 0x100056c
	KEY_Armenian_KHE                = 0x100053d
	KEY_Armenian_khe                = 0x100056d
	KEY_Armenian_TSA                = 0x100053e
	KEY_Armenian_tsa                = 0x100056e
	KEY_Armenian_KEN                = 0x100053f
	KEY_Armenian_ken                = 0x100056f
	KEY_Armenian_HO                 = 0x1000540
	KEY_Armenian_ho                 = 0x1000570
	KEY_Armenian_DZA                = 0x1000541
	KEY_Armenian_dza                = 0x1000571
	KEY_Armenian_GHAT               = 0x1000542
	KEY_Armenian_ghat               = 0x1000572
	KEY_Armenian_TCHE               = 0x1000543
	KEY_Armenian_tche               = 0x1000573
	KEY_Armenian_MEN                = 0x1000544
	KEY_Armenian_men                = 0x1000574
	KEY_Armenian_HI                 = 0x1000545
	KEY_Armenian_hi                 = 0x1000575
	KEY_Armenian_NU                 = 0x1000546
	KEY_Armenian_nu                 = 0x1000576
	KEY_Armenian_SHA                = 0x1000547
	KEY_Armenian_sha                = 0x1000577
	KEY_Armenian_VO                 = 0x1000548
	KEY_Armenian_vo                 = 0x1000578
	KEY_Armenian_CHA                = 0x1000549
	KEY_Armenian_cha                = 0x1000579
	KEY_Armenian_PE                 = 0x100054a
	KEY_Armenian_pe                 = 0x100057a
	KEY_Armenian_JE                 = 0x100054b
	KEY_Armenian_je                 = 0x100057b
	KEY_Armenian_RA                 = 0x100054c
	KEY_Armenian_ra                 = 0x100057c
	KEY_Armenian_SE                 = 0x100054d
	KEY_Armenian_se                 = 0x100057d
	KEY_Armenian_VEV                = 0x100054e
	KEY_Armenian_vev                = 0x100057e
	KEY_Armenian_TYUN               = 0x100054f
	KEY_Armenian_tyun               = 0x100057f
	KEY_Armenian_RE                 = 0x1000550
	KEY_Armenian_re                 = 0x1000580
	KEY_Armenian_TSO                = 0x1000551
	KEY_Armenian_tso                = 0x1000581
	KEY_Armenian_VYUN               = 0x1000552
	KEY_Armenian_vyun               = 0x1000582
	KEY_Armenian_PYUR               = 0x1000553
	KEY_Armenian_pyur               = 0x1000583
	KEY_Armenian_KE                 = 0x1000554
	KEY_Armenian_ke                 = 0x1000584
	KEY_Armenian_O                  = 0x1000555
	KEY_Armenian_o                  = 0x1000585
	KEY_Armenian_FE                 = 0x1000556
	KEY_Armenian_fe                 = 0x1000586
	KEY_Armenian_apostrophe         = 0x100055a
	KEY_Georgian_an                 = 0x10010d0
	KEY_Georgian_ban                = 0x10010d1
	KEY_Georgian_gan                = 0x10010d2
	KEY_Georgian_don                = 0x10010d3
	KEY_Georgian_en                 = 0x10010d4
	KEY_Georgian_vin                = 0x10010d5
	KEY_Georgian_zen                = 0x10010d6
	KEY_Georgian_tan                = 0x10010d7
	KEY_Georgian_in                 = 0x10010d8
	KEY_Georgian_kan                = 0x10010d9
	KEY_Georgian_las                = 0x10010da
	KEY_Georgian_man                = 0x10010db
	KEY_Georgian_nar                = 0x10010dc
	KEY_Georgian_on                 = 0x10010dd
	KEY_Georgian_par                = 0x10010de
	KEY_Georgian_zhar               = 0x10010df
	KEY_Georgian_rae                = 0x10010e0
	KEY_Georgian_san                = 0x10010e1
	KEY_Georgian_tar                = 0x10010e2
	KEY_Georgian_un                 = 0x10010e3
	KEY_Georgian_phar               = 0x10010e4
	KEY_Georgian_khar               = 0x10010e5
	KEY_Georgian_ghan               = 0x10010e6
	KEY_Georgian_qar                = 0x10010e7
	KEY_Georgian_shin               = 0x10010e8
	KEY_Georgian_chin               = 0x10010e9
	KEY_Georgian_can                = 0x10010ea
	KEY_Georgian_jil                = 0x10010eb
	KEY_Georgian_cil                = 0x10010ec
	KEY_Georgian_char               = 0x10010ed
	KEY_Georgian_xan                = 0x10010ee
	KEY_Georgian_jhan               = 0x10010ef
	KEY_Georgian_hae                = 0x10010f0
	KEY_Georgian_he                 = 0x10010f1
	KEY_Georgian_hie                = 0x10010f2
	KEY_Georgian_we                 = 0x10010f3
	KEY_Georgian_har                = 0x10010f4
	KEY_Georgian_hoe                = 0x10010f5
	KEY_Georgian_fi                 = 0x10010f6
	KEY_Xabovedot                   = 0x1001e8a
	KEY_Ibreve                      = 0x100012c
	KEY_Zstroke                     = 0x10001b5
	KEY_Gcaron                      = 0x10001e6
	KEY_Ocaron                      = 0x10001d1
	KEY_Obarred                     = 0x100019f
	KEY_xabovedot                   = 0x1001e8b
	KEY_ibreve                      = 0x100012d
	KEY_zstroke                     = 0x10001b6
	KEY_gcaron                      = 0x10001e7
	KEY_ocaron                      = 0x10001d2
	KEY_obarred                     = 0x1000275
	KEY_SCHWA                       = 0x100018f
	KEY_schwa                       = 0x1000259
	KEY_Lbelowdot                   = 0x1001e36
	KEY_lbelowdot                   = 0x1001e37
	KEY_Abelowdot                   = 0x1001ea0
	KEY_abelowdot                   = 0x1001ea1
	KEY_Ahook                       = 0x1001ea2
	KEY_ahook                       = 0x1001ea3
	KEY_Acircumflexacute            = 0x1001ea4
	KEY_acircumflexacute            = 0x1001ea5
	KEY_Acircumflexgrave            = 0x1001ea6
	KEY_acircumflexgrave            = 0x1001ea7
	KEY_Acircumflexhook             = 0x1001ea8
	KEY_acircumflexhook             = 0x1001ea9
	KEY_Acircumflextilde            = 0x1001eaa
	KEY_acircumflextilde            = 0x1001eab
	KEY_Acircumflexbelowdot         = 0x1001eac
	KEY_acircumflexbelowdot         = 0x1001ead
	KEY_Abreveacute                 = 0x1001eae
	KEY_abreveacute                 = 0x1001eaf
	KEY_Abrevegrave                 = 0x1001eb0
	KEY_abrevegrave                 = 0x1001eb1
	KEY_Abrevehook                  = 0x1001eb2
	KEY_abrevehook                  = 0x1001eb3
	KEY_Abrevetilde                 = 0x1001eb4
	KEY_abrevetilde                 = 0x1001eb5
	KEY_Abrevebelowdot              = 0x1001eb6
	KEY_abrevebelowdot              = 0x1001eb7
	KEY_Ebelowdot                   = 0x1001eb8
	KEY_ebelowdot                   = 0x1001eb9
	KEY_Ehook                       = 0x1001eba
	KEY_ehook                       = 0x1001ebb
	KEY_Etilde                      = 0x1001ebc
	KEY_etilde                      = 0x1001ebd
	KEY_Ecircumflexacute            = 0x1001ebe
	KEY_ecircumflexacute            = 0x1001ebf
	KEY_Ecircumflexgrave            = 0x1001ec0
	KEY_ecircumflexgrave            = 0x1001ec1
	KEY_Ecircumflexhook             = 0x1001ec2
	KEY_ecircumflexhook             = 0x1001ec3
	KEY_Ecircumflextilde            = 0x1001ec4
	KEY_ecircumflextilde            = 0x1001ec5
	KEY_Ecircumflexbelowdot         = 0x1001ec6
	KEY_ecircumflexbelowdot         = 0x1001ec7
	KEY_Ihook                       = 0x1001ec8
	KEY_ihook                       = 0x1001ec9
	KEY_Ibelowdot                   = 0x1001eca
	KEY_ibelowdot                   = 0x1001ecb
	KEY_Obelowdot                   = 0x1001ecc
	KEY_obelowdot                   = 0x1001ecd
	KEY_Ohook                       = 0x1001ece
	KEY_ohook                       = 0x1001ecf
	KEY_Ocircumflexacute            = 0x1001ed0
	KEY_ocircumflexacute            = 0x1001ed1
	KEY_Ocircumflexgrave            = 0x1001ed2
	KEY_ocircumflexgrave            = 0x1001ed3
	KEY_Ocircumflexhook             = 0x1001ed4
	KEY_ocircumflexhook             = 0x1001ed5
	KEY_Ocircumflextilde            = 0x1001ed6
	KEY_ocircumflextilde            = 0x1001ed7
	KEY_Ocircumflexbelowdot         = 0x1001ed8
	KEY_ocircumflexbelowdot         = 0x1001ed9
	KEY_Ohornacute                  = 0x1001eda
	KEY_ohornacute                  = 0x1001edb
	KEY_Ohorngrave                  = 0x1001edc
	KEY_ohorngrave                  = 0x1001edd
	KEY_Ohornhook                   = 0x1001ede
	KEY_ohornhook                   = 0x1001edf
	KEY_Ohorntilde                  = 0x1001ee0
	KEY_ohorntilde                  = 0x1001ee1
	KEY_Ohornbelowdot               = 0x1001ee2
	KEY_ohornbelowdot               = 0x1001ee3
	KEY_Ubelowdot                   = 0x1001ee4
	KEY_ubelowdot                   = 0x1001ee5
	KEY_Uhook                       = 0x1001ee6
	KEY_uhook                       = 0x1001ee7
	KEY_Uhornacute                  = 0x1001ee8
	KEY_uhornacute                  = 0x1001ee9
	KEY_Uhorngrave                  = 0x1001eea
	KEY_uhorngrave                  = 0x1001eeb
	KEY_Uhornhook                   = 0x1001eec
	KEY_uhornhook                   = 0x1001eed
	KEY_Uhorntilde                  = 0x1001eee
	KEY_uhorntilde                  = 0x1001eef
	KEY_Uhornbelowdot               = 0x1001ef0
	KEY_uhornbelowdot               = 0x1001ef1
	KEY_Ybelowdot                   = 0x1001ef4
	KEY_ybelowdot                   = 0x1001ef5
	KEY_Yhook                       = 0x1001ef6
	KEY_yhook                       = 0x1001ef7
	KEY_Ytilde                      = 0x1001ef8
	KEY_ytilde                      = 0x1001ef9
	KEY_Ohorn                       = 0x10001a0
	KEY_ohorn                       = 0x10001a1
	KEY_Uhorn                       = 0x10001af
	KEY_uhorn                       = 0x10001b0
	KEY_EcuSign                     = 0x10020a0
	KEY_ColonSign                   = 0x10020a1
	KEY_CruzeiroSign                = 0x10020a2
	KEY_FFrancSign                  = 0x10020a3
	KEY_LiraSign                    = 0x10020a4
	KEY_MillSign                    = 0x10020a5
	KEY_NairaSign                   = 0x10020a6
	KEY_PesetaSign                  = 0x10020a7
	KEY_RupeeSign                   = 0x10020a8
	KEY_WonSign                     = 0x10020a9
	KEY_NewSheqelSign               = 0x10020aa
	KEY_DongSign                    = 0x10020ab
	KEY_EuroSign                    = 0x20ac
	KEY_zerosuperior                = 0x1002070
	KEY_foursuperior                = 0x1002074
	KEY_fivesuperior                = 0x1002075
	KEY_sixsuperior                 = 0x1002076
	KEY_sevensuperior               = 0x1002077
	KEY_eightsuperior               = 0x1002078
	KEY_ninesuperior                = 0x1002079
	KEY_zerosubscript               = 0x1002080
	KEY_onesubscript                = 0x1002081
	KEY_twosubscript                = 0x1002082
	KEY_threesubscript              = 0x1002083
	KEY_foursubscript               = 0x1002084
	KEY_fivesubscript               = 0x1002085
	KEY_sixsubscript                = 0x1002086
	KEY_sevensubscript              = 0x1002087
	KEY_eightsubscript              = 0x1002088
	KEY_ninesubscript               = 0x1002089
	KEY_partdifferential            = 0x1002202
	KEY_emptyset                    = 0x1002205
	KEY_elementof                   = 0x1002208
	KEY_notelementof                = 0x1002209
	KEY_containsas                  = 0x100220b
	KEY_squareroot                  = 0x100221a
	KEY_cuberoot                    = 0x100221b
	KEY_fourthroot                  = 0x100221c
	KEY_dintegral                   = 0x100222c
	KEY_tintegral                   = 0x100222d
	KEY_because                     = 0x1002235
	KEY_approxeq                    = 0x1002248
	KEY_notapproxeq                 = 0x1002247
	KEY_notidentical                = 0x1002262
	KEY_stricteq                    = 0x1002263
	KEY_braille_dot_1               = 0xfff1
	KEY_braille_dot_2               = 0xfff2
	KEY_braille_dot_3               = 0xfff3
	KEY_braille_dot_4               = 0xfff4
	KEY_braille_dot_5               = 0xfff5
	KEY_braille_dot_6               = 0xfff6
	KEY_braille_dot_7               = 0xfff7
	KEY_braille_dot_8               = 0xfff8
	KEY_braille_dot_9               = 0xfff9
	KEY_braille_dot_10              = 0xfffa
	KEY_braille_blank               = 0x1002800
	KEY_braille_dots_1              = 0x1002801
	KEY_braille_dots_2              = 0x1002802
	KEY_braille_dots_12             = 0x1002803
	KEY_braille_dots_3              = 0x1002804
	KEY_braille_dots_13             = 0x1002805
	KEY_braille_dots_23             = 0x1002806
	KEY_braille_dots_123            = 0x1002807
	KEY_braille_dots_4              = 0x1002808
	KEY_braille_dots_14             = 0x1002809
	KEY_braille_dots_24             = 0x100280a
	KEY_braille_dots_124            = 0x100280b
	KEY_braille_dots_34             = 0x100280c
	KEY_braille_dots_134            = 0x100280d
	KEY_braille_dots_234            = 0x100280e
	KEY_braille_dots_1234           = 0x100280f
	KEY_braille_dots_5              = 0x1002810
	KEY_braille_dots_15             = 0x1002811
	KEY_braille_dots_25             = 0x1002812
	KEY_braille_dots_125            = 0x1002813
	KEY_braille_dots_35             = 0x1002814
	KEY_braille_dots_135            = 0x1002815
	KEY_braille_dots_235            = 0x1002816
	KEY_braille_dots_1235           = 0x1002817
	KEY_braille_dots_45             = 0x1002818
	KEY_braille_dots_145            = 0x1002819
	KEY_braille_dots_245            = 0x100281a
	KEY_braille_dots_1245           = 0x100281b
	KEY_braille_dots_345            = 0x100281c
	KEY_braille_dots_1345           = 0x100281d
	KEY_braille_dots_2345           = 0x100281e
	KEY_braille_dots_12345          = 0x100281f
	KEY_braille_dots_6              = 0x1002820
	KEY_braille_dots_16             = 0x1002821
	KEY_braille_dots_26             = 0x1002822
	KEY_braille_dots_126            = 0x1002823
	KEY_braille_dots_36             = 0x1002824
	KEY_braille_dots_136            = 0x1002825
	KEY_braille_dots_236            = 0x1002826
	KEY_braille_dots_1236           = 0x1002827
	KEY_braille_dots_46             = 0x1002828
	KEY_braille_dots_146            = 0x1002829
	KEY_braille_dots_246            = 0x100282a
	KEY_braille_dots_1246           = 0x100282b
	KEY_braille_dots_346            = 0x100282c
	KEY_braille_dots_1346           = 0x100282d
	KEY_braille_dots_2346           = 0x100282e
	KEY_braille_dots_12346          = 0x100282f
	KEY_braille_dots_56             = 0x1002830
	KEY_braille_dots_156            = 0x1002831
	KEY_braille_dots_256            = 0x1002832
	KEY_braille_dots_1256           = 0x1002833
	KEY_braille_dots_356            = 0x1002834
	KEY_braille_dots_1356           = 0x1002835
	KEY_braille_dots_2356           = 0x1002836
	KEY_braille_dots_12356          = 0x1002837
	KEY_braille_dots_456            = 0x1002838
	KEY_braille_dots_1456           = 0x1002839
	KEY_braille_dots_2456           = 0x100283a
	KEY_braille_dots_12456          = 0x100283b
	KEY_braille_dots_3456           = 0x100283c
	KEY_braille_dots_13456          = 0x100283d
	KEY_braille_dots_23456          = 0x100283e
	KEY_braille_dots_123456         = 0x100283f
	KEY_braille_dots_7              = 0x1002840
	KEY_braille_dots_17             = 0x1002841
	KEY_braille_dots_27             = 0x1002842
	KEY_braille_dots_127            = 0x1002843
	KEY_braille_dots_37             = 0x1002844
	KEY_braille_dots_137            = 0x1002845
	KEY_braille_dots_237            = 0x1002846
	KEY_braille_dots_1237           = 0x1002847
	KEY_braille_dots_47             = 0x1002848
	KEY_braille_dots_147            = 0x1002849
	KEY_braille_dots_247            = 0x100284a
	KEY_braille_dots_1247           = 0x100284b
	KEY_braille_dots_347            = 0x100284c
	KEY_braille_dots_1347           = 0x100284d
	KEY_braille_dots_2347           = 0x100284e
	KEY_braille_dots_12347          = 0x100284f
	KEY_braille_dots_57             = 0x1002850
	KEY_braille_dots_157            = 0x1002851
	KEY_braille_dots_257            = 0x1002852
	KEY_braille_dots_1257           = 0x1002853
	KEY_braille_dots_357            = 0x1002854
	KEY_braille_dots_1357           = 0x1002855
	KEY_braille_dots_2357           = 0x1002856
	KEY_braille_dots_12357          = 0x1002857
	KEY_braille_dots_457            = 0x1002858
	KEY_braille_dots_1457           = 0x1002859
	KEY_braille_dots_2457           = 0x100285a
	KEY_braille_dots_12457          = 0x100285b
	KEY_braille_dots_3457           = 0x100285c
	KEY_braille_dots_13457          = 0x100285d
	KEY_braille_dots_23457          = 0x100285e
	KEY_braille_dots_123457         = 0x100285f
	KEY_braille_dots_67             = 0x1002860
	KEY_braille_dots_167            = 0x1002861
	KEY_braille_dots_267            = 0x1002862
	KEY_braille_dots_1267           = 0x1002863
	KEY_braille_dots_367            = 0x1002864
	KEY_braille_dots_1367           = 0x1002865
	KEY_braille_dots_2367           = 0x1002866
	KEY_braille_dots_12367          = 0x1002867
	KEY_braille_dots_467            = 0x1002868
	KEY_braille_dots_1467           = 0x1002869
	KEY_braille_dots_2467           = 0x100286a
	KEY_braille_dots_12467          = 0x100286b
	KEY_braille_dots_3467           = 0x100286c
	KEY_braille_dots_13467          = 0x100286d
	KEY_braille_dots_23467          = 0x100286e
	KEY_braille_dots_123467         = 0x100286f
	KEY_braille_dots_567            = 0x1002870
	KEY_braille_dots_1567           = 0x1002871
	KEY_braille_dots_2567           = 0x1002872
	KEY_braille_dots_12567          = 0x1002873
	KEY_braille_dots_3567           = 0x1002874
	KEY_braille_dots_13567          = 0x1002875
	KEY_braille_dots_23567          = 0x1002876
	KEY_braille_dots_123567         = 0x1002877
	KEY_braille_dots_4567           = 0x1002878
	KEY_braille_dots_14567          = 0x1002879
	KEY_braille_dots_24567          = 0x100287a
	KEY_braille_dots_124567         = 0x100287b
	KEY_braille_dots_34567          = 0x100287c
	KEY_braille_dots_134567         = 0x100287d
	KEY_braille_dots_234567         = 0x100287e
	KEY_braille_dots_1234567        = 0x100287f
	KEY_braille_dots_8              = 0x1002880
	KEY_braille_dots_18             = 0x1002881
	KEY_braille_dots_28             = 0x1002882
	KEY_braille_dots_128            = 0x1002883
	KEY_braille_dots_38             = 0x1002884
	KEY_braille_dots_138            = 0x1002885
	KEY_braille_dots_238            = 0x1002886
	KEY_braille_dots_1238           = 0x1002887
	KEY_braille_dots_48             = 0x1002888
	KEY_braille_dots_148            = 0x1002889
	KEY_braille_dots_248            = 0x100288a
	KEY_braille_dots_1248           = 0x100288b
	KEY_braille_dots_348            = 0x100288c
	KEY_braille_dots_1348           = 0x100288d
	KEY_braille_dots_2348           = 0x100288e
	KEY_braille_dots_12348          = 0x100288f
	KEY_braille_dots_58             = 0x1002890
	KEY_braille_dots_158            = 0x1002891
	KEY_braille_dots_258            = 0x1002892
	KEY_braille_dots_1258           = 0x1002893
	KEY_braille_dots_358            = 0x1002894
	KEY_braille_dots_1358           = 0x1002895
	KEY_braille_dots_2358           = 0x1002896
	KEY_braille_dots_12358          = 0x1002897
	KEY_braille_dots_458            = 0x1002898
	KEY_braille_dots_1458           = 0x1002899
	KEY_braille_dots_2458           = 0x100289a
	KEY_braille_dots_12458          = 0x100289b
	KEY_braille_dots_3458           = 0x100289c
	KEY_braille_dots_13458          = 0x100289d
	KEY_braille_dots_23458          = 0x100289e
	KEY_braille_dots_123458         = 0x100289f
	KEY_braille_dots_68             = 0x10028a0
	KEY_braille_dots_168            = 0x10028a1
	KEY_braille_dots_268            = 0x10028a2
	KEY_braille_dots_1268           = 0x10028a3
	KEY_braille_dots_368            = 0x10028a4
	KEY_braille_dots_1368           = 0x10028a5
	KEY_braille_dots_2368           = 0x10028a6
	KEY_braille_dots_12368          = 0x10028a7
	KEY_braille_dots_468            = 0x10028a8
	KEY_braille_dots_1468           = 0x10028a9
	KEY_braille_dots_2468           = 0x10028aa
	KEY_braille_dots_12468          = 0x10028ab
	KEY_braille_dots_3468           = 0x10028ac
	KEY_braille_dots_13468          = 0x10028ad
	KEY_braille_dots_23468          = 0x10028ae
	KEY_braille_dots_123468         = 0x10028af
	KEY_braille_dots_568            = 0x10028b0
	KEY_braille_dots_1568           = 0x10028b1
	KEY_braille_dots_2568           = 0x10028b2
	KEY_braille_dots_12568          = 0x10028b3
	KEY_braille_dots_3568           = 0x10028b4
	KEY_braille_dots_13568          = 0x10028b5
	KEY_braille_dots_23568          = 0x10028b6
	KEY_braille_dots_123568         = 0x10028b7
	KEY_braille_dots_4568           = 0x10028b8
	KEY_braille_dots_14568          = 0x10028b9
	KEY_braille_dots_24568          = 0x10028ba
	KEY_braille_dots_124568         = 0x10028bb
	KEY_braille_dots_34568          = 0x10028bc
	KEY_braille_dots_134568         = 0x10028bd
	KEY_braille_dots_234568         = 0x10028be
	KEY_braille_dots_1234568        = 0x10028bf
	KEY_braille_dots_78             = 0x10028c0
	KEY_braille_dots_178            = 0x10028c1
	KEY_braille_dots_278            = 0x10028c2
	KEY_braille_dots_1278           = 0x10028c3
	KEY_braille_dots_378            = 0x10028c4
	KEY_braille_dots_1378           = 0x10028c5
	KEY_braille_dots_2378           = 0x10028c6
	KEY_braille_dots_12378          = 0x10028c7
	KEY_braille_dots_478            = 0x10028c8
	KEY_braille_dots_1478           = 0x10028c9
	KEY_braille_dots_2478           = 0x10028ca
	KEY_braille_dots_12478          = 0x10028cb
	KEY_braille_dots_3478           = 0x10028cc
	KEY_braille_dots_13478          = 0x10028cd
	KEY_braille_dots_23478          = 0x10028ce
	KEY_braille_dots_123478         = 0x10028cf
	KEY_braille_dots_578            = 0x10028d0
	KEY_braille_dots_1578           = 0x10028d1
	KEY_braille_dots_2578           = 0x10028d2
	KEY_braille_dots_12578          = 0x10028d3
	KEY_braille_dots_3578           = 0x10028d4
	KEY_braille_dots_13578          = 0x10028d5
	KEY_braille_dots_23578          = 0x10028d6
	KEY_braille_dots_123578         = 0x10028d7
	KEY_braille_dots_4578           = 0x10028d8
	KEY_braille_dots_14578          = 0x10028d9
	KEY_braille_dots_24578          = 0x10028da
	KEY_braille_dots_124578         = 0x10028db
	KEY_braille_dots_34578          = 0x10028dc
	KEY_braille_dots_134578         = 0x10028dd
	KEY_braille_dots_234578         = 0x10028de
	KEY_braille_dots_1234578        = 0x10028df
	KEY_braille_dots_678            = 0x10028e0
	KEY_braille_dots_1678           = 0x10028e1
	KEY_braille_dots_2678           = 0x10028e2
	KEY_braille_dots_12678          = 0x10028e3
	KEY_braille_dots_3678           = 0x10028e4
	KEY_braille_dots_13678          = 0x10028e5
	KEY_braille_dots_23678          = 0x10028e6
	KEY_braille_dots_123678         = 0x10028e7
	KEY_braille_dots_4678           = 0x10028e8
	KEY_braille_dots_14678          = 0x10028e9
	KEY_braille_dots_24678          = 0x10028ea
	KEY_braille_dots_124678         = 0x10028eb
	KEY_braille_dots_34678          = 0x10028ec
	KEY_braille_dots_134678         = 0x10028ed
	KEY_braille_dots_234678         = 0x10028ee
	KEY_braille_dots_1234678        = 0x10028ef
	KEY_braille_dots_5678           = 0x10028f0
	KEY_braille_dots_15678          = 0x10028f1
	KEY_braille_dots_25678          = 0x10028f2
	KEY_braille_dots_125678         = 0x10028f3
	KEY_braille_dots_35678          = 0x10028f4
	KEY_braille_dots_135678         = 0x10028f5
	KEY_braille_dots_235678         = 0x10028f6
	KEY_braille_dots_1235678        = 0x10028f7
	KEY_braille_dots_45678          = 0x10028f8
	KEY_braille_dots_145678         = 0x10028f9
	KEY_braille_dots_245678         = 0x10028fa
	KEY_braille_dots_1245678        = 0x10028fb
	KEY_braille_dots_345678         = 0x10028fc
	KEY_braille_dots_1345678        = 0x10028fd
	KEY_braille_dots_2345678        = 0x10028fe
	KEY_braille_dots_12345678       = 0x10028ff
	KEY_ModeLock                    = 0x1008ff01
	KEY_MonBrightnessUp             = 0x1008ff02
	KEY_MonBrightnessDown           = 0x1008ff03
	KEY_KbdLightOnOff               = 0x1008ff04
	KEY_KbdBrightnessUp             = 0x1008ff05
	KEY_KbdBrightnessDown           = 0x1008ff06
	KEY_Standby                     = 0x1008ff10
	KEY_AudioLowerVolume            = 0x1008ff11
	KEY_AudioMute                   = 0x1008ff12
	KEY_AudioRaiseVolume            = 0x1008ff13
	KEY_AudioPlay                   = 0x1008ff14
	KEY_AudioStop                   = 0x1008ff15
	KEY_AudioPrev                   = 0x1008ff16
	KEY_AudioNext                   = 0x1008ff17
	KEY_HomePage                    = 0x1008ff18
	KEY_Mail                        = 0x1008ff19
	KEY_Start                       = 0x1008ff1a
	KEY_Search                      = 0x1008ff1b
	KEY_AudioRecord                 = 0x1008ff1c
	KEY_Calculator                  = 0x1008ff1d
	KEY_Memo                        = 0x1008ff1e
	KEY_ToDoList                    = 0x1008ff1f
	KEY_Calendar                    = 0x1008ff20
	KEY_PowerDown                   = 0x1008ff21
	KEY_ContrastAdjust              = 0x1008ff22
	KEY_RockerUp                    = 0x1008ff23
	KEY_RockerDown                  = 0x1008ff24
	KEY_RockerEnter                 = 0x1008ff25
	KEY_Back                        = 0x1008ff26
	KEY_Forward                     = 0x1008ff27
	KEY_Stop                        = 0x1008ff28
	KEY_Refresh                     = 0x1008ff29
	KEY_PowerOff                    = 0x1008ff2a
	KEY_WakeUp                      = 0x1008ff2b
	KEY_Eject                       = 0x1008ff2c
	KEY_ScreenSaver                 = 0x1008ff2d
	KEY_WWW                         = 0x1008ff2e
	KEY_Sleep                       = 0x1008ff2f
	KEY_Favorites                   = 0x1008ff30
	KEY_AudioPause                  = 0x1008ff31
	KEY_AudioMedia                  = 0x1008ff32
	KEY_MyComputer                  = 0x1008ff33
	KEY_VendorHome                  = 0x1008ff34
	KEY_LightBulb                   = 0x1008ff35
	KEY_Shop                        = 0x1008ff36
	KEY_History                     = 0x1008ff37
	KEY_OpenURL                     = 0x1008ff38
	KEY_AddFavorite                 = 0x1008ff39
	KEY_HotLinks                    = 0x1008ff3a
	KEY_BrightnessAdjust            = 0x1008ff3b
	KEY_Finance                     = 0x1008ff3c
	KEY_Community                   = 0x1008ff3d
	KEY_AudioRewind                 = 0x1008ff3e
	KEY_BackForward                 = 0x1008ff3f
	KEY_Launch0                     = 0x1008ff40
	KEY_Launch1                     = 0x1008ff41
	KEY_Launch2                     = 0x1008ff42
	KEY_Launch3                     = 0x1008ff43
	KEY_Launch4                     = 0x1008ff44
	KEY_Launch5                     = 0x1008ff45
	KEY_Launch6                     = 0x1008ff46
	KEY_Launch7                     = 0x1008ff47
	KEY_Launch8                     = 0x1008ff48
	KEY_Launch9                     = 0x1008ff49
	KEY_LaunchA                     = 0x1008ff4a
	KEY_LaunchB                     = 0x1008ff4b
	KEY_LaunchC                     = 0x1008ff4c
	KEY_LaunchD                     = 0x1008ff4d
	KEY_LaunchE                     = 0x1008ff4e
	KEY_LaunchF                     = 0x1008ff4f
	KEY_ApplicationLeft             = 0x1008ff50
	KEY_ApplicationRight            = 0x1008ff51
	KEY_Book                        = 0x1008ff52
	KEY_CD                          = 0x1008ff53
	KEY_WindowClear                 = 0x1008ff55
	KEY_Close                       = 0x1008ff56
	KEY_Copy                        = 0x1008ff57
	KEY_Cut                         = 0x1008ff58
	KEY_Display                     = 0x1008ff59
	KEY_DOS                         = 0x1008ff5a
	KEY_Documents                   = 0x1008ff5b
	KEY_Excel                       = 0x1008ff5c
	KEY_Explorer                    = 0x1008ff5d
	KEY_Game                        = 0x1008ff5e
	KEY_Go                          = 0x1008ff5f
	KEY_iTouch                      = 0x1008ff60
	KEY_LogOff                      = 0x1008ff61
	KEY_Market                      = 0x1008ff62
	KEY_Meeting                     = 0x1008ff63
	KEY_MenuKB                      = 0x1008ff65
	KEY_MenuPB                      = 0x1008ff66
	KEY_MySites                     = 0x1008ff67
	KEY_New                         = 0x1008ff68
	KEY_News                        = 0x1008ff69
	KEY_OfficeHome                  = 0x1008ff6a
	KEY_Open                        = 0x1008ff6b
	KEY_Option                      = 0x1008ff6c
	KEY_Paste                       = 0x1008ff6d
	KEY_Phone                       = 0x1008ff6e
	KEY_Reply                       = 0x1008ff72
	KEY_Reload                      = 0x1008ff73
	KEY_RotateWindows               = 0x1008ff74
	KEY_RotationPB                  = 0x1008ff75
	KEY_RotationKB                  = 0x1008ff76
	KEY_Save                        = 0x1008ff77
	KEY_ScrollUp                    = 0x1008ff78
	KEY_ScrollDown                  = 0x1008ff79
	KEY_ScrollClick                 = 0x1008ff7a
	KEY_Send                        = 0x1008ff7b
	KEY_Spell                       = 0x1008ff7c
	KEY_SplitScreen                 = 0x1008ff7d
	KEY_Support                     = 0x1008ff7e
	KEY_TaskPane                    = 0x1008ff7f
	KEY_Terminal                    = 0x1008ff80
	KEY_Tools                       = 0x1008ff81
	KEY_Travel                      = 0x1008ff82
	KEY_UserPB                      = 0x1008ff84
	KEY_User1KB                     = 0x1008ff85
	KEY_User2KB                     = 0x1008ff86
	KEY_Video                       = 0x1008ff87
	KEY_WheelButton                 = 0x1008ff88
	KEY_Word                        = 0x1008ff89
	KEY_Xfer                        = 0x1008ff8a
	KEY_ZoomIn                      = 0x1008ff8b
	KEY_ZoomOut                     = 0x1008ff8c
	KEY_Away                        = 0x1008ff8d
	KEY_Messenger                   = 0x1008ff8e
	KEY_WebCam                      = 0x1008ff8f
	KEY_MailForward                 = 0x1008ff90
	KEY_Pictures                    = 0x1008ff91
	KEY_Music                       = 0x1008ff92
	KEY_Battery                     = 0x1008ff93
	KEY_Bluetooth                   = 0x1008ff94
	KEY_WLAN                        = 0x1008ff95
	KEY_UWB                         = 0x1008ff96
	KEY_AudioForward                = 0x1008ff97
	KEY_AudioRepeat                 = 0x1008ff98
	KEY_AudioRandomPlay             = 0x1008ff99
	KEY_Subtitle                    = 0x1008ff9a
	KEY_AudioCycleTrack             = 0x1008ff9b
	KEY_CycleAngle                  = 0x1008ff9c
	KEY_FrameBack                   = 0x1008ff9d
	KEY_FrameForward                = 0x1008ff9e
	KEY_Time                        = 0x1008ff9f
	KEY_SelectButton                = 0x1008ffa0
	KEY_View                        = 0x1008ffa1
	KEY_TopMenu                     = 0x1008ffa2
	KEY_Red                         = 0x1008ffa3
	KEY_Green                       = 0x1008ffa4
	KEY_Yellow                      = 0x1008ffa5
	KEY_Blue                        = 0x1008ffa6
	KEY_Suspend                     = 0x1008ffa7
	KEY_Hibernate                   = 0x1008ffa8
	KEY_TouchpadToggle              = 0x1008ffa9
	KEY_Switch_VT_1                 = 0x1008fe01
	KEY_Switch_VT_2                 = 0x1008fe02
	KEY_Switch_VT_3                 = 0x1008fe03
	KEY_Switch_VT_4                 = 0x1008fe04
	KEY_Switch_VT_5                 = 0x1008fe05
	KEY_Switch_VT_6                 = 0x1008fe06
	KEY_Switch_VT_7                 = 0x1008fe07
	KEY_Switch_VT_8                 = 0x1008fe08
	KEY_Switch_VT_9                 = 0x1008fe09
	KEY_Switch_VT_10                = 0x1008fe0a
	KEY_Switch_VT_11                = 0x1008fe0b
	KEY_Switch_VT_12                = 0x1008fe0c
	KEY_Ungrab                      = 0x1008fe20
	KEY_ClearGrab                   = 0x1008fe21
	KEY_Next_VMode                  = 0x1008fe22
	KEY_Prev_VMode                  = 0x1008fe23
)

type EventAny struct {
	Type      EventType
	Window    unsafe.Pointer
	SendEvent int8
}

type EventKey struct {
	Type            C.int
	Window          unsafe.Pointer
	SendEvent       int8
	Time            uint32
	State           C.unsigned
	Keyval          C.unsigned
	Length          int
	String          *uint8
	HardwareKeycode uint16
	Group           uint8
	IsModifier      C.unsigned
}

type EventButton struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	Axes      *float64
	State     uint32
	Button    uint16
	Device    uintptr
	XRoot     float64
	YRoot     float64
}

type EventScroll struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	State     uint
	Direction int
	Device    uintptr
	XRoot     float64
	YRoot     float64
}

type EventMotion struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	Axes      *float64
	State     uint
	IsHint    uint16
	Device    uintptr
	XRoot     float64
	YRoot     float64
}

type EventExpose struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
	Area      Rectangle
	Region    uintptr
	Count     int
}

type EventVisibility struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
	State     int
}

type EventCrossing struct {
	// TODO
}

type EventFocus struct {
	// TODO
}

type EventConfigure struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent uint8
	X         int
	Y         int
	Width     int
	Height    int
}

type EventProperty struct {
	// TODO
}

type EventProximity struct {
	// TODO
}

type EventClient struct {
	// TODO
}

type EventNoExpose struct {
	// TODO
}

type EventWindowState struct {
	// TODO
}

type EventSetting struct {
	// TODO
}

type EventOwnerChange struct {
	// TODO
}

type EventGrabBroken struct {
	// TODO
}

type DragContext struct {
	DragContext *C.GdkDragContext
}

func DragContextFromNative(l unsafe.Pointer) *DragContext {
	return &DragContext{C.toGdkDragContext(l)}
}

//-----------------------------------------------------------------------
// GdkAtom
//-----------------------------------------------------------------------
type Atom uintptr

//-----------------------------------------------------------------------
// GdkDisplay
//-----------------------------------------------------------------------
type Display struct {
	GDisplay unsafe.Pointer
}

func DisplayGetDefault() *Display {
	return &Display{C._gdk_display_get_default()}
}
