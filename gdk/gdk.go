package gdk

/*
#include <gdk/gdk.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

static gchar* to_gcharptr(char* s) { return (gchar*)s; }

static void free_string(char* s) { free(s); }

static GdkWindow* to_GdkWindow(void* w) {
	return GDK_WINDOW(w);
}
static void _g_thread_init(GThreadFunctions *vtable) {
#ifdef	G_THREADS_ENABLED
	g_thread_init(vtable);
#endif
}
*/
import "C"
import "unsafe"

func bool2gboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gboolean2bool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}
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
	if first {
		first = false
		C._g_thread_init(nil)
	}
	C.gdk_threads_init()
}

func ThreadsEnter() {
	C.gdk_threads_enter()
}

func ThreadsLeave() {
	C.gdk_threads_leave()
}

//-----------------------------------------------------------------------
// GdkCursor
//-----------------------------------------------------------------------
type GdkCursorType int

const (
	GDK_X_CURSOR            GdkCursorType = 0
	GDK_ARROW               GdkCursorType = 2
	GDK_BASED_ARROW_DOWN    GdkCursorType = 4
	GDK_BASED_ARROW_UP      GdkCursorType = 6
	GDK_BOAT                GdkCursorType = 8
	GDK_BOGOSITY            GdkCursorType = 10
	GDK_BOTTOM_LEFT_CORNER  GdkCursorType = 12
	GDK_BOTTOM_RIGHT_CORNER GdkCursorType = 14
	GDK_BOTTOM_SIDE         GdkCursorType = 16
	GDK_BOTTOM_TEE          GdkCursorType = 18
	GDK_BOX_SPIRAL          GdkCursorType = 20
	GDK_CENTER_PTR          GdkCursorType = 22
	GDK_CIRCLE              GdkCursorType = 24
	GDK_CLOCK               GdkCursorType = 26
	GDK_COFFEE_MUG          GdkCursorType = 28
	GDK_CROSS               GdkCursorType = 30
	GDK_CROSS_REVERSE       GdkCursorType = 32
	GDK_CROSSHAIR           GdkCursorType = 34
	GDK_DIAMOND_CROSS       GdkCursorType = 36
	GDK_DOT                 GdkCursorType = 38
	GDK_DOTBOX              GdkCursorType = 40
	GDK_DOUBLE_ARROW        GdkCursorType = 42
	GDK_DRAFT_LARGE         GdkCursorType = 44
	GDK_DRAFT_SMALL         GdkCursorType = 46
	GDK_DRAPED_BOX          GdkCursorType = 48
	GDK_EXCHANGE            GdkCursorType = 50
	GDK_FLEUR               GdkCursorType = 52
	GDK_GOBBLER             GdkCursorType = 54
	GDK_GUMBY               GdkCursorType = 56
	GDK_HAND1               GdkCursorType = 58
	GDK_HAND2               GdkCursorType = 60
	GDK_HEART               GdkCursorType = 62
	GDK_ICON                GdkCursorType = 64
	GDK_IRON_CROSS          GdkCursorType = 66
	GDK_LEFT_PTR            GdkCursorType = 68
	GDK_LEFT_SIDE           GdkCursorType = 70
	GDK_LEFT_TEE            GdkCursorType = 72
	GDK_LEFTBUTTON          GdkCursorType = 74
	GDK_LL_ANGLE            GdkCursorType = 76
	GDK_LR_ANGLE            GdkCursorType = 78
	GDK_MAN                 GdkCursorType = 80
	GDK_MIDDLEBUTTON        GdkCursorType = 82
	GDK_MOUSE               GdkCursorType = 84
	GDK_PENCIL              GdkCursorType = 86
	GDK_PIRATE              GdkCursorType = 88
	GDK_PLUS                GdkCursorType = 90
	GDK_QUESTION_ARROW      GdkCursorType = 92
	GDK_RIGHT_PTR           GdkCursorType = 94
	GDK_RIGHT_SIDE          GdkCursorType = 96
	GDK_RIGHT_TEE           GdkCursorType = 98
	GDK_RIGHTBUTTON         GdkCursorType = 100
	GDK_RTL_LOGO            GdkCursorType = 102
	GDK_SAILBOAT            GdkCursorType = 104
	GDK_SB_DOWN_ARROW       GdkCursorType = 106
	GDK_SB_H_DOUBLE_ARROW   GdkCursorType = 108
	GDK_SB_LEFT_ARROW       GdkCursorType = 110
	GDK_SB_RIGHT_ARROW      GdkCursorType = 112
	GDK_SB_UP_ARROW         GdkCursorType = 114
	GDK_SB_V_DOUBLE_ARROW   GdkCursorType = 116
	GDK_SHUTTLE             GdkCursorType = 118
	GDK_SIZING              GdkCursorType = 120
	GDK_SPIDER              GdkCursorType = 122
	GDK_SPRAYCAN            GdkCursorType = 124
	GDK_STAR                GdkCursorType = 126
	GDK_TARGET              GdkCursorType = 128
	GDK_TCROSS              GdkCursorType = 130
	GDK_TOP_LEFT_ARROW      GdkCursorType = 132
	GDK_TOP_LEFT_CORNER     GdkCursorType = 134
	GDK_TOP_RIGHT_CORNER    GdkCursorType = 136
	GDK_TOP_SIDE            GdkCursorType = 138
	GDK_TOP_TEE             GdkCursorType = 140
	GDK_TREK                GdkCursorType = 142
	GDK_UL_ANGLE            GdkCursorType = 144
	GDK_UMBRELLA            GdkCursorType = 146
	GDK_UR_ANGLE            GdkCursorType = 148
	GDK_WATCH               GdkCursorType = 150
	GDK_XTERM               GdkCursorType = 152
	GDK_LAST_CURSOR         GdkCursorType = 153
	GDK_BLANK_CURSOR        GdkCursorType = -2
	GDK_CURSOR_IS_PIXMAP    GdkCursorType = -1
)

var (
	first = true
)

type GdkCursor struct {
	Cursor *C.GdkCursor
}

func Cursor(cursor_type GdkCursorType) *GdkCursor {
	return &GdkCursor{
		C.gdk_cursor_new(C.GdkCursorType(cursor_type))}
}

//-----------------------------------------------------------------------
// GdkColor
//-----------------------------------------------------------------------
type GdkColor struct {
	Color C.GdkColor
}

func Color(name string) *GdkColor {
	var color C.GdkColor
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gdk_color_parse(C.to_gcharptr(ptr), &color)
	return &GdkColor{color}
}

//-----------------------------------------------------------------------
// GdkFont
//-----------------------------------------------------------------------
type GdkFont struct {
	Font *C.GdkFont
}

func FontLoad(name string) *GdkFont {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &GdkFont{
		C.gdk_font_load(C.to_gcharptr(ptr))}
}

func FontsetLoad(name string) *GdkFont {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &GdkFont{
		C.gdk_fontset_load(C.to_gcharptr(ptr))}
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
type GdkGC struct {
	GC *C.GdkGC
}

func GC(drawable *GdkDrawable) *GdkGC {
	return &GdkGC{
		C.gdk_gc_new(drawable.Drawable)}
}

// GdkGC *gdk_gc_new_with_values (GdkDrawable *drawable, GdkGCValues *values, GdkGCValuesMask values_mask);
// GdkGC *gdk_gc_ref (GdkGC *gc);
// void gdk_gc_unref (GdkGC *gc);
// void gdk_gc_get_values (GdkGC *gc, GdkGCValues *values);
// void gdk_gc_set_values (GdkGC *gc, GdkGCValues *values, GdkGCValuesMask values_mask);
func (v *GdkGC) SetForeground(color *GdkColor) {
	C.gdk_gc_set_foreground(v.GC, &color.Color)
}
func (v *GdkGC) SetBackground(color *GdkColor) {
	C.gdk_gc_set_background(v.GC, &color.Color)
}
// void gdk_gc_set_font (GdkGC *gc, GdkFont *font);
// void gdk_gc_set_function (GdkGC *gc, GdkFunction function);
// void gdk_gc_set_fill (GdkGC *gc, GdkFill fill);
// void gdk_gc_set_tile (GdkGC *gc, GdkPixmap *tile);
// void gdk_gc_set_stipple (GdkGC *gc, GdkPixmap *stipple);
// void gdk_gc_set_ts_origin (GdkGC *gc, gint x, gint y);
// void gdk_gc_set_clip_origin (GdkGC *gc, gint x, gint y);
// void gdk_gc_set_clip_mask (GdkGC *gc, GdkBitmap *mask);
// void gdk_gc_set_clip_rectangle (GdkGC *gc, const GdkRectangle *rectangle);
// void gdk_gc_set_clip_region (GdkGC *gc, const GdkRegion *region);
// void gdk_gc_set_subwindow (GdkGC *gc, GdkSubwindowMode mode);
// void gdk_gc_set_exposures (GdkGC *gc, gboolean exposures);
// void gdk_gc_set_line_attributes (GdkGC *gc, gint line_width, GdkLineStyle line_style, GdkCapStyle cap_style, GdkJoinStyle join_style);
// void gdk_gc_set_dashes (GdkGC *gc, gint dash_offset, gint8 dash_list[], gint n);
// void gdk_gc_offset (GdkGC *gc, gint x_offset, gint y_offset);
// void gdk_gc_copy (GdkGC *dst_gc, GdkGC *src_gc);
// void gdk_gc_set_colormap (GdkGC *gc, GdkColormap *colormap);
// GdkColormap *gdk_gc_get_colormap (GdkGC *gc);
func (v *GdkGC) SetRgbFgColor(color *GdkColor) {
	C.gdk_gc_set_rgb_fg_color(v.GC, &color.Color)
}
func (v *GdkGC) SetRgbBgColor(color *GdkColor) {
	C.gdk_gc_set_rgb_bg_color(v.GC, &color.Color)
}
// GdkScreen * gdk_gc_get_screen (GdkGC *gc);

//-----------------------------------------------------------------------
// GdkDrawable
//-----------------------------------------------------------------------
type GdkDrawable struct {
	Drawable *C.GdkDrawable
}

func (v *GdkDrawable) DrawPoint(gc *GdkGC, x int, y int) {
	C.gdk_draw_point(v.Drawable, gc.GC, C.gint(x), C.gint(y))
}
func (v *GdkDrawable) DrawLine(gc *GdkGC, x1 int, y1 int, x2 int, y2 int) {
	C.gdk_draw_line(v.Drawable, gc.GC, C.gint(x1), C.gint(y1), C.gint(x2), C.gint(y2))
}
func (v *GdkDrawable) DrawRectangle(gc *GdkGC, filled bool, x int, y int, width int, height int) {
	C.gdk_draw_rectangle(v.Drawable, gc.GC, bool2gboolean(filled), C.gint(x), C.gint(y), C.gint(width), C.gint(height))
}
func (v *GdkDrawable) DrawArc(gc *GdkGC, filled bool, x int, y int, width int, height int, angle1 int, angle2 int) {
	C.gdk_draw_arc(v.Drawable, gc.GC, bool2gboolean(filled), C.gint(x), C.gint(y), C.gint(width), C.gint(height), C.gint(angle1), C.gint(angle2))
}
// void gdk_draw_polygon (GdkDrawable *drawable, GdkGC *gc, gboolean filled, const GdkPoint *points, gint n_points);
// void gdk_draw_string (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *string);
func (v *GdkDrawable) DrawString(font *GdkFont, gc *GdkGC, x int, y int, str string) {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	C.gdk_draw_string(v.Drawable, font.Font, gc.GC, C.gint(x), C.gint(y), C.to_gcharptr(ptr))
}
// void gdk_draw_text (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *text, gint text_length);
// void gdk_draw_text_wc (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const GdkWChar *text, gint text_length);
func (v *GdkDrawable) DrawDrawable(gc *GdkGC, src *GdkDrawable, xsrc int, ysrc int, xdest int, ydest int, width int, height int) {
	C.gdk_draw_drawable(v.Drawable, gc.GC, src.Drawable, C.gint(xsrc), C.gint(ysrc), C.gint(xdest), C.gint(ydest), C.gint(width), C.gint(height))
}
// void gdk_draw_image (GdkDrawable *drawable, GdkGC *gc, GdkImage *image, gint xsrc, gint ysrc, gint xdest, gint ydest, gint width, gint height);
// void gdk_draw_points (GdkDrawable *drawable, GdkGC *gc, const GdkPoint *points, gint n_points);
// void gdk_draw_segments (GdkDrawable *drawable, GdkGC *gc, const GdkSegment *segs, gint n_segs);
// void gdk_draw_lines (GdkDrawable *drawable, GdkGC *gc, const GdkPoint *points, gint n_points);
// void gdk_draw_pixbuf (GdkDrawable *drawable, GdkGC *gc, const GdkPixbuf *pixbuf, gint src_x, gint src_y, gint dest_x, gint dest_y, gint width, gint height, GdkRgbDither dither, gint x_dither, gint y_dither);
// void gdk_draw_glyphs (GdkDrawable *drawable, GdkGC *gc, PangoFont *font, gint x, gint y, PangoGlyphString *glyphs);
// void gdk_draw_layout_line (GdkDrawable *drawable, GdkGC *gc, gint x, gint y, PangoLayoutLine *line);
// void gdk_draw_layout (GdkDrawable *drawable, GdkGC *gc, gint x, gint y, PangoLayout *layout);
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
type GdkModifierType int

const (
	GDK_SHIFT_MASK   GdkModifierType = 1 << 0
	GDK_LOCK_MASK    GdkModifierType = 1 << 1
	GDK_CONTROL_MASK GdkModifierType = 1 << 2
	GDK_MOD1_MASK    GdkModifierType = 1 << 3
	GDK_MOD2_MASK    GdkModifierType = 1 << 4
	GDK_MOD3_MASK    GdkModifierType = 1 << 5
	GDK_MOD4_MASK    GdkModifierType = 1 << 6
	GDK_MOD5_MASK    GdkModifierType = 1 << 7
	GDK_BUTTON1_MASK GdkModifierType = 1 << 8
	GDK_BUTTON2_MASK GdkModifierType = 1 << 9
	GDK_BUTTON3_MASK GdkModifierType = 1 << 10
	GDK_BUTTON4_MASK GdkModifierType = 1 << 11
	GDK_BUTTON5_MASK GdkModifierType = 1 << 12

	/* The next few modifiers are used by XKB, so we skip to the end.
	 * Bits 15 - 25 are currently unused. Bit 29 is used internally.
	 */

	GDK_SUPER_MASK GdkModifierType = 1 << 26
	GDK_HYPER_MASK GdkModifierType = 1 << 27
	GDK_META_MASK  GdkModifierType = 1 << 28

	GDK_RELEASE_MASK GdkModifierType = 1 << 30

	GDK_MODIFIER_MASK GdkModifierType = 0x5c001fff
)

type GdkEventType int

const (
	GDK_NOTHING           GdkEventType = -1
	GDK_DELETE            GdkEventType = 0
	GDK_DESTROY           GdkEventType = 1
	GDK_EXPOSE            GdkEventType = 2
	GDK_MOTION_NOTIFY     GdkEventType = 3
	GDK_BUTTON_PRESS      GdkEventType = 4
	GDK_2BUTTON_PRESS     GdkEventType = 5
	GDK_3BUTTON_PRESS     GdkEventType = 6
	GDK_BUTTON_RELEASE    GdkEventType = 7
	GDK_KEY_PRESS         GdkEventType = 8
	GDK_KEY_RELEASE       GdkEventType = 9
	GDK_ENTER_NOTIFY      GdkEventType = 10
	GDK_LEAVE_NOTIFY      GdkEventType = 11
	GDK_FOCUS_CHANGE      GdkEventType = 12
	GDK_CONFIGURE         GdkEventType = 13
	GDK_MAP               GdkEventType = 14
	GDK_UNMAP             GdkEventType = 15
	GDK_PROPERTY_NOTIFY   GdkEventType = 16
	GDK_SELECTION_CLEAR   GdkEventType = 17
	GDK_SELECTION_REQUEST GdkEventType = 18
	GDK_SELECTION_NOTIFY  GdkEventType = 19
	GDK_PROXIMITY_IN      GdkEventType = 20
	GDK_PROXIMITY_OUT     GdkEventType = 21
	GDK_DRAG_ENTER        GdkEventType = 22
	GDK_DRAG_LEAVE        GdkEventType = 23
	GDK_DRAG_MOTION       GdkEventType = 24
	GDK_DRAG_STATUS       GdkEventType = 25
	GDK_DROP_START        GdkEventType = 26
	GDK_DROP_FINISHED     GdkEventType = 27
	GDK_CLIENT_EVENT      GdkEventType = 28
	GDK_VISIBILITY_NOTIFY GdkEventType = 29
	GDK_NO_EXPOSE         GdkEventType = 30
	GDK_SCROLL            GdkEventType = 31
	GDK_WINDOW_STATE      GdkEventType = 32
	GDK_SETTING           GdkEventType = 33
	GDK_OWNER_CHANGE      GdkEventType = 34
	GDK_GRAB_BROKEN       GdkEventType = 35
	GDK_DAMAGE            GdkEventType = 36
	GDK_EVENT_LAST        GdkEventType = 37 /* helper variable for decls */
)

type GdkEventMask int

const (
	GDK_EXPOSURE_MASK            GdkEventMask = 1 << 1
	GDK_POINTER_MOTION_MASK      GdkEventMask = 1 << 2
	GDK_POINTER_MOTION_HINT_MASK GdkEventMask = 1 << 3
	GDK_BUTTON_MOTION_MASK       GdkEventMask = 1 << 4
	GDK_BUTTON1_MOTION_MASK      GdkEventMask = 1 << 5
	GDK_BUTTON2_MOTION_MASK      GdkEventMask = 1 << 6
	GDK_BUTTON3_MOTION_MASK      GdkEventMask = 1 << 7
	GDK_BUTTON_PRESS_MASK        GdkEventMask = 1 << 8
	GDK_BUTTON_RELEASE_MASK      GdkEventMask = 1 << 9
	GDK_KEY_PRESS_MASK           GdkEventMask = 1 << 10
	GDK_KEY_RELEASE_MASK         GdkEventMask = 1 << 11
	GDK_ENTER_NOTIFY_MASK        GdkEventMask = 1 << 12
	GDK_LEAVE_NOTIFY_MASK        GdkEventMask = 1 << 13
	GDK_FOCUS_CHANGE_MASK        GdkEventMask = 1 << 14
	GDK_STRUCTURE_MASK           GdkEventMask = 1 << 15
	GDK_PROPERTY_CHANGE_MASK     GdkEventMask = 1 << 16
	GDK_VISIBILITY_NOTIFY_MASK   GdkEventMask = 1 << 17
	GDK_PROXIMITY_IN_MASK        GdkEventMask = 1 << 18
	GDK_PROXIMITY_OUT_MASK       GdkEventMask = 1 << 19
	GDK_SUBSTRUCTURE_MASK        GdkEventMask = 1 << 20
	GDK_SCROLL_MASK              GdkEventMask = 1 << 21
	GDK_ALL_EVENTS_MASK          GdkEventMask = 0x3FFFFE
)

type GdkWindow struct {
	Window *C.GdkWindow
}

func WindowFromUnsafe(window unsafe.Pointer) *GdkWindow {
	return &GdkWindow{
		C.to_GdkWindow(window)}
}

func (v *GdkWindow) GetPointer(x *int, y *int, mask *GdkModifierType) *GdkWindow {
	var cx, cy C.gint
	var mt C.GdkModifierType
	ret := &GdkWindow{
		C.gdk_window_get_pointer(v.Window, &cx, &cy, &mt)}
	*x = int(cx)
	*y = int(cy)
	*mask = GdkModifierType(mt)
	return ret
}

func (v *GdkWindow) GetDrawable() *GdkDrawable {
	return &GdkDrawable{
		(*C.GdkDrawable)(v.Window)}
}

func (v *GdkWindow) Invalidate(rect *Rectangle, invalidate_children bool) {
	if rect != nil {
		var _rect C.GdkRectangle
		_rect.x = C.gint(rect.X)
		_rect.y = C.gint(rect.Y)
		_rect.width = C.gint(rect.Width)
		_rect.height = C.gint(rect.Height)
		C.gdk_window_invalidate_rect(v.Window, &_rect, bool2gboolean(invalidate_children))
	} else {
		C.gdk_window_invalidate_rect(v.Window, nil, bool2gboolean(invalidate_children))
	}
}

func (v *GdkWindow) Show() {
	C.gdk_window_show(v.Window)
}

func (v *GdkWindow) Raise() {
	C.gdk_window_raise(v.Window)
}

//-----------------------------------------------------------------------
// GdkPixmap
//-----------------------------------------------------------------------
type GdkPixmap struct {
	Pixmap *C.GdkPixmap
}

func Pixmap(drawable *GdkDrawable, width int, height int, depth int) *GdkPixmap {
	return &GdkPixmap{
		C.gdk_pixmap_new(drawable.Drawable, C.gint(width), C.gint(height), C.gint(depth))}
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
func (v *GdkPixmap) Ref() {
	C.g_object_ref(C.gpointer(v.Pixmap))
}
func (v *GdkPixmap) Unref() {
	C.g_object_unref(C.gpointer(v.Pixmap))
}
func (v *GdkPixmap) GetDrawable() *GdkDrawable {
	return &GdkDrawable{
		(*C.GdkDrawable)(v.Pixmap)}
}

// Subset of gdkkeysyms.h
const (
	GDK_0            = 0x030
	GDK_1            = 0x031
	GDK_2            = 0x032
	GDK_3            = 0x033
	GDK_4            = 0x034
	GDK_5            = 0x035
	GDK_6            = 0x036
	GDK_7            = 0x037
	GDK_8            = 0x038
	GDK_9            = 0x039
	GDK_colon        = 0x03a
	GDK_semicolon    = 0x03b
	GDK_less         = 0x03c
	GDK_equal        = 0x03d
	GDK_greater      = 0x03e
	GDK_question     = 0x03f
	GDK_at           = 0x040
	GDK_A            = 0x041
	GDK_B            = 0x042
	GDK_C            = 0x043
	GDK_D            = 0x044
	GDK_E            = 0x045
	GDK_F            = 0x046
	GDK_G            = 0x047
	GDK_H            = 0x048
	GDK_I            = 0x049
	GDK_J            = 0x04a
	GDK_K            = 0x04b
	GDK_L            = 0x04c
	GDK_M            = 0x04d
	GDK_N            = 0x04e
	GDK_O            = 0x04f
	GDK_P            = 0x050
	GDK_Q            = 0x051
	GDK_R            = 0x052
	GDK_S            = 0x053
	GDK_T            = 0x054
	GDK_U            = 0x055
	GDK_V            = 0x056
	GDK_W            = 0x057
	GDK_X            = 0x058
	GDK_Y            = 0x059
	GDK_Z            = 0x05a
	GDK_bracketleft  = 0x05b
	GDK_backslash    = 0x05c
	GDK_bracketright = 0x05d
	GDK_asciicircum  = 0x05e
	GDK_underscore   = 0x05f
	GDK_grave        = 0x060
	GDK_quoteleft    = 0x060
	GDK_a            = 0x061
	GDK_b            = 0x062
	GDK_c            = 0x063
	GDK_d            = 0x064
	GDK_e            = 0x065
	GDK_f            = 0x066
	GDK_g            = 0x067
	GDK_h            = 0x068
	GDK_i            = 0x069
	GDK_j            = 0x06a
	GDK_k            = 0x06b
	GDK_l            = 0x06c
	GDK_m            = 0x06d
	GDK_n            = 0x06e
	GDK_o            = 0x06f
	GDK_p            = 0x070
	GDK_q            = 0x071
	GDK_r            = 0x072
	GDK_s            = 0x073
	GDK_t            = 0x074
	GDK_u            = 0x075
	GDK_v            = 0x076
	GDK_w            = 0x077
	GDK_x            = 0x078
	GDK_y            = 0x079
	GDK_z            = 0x07a
	GDK_F1           = 0xffbe
	GDK_F2           = 0xffbf
	GDK_F3           = 0xffc0
	GDK_F4           = 0xffc1
	GDK_F5           = 0xffc2
	GDK_F6           = 0xffc3
	GDK_F7           = 0xffc4
	GDK_F8           = 0xffc5
	GDK_F9           = 0xffc6
	GDK_F10          = 0xffc7
	GDK_F11          = 0xffc8
	GDK_Return       = 0xff0d
	GDK_KP_Enter     = 0xff8d
	GDK_ISO_Enter    = 0xfe34
	GDK_bar          = 0x07c
)

type EventAny struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
}

type EventKey struct {
	Type            int
	Window          unsafe.Pointer
	SendEvent       int8
	Time            uint32
	State           uint
	Keyval          uint
	Length          int
	String          *uint8
	HardwareKeycode uint16
	Group           uint8
	IsModifier      uint
}

type EventButton struct {
	Type      int
	Window    unsafe.Pointer
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	Axes      *float64
	State     uint
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
	Axes      *float64
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
	// TODO
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
