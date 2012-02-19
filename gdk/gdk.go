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

static GdkDragContext* to_GtkDragContext(void* l) {
	return (GdkDragContext*)l;
}

static void* _gdk_display_get_default() {
	return (void*) gdk_display_get_default();
}

static GdkWindow* to_GdkWindow(void* w) {
	return GDK_WINDOW(w);
}

static void _g_thread_init(GThreadFunctions *vtable) {
#ifdef	G_THREADS_ENABLED
	g_thread_init(vtable);
#endif
}
*/
// #cgo pkg-config: gdk-2.0 gthread-2.0
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
// GdkSelection
//-----------------------------------------------------------------------
const (
	GDK_SELECTION_PRIMARY GdkAtom = GdkAtom(uintptr(1))
	GDK_SELECTION_SECONDARY GdkAtom = GdkAtom(uintptr(2))
	GDK_SELECTION_CLIPBOARD GdkAtom = GdkAtom(uintptr(69))
)

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

type GdkDragAction int

const (
	GDK_ACTION_DEFAULT GdkDragAction = 1 << 0
	GDK_ACTION_COPY    GdkDragAction = 1 << 1
	GDK_ACTION_MOVE    GdkDragAction = 1 << 2
	GDK_ACTION_LINK    GdkDragAction = 1 << 3
	GDK_ACTION_PRIVATE GdkDragAction = 1 << 4
	GDK_ACTION_ASK     GdkDragAction = 1 << 5
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
	GDK_KEY_VoidSymbol                  = 0xffffff
	GDK_KEY_BackSpace                   = 0xff08
	GDK_KEY_Tab                         = 0xff09
	GDK_KEY_Linefeed                    = 0xff0a
	GDK_KEY_Clear                       = 0xff0b
	GDK_KEY_Return                      = 0xff0d
	GDK_KEY_Pause                       = 0xff13
	GDK_KEY_Scroll_Lock                 = 0xff14
	GDK_KEY_Sys_Req                     = 0xff15
	GDK_KEY_Escape                      = 0xff1b
	GDK_KEY_Delete                      = 0xffff
	GDK_KEY_Multi_key                   = 0xff20
	GDK_KEY_Codeinput                   = 0xff37
	GDK_KEY_SingleCandidate             = 0xff3c
	GDK_KEY_MultipleCandidate           = 0xff3d
	GDK_KEY_PreviousCandidate           = 0xff3e
	GDK_KEY_Kanji                       = 0xff21
	GDK_KEY_Muhenkan                    = 0xff22
	GDK_KEY_Henkan_Mode                 = 0xff23
	GDK_KEY_Henkan                      = 0xff23
	GDK_KEY_Romaji                      = 0xff24
	GDK_KEY_Hiragana                    = 0xff25
	GDK_KEY_Katakana                    = 0xff26
	GDK_KEY_Hiragana_Katakana           = 0xff27
	GDK_KEY_Zenkaku                     = 0xff28
	GDK_KEY_Hankaku                     = 0xff29
	GDK_KEY_Zenkaku_Hankaku             = 0xff2a
	GDK_KEY_Touroku                     = 0xff2b
	GDK_KEY_Massyo                      = 0xff2c
	GDK_KEY_Kana_Lock                   = 0xff2d
	GDK_KEY_Kana_Shift                  = 0xff2e
	GDK_KEY_Eisu_Shift                  = 0xff2f
	GDK_KEY_Eisu_toggle                 = 0xff30
	GDK_KEY_Kanji_Bangou                = 0xff37
	GDK_KEY_Zen_Koho                    = 0xff3d
	GDK_KEY_Mae_Koho                    = 0xff3e
	GDK_KEY_Home                        = 0xff50
	GDK_KEY_Left                        = 0xff51
	GDK_KEY_Up                          = 0xff52
	GDK_KEY_Right                       = 0xff53
	GDK_KEY_Down                        = 0xff54
	GDK_KEY_Prior                       = 0xff55
	GDK_KEY_Page_Up                     = 0xff55
	GDK_KEY_Next                        = 0xff56
	GDK_KEY_Page_Down                   = 0xff56
	GDK_KEY_End                         = 0xff57
	GDK_KEY_Begin                       = 0xff58
	GDK_KEY_Select                      = 0xff60
	GDK_KEY_Print                       = 0xff61
	GDK_KEY_Execute                     = 0xff62
	GDK_KEY_Insert                      = 0xff63
	GDK_KEY_Undo                        = 0xff65
	GDK_KEY_Redo                        = 0xff66
	GDK_KEY_Menu                        = 0xff67
	GDK_KEY_Find                        = 0xff68
	GDK_KEY_Cancel                      = 0xff69
	GDK_KEY_Help                        = 0xff6a
	GDK_KEY_Break                       = 0xff6b
	GDK_KEY_Mode_switch                 = 0xff7e
	GDK_KEY_script_switch               = 0xff7e
	GDK_KEY_Num_Lock                    = 0xff7f
	GDK_KEY_KP_Space                    = 0xff80
	GDK_KEY_KP_Tab                      = 0xff89
	GDK_KEY_KP_Enter                    = 0xff8d
	GDK_KEY_KP_F1                       = 0xff91
	GDK_KEY_KP_F2                       = 0xff92
	GDK_KEY_KP_F3                       = 0xff93
	GDK_KEY_KP_F4                       = 0xff94
	GDK_KEY_KP_Home                     = 0xff95
	GDK_KEY_KP_Left                     = 0xff96
	GDK_KEY_KP_Up                       = 0xff97
	GDK_KEY_KP_Right                    = 0xff98
	GDK_KEY_KP_Down                     = 0xff99
	GDK_KEY_KP_Prior                    = 0xff9a
	GDK_KEY_KP_Page_Up                  = 0xff9a
	GDK_KEY_KP_Next                     = 0xff9b
	GDK_KEY_KP_Page_Down                = 0xff9b
	GDK_KEY_KP_End                      = 0xff9c
	GDK_KEY_KP_Begin                    = 0xff9d
	GDK_KEY_KP_Insert                   = 0xff9e
	GDK_KEY_KP_Delete                   = 0xff9f
	GDK_KEY_KP_Equal                    = 0xffbd
	GDK_KEY_KP_Multiply                 = 0xffaa
	GDK_KEY_KP_Add                      = 0xffab
	GDK_KEY_KP_Separator                = 0xffac
	GDK_KEY_KP_Subtract                 = 0xffad
	GDK_KEY_KP_Decimal                  = 0xffae
	GDK_KEY_KP_Divide                   = 0xffaf
	GDK_KEY_KP_0                        = 0xffb0
	GDK_KEY_KP_1                        = 0xffb1
	GDK_KEY_KP_2                        = 0xffb2
	GDK_KEY_KP_3                        = 0xffb3
	GDK_KEY_KP_4                        = 0xffb4
	GDK_KEY_KP_5                        = 0xffb5
	GDK_KEY_KP_6                        = 0xffb6
	GDK_KEY_KP_7                        = 0xffb7
	GDK_KEY_KP_8                        = 0xffb8
	GDK_KEY_KP_9                        = 0xffb9
	GDK_KEY_F1                          = 0xffbe
	GDK_KEY_F2                          = 0xffbf
	GDK_KEY_F3                          = 0xffc0
	GDK_KEY_F4                          = 0xffc1
	GDK_KEY_F5                          = 0xffc2
	GDK_KEY_F6                          = 0xffc3
	GDK_KEY_F7                          = 0xffc4
	GDK_KEY_F8                          = 0xffc5
	GDK_KEY_F9                          = 0xffc6
	GDK_KEY_F10                         = 0xffc7
	GDK_KEY_F11                         = 0xffc8
	GDK_KEY_L1                          = 0xffc8
	GDK_KEY_F12                         = 0xffc9
	GDK_KEY_L2                          = 0xffc9
	GDK_KEY_F13                         = 0xffca
	GDK_KEY_L3                          = 0xffca
	GDK_KEY_F14                         = 0xffcb
	GDK_KEY_L4                          = 0xffcb
	GDK_KEY_F15                         = 0xffcc
	GDK_KEY_L5                          = 0xffcc
	GDK_KEY_F16                         = 0xffcd
	GDK_KEY_L6                          = 0xffcd
	GDK_KEY_F17                         = 0xffce
	GDK_KEY_L7                          = 0xffce
	GDK_KEY_F18                         = 0xffcf
	GDK_KEY_L8                          = 0xffcf
	GDK_KEY_F19                         = 0xffd0
	GDK_KEY_L9                          = 0xffd0
	GDK_KEY_F20                         = 0xffd1
	GDK_KEY_L10                         = 0xffd1
	GDK_KEY_F21                         = 0xffd2
	GDK_KEY_R1                          = 0xffd2
	GDK_KEY_F22                         = 0xffd3
	GDK_KEY_R2                          = 0xffd3
	GDK_KEY_F23                         = 0xffd4
	GDK_KEY_R3                          = 0xffd4
	GDK_KEY_F24                         = 0xffd5
	GDK_KEY_R4                          = 0xffd5
	GDK_KEY_F25                         = 0xffd6
	GDK_KEY_R5                          = 0xffd6
	GDK_KEY_F26                         = 0xffd7
	GDK_KEY_R6                          = 0xffd7
	GDK_KEY_F27                         = 0xffd8
	GDK_KEY_R7                          = 0xffd8
	GDK_KEY_F28                         = 0xffd9
	GDK_KEY_R8                          = 0xffd9
	GDK_KEY_F29                         = 0xffda
	GDK_KEY_R9                          = 0xffda
	GDK_KEY_F30                         = 0xffdb
	GDK_KEY_R10                         = 0xffdb
	GDK_KEY_F31                         = 0xffdc
	GDK_KEY_R11                         = 0xffdc
	GDK_KEY_F32                         = 0xffdd
	GDK_KEY_R12                         = 0xffdd
	GDK_KEY_F33                         = 0xffde
	GDK_KEY_R13                         = 0xffde
	GDK_KEY_F34                         = 0xffdf
	GDK_KEY_R14                         = 0xffdf
	GDK_KEY_F35                         = 0xffe0
	GDK_KEY_R15                         = 0xffe0
	GDK_KEY_Shift_L                     = 0xffe1
	GDK_KEY_Shift_R                     = 0xffe2
	GDK_KEY_Control_L                   = 0xffe3
	GDK_KEY_Control_R                   = 0xffe4
	GDK_KEY_Caps_Lock                   = 0xffe5
	GDK_KEY_Shift_Lock                  = 0xffe6
	GDK_KEY_Meta_L                      = 0xffe7
	GDK_KEY_Meta_R                      = 0xffe8
	GDK_KEY_Alt_L                       = 0xffe9
	GDK_KEY_Alt_R                       = 0xffea
	GDK_KEY_Super_L                     = 0xffeb
	GDK_KEY_Super_R                     = 0xffec
	GDK_KEY_Hyper_L                     = 0xffed
	GDK_KEY_Hyper_R                     = 0xffee
	GDK_KEY_ISO_Lock                    = 0xfe01
	GDK_KEY_ISO_Level2_Latch            = 0xfe02
	GDK_KEY_ISO_Level3_Shift            = 0xfe03
	GDK_KEY_ISO_Level3_Latch            = 0xfe04
	GDK_KEY_ISO_Level3_Lock             = 0xfe05
	GDK_KEY_ISO_Level5_Shift            = 0xfe11
	GDK_KEY_ISO_Level5_Latch            = 0xfe12
	GDK_KEY_ISO_Level5_Lock             = 0xfe13
	GDK_KEY_ISO_Group_Shift             = 0xff7e
	GDK_KEY_ISO_Group_Latch             = 0xfe06
	GDK_KEY_ISO_Group_Lock              = 0xfe07
	GDK_KEY_ISO_Next_Group              = 0xfe08
	GDK_KEY_ISO_Next_Group_Lock         = 0xfe09
	GDK_KEY_ISO_Prev_Group              = 0xfe0a
	GDK_KEY_ISO_Prev_Group_Lock         = 0xfe0b
	GDK_KEY_ISO_First_Group             = 0xfe0c
	GDK_KEY_ISO_First_Group_Lock        = 0xfe0d
	GDK_KEY_ISO_Last_Group              = 0xfe0e
	GDK_KEY_ISO_Last_Group_Lock         = 0xfe0f
	GDK_KEY_ISO_Left_Tab                = 0xfe20
	GDK_KEY_ISO_Move_Line_Up            = 0xfe21
	GDK_KEY_ISO_Move_Line_Down          = 0xfe22
	GDK_KEY_ISO_Partial_Line_Up         = 0xfe23
	GDK_KEY_ISO_Partial_Line_Down       = 0xfe24
	GDK_KEY_ISO_Partial_Space_Left      = 0xfe25
	GDK_KEY_ISO_Partial_Space_Right     = 0xfe26
	GDK_KEY_ISO_Set_Margin_Left         = 0xfe27
	GDK_KEY_ISO_Set_Margin_Right        = 0xfe28
	GDK_KEY_ISO_Release_Margin_Left     = 0xfe29
	GDK_KEY_ISO_Release_Margin_Right    = 0xfe2a
	GDK_KEY_ISO_Release_Both_Margins    = 0xfe2b
	GDK_KEY_ISO_Fast_Cursor_Left        = 0xfe2c
	GDK_KEY_ISO_Fast_Cursor_Right       = 0xfe2d
	GDK_KEY_ISO_Fast_Cursor_Up          = 0xfe2e
	GDK_KEY_ISO_Fast_Cursor_Down        = 0xfe2f
	GDK_KEY_ISO_Continuous_Underline    = 0xfe30
	GDK_KEY_ISO_Discontinuous_Underline = 0xfe31
	GDK_KEY_ISO_Emphasize               = 0xfe32
	GDK_KEY_ISO_Center_Object           = 0xfe33
	GDK_KEY_ISO_Enter                   = 0xfe34
	GDK_KEY_dead_grave                  = 0xfe50
	GDK_KEY_dead_acute                  = 0xfe51
	GDK_KEY_dead_circumflex             = 0xfe52
	GDK_KEY_dead_tilde                  = 0xfe53
	GDK_KEY_dead_perispomeni            = 0xfe53
	GDK_KEY_dead_macron                 = 0xfe54
	GDK_KEY_dead_breve                  = 0xfe55
	GDK_KEY_dead_abovedot               = 0xfe56
	GDK_KEY_dead_diaeresis              = 0xfe57
	GDK_KEY_dead_abovering              = 0xfe58
	GDK_KEY_dead_doubleacute            = 0xfe59
	GDK_KEY_dead_caron                  = 0xfe5a
	GDK_KEY_dead_cedilla                = 0xfe5b
	GDK_KEY_dead_ogonek                 = 0xfe5c
	GDK_KEY_dead_iota                   = 0xfe5d
	GDK_KEY_dead_voiced_sound           = 0xfe5e
	GDK_KEY_dead_semivoiced_sound       = 0xfe5f
	GDK_KEY_dead_belowdot               = 0xfe60
	GDK_KEY_dead_hook                   = 0xfe61
	GDK_KEY_dead_horn                   = 0xfe62
	GDK_KEY_dead_stroke                 = 0xfe63
	GDK_KEY_dead_abovecomma             = 0xfe64
	GDK_KEY_dead_psili                  = 0xfe64
	GDK_KEY_dead_abovereversedcomma     = 0xfe65
	GDK_KEY_dead_dasia                  = 0xfe65
	GDK_KEY_dead_doublegrave            = 0xfe66
	GDK_KEY_dead_belowring              = 0xfe67
	GDK_KEY_dead_belowmacron            = 0xfe68
	GDK_KEY_dead_belowcircumflex        = 0xfe69
	GDK_KEY_dead_belowtilde             = 0xfe6a
	GDK_KEY_dead_belowbreve             = 0xfe6b
	GDK_KEY_dead_belowdiaeresis         = 0xfe6c
	GDK_KEY_dead_invertedbreve          = 0xfe6d
	GDK_KEY_dead_belowcomma             = 0xfe6e
	GDK_KEY_dead_currency               = 0xfe6f
	GDK_KEY_dead_a                      = 0xfe80
	GDK_KEY_dead_A                      = 0xfe81
	GDK_KEY_dead_e                      = 0xfe82
	GDK_KEY_dead_E                      = 0xfe83
	GDK_KEY_dead_i                      = 0xfe84
	GDK_KEY_dead_I                      = 0xfe85
	GDK_KEY_dead_o                      = 0xfe86
	GDK_KEY_dead_O                      = 0xfe87
	GDK_KEY_dead_u                      = 0xfe88
	GDK_KEY_dead_U                      = 0xfe89
	GDK_KEY_dead_small_schwa            = 0xfe8a
	GDK_KEY_dead_capital_schwa          = 0xfe8b
	GDK_KEY_First_Virtual_Screen        = 0xfed0
	GDK_KEY_Prev_Virtual_Screen         = 0xfed1
	GDK_KEY_Next_Virtual_Screen         = 0xfed2
	GDK_KEY_Last_Virtual_Screen         = 0xfed4
	GDK_KEY_Terminate_Server            = 0xfed5
	GDK_KEY_AccessX_Enable              = 0xfe70
	GDK_KEY_AccessX_Feedback_Enable     = 0xfe71
	GDK_KEY_RepeatKeys_Enable           = 0xfe72
	GDK_KEY_SlowKeys_Enable             = 0xfe73
	GDK_KEY_BounceKeys_Enable           = 0xfe74
	GDK_KEY_StickyKeys_Enable           = 0xfe75
	GDK_KEY_MouseKeys_Enable            = 0xfe76
	GDK_KEY_MouseKeys_Accel_Enable      = 0xfe77
	GDK_KEY_Overlay1_Enable             = 0xfe78
	GDK_KEY_Overlay2_Enable             = 0xfe79
	GDK_KEY_AudibleBell_Enable          = 0xfe7a
	GDK_KEY_Pointer_Left                = 0xfee0
	GDK_KEY_Pointer_Right               = 0xfee1
	GDK_KEY_Pointer_Up                  = 0xfee2
	GDK_KEY_Pointer_Down                = 0xfee3
	GDK_KEY_Pointer_UpLeft              = 0xfee4
	GDK_KEY_Pointer_UpRight             = 0xfee5
	GDK_KEY_Pointer_DownLeft            = 0xfee6
	GDK_KEY_Pointer_DownRight           = 0xfee7
	GDK_KEY_Pointer_Button_Dflt         = 0xfee8
	GDK_KEY_Pointer_Button1             = 0xfee9
	GDK_KEY_Pointer_Button2             = 0xfeea
	GDK_KEY_Pointer_Button3             = 0xfeeb
	GDK_KEY_Pointer_Button4             = 0xfeec
	GDK_KEY_Pointer_Button5             = 0xfeed
	GDK_KEY_Pointer_DblClick_Dflt       = 0xfeee
	GDK_KEY_Pointer_DblClick1           = 0xfeef
	GDK_KEY_Pointer_DblClick2           = 0xfef0
	GDK_KEY_Pointer_DblClick3           = 0xfef1
	GDK_KEY_Pointer_DblClick4           = 0xfef2
	GDK_KEY_Pointer_DblClick5           = 0xfef3
	GDK_KEY_Pointer_Drag_Dflt           = 0xfef4
	GDK_KEY_Pointer_Drag1               = 0xfef5
	GDK_KEY_Pointer_Drag2               = 0xfef6
	GDK_KEY_Pointer_Drag3               = 0xfef7
	GDK_KEY_Pointer_Drag4               = 0xfef8
	GDK_KEY_Pointer_Drag5               = 0xfefd
	GDK_KEY_Pointer_EnableKeys          = 0xfef9
	GDK_KEY_Pointer_Accelerate          = 0xfefa
	GDK_KEY_Pointer_DfltBtnNext         = 0xfefb
	GDK_KEY_Pointer_DfltBtnPrev         = 0xfefc
	GDK_KEY_3270_Duplicate              = 0xfd01
	GDK_KEY_3270_FieldMark              = 0xfd02
	GDK_KEY_3270_Right2                 = 0xfd03
	GDK_KEY_3270_Left2                  = 0xfd04
	GDK_KEY_3270_BackTab                = 0xfd05
	GDK_KEY_3270_EraseEOF               = 0xfd06
	GDK_KEY_3270_EraseInput             = 0xfd07
	GDK_KEY_3270_Reset                  = 0xfd08
	GDK_KEY_3270_Quit                   = 0xfd09
	GDK_KEY_3270_PA1                    = 0xfd0a
	GDK_KEY_3270_PA2                    = 0xfd0b
	GDK_KEY_3270_PA3                    = 0xfd0c
	GDK_KEY_3270_Test                   = 0xfd0d
	GDK_KEY_3270_Attn                   = 0xfd0e
	GDK_KEY_3270_CursorBlink            = 0xfd0f
	GDK_KEY_3270_AltCursor              = 0xfd10
	GDK_KEY_3270_KeyClick               = 0xfd11
	GDK_KEY_3270_Jump                   = 0xfd12
	GDK_KEY_3270_Ident                  = 0xfd13
	GDK_KEY_3270_Rule                   = 0xfd14
	GDK_KEY_3270_Copy                   = 0xfd15
	GDK_KEY_3270_Play                   = 0xfd16
	GDK_KEY_3270_Setup                  = 0xfd17
	GDK_KEY_3270_Record                 = 0xfd18
	GDK_KEY_3270_ChangeScreen           = 0xfd19
	GDK_KEY_3270_DeleteWord             = 0xfd1a
	GDK_KEY_3270_ExSelect               = 0xfd1b
	GDK_KEY_3270_CursorSelect           = 0xfd1c
	GDK_KEY_3270_PrintScreen            = 0xfd1d
	GDK_KEY_3270_Enter                  = 0xfd1e
	GDK_KEY_space                       = 0x020
	GDK_KEY_exclam                      = 0x021
	GDK_KEY_quotedbl                    = 0x022
	GDK_KEY_numbersign                  = 0x023
	GDK_KEY_dollar                      = 0x024
	GDK_KEY_percent                     = 0x025
	GDK_KEY_ampersand                   = 0x026
	GDK_KEY_apostrophe                  = 0x027
	GDK_KEY_quoteright                  = 0x027
	GDK_KEY_parenleft                   = 0x028
	GDK_KEY_parenright                  = 0x029
	GDK_KEY_asterisk                    = 0x02a
	GDK_KEY_plus                        = 0x02b
	GDK_KEY_comma                       = 0x02c
	GDK_KEY_minus                       = 0x02d
	GDK_KEY_period                      = 0x02e
	GDK_KEY_slash                       = 0x02f
	GDK_KEY_0                           = 0x030
	GDK_KEY_1                           = 0x031
	GDK_KEY_2                           = 0x032
	GDK_KEY_3                           = 0x033
	GDK_KEY_4                           = 0x034
	GDK_KEY_5                           = 0x035
	GDK_KEY_6                           = 0x036
	GDK_KEY_7                           = 0x037
	GDK_KEY_8                           = 0x038
	GDK_KEY_9                           = 0x039
	GDK_KEY_colon                       = 0x03a
	GDK_KEY_semicolon                   = 0x03b
	GDK_KEY_less                        = 0x03c
	GDK_KEY_equal                       = 0x03d
	GDK_KEY_greater                     = 0x03e
	GDK_KEY_question                    = 0x03f
	GDK_KEY_at                          = 0x040
	GDK_KEY_A                           = 0x041
	GDK_KEY_B                           = 0x042
	GDK_KEY_C                           = 0x043
	GDK_KEY_D                           = 0x044
	GDK_KEY_E                           = 0x045
	GDK_KEY_F                           = 0x046
	GDK_KEY_G                           = 0x047
	GDK_KEY_H                           = 0x048
	GDK_KEY_I                           = 0x049
	GDK_KEY_J                           = 0x04a
	GDK_KEY_K                           = 0x04b
	GDK_KEY_L                           = 0x04c
	GDK_KEY_M                           = 0x04d
	GDK_KEY_N                           = 0x04e
	GDK_KEY_O                           = 0x04f
	GDK_KEY_P                           = 0x050
	GDK_KEY_Q                           = 0x051
	GDK_KEY_R                           = 0x052
	GDK_KEY_S                           = 0x053
	GDK_KEY_T                           = 0x054
	GDK_KEY_U                           = 0x055
	GDK_KEY_V                           = 0x056
	GDK_KEY_W                           = 0x057
	GDK_KEY_X                           = 0x058
	GDK_KEY_Y                           = 0x059
	GDK_KEY_Z                           = 0x05a
	GDK_KEY_bracketleft                 = 0x05b
	GDK_KEY_backslash                   = 0x05c
	GDK_KEY_bracketright                = 0x05d
	GDK_KEY_asciicircum                 = 0x05e
	GDK_KEY_underscore                  = 0x05f
	GDK_KEY_grave                       = 0x060
	GDK_KEY_quoteleft                   = 0x060
	GDK_KEY_a                           = 0x061
	GDK_KEY_b                           = 0x062
	GDK_KEY_c                           = 0x063
	GDK_KEY_d                           = 0x064
	GDK_KEY_e                           = 0x065
	GDK_KEY_f                           = 0x066
	GDK_KEY_g                           = 0x067
	GDK_KEY_h                           = 0x068
	GDK_KEY_i                           = 0x069
	GDK_KEY_j                           = 0x06a
	GDK_KEY_k                           = 0x06b
	GDK_KEY_l                           = 0x06c
	GDK_KEY_m                           = 0x06d
	GDK_KEY_n                           = 0x06e
	GDK_KEY_o                           = 0x06f
	GDK_KEY_p                           = 0x070
	GDK_KEY_q                           = 0x071
	GDK_KEY_r                           = 0x072
	GDK_KEY_s                           = 0x073
	GDK_KEY_t                           = 0x074
	GDK_KEY_u                           = 0x075
	GDK_KEY_v                           = 0x076
	GDK_KEY_w                           = 0x077
	GDK_KEY_x                           = 0x078
	GDK_KEY_y                           = 0x079
	GDK_KEY_z                           = 0x07a
	GDK_KEY_braceleft                   = 0x07b
	GDK_KEY_bar                         = 0x07c
	GDK_KEY_braceright                  = 0x07d
	GDK_KEY_asciitilde                  = 0x07e
	GDK_KEY_nobreakspace                = 0x0a0
	GDK_KEY_exclamdown                  = 0x0a1
	GDK_KEY_cent                        = 0x0a2
	GDK_KEY_sterling                    = 0x0a3
	GDK_KEY_currency                    = 0x0a4
	GDK_KEY_yen                         = 0x0a5
	GDK_KEY_brokenbar                   = 0x0a6
	GDK_KEY_section                     = 0x0a7
	GDK_KEY_diaeresis                   = 0x0a8
	GDK_KEY_copyright                   = 0x0a9
	GDK_KEY_ordfeminine                 = 0x0aa
	GDK_KEY_guillemotleft               = 0x0ab
	GDK_KEY_notsign                     = 0x0ac
	GDK_KEY_hyphen                      = 0x0ad
	GDK_KEY_registered                  = 0x0ae
	GDK_KEY_macron                      = 0x0af
	GDK_KEY_degree                      = 0x0b0
	GDK_KEY_plusminus                   = 0x0b1
	GDK_KEY_twosuperior                 = 0x0b2
	GDK_KEY_threesuperior               = 0x0b3
	GDK_KEY_acute                       = 0x0b4
	GDK_KEY_mu                          = 0x0b5
	GDK_KEY_paragraph                   = 0x0b6
	GDK_KEY_periodcentered              = 0x0b7
	GDK_KEY_cedilla                     = 0x0b8
	GDK_KEY_onesuperior                 = 0x0b9
	GDK_KEY_masculine                   = 0x0ba
	GDK_KEY_guillemotright              = 0x0bb
	GDK_KEY_onequarter                  = 0x0bc
	GDK_KEY_onehalf                     = 0x0bd
	GDK_KEY_threequarters               = 0x0be
	GDK_KEY_questiondown                = 0x0bf
	GDK_KEY_Agrave                      = 0x0c0
	GDK_KEY_Aacute                      = 0x0c1
	GDK_KEY_Acircumflex                 = 0x0c2
	GDK_KEY_Atilde                      = 0x0c3
	GDK_KEY_Adiaeresis                  = 0x0c4
	GDK_KEY_Aring                       = 0x0c5
	GDK_KEY_AE                          = 0x0c6
	GDK_KEY_Ccedilla                    = 0x0c7
	GDK_KEY_Egrave                      = 0x0c8
	GDK_KEY_Eacute                      = 0x0c9
	GDK_KEY_Ecircumflex                 = 0x0ca
	GDK_KEY_Ediaeresis                  = 0x0cb
	GDK_KEY_Igrave                      = 0x0cc
	GDK_KEY_Iacute                      = 0x0cd
	GDK_KEY_Icircumflex                 = 0x0ce
	GDK_KEY_Idiaeresis                  = 0x0cf
	GDK_KEY_ETH                         = 0x0d0
	GDK_KEY_Eth                         = 0x0d0
	GDK_KEY_Ntilde                      = 0x0d1
	GDK_KEY_Ograve                      = 0x0d2
	GDK_KEY_Oacute                      = 0x0d3
	GDK_KEY_Ocircumflex                 = 0x0d4
	GDK_KEY_Otilde                      = 0x0d5
	GDK_KEY_Odiaeresis                  = 0x0d6
	GDK_KEY_multiply                    = 0x0d7
	GDK_KEY_Oslash                      = 0x0d8
	GDK_KEY_Ooblique                    = 0x0d8
	GDK_KEY_Ugrave                      = 0x0d9
	GDK_KEY_Uacute                      = 0x0da
	GDK_KEY_Ucircumflex                 = 0x0db
	GDK_KEY_Udiaeresis                  = 0x0dc
	GDK_KEY_Yacute                      = 0x0dd
	GDK_KEY_THORN                       = 0x0de
	GDK_KEY_Thorn                       = 0x0de
	GDK_KEY_ssharp                      = 0x0df
	GDK_KEY_agrave                      = 0x0e0
	GDK_KEY_aacute                      = 0x0e1
	GDK_KEY_acircumflex                 = 0x0e2
	GDK_KEY_atilde                      = 0x0e3
	GDK_KEY_adiaeresis                  = 0x0e4
	GDK_KEY_aring                       = 0x0e5
	GDK_KEY_ae                          = 0x0e6
	GDK_KEY_ccedilla                    = 0x0e7
	GDK_KEY_egrave                      = 0x0e8
	GDK_KEY_eacute                      = 0x0e9
	GDK_KEY_ecircumflex                 = 0x0ea
	GDK_KEY_ediaeresis                  = 0x0eb
	GDK_KEY_igrave                      = 0x0ec
	GDK_KEY_iacute                      = 0x0ed
	GDK_KEY_icircumflex                 = 0x0ee
	GDK_KEY_idiaeresis                  = 0x0ef
	GDK_KEY_eth                         = 0x0f0
	GDK_KEY_ntilde                      = 0x0f1
	GDK_KEY_ograve                      = 0x0f2
	GDK_KEY_oacute                      = 0x0f3
	GDK_KEY_ocircumflex                 = 0x0f4
	GDK_KEY_otilde                      = 0x0f5
	GDK_KEY_odiaeresis                  = 0x0f6
	GDK_KEY_division                    = 0x0f7
	GDK_KEY_oslash                      = 0x0f8
	GDK_KEY_ooblique                    = 0x0f8
	GDK_KEY_ugrave                      = 0x0f9
	GDK_KEY_uacute                      = 0x0fa
	GDK_KEY_ucircumflex                 = 0x0fb
	GDK_KEY_udiaeresis                  = 0x0fc
	GDK_KEY_yacute                      = 0x0fd
	GDK_KEY_thorn                       = 0x0fe
	GDK_KEY_ydiaeresis                  = 0x0ff
	GDK_KEY_Aogonek                     = 0x1a1
	GDK_KEY_breve                       = 0x1a2
	GDK_KEY_Lstroke                     = 0x1a3
	GDK_KEY_Lcaron                      = 0x1a5
	GDK_KEY_Sacute                      = 0x1a6
	GDK_KEY_Scaron                      = 0x1a9
	GDK_KEY_Scedilla                    = 0x1aa
	GDK_KEY_Tcaron                      = 0x1ab
	GDK_KEY_Zacute                      = 0x1ac
	GDK_KEY_Zcaron                      = 0x1ae
	GDK_KEY_Zabovedot                   = 0x1af
	GDK_KEY_aogonek                     = 0x1b1
	GDK_KEY_ogonek                      = 0x1b2
	GDK_KEY_lstroke                     = 0x1b3
	GDK_KEY_lcaron                      = 0x1b5
	GDK_KEY_sacute                      = 0x1b6
	GDK_KEY_caron                       = 0x1b7
	GDK_KEY_scaron                      = 0x1b9
	GDK_KEY_scedilla                    = 0x1ba
	GDK_KEY_tcaron                      = 0x1bb
	GDK_KEY_zacute                      = 0x1bc
	GDK_KEY_doubleacute                 = 0x1bd
	GDK_KEY_zcaron                      = 0x1be
	GDK_KEY_zabovedot                   = 0x1bf
	GDK_KEY_Racute                      = 0x1c0
	GDK_KEY_Abreve                      = 0x1c3
	GDK_KEY_Lacute                      = 0x1c5
	GDK_KEY_Cacute                      = 0x1c6
	GDK_KEY_Ccaron                      = 0x1c8
	GDK_KEY_Eogonek                     = 0x1ca
	GDK_KEY_Ecaron                      = 0x1cc
	GDK_KEY_Dcaron                      = 0x1cf
	GDK_KEY_Dstroke                     = 0x1d0
	GDK_KEY_Nacute                      = 0x1d1
	GDK_KEY_Ncaron                      = 0x1d2
	GDK_KEY_Odoubleacute                = 0x1d5
	GDK_KEY_Rcaron                      = 0x1d8
	GDK_KEY_Uring                       = 0x1d9
	GDK_KEY_Udoubleacute                = 0x1db
	GDK_KEY_Tcedilla                    = 0x1de
	GDK_KEY_racute                      = 0x1e0
	GDK_KEY_abreve                      = 0x1e3
	GDK_KEY_lacute                      = 0x1e5
	GDK_KEY_cacute                      = 0x1e6
	GDK_KEY_ccaron                      = 0x1e8
	GDK_KEY_eogonek                     = 0x1ea
	GDK_KEY_ecaron                      = 0x1ec
	GDK_KEY_dcaron                      = 0x1ef
	GDK_KEY_dstroke                     = 0x1f0
	GDK_KEY_nacute                      = 0x1f1
	GDK_KEY_ncaron                      = 0x1f2
	GDK_KEY_odoubleacute                = 0x1f5
	GDK_KEY_udoubleacute                = 0x1fb
	GDK_KEY_rcaron                      = 0x1f8
	GDK_KEY_uring                       = 0x1f9
	GDK_KEY_tcedilla                    = 0x1fe
	GDK_KEY_abovedot                    = 0x1ff
	GDK_KEY_Hstroke                     = 0x2a1
	GDK_KEY_Hcircumflex                 = 0x2a6
	GDK_KEY_Iabovedot                   = 0x2a9
	GDK_KEY_Gbreve                      = 0x2ab
	GDK_KEY_Jcircumflex                 = 0x2ac
	GDK_KEY_hstroke                     = 0x2b1
	GDK_KEY_hcircumflex                 = 0x2b6
	GDK_KEY_idotless                    = 0x2b9
	GDK_KEY_gbreve                      = 0x2bb
	GDK_KEY_jcircumflex                 = 0x2bc
	GDK_KEY_Cabovedot                   = 0x2c5
	GDK_KEY_Ccircumflex                 = 0x2c6
	GDK_KEY_Gabovedot                   = 0x2d5
	GDK_KEY_Gcircumflex                 = 0x2d8
	GDK_KEY_Ubreve                      = 0x2dd
	GDK_KEY_Scircumflex                 = 0x2de
	GDK_KEY_cabovedot                   = 0x2e5
	GDK_KEY_ccircumflex                 = 0x2e6
	GDK_KEY_gabovedot                   = 0x2f5
	GDK_KEY_gcircumflex                 = 0x2f8
	GDK_KEY_ubreve                      = 0x2fd
	GDK_KEY_scircumflex                 = 0x2fe
	GDK_KEY_kra                         = 0x3a2
	GDK_KEY_kappa                       = 0x3a2
	GDK_KEY_Rcedilla                    = 0x3a3
	GDK_KEY_Itilde                      = 0x3a5
	GDK_KEY_Lcedilla                    = 0x3a6
	GDK_KEY_Emacron                     = 0x3aa
	GDK_KEY_Gcedilla                    = 0x3ab
	GDK_KEY_Tslash                      = 0x3ac
	GDK_KEY_rcedilla                    = 0x3b3
	GDK_KEY_itilde                      = 0x3b5
	GDK_KEY_lcedilla                    = 0x3b6
	GDK_KEY_emacron                     = 0x3ba
	GDK_KEY_gcedilla                    = 0x3bb
	GDK_KEY_tslash                      = 0x3bc
	GDK_KEY_ENG                         = 0x3bd
	GDK_KEY_eng                         = 0x3bf
	GDK_KEY_Amacron                     = 0x3c0
	GDK_KEY_Iogonek                     = 0x3c7
	GDK_KEY_Eabovedot                   = 0x3cc
	GDK_KEY_Imacron                     = 0x3cf
	GDK_KEY_Ncedilla                    = 0x3d1
	GDK_KEY_Omacron                     = 0x3d2
	GDK_KEY_Kcedilla                    = 0x3d3
	GDK_KEY_Uogonek                     = 0x3d9
	GDK_KEY_Utilde                      = 0x3dd
	GDK_KEY_Umacron                     = 0x3de
	GDK_KEY_amacron                     = 0x3e0
	GDK_KEY_iogonek                     = 0x3e7
	GDK_KEY_eabovedot                   = 0x3ec
	GDK_KEY_imacron                     = 0x3ef
	GDK_KEY_ncedilla                    = 0x3f1
	GDK_KEY_omacron                     = 0x3f2
	GDK_KEY_kcedilla                    = 0x3f3
	GDK_KEY_uogonek                     = 0x3f9
	GDK_KEY_utilde                      = 0x3fd
	GDK_KEY_umacron                     = 0x3fe
	GDK_KEY_Babovedot                   = 0x1001e02
	GDK_KEY_babovedot                   = 0x1001e03
	GDK_KEY_Dabovedot                   = 0x1001e0a
	GDK_KEY_Wgrave                      = 0x1001e80
	GDK_KEY_Wacute                      = 0x1001e82
	GDK_KEY_dabovedot                   = 0x1001e0b
	GDK_KEY_Ygrave                      = 0x1001ef2
	GDK_KEY_Fabovedot                   = 0x1001e1e
	GDK_KEY_fabovedot                   = 0x1001e1f
	GDK_KEY_Mabovedot                   = 0x1001e40
	GDK_KEY_mabovedot                   = 0x1001e41
	GDK_KEY_Pabovedot                   = 0x1001e56
	GDK_KEY_wgrave                      = 0x1001e81
	GDK_KEY_pabovedot                   = 0x1001e57
	GDK_KEY_wacute                      = 0x1001e83
	GDK_KEY_Sabovedot                   = 0x1001e60
	GDK_KEY_ygrave                      = 0x1001ef3
	GDK_KEY_Wdiaeresis                  = 0x1001e84
	GDK_KEY_wdiaeresis                  = 0x1001e85
	GDK_KEY_sabovedot                   = 0x1001e61
	GDK_KEY_Wcircumflex                 = 0x1000174
	GDK_KEY_Tabovedot                   = 0x1001e6a
	GDK_KEY_Ycircumflex                 = 0x1000176
	GDK_KEY_wcircumflex                 = 0x1000175
	GDK_KEY_tabovedot                   = 0x1001e6b
	GDK_KEY_ycircumflex                 = 0x1000177
	GDK_KEY_OE                          = 0x13bc
	GDK_KEY_oe                          = 0x13bd
	GDK_KEY_Ydiaeresis                  = 0x13be
	GDK_KEY_overline                    = 0x47e
	GDK_KEY_kana_fullstop               = 0x4a1
	GDK_KEY_kana_openingbracket         = 0x4a2
	GDK_KEY_kana_closingbracket         = 0x4a3
	GDK_KEY_kana_comma                  = 0x4a4
	GDK_KEY_kana_conjunctive            = 0x4a5
	GDK_KEY_kana_middledot              = 0x4a5
	GDK_KEY_kana_WO                     = 0x4a6
	GDK_KEY_kana_a                      = 0x4a7
	GDK_KEY_kana_i                      = 0x4a8
	GDK_KEY_kana_u                      = 0x4a9
	GDK_KEY_kana_e                      = 0x4aa
	GDK_KEY_kana_o                      = 0x4ab
	GDK_KEY_kana_ya                     = 0x4ac
	GDK_KEY_kana_yu                     = 0x4ad
	GDK_KEY_kana_yo                     = 0x4ae
	GDK_KEY_kana_tsu                    = 0x4af
	GDK_KEY_kana_tu                     = 0x4af
	GDK_KEY_prolongedsound              = 0x4b0
	GDK_KEY_kana_A                      = 0x4b1
	GDK_KEY_kana_I                      = 0x4b2
	GDK_KEY_kana_U                      = 0x4b3
	GDK_KEY_kana_E                      = 0x4b4
	GDK_KEY_kana_O                      = 0x4b5
	GDK_KEY_kana_KA                     = 0x4b6
	GDK_KEY_kana_KI                     = 0x4b7
	GDK_KEY_kana_KU                     = 0x4b8
	GDK_KEY_kana_KE                     = 0x4b9
	GDK_KEY_kana_KO                     = 0x4ba
	GDK_KEY_kana_SA                     = 0x4bb
	GDK_KEY_kana_SHI                    = 0x4bc
	GDK_KEY_kana_SU                     = 0x4bd
	GDK_KEY_kana_SE                     = 0x4be
	GDK_KEY_kana_SO                     = 0x4bf
	GDK_KEY_kana_TA                     = 0x4c0
	GDK_KEY_kana_CHI                    = 0x4c1
	GDK_KEY_kana_TI                     = 0x4c1
	GDK_KEY_kana_TSU                    = 0x4c2
	GDK_KEY_kana_TU                     = 0x4c2
	GDK_KEY_kana_TE                     = 0x4c3
	GDK_KEY_kana_TO                     = 0x4c4
	GDK_KEY_kana_NA                     = 0x4c5
	GDK_KEY_kana_NI                     = 0x4c6
	GDK_KEY_kana_NU                     = 0x4c7
	GDK_KEY_kana_NE                     = 0x4c8
	GDK_KEY_kana_NO                     = 0x4c9
	GDK_KEY_kana_HA                     = 0x4ca
	GDK_KEY_kana_HI                     = 0x4cb
	GDK_KEY_kana_FU                     = 0x4cc
	GDK_KEY_kana_HU                     = 0x4cc
	GDK_KEY_kana_HE                     = 0x4cd
	GDK_KEY_kana_HO                     = 0x4ce
	GDK_KEY_kana_MA                     = 0x4cf
	GDK_KEY_kana_MI                     = 0x4d0
	GDK_KEY_kana_MU                     = 0x4d1
	GDK_KEY_kana_ME                     = 0x4d2
	GDK_KEY_kana_MO                     = 0x4d3
	GDK_KEY_kana_YA                     = 0x4d4
	GDK_KEY_kana_YU                     = 0x4d5
	GDK_KEY_kana_YO                     = 0x4d6
	GDK_KEY_kana_RA                     = 0x4d7
	GDK_KEY_kana_RI                     = 0x4d8
	GDK_KEY_kana_RU                     = 0x4d9
	GDK_KEY_kana_RE                     = 0x4da
	GDK_KEY_kana_RO                     = 0x4db
	GDK_KEY_kana_WA                     = 0x4dc
	GDK_KEY_kana_N                      = 0x4dd
	GDK_KEY_voicedsound                 = 0x4de
	GDK_KEY_semivoicedsound             = 0x4df
	GDK_KEY_kana_switch                 = 0xff7e
	GDK_KEY_Farsi_0                     = 0x10006f0
	GDK_KEY_Farsi_1                     = 0x10006f1
	GDK_KEY_Farsi_2                     = 0x10006f2
	GDK_KEY_Farsi_3                     = 0x10006f3
	GDK_KEY_Farsi_4                     = 0x10006f4
	GDK_KEY_Farsi_5                     = 0x10006f5
	GDK_KEY_Farsi_6                     = 0x10006f6
	GDK_KEY_Farsi_7                     = 0x10006f7
	GDK_KEY_Farsi_8                     = 0x10006f8
	GDK_KEY_Farsi_9                     = 0x10006f9
	GDK_KEY_Arabic_percent              = 0x100066a
	GDK_KEY_Arabic_superscript_alef     = 0x1000670
	GDK_KEY_Arabic_tteh                 = 0x1000679
	GDK_KEY_Arabic_peh                  = 0x100067e
	GDK_KEY_Arabic_tcheh                = 0x1000686
	GDK_KEY_Arabic_ddal                 = 0x1000688
	GDK_KEY_Arabic_rreh                 = 0x1000691
	GDK_KEY_Arabic_comma                = 0x5ac
	GDK_KEY_Arabic_fullstop             = 0x10006d4
	GDK_KEY_Arabic_0                    = 0x1000660
	GDK_KEY_Arabic_1                    = 0x1000661
	GDK_KEY_Arabic_2                    = 0x1000662
	GDK_KEY_Arabic_3                    = 0x1000663
	GDK_KEY_Arabic_4                    = 0x1000664
	GDK_KEY_Arabic_5                    = 0x1000665
	GDK_KEY_Arabic_6                    = 0x1000666
	GDK_KEY_Arabic_7                    = 0x1000667
	GDK_KEY_Arabic_8                    = 0x1000668
	GDK_KEY_Arabic_9                    = 0x1000669
	GDK_KEY_Arabic_semicolon            = 0x5bb
	GDK_KEY_Arabic_question_mark        = 0x5bf
	GDK_KEY_Arabic_hamza                = 0x5c1
	GDK_KEY_Arabic_maddaonalef          = 0x5c2
	GDK_KEY_Arabic_hamzaonalef          = 0x5c3
	GDK_KEY_Arabic_hamzaonwaw           = 0x5c4
	GDK_KEY_Arabic_hamzaunderalef       = 0x5c5
	GDK_KEY_Arabic_hamzaonyeh           = 0x5c6
	GDK_KEY_Arabic_alef                 = 0x5c7
	GDK_KEY_Arabic_beh                  = 0x5c8
	GDK_KEY_Arabic_tehmarbuta           = 0x5c9
	GDK_KEY_Arabic_teh                  = 0x5ca
	GDK_KEY_Arabic_theh                 = 0x5cb
	GDK_KEY_Arabic_jeem                 = 0x5cc
	GDK_KEY_Arabic_hah                  = 0x5cd
	GDK_KEY_Arabic_khah                 = 0x5ce
	GDK_KEY_Arabic_dal                  = 0x5cf
	GDK_KEY_Arabic_thal                 = 0x5d0
	GDK_KEY_Arabic_ra                   = 0x5d1
	GDK_KEY_Arabic_zain                 = 0x5d2
	GDK_KEY_Arabic_seen                 = 0x5d3
	GDK_KEY_Arabic_sheen                = 0x5d4
	GDK_KEY_Arabic_sad                  = 0x5d5
	GDK_KEY_Arabic_dad                  = 0x5d6
	GDK_KEY_Arabic_tah                  = 0x5d7
	GDK_KEY_Arabic_zah                  = 0x5d8
	GDK_KEY_Arabic_ain                  = 0x5d9
	GDK_KEY_Arabic_ghain                = 0x5da
	GDK_KEY_Arabic_tatweel              = 0x5e0
	GDK_KEY_Arabic_feh                  = 0x5e1
	GDK_KEY_Arabic_qaf                  = 0x5e2
	GDK_KEY_Arabic_kaf                  = 0x5e3
	GDK_KEY_Arabic_lam                  = 0x5e4
	GDK_KEY_Arabic_meem                 = 0x5e5
	GDK_KEY_Arabic_noon                 = 0x5e6
	GDK_KEY_Arabic_ha                   = 0x5e7
	GDK_KEY_Arabic_heh                  = 0x5e7
	GDK_KEY_Arabic_waw                  = 0x5e8
	GDK_KEY_Arabic_alefmaksura          = 0x5e9
	GDK_KEY_Arabic_yeh                  = 0x5ea
	GDK_KEY_Arabic_fathatan             = 0x5eb
	GDK_KEY_Arabic_dammatan             = 0x5ec
	GDK_KEY_Arabic_kasratan             = 0x5ed
	GDK_KEY_Arabic_fatha                = 0x5ee
	GDK_KEY_Arabic_damma                = 0x5ef
	GDK_KEY_Arabic_kasra                = 0x5f0
	GDK_KEY_Arabic_shadda               = 0x5f1
	GDK_KEY_Arabic_sukun                = 0x5f2
	GDK_KEY_Arabic_madda_above          = 0x1000653
	GDK_KEY_Arabic_hamza_above          = 0x1000654
	GDK_KEY_Arabic_hamza_below          = 0x1000655
	GDK_KEY_Arabic_jeh                  = 0x1000698
	GDK_KEY_Arabic_veh                  = 0x10006a4
	GDK_KEY_Arabic_keheh                = 0x10006a9
	GDK_KEY_Arabic_gaf                  = 0x10006af
	GDK_KEY_Arabic_noon_ghunna          = 0x10006ba
	GDK_KEY_Arabic_heh_doachashmee      = 0x10006be
	GDK_KEY_Farsi_yeh                   = 0x10006cc
	GDK_KEY_Arabic_farsi_yeh            = 0x10006cc
	GDK_KEY_Arabic_yeh_baree            = 0x10006d2
	GDK_KEY_Arabic_heh_goal             = 0x10006c1
	GDK_KEY_Arabic_switch               = 0xff7e
	GDK_KEY_Cyrillic_GHE_bar            = 0x1000492
	GDK_KEY_Cyrillic_ghe_bar            = 0x1000493
	GDK_KEY_Cyrillic_ZHE_descender      = 0x1000496
	GDK_KEY_Cyrillic_zhe_descender      = 0x1000497
	GDK_KEY_Cyrillic_KA_descender       = 0x100049a
	GDK_KEY_Cyrillic_ka_descender       = 0x100049b
	GDK_KEY_Cyrillic_KA_vertstroke      = 0x100049c
	GDK_KEY_Cyrillic_ka_vertstroke      = 0x100049d
	GDK_KEY_Cyrillic_EN_descender       = 0x10004a2
	GDK_KEY_Cyrillic_en_descender       = 0x10004a3
	GDK_KEY_Cyrillic_U_straight         = 0x10004ae
	GDK_KEY_Cyrillic_u_straight         = 0x10004af
	GDK_KEY_Cyrillic_U_straight_bar     = 0x10004b0
	GDK_KEY_Cyrillic_u_straight_bar     = 0x10004b1
	GDK_KEY_Cyrillic_HA_descender       = 0x10004b2
	GDK_KEY_Cyrillic_ha_descender       = 0x10004b3
	GDK_KEY_Cyrillic_CHE_descender      = 0x10004b6
	GDK_KEY_Cyrillic_che_descender      = 0x10004b7
	GDK_KEY_Cyrillic_CHE_vertstroke     = 0x10004b8
	GDK_KEY_Cyrillic_che_vertstroke     = 0x10004b9
	GDK_KEY_Cyrillic_SHHA               = 0x10004ba
	GDK_KEY_Cyrillic_shha               = 0x10004bb
	GDK_KEY_Cyrillic_SCHWA              = 0x10004d8
	GDK_KEY_Cyrillic_schwa              = 0x10004d9
	GDK_KEY_Cyrillic_I_macron           = 0x10004e2
	GDK_KEY_Cyrillic_i_macron           = 0x10004e3
	GDK_KEY_Cyrillic_O_bar              = 0x10004e8
	GDK_KEY_Cyrillic_o_bar              = 0x10004e9
	GDK_KEY_Cyrillic_U_macron           = 0x10004ee
	GDK_KEY_Cyrillic_u_macron           = 0x10004ef
	GDK_KEY_Serbian_dje                 = 0x6a1
	GDK_KEY_Macedonia_gje               = 0x6a2
	GDK_KEY_Cyrillic_io                 = 0x6a3
	GDK_KEY_Ukrainian_ie                = 0x6a4
	GDK_KEY_Ukranian_je                 = 0x6a4
	GDK_KEY_Macedonia_dse               = 0x6a5
	GDK_KEY_Ukrainian_i                 = 0x6a6
	GDK_KEY_Ukranian_i                  = 0x6a6
	GDK_KEY_Ukrainian_yi                = 0x6a7
	GDK_KEY_Ukranian_yi                 = 0x6a7
	GDK_KEY_Cyrillic_je                 = 0x6a8
	GDK_KEY_Serbian_je                  = 0x6a8
	GDK_KEY_Cyrillic_lje                = 0x6a9
	GDK_KEY_Serbian_lje                 = 0x6a9
	GDK_KEY_Cyrillic_nje                = 0x6aa
	GDK_KEY_Serbian_nje                 = 0x6aa
	GDK_KEY_Serbian_tshe                = 0x6ab
	GDK_KEY_Macedonia_kje               = 0x6ac
	GDK_KEY_Ukrainian_ghe_with_upturn   = 0x6ad
	GDK_KEY_Byelorussian_shortu         = 0x6ae
	GDK_KEY_Cyrillic_dzhe               = 0x6af
	GDK_KEY_Serbian_dze                 = 0x6af
	GDK_KEY_numerosign                  = 0x6b0
	GDK_KEY_Serbian_DJE                 = 0x6b1
	GDK_KEY_Macedonia_GJE               = 0x6b2
	GDK_KEY_Cyrillic_IO                 = 0x6b3
	GDK_KEY_Ukrainian_IE                = 0x6b4
	GDK_KEY_Ukranian_JE                 = 0x6b4
	GDK_KEY_Macedonia_DSE               = 0x6b5
	GDK_KEY_Ukrainian_I                 = 0x6b6
	GDK_KEY_Ukranian_I                  = 0x6b6
	GDK_KEY_Ukrainian_YI                = 0x6b7
	GDK_KEY_Ukranian_YI                 = 0x6b7
	GDK_KEY_Cyrillic_JE                 = 0x6b8
	GDK_KEY_Serbian_JE                  = 0x6b8
	GDK_KEY_Cyrillic_LJE                = 0x6b9
	GDK_KEY_Serbian_LJE                 = 0x6b9
	GDK_KEY_Cyrillic_NJE                = 0x6ba
	GDK_KEY_Serbian_NJE                 = 0x6ba
	GDK_KEY_Serbian_TSHE                = 0x6bb
	GDK_KEY_Macedonia_KJE               = 0x6bc
	GDK_KEY_Ukrainian_GHE_WITH_UPTURN   = 0x6bd
	GDK_KEY_Byelorussian_SHORTU         = 0x6be
	GDK_KEY_Cyrillic_DZHE               = 0x6bf
	GDK_KEY_Serbian_DZE                 = 0x6bf
	GDK_KEY_Cyrillic_yu                 = 0x6c0
	GDK_KEY_Cyrillic_a                  = 0x6c1
	GDK_KEY_Cyrillic_be                 = 0x6c2
	GDK_KEY_Cyrillic_tse                = 0x6c3
	GDK_KEY_Cyrillic_de                 = 0x6c4
	GDK_KEY_Cyrillic_ie                 = 0x6c5
	GDK_KEY_Cyrillic_ef                 = 0x6c6
	GDK_KEY_Cyrillic_ghe                = 0x6c7
	GDK_KEY_Cyrillic_ha                 = 0x6c8
	GDK_KEY_Cyrillic_i                  = 0x6c9
	GDK_KEY_Cyrillic_shorti             = 0x6ca
	GDK_KEY_Cyrillic_ka                 = 0x6cb
	GDK_KEY_Cyrillic_el                 = 0x6cc
	GDK_KEY_Cyrillic_em                 = 0x6cd
	GDK_KEY_Cyrillic_en                 = 0x6ce
	GDK_KEY_Cyrillic_o                  = 0x6cf
	GDK_KEY_Cyrillic_pe                 = 0x6d0
	GDK_KEY_Cyrillic_ya                 = 0x6d1
	GDK_KEY_Cyrillic_er                 = 0x6d2
	GDK_KEY_Cyrillic_es                 = 0x6d3
	GDK_KEY_Cyrillic_te                 = 0x6d4
	GDK_KEY_Cyrillic_u                  = 0x6d5
	GDK_KEY_Cyrillic_zhe                = 0x6d6
	GDK_KEY_Cyrillic_ve                 = 0x6d7
	GDK_KEY_Cyrillic_softsign           = 0x6d8
	GDK_KEY_Cyrillic_yeru               = 0x6d9
	GDK_KEY_Cyrillic_ze                 = 0x6da
	GDK_KEY_Cyrillic_sha                = 0x6db
	GDK_KEY_Cyrillic_e                  = 0x6dc
	GDK_KEY_Cyrillic_shcha              = 0x6dd
	GDK_KEY_Cyrillic_che                = 0x6de
	GDK_KEY_Cyrillic_hardsign           = 0x6df
	GDK_KEY_Cyrillic_YU                 = 0x6e0
	GDK_KEY_Cyrillic_A                  = 0x6e1
	GDK_KEY_Cyrillic_BE                 = 0x6e2
	GDK_KEY_Cyrillic_TSE                = 0x6e3
	GDK_KEY_Cyrillic_DE                 = 0x6e4
	GDK_KEY_Cyrillic_IE                 = 0x6e5
	GDK_KEY_Cyrillic_EF                 = 0x6e6
	GDK_KEY_Cyrillic_GHE                = 0x6e7
	GDK_KEY_Cyrillic_HA                 = 0x6e8
	GDK_KEY_Cyrillic_I                  = 0x6e9
	GDK_KEY_Cyrillic_SHORTI             = 0x6ea
	GDK_KEY_Cyrillic_KA                 = 0x6eb
	GDK_KEY_Cyrillic_EL                 = 0x6ec
	GDK_KEY_Cyrillic_EM                 = 0x6ed
	GDK_KEY_Cyrillic_EN                 = 0x6ee
	GDK_KEY_Cyrillic_O                  = 0x6ef
	GDK_KEY_Cyrillic_PE                 = 0x6f0
	GDK_KEY_Cyrillic_YA                 = 0x6f1
	GDK_KEY_Cyrillic_ER                 = 0x6f2
	GDK_KEY_Cyrillic_ES                 = 0x6f3
	GDK_KEY_Cyrillic_TE                 = 0x6f4
	GDK_KEY_Cyrillic_U                  = 0x6f5
	GDK_KEY_Cyrillic_ZHE                = 0x6f6
	GDK_KEY_Cyrillic_VE                 = 0x6f7
	GDK_KEY_Cyrillic_SOFTSIGN           = 0x6f8
	GDK_KEY_Cyrillic_YERU               = 0x6f9
	GDK_KEY_Cyrillic_ZE                 = 0x6fa
	GDK_KEY_Cyrillic_SHA                = 0x6fb
	GDK_KEY_Cyrillic_E                  = 0x6fc
	GDK_KEY_Cyrillic_SHCHA              = 0x6fd
	GDK_KEY_Cyrillic_CHE                = 0x6fe
	GDK_KEY_Cyrillic_HARDSIGN           = 0x6ff
	GDK_KEY_Greek_ALPHAaccent           = 0x7a1
	GDK_KEY_Greek_EPSILONaccent         = 0x7a2
	GDK_KEY_Greek_ETAaccent             = 0x7a3
	GDK_KEY_Greek_IOTAaccent            = 0x7a4
	GDK_KEY_Greek_IOTAdieresis          = 0x7a5
	GDK_KEY_Greek_IOTAdiaeresis         = 0x7a5
	GDK_KEY_Greek_OMICRONaccent         = 0x7a7
	GDK_KEY_Greek_UPSILONaccent         = 0x7a8
	GDK_KEY_Greek_UPSILONdieresis       = 0x7a9
	GDK_KEY_Greek_OMEGAaccent           = 0x7ab
	GDK_KEY_Greek_accentdieresis        = 0x7ae
	GDK_KEY_Greek_horizbar              = 0x7af
	GDK_KEY_Greek_alphaaccent           = 0x7b1
	GDK_KEY_Greek_epsilonaccent         = 0x7b2
	GDK_KEY_Greek_etaaccent             = 0x7b3
	GDK_KEY_Greek_iotaaccent            = 0x7b4
	GDK_KEY_Greek_iotadieresis          = 0x7b5
	GDK_KEY_Greek_iotaaccentdieresis    = 0x7b6
	GDK_KEY_Greek_omicronaccent         = 0x7b7
	GDK_KEY_Greek_upsilonaccent         = 0x7b8
	GDK_KEY_Greek_upsilondieresis       = 0x7b9
	GDK_KEY_Greek_upsilonaccentdieresis = 0x7ba
	GDK_KEY_Greek_omegaaccent           = 0x7bb
	GDK_KEY_Greek_ALPHA                 = 0x7c1
	GDK_KEY_Greek_BETA                  = 0x7c2
	GDK_KEY_Greek_GAMMA                 = 0x7c3
	GDK_KEY_Greek_DELTA                 = 0x7c4
	GDK_KEY_Greek_EPSILON               = 0x7c5
	GDK_KEY_Greek_ZETA                  = 0x7c6
	GDK_KEY_Greek_ETA                   = 0x7c7
	GDK_KEY_Greek_THETA                 = 0x7c8
	GDK_KEY_Greek_IOTA                  = 0x7c9
	GDK_KEY_Greek_KAPPA                 = 0x7ca
	GDK_KEY_Greek_LAMDA                 = 0x7cb
	GDK_KEY_Greek_LAMBDA                = 0x7cb
	GDK_KEY_Greek_MU                    = 0x7cc
	GDK_KEY_Greek_NU                    = 0x7cd
	GDK_KEY_Greek_XI                    = 0x7ce
	GDK_KEY_Greek_OMICRON               = 0x7cf
	GDK_KEY_Greek_PI                    = 0x7d0
	GDK_KEY_Greek_RHO                   = 0x7d1
	GDK_KEY_Greek_SIGMA                 = 0x7d2
	GDK_KEY_Greek_TAU                   = 0x7d4
	GDK_KEY_Greek_UPSILON               = 0x7d5
	GDK_KEY_Greek_PHI                   = 0x7d6
	GDK_KEY_Greek_CHI                   = 0x7d7
	GDK_KEY_Greek_PSI                   = 0x7d8
	GDK_KEY_Greek_OMEGA                 = 0x7d9
	GDK_KEY_Greek_alpha                 = 0x7e1
	GDK_KEY_Greek_beta                  = 0x7e2
	GDK_KEY_Greek_gamma                 = 0x7e3
	GDK_KEY_Greek_delta                 = 0x7e4
	GDK_KEY_Greek_epsilon               = 0x7e5
	GDK_KEY_Greek_zeta                  = 0x7e6
	GDK_KEY_Greek_eta                   = 0x7e7
	GDK_KEY_Greek_theta                 = 0x7e8
	GDK_KEY_Greek_iota                  = 0x7e9
	GDK_KEY_Greek_kappa                 = 0x7ea
	GDK_KEY_Greek_lamda                 = 0x7eb
	GDK_KEY_Greek_lambda                = 0x7eb
	GDK_KEY_Greek_mu                    = 0x7ec
	GDK_KEY_Greek_nu                    = 0x7ed
	GDK_KEY_Greek_xi                    = 0x7ee
	GDK_KEY_Greek_omicron               = 0x7ef
	GDK_KEY_Greek_pi                    = 0x7f0
	GDK_KEY_Greek_rho                   = 0x7f1
	GDK_KEY_Greek_sigma                 = 0x7f2
	GDK_KEY_Greek_finalsmallsigma       = 0x7f3
	GDK_KEY_Greek_tau                   = 0x7f4
	GDK_KEY_Greek_upsilon               = 0x7f5
	GDK_KEY_Greek_phi                   = 0x7f6
	GDK_KEY_Greek_chi                   = 0x7f7
	GDK_KEY_Greek_psi                   = 0x7f8
	GDK_KEY_Greek_omega                 = 0x7f9
	GDK_KEY_Greek_switch                = 0xff7e
	GDK_KEY_leftradical                 = 0x8a1
	GDK_KEY_topleftradical              = 0x8a2
	GDK_KEY_horizconnector              = 0x8a3
	GDK_KEY_topintegral                 = 0x8a4
	GDK_KEY_botintegral                 = 0x8a5
	GDK_KEY_vertconnector               = 0x8a6
	GDK_KEY_topleftsqbracket            = 0x8a7
	GDK_KEY_botleftsqbracket            = 0x8a8
	GDK_KEY_toprightsqbracket           = 0x8a9
	GDK_KEY_botrightsqbracket           = 0x8aa
	GDK_KEY_topleftparens               = 0x8ab
	GDK_KEY_botleftparens               = 0x8ac
	GDK_KEY_toprightparens              = 0x8ad
	GDK_KEY_botrightparens              = 0x8ae
	GDK_KEY_leftmiddlecurlybrace        = 0x8af
	GDK_KEY_rightmiddlecurlybrace       = 0x8b0
	GDK_KEY_topleftsummation            = 0x8b1
	GDK_KEY_botleftsummation            = 0x8b2
	GDK_KEY_topvertsummationconnector   = 0x8b3
	GDK_KEY_botvertsummationconnector   = 0x8b4
	GDK_KEY_toprightsummation           = 0x8b5
	GDK_KEY_botrightsummation           = 0x8b6
	GDK_KEY_rightmiddlesummation        = 0x8b7
	GDK_KEY_lessthanequal               = 0x8bc
	GDK_KEY_notequal                    = 0x8bd
	GDK_KEY_greaterthanequal            = 0x8be
	GDK_KEY_integral                    = 0x8bf
	GDK_KEY_therefore                   = 0x8c0
	GDK_KEY_variation                   = 0x8c1
	GDK_KEY_infinity                    = 0x8c2
	GDK_KEY_nabla                       = 0x8c5
	GDK_KEY_approximate                 = 0x8c8
	GDK_KEY_similarequal                = 0x8c9
	GDK_KEY_ifonlyif                    = 0x8cd
	GDK_KEY_implies                     = 0x8ce
	GDK_KEY_identical                   = 0x8cf
	GDK_KEY_radical                     = 0x8d6
	GDK_KEY_includedin                  = 0x8da
	GDK_KEY_includes                    = 0x8db
	GDK_KEY_intersection                = 0x8dc
	GDK_KEY_union                       = 0x8dd
	GDK_KEY_logicaland                  = 0x8de
	GDK_KEY_logicalor                   = 0x8df
	GDK_KEY_partialderivative           = 0x8ef
	GDK_KEY_function                    = 0x8f6
	GDK_KEY_leftarrow                   = 0x8fb
	GDK_KEY_uparrow                     = 0x8fc
	GDK_KEY_rightarrow                  = 0x8fd
	GDK_KEY_downarrow                   = 0x8fe
	GDK_KEY_blank                       = 0x9df
	GDK_KEY_soliddiamond                = 0x9e0
	GDK_KEY_checkerboard                = 0x9e1
	GDK_KEY_ht                          = 0x9e2
	GDK_KEY_ff                          = 0x9e3
	GDK_KEY_cr                          = 0x9e4
	GDK_KEY_lf                          = 0x9e5
	GDK_KEY_nl                          = 0x9e8
	GDK_KEY_vt                          = 0x9e9
	GDK_KEY_lowrightcorner              = 0x9ea
	GDK_KEY_uprightcorner               = 0x9eb
	GDK_KEY_upleftcorner                = 0x9ec
	GDK_KEY_lowleftcorner               = 0x9ed
	GDK_KEY_crossinglines               = 0x9ee
	GDK_KEY_horizlinescan1              = 0x9ef
	GDK_KEY_horizlinescan3              = 0x9f0
	GDK_KEY_horizlinescan5              = 0x9f1
	GDK_KEY_horizlinescan7              = 0x9f2
	GDK_KEY_horizlinescan9              = 0x9f3
	GDK_KEY_leftt                       = 0x9f4
	GDK_KEY_rightt                      = 0x9f5
	GDK_KEY_bott                        = 0x9f6
	GDK_KEY_topt                        = 0x9f7
	GDK_KEY_vertbar                     = 0x9f8
	GDK_KEY_emspace                     = 0xaa1
	GDK_KEY_enspace                     = 0xaa2
	GDK_KEY_em3space                    = 0xaa3
	GDK_KEY_em4space                    = 0xaa4
	GDK_KEY_digitspace                  = 0xaa5
	GDK_KEY_punctspace                  = 0xaa6
	GDK_KEY_thinspace                   = 0xaa7
	GDK_KEY_hairspace                   = 0xaa8
	GDK_KEY_emdash                      = 0xaa9
	GDK_KEY_endash                      = 0xaaa
	GDK_KEY_signifblank                 = 0xaac
	GDK_KEY_ellipsis                    = 0xaae
	GDK_KEY_doubbaselinedot             = 0xaaf
	GDK_KEY_onethird                    = 0xab0
	GDK_KEY_twothirds                   = 0xab1
	GDK_KEY_onefifth                    = 0xab2
	GDK_KEY_twofifths                   = 0xab3
	GDK_KEY_threefifths                 = 0xab4
	GDK_KEY_fourfifths                  = 0xab5
	GDK_KEY_onesixth                    = 0xab6
	GDK_KEY_fivesixths                  = 0xab7
	GDK_KEY_careof                      = 0xab8
	GDK_KEY_figdash                     = 0xabb
	GDK_KEY_leftanglebracket            = 0xabc
	GDK_KEY_decimalpoint                = 0xabd
	GDK_KEY_rightanglebracket           = 0xabe
	GDK_KEY_marker                      = 0xabf
	GDK_KEY_oneeighth                   = 0xac3
	GDK_KEY_threeeighths                = 0xac4
	GDK_KEY_fiveeighths                 = 0xac5
	GDK_KEY_seveneighths                = 0xac6
	GDK_KEY_trademark                   = 0xac9
	GDK_KEY_signaturemark               = 0xaca
	GDK_KEY_trademarkincircle           = 0xacb
	GDK_KEY_leftopentriangle            = 0xacc
	GDK_KEY_rightopentriangle           = 0xacd
	GDK_KEY_emopencircle                = 0xace
	GDK_KEY_emopenrectangle             = 0xacf
	GDK_KEY_leftsinglequotemark         = 0xad0
	GDK_KEY_rightsinglequotemark        = 0xad1
	GDK_KEY_leftdoublequotemark         = 0xad2
	GDK_KEY_rightdoublequotemark        = 0xad3
	GDK_KEY_prescription                = 0xad4
	GDK_KEY_minutes                     = 0xad6
	GDK_KEY_seconds                     = 0xad7
	GDK_KEY_latincross                  = 0xad9
	GDK_KEY_hexagram                    = 0xada
	GDK_KEY_filledrectbullet            = 0xadb
	GDK_KEY_filledlefttribullet         = 0xadc
	GDK_KEY_filledrighttribullet        = 0xadd
	GDK_KEY_emfilledcircle              = 0xade
	GDK_KEY_emfilledrect                = 0xadf
	GDK_KEY_enopencircbullet            = 0xae0
	GDK_KEY_enopensquarebullet          = 0xae1
	GDK_KEY_openrectbullet              = 0xae2
	GDK_KEY_opentribulletup             = 0xae3
	GDK_KEY_opentribulletdown           = 0xae4
	GDK_KEY_openstar                    = 0xae5
	GDK_KEY_enfilledcircbullet          = 0xae6
	GDK_KEY_enfilledsqbullet            = 0xae7
	GDK_KEY_filledtribulletup           = 0xae8
	GDK_KEY_filledtribulletdown         = 0xae9
	GDK_KEY_leftpointer                 = 0xaea
	GDK_KEY_rightpointer                = 0xaeb
	GDK_KEY_club                        = 0xaec
	GDK_KEY_diamond                     = 0xaed
	GDK_KEY_heart                       = 0xaee
	GDK_KEY_maltesecross                = 0xaf0
	GDK_KEY_dagger                      = 0xaf1
	GDK_KEY_doubledagger                = 0xaf2
	GDK_KEY_checkmark                   = 0xaf3
	GDK_KEY_ballotcross                 = 0xaf4
	GDK_KEY_musicalsharp                = 0xaf5
	GDK_KEY_musicalflat                 = 0xaf6
	GDK_KEY_malesymbol                  = 0xaf7
	GDK_KEY_femalesymbol                = 0xaf8
	GDK_KEY_telephone                   = 0xaf9
	GDK_KEY_telephonerecorder           = 0xafa
	GDK_KEY_phonographcopyright         = 0xafb
	GDK_KEY_caret                       = 0xafc
	GDK_KEY_singlelowquotemark          = 0xafd
	GDK_KEY_doublelowquotemark          = 0xafe
	GDK_KEY_cursor                      = 0xaff
	GDK_KEY_leftcaret                   = 0xba3
	GDK_KEY_rightcaret                  = 0xba6
	GDK_KEY_downcaret                   = 0xba8
	GDK_KEY_upcaret                     = 0xba9
	GDK_KEY_overbar                     = 0xbc0
	GDK_KEY_downtack                    = 0xbc2
	GDK_KEY_upshoe                      = 0xbc3
	GDK_KEY_downstile                   = 0xbc4
	GDK_KEY_underbar                    = 0xbc6
	GDK_KEY_jot                         = 0xbca
	GDK_KEY_quad                        = 0xbcc
	GDK_KEY_uptack                      = 0xbce
	GDK_KEY_circle                      = 0xbcf
	GDK_KEY_upstile                     = 0xbd3
	GDK_KEY_downshoe                    = 0xbd6
	GDK_KEY_rightshoe                   = 0xbd8
	GDK_KEY_leftshoe                    = 0xbda
	GDK_KEY_lefttack                    = 0xbdc
	GDK_KEY_righttack                   = 0xbfc
	GDK_KEY_hebrew_doublelowline        = 0xcdf
	GDK_KEY_hebrew_aleph                = 0xce0
	GDK_KEY_hebrew_bet                  = 0xce1
	GDK_KEY_hebrew_beth                 = 0xce1
	GDK_KEY_hebrew_gimel                = 0xce2
	GDK_KEY_hebrew_gimmel               = 0xce2
	GDK_KEY_hebrew_dalet                = 0xce3
	GDK_KEY_hebrew_daleth               = 0xce3
	GDK_KEY_hebrew_he                   = 0xce4
	GDK_KEY_hebrew_waw                  = 0xce5
	GDK_KEY_hebrew_zain                 = 0xce6
	GDK_KEY_hebrew_zayin                = 0xce6
	GDK_KEY_hebrew_chet                 = 0xce7
	GDK_KEY_hebrew_het                  = 0xce7
	GDK_KEY_hebrew_tet                  = 0xce8
	GDK_KEY_hebrew_teth                 = 0xce8
	GDK_KEY_hebrew_yod                  = 0xce9
	GDK_KEY_hebrew_finalkaph            = 0xcea
	GDK_KEY_hebrew_kaph                 = 0xceb
	GDK_KEY_hebrew_lamed                = 0xcec
	GDK_KEY_hebrew_finalmem             = 0xced
	GDK_KEY_hebrew_mem                  = 0xcee
	GDK_KEY_hebrew_finalnun             = 0xcef
	GDK_KEY_hebrew_nun                  = 0xcf0
	GDK_KEY_hebrew_samech               = 0xcf1
	GDK_KEY_hebrew_samekh               = 0xcf1
	GDK_KEY_hebrew_ayin                 = 0xcf2
	GDK_KEY_hebrew_finalpe              = 0xcf3
	GDK_KEY_hebrew_pe                   = 0xcf4
	GDK_KEY_hebrew_finalzade            = 0xcf5
	GDK_KEY_hebrew_finalzadi            = 0xcf5
	GDK_KEY_hebrew_zade                 = 0xcf6
	GDK_KEY_hebrew_zadi                 = 0xcf6
	GDK_KEY_hebrew_qoph                 = 0xcf7
	GDK_KEY_hebrew_kuf                  = 0xcf7
	GDK_KEY_hebrew_resh                 = 0xcf8
	GDK_KEY_hebrew_shin                 = 0xcf9
	GDK_KEY_hebrew_taw                  = 0xcfa
	GDK_KEY_hebrew_taf                  = 0xcfa
	GDK_KEY_Hebrew_switch               = 0xff7e
	GDK_KEY_Thai_kokai                  = 0xda1
	GDK_KEY_Thai_khokhai                = 0xda2
	GDK_KEY_Thai_khokhuat               = 0xda3
	GDK_KEY_Thai_khokhwai               = 0xda4
	GDK_KEY_Thai_khokhon                = 0xda5
	GDK_KEY_Thai_khorakhang             = 0xda6
	GDK_KEY_Thai_ngongu                 = 0xda7
	GDK_KEY_Thai_chochan                = 0xda8
	GDK_KEY_Thai_choching               = 0xda9
	GDK_KEY_Thai_chochang               = 0xdaa
	GDK_KEY_Thai_soso                   = 0xdab
	GDK_KEY_Thai_chochoe                = 0xdac
	GDK_KEY_Thai_yoying                 = 0xdad
	GDK_KEY_Thai_dochada                = 0xdae
	GDK_KEY_Thai_topatak                = 0xdaf
	GDK_KEY_Thai_thothan                = 0xdb0
	GDK_KEY_Thai_thonangmontho          = 0xdb1
	GDK_KEY_Thai_thophuthao             = 0xdb2
	GDK_KEY_Thai_nonen                  = 0xdb3
	GDK_KEY_Thai_dodek                  = 0xdb4
	GDK_KEY_Thai_totao                  = 0xdb5
	GDK_KEY_Thai_thothung               = 0xdb6
	GDK_KEY_Thai_thothahan              = 0xdb7
	GDK_KEY_Thai_thothong               = 0xdb8
	GDK_KEY_Thai_nonu                   = 0xdb9
	GDK_KEY_Thai_bobaimai               = 0xdba
	GDK_KEY_Thai_popla                  = 0xdbb
	GDK_KEY_Thai_phophung               = 0xdbc
	GDK_KEY_Thai_fofa                   = 0xdbd
	GDK_KEY_Thai_phophan                = 0xdbe
	GDK_KEY_Thai_fofan                  = 0xdbf
	GDK_KEY_Thai_phosamphao             = 0xdc0
	GDK_KEY_Thai_moma                   = 0xdc1
	GDK_KEY_Thai_yoyak                  = 0xdc2
	GDK_KEY_Thai_rorua                  = 0xdc3
	GDK_KEY_Thai_ru                     = 0xdc4
	GDK_KEY_Thai_loling                 = 0xdc5
	GDK_KEY_Thai_lu                     = 0xdc6
	GDK_KEY_Thai_wowaen                 = 0xdc7
	GDK_KEY_Thai_sosala                 = 0xdc8
	GDK_KEY_Thai_sorusi                 = 0xdc9
	GDK_KEY_Thai_sosua                  = 0xdca
	GDK_KEY_Thai_hohip                  = 0xdcb
	GDK_KEY_Thai_lochula                = 0xdcc
	GDK_KEY_Thai_oang                   = 0xdcd
	GDK_KEY_Thai_honokhuk               = 0xdce
	GDK_KEY_Thai_paiyannoi              = 0xdcf
	GDK_KEY_Thai_saraa                  = 0xdd0
	GDK_KEY_Thai_maihanakat             = 0xdd1
	GDK_KEY_Thai_saraaa                 = 0xdd2
	GDK_KEY_Thai_saraam                 = 0xdd3
	GDK_KEY_Thai_sarai                  = 0xdd4
	GDK_KEY_Thai_saraii                 = 0xdd5
	GDK_KEY_Thai_saraue                 = 0xdd6
	GDK_KEY_Thai_sarauee                = 0xdd7
	GDK_KEY_Thai_sarau                  = 0xdd8
	GDK_KEY_Thai_sarauu                 = 0xdd9
	GDK_KEY_Thai_phinthu                = 0xdda
	GDK_KEY_Thai_maihanakat_maitho      = 0xdde
	GDK_KEY_Thai_baht                   = 0xddf
	GDK_KEY_Thai_sarae                  = 0xde0
	GDK_KEY_Thai_saraae                 = 0xde1
	GDK_KEY_Thai_sarao                  = 0xde2
	GDK_KEY_Thai_saraaimaimuan          = 0xde3
	GDK_KEY_Thai_saraaimaimalai         = 0xde4
	GDK_KEY_Thai_lakkhangyao            = 0xde5
	GDK_KEY_Thai_maiyamok               = 0xde6
	GDK_KEY_Thai_maitaikhu              = 0xde7
	GDK_KEY_Thai_maiek                  = 0xde8
	GDK_KEY_Thai_maitho                 = 0xde9
	GDK_KEY_Thai_maitri                 = 0xdea
	GDK_KEY_Thai_maichattawa            = 0xdeb
	GDK_KEY_Thai_thanthakhat            = 0xdec
	GDK_KEY_Thai_nikhahit               = 0xded
	GDK_KEY_Thai_leksun                 = 0xdf0
	GDK_KEY_Thai_leknung                = 0xdf1
	GDK_KEY_Thai_leksong                = 0xdf2
	GDK_KEY_Thai_leksam                 = 0xdf3
	GDK_KEY_Thai_leksi                  = 0xdf4
	GDK_KEY_Thai_lekha                  = 0xdf5
	GDK_KEY_Thai_lekhok                 = 0xdf6
	GDK_KEY_Thai_lekchet                = 0xdf7
	GDK_KEY_Thai_lekpaet                = 0xdf8
	GDK_KEY_Thai_lekkao                 = 0xdf9
	GDK_KEY_Hangul                      = 0xff31
	GDK_KEY_Hangul_Start                = 0xff32
	GDK_KEY_Hangul_End                  = 0xff33
	GDK_KEY_Hangul_Hanja                = 0xff34
	GDK_KEY_Hangul_Jamo                 = 0xff35
	GDK_KEY_Hangul_Romaja               = 0xff36
	GDK_KEY_Hangul_Codeinput            = 0xff37
	GDK_KEY_Hangul_Jeonja               = 0xff38
	GDK_KEY_Hangul_Banja                = 0xff39
	GDK_KEY_Hangul_PreHanja             = 0xff3a
	GDK_KEY_Hangul_PostHanja            = 0xff3b
	GDK_KEY_Hangul_SingleCandidate      = 0xff3c
	GDK_KEY_Hangul_MultipleCandidate    = 0xff3d
	GDK_KEY_Hangul_PreviousCandidate    = 0xff3e
	GDK_KEY_Hangul_Special              = 0xff3f
	GDK_KEY_Hangul_switch               = 0xff7e
	GDK_KEY_Hangul_Kiyeog               = 0xea1
	GDK_KEY_Hangul_SsangKiyeog          = 0xea2
	GDK_KEY_Hangul_KiyeogSios           = 0xea3
	GDK_KEY_Hangul_Nieun                = 0xea4
	GDK_KEY_Hangul_NieunJieuj           = 0xea5
	GDK_KEY_Hangul_NieunHieuh           = 0xea6
	GDK_KEY_Hangul_Dikeud               = 0xea7
	GDK_KEY_Hangul_SsangDikeud          = 0xea8
	GDK_KEY_Hangul_Rieul                = 0xea9
	GDK_KEY_Hangul_RieulKiyeog          = 0xeaa
	GDK_KEY_Hangul_RieulMieum           = 0xeab
	GDK_KEY_Hangul_RieulPieub           = 0xeac
	GDK_KEY_Hangul_RieulSios            = 0xead
	GDK_KEY_Hangul_RieulTieut           = 0xeae
	GDK_KEY_Hangul_RieulPhieuf          = 0xeaf
	GDK_KEY_Hangul_RieulHieuh           = 0xeb0
	GDK_KEY_Hangul_Mieum                = 0xeb1
	GDK_KEY_Hangul_Pieub                = 0xeb2
	GDK_KEY_Hangul_SsangPieub           = 0xeb3
	GDK_KEY_Hangul_PieubSios            = 0xeb4
	GDK_KEY_Hangul_Sios                 = 0xeb5
	GDK_KEY_Hangul_SsangSios            = 0xeb6
	GDK_KEY_Hangul_Ieung                = 0xeb7
	GDK_KEY_Hangul_Jieuj                = 0xeb8
	GDK_KEY_Hangul_SsangJieuj           = 0xeb9
	GDK_KEY_Hangul_Cieuc                = 0xeba
	GDK_KEY_Hangul_Khieuq               = 0xebb
	GDK_KEY_Hangul_Tieut                = 0xebc
	GDK_KEY_Hangul_Phieuf               = 0xebd
	GDK_KEY_Hangul_Hieuh                = 0xebe
	GDK_KEY_Hangul_A                    = 0xebf
	GDK_KEY_Hangul_AE                   = 0xec0
	GDK_KEY_Hangul_YA                   = 0xec1
	GDK_KEY_Hangul_YAE                  = 0xec2
	GDK_KEY_Hangul_EO                   = 0xec3
	GDK_KEY_Hangul_E                    = 0xec4
	GDK_KEY_Hangul_YEO                  = 0xec5
	GDK_KEY_Hangul_YE                   = 0xec6
	GDK_KEY_Hangul_O                    = 0xec7
	GDK_KEY_Hangul_WA                   = 0xec8
	GDK_KEY_Hangul_WAE                  = 0xec9
	GDK_KEY_Hangul_OE                   = 0xeca
	GDK_KEY_Hangul_YO                   = 0xecb
	GDK_KEY_Hangul_U                    = 0xecc
	GDK_KEY_Hangul_WEO                  = 0xecd
	GDK_KEY_Hangul_WE                   = 0xece
	GDK_KEY_Hangul_WI                   = 0xecf
	GDK_KEY_Hangul_YU                   = 0xed0
	GDK_KEY_Hangul_EU                   = 0xed1
	GDK_KEY_Hangul_YI                   = 0xed2
	GDK_KEY_Hangul_I                    = 0xed3
	GDK_KEY_Hangul_J_Kiyeog             = 0xed4
	GDK_KEY_Hangul_J_SsangKiyeog        = 0xed5
	GDK_KEY_Hangul_J_KiyeogSios         = 0xed6
	GDK_KEY_Hangul_J_Nieun              = 0xed7
	GDK_KEY_Hangul_J_NieunJieuj         = 0xed8
	GDK_KEY_Hangul_J_NieunHieuh         = 0xed9
	GDK_KEY_Hangul_J_Dikeud             = 0xeda
	GDK_KEY_Hangul_J_Rieul              = 0xedb
	GDK_KEY_Hangul_J_RieulKiyeog        = 0xedc
	GDK_KEY_Hangul_J_RieulMieum         = 0xedd
	GDK_KEY_Hangul_J_RieulPieub         = 0xede
	GDK_KEY_Hangul_J_RieulSios          = 0xedf
	GDK_KEY_Hangul_J_RieulTieut         = 0xee0
	GDK_KEY_Hangul_J_RieulPhieuf        = 0xee1
	GDK_KEY_Hangul_J_RieulHieuh         = 0xee2
	GDK_KEY_Hangul_J_Mieum              = 0xee3
	GDK_KEY_Hangul_J_Pieub              = 0xee4
	GDK_KEY_Hangul_J_PieubSios          = 0xee5
	GDK_KEY_Hangul_J_Sios               = 0xee6
	GDK_KEY_Hangul_J_SsangSios          = 0xee7
	GDK_KEY_Hangul_J_Ieung              = 0xee8
	GDK_KEY_Hangul_J_Jieuj              = 0xee9
	GDK_KEY_Hangul_J_Cieuc              = 0xeea
	GDK_KEY_Hangul_J_Khieuq             = 0xeeb
	GDK_KEY_Hangul_J_Tieut              = 0xeec
	GDK_KEY_Hangul_J_Phieuf             = 0xeed
	GDK_KEY_Hangul_J_Hieuh              = 0xeee
	GDK_KEY_Hangul_RieulYeorinHieuh     = 0xeef
	GDK_KEY_Hangul_SunkyeongeumMieum    = 0xef0
	GDK_KEY_Hangul_SunkyeongeumPieub    = 0xef1
	GDK_KEY_Hangul_PanSios              = 0xef2
	GDK_KEY_Hangul_KkogjiDalrinIeung    = 0xef3
	GDK_KEY_Hangul_SunkyeongeumPhieuf   = 0xef4
	GDK_KEY_Hangul_YeorinHieuh          = 0xef5
	GDK_KEY_Hangul_AraeA                = 0xef6
	GDK_KEY_Hangul_AraeAE               = 0xef7
	GDK_KEY_Hangul_J_PanSios            = 0xef8
	GDK_KEY_Hangul_J_KkogjiDalrinIeung  = 0xef9
	GDK_KEY_Hangul_J_YeorinHieuh        = 0xefa
	GDK_KEY_Korean_Won                  = 0xeff
	GDK_KEY_Armenian_ligature_ew        = 0x1000587
	GDK_KEY_Armenian_full_stop          = 0x1000589
	GDK_KEY_Armenian_verjaket           = 0x1000589
	GDK_KEY_Armenian_separation_mark    = 0x100055d
	GDK_KEY_Armenian_but                = 0x100055d
	GDK_KEY_Armenian_hyphen             = 0x100058a
	GDK_KEY_Armenian_yentamna           = 0x100058a
	GDK_KEY_Armenian_exclam             = 0x100055c
	GDK_KEY_Armenian_amanak             = 0x100055c
	GDK_KEY_Armenian_accent             = 0x100055b
	GDK_KEY_Armenian_shesht             = 0x100055b
	GDK_KEY_Armenian_question           = 0x100055e
	GDK_KEY_Armenian_paruyk             = 0x100055e
	GDK_KEY_Armenian_AYB                = 0x1000531
	GDK_KEY_Armenian_ayb                = 0x1000561
	GDK_KEY_Armenian_BEN                = 0x1000532
	GDK_KEY_Armenian_ben                = 0x1000562
	GDK_KEY_Armenian_GIM                = 0x1000533
	GDK_KEY_Armenian_gim                = 0x1000563
	GDK_KEY_Armenian_DA                 = 0x1000534
	GDK_KEY_Armenian_da                 = 0x1000564
	GDK_KEY_Armenian_YECH               = 0x1000535
	GDK_KEY_Armenian_yech               = 0x1000565
	GDK_KEY_Armenian_ZA                 = 0x1000536
	GDK_KEY_Armenian_za                 = 0x1000566
	GDK_KEY_Armenian_E                  = 0x1000537
	GDK_KEY_Armenian_e                  = 0x1000567
	GDK_KEY_Armenian_AT                 = 0x1000538
	GDK_KEY_Armenian_at                 = 0x1000568
	GDK_KEY_Armenian_TO                 = 0x1000539
	GDK_KEY_Armenian_to                 = 0x1000569
	GDK_KEY_Armenian_ZHE                = 0x100053a
	GDK_KEY_Armenian_zhe                = 0x100056a
	GDK_KEY_Armenian_INI                = 0x100053b
	GDK_KEY_Armenian_ini                = 0x100056b
	GDK_KEY_Armenian_LYUN               = 0x100053c
	GDK_KEY_Armenian_lyun               = 0x100056c
	GDK_KEY_Armenian_KHE                = 0x100053d
	GDK_KEY_Armenian_khe                = 0x100056d
	GDK_KEY_Armenian_TSA                = 0x100053e
	GDK_KEY_Armenian_tsa                = 0x100056e
	GDK_KEY_Armenian_KEN                = 0x100053f
	GDK_KEY_Armenian_ken                = 0x100056f
	GDK_KEY_Armenian_HO                 = 0x1000540
	GDK_KEY_Armenian_ho                 = 0x1000570
	GDK_KEY_Armenian_DZA                = 0x1000541
	GDK_KEY_Armenian_dza                = 0x1000571
	GDK_KEY_Armenian_GHAT               = 0x1000542
	GDK_KEY_Armenian_ghat               = 0x1000572
	GDK_KEY_Armenian_TCHE               = 0x1000543
	GDK_KEY_Armenian_tche               = 0x1000573
	GDK_KEY_Armenian_MEN                = 0x1000544
	GDK_KEY_Armenian_men                = 0x1000574
	GDK_KEY_Armenian_HI                 = 0x1000545
	GDK_KEY_Armenian_hi                 = 0x1000575
	GDK_KEY_Armenian_NU                 = 0x1000546
	GDK_KEY_Armenian_nu                 = 0x1000576
	GDK_KEY_Armenian_SHA                = 0x1000547
	GDK_KEY_Armenian_sha                = 0x1000577
	GDK_KEY_Armenian_VO                 = 0x1000548
	GDK_KEY_Armenian_vo                 = 0x1000578
	GDK_KEY_Armenian_CHA                = 0x1000549
	GDK_KEY_Armenian_cha                = 0x1000579
	GDK_KEY_Armenian_PE                 = 0x100054a
	GDK_KEY_Armenian_pe                 = 0x100057a
	GDK_KEY_Armenian_JE                 = 0x100054b
	GDK_KEY_Armenian_je                 = 0x100057b
	GDK_KEY_Armenian_RA                 = 0x100054c
	GDK_KEY_Armenian_ra                 = 0x100057c
	GDK_KEY_Armenian_SE                 = 0x100054d
	GDK_KEY_Armenian_se                 = 0x100057d
	GDK_KEY_Armenian_VEV                = 0x100054e
	GDK_KEY_Armenian_vev                = 0x100057e
	GDK_KEY_Armenian_TYUN               = 0x100054f
	GDK_KEY_Armenian_tyun               = 0x100057f
	GDK_KEY_Armenian_RE                 = 0x1000550
	GDK_KEY_Armenian_re                 = 0x1000580
	GDK_KEY_Armenian_TSO                = 0x1000551
	GDK_KEY_Armenian_tso                = 0x1000581
	GDK_KEY_Armenian_VYUN               = 0x1000552
	GDK_KEY_Armenian_vyun               = 0x1000582
	GDK_KEY_Armenian_PYUR               = 0x1000553
	GDK_KEY_Armenian_pyur               = 0x1000583
	GDK_KEY_Armenian_KE                 = 0x1000554
	GDK_KEY_Armenian_ke                 = 0x1000584
	GDK_KEY_Armenian_O                  = 0x1000555
	GDK_KEY_Armenian_o                  = 0x1000585
	GDK_KEY_Armenian_FE                 = 0x1000556
	GDK_KEY_Armenian_fe                 = 0x1000586
	GDK_KEY_Armenian_apostrophe         = 0x100055a
	GDK_KEY_Georgian_an                 = 0x10010d0
	GDK_KEY_Georgian_ban                = 0x10010d1
	GDK_KEY_Georgian_gan                = 0x10010d2
	GDK_KEY_Georgian_don                = 0x10010d3
	GDK_KEY_Georgian_en                 = 0x10010d4
	GDK_KEY_Georgian_vin                = 0x10010d5
	GDK_KEY_Georgian_zen                = 0x10010d6
	GDK_KEY_Georgian_tan                = 0x10010d7
	GDK_KEY_Georgian_in                 = 0x10010d8
	GDK_KEY_Georgian_kan                = 0x10010d9
	GDK_KEY_Georgian_las                = 0x10010da
	GDK_KEY_Georgian_man                = 0x10010db
	GDK_KEY_Georgian_nar                = 0x10010dc
	GDK_KEY_Georgian_on                 = 0x10010dd
	GDK_KEY_Georgian_par                = 0x10010de
	GDK_KEY_Georgian_zhar               = 0x10010df
	GDK_KEY_Georgian_rae                = 0x10010e0
	GDK_KEY_Georgian_san                = 0x10010e1
	GDK_KEY_Georgian_tar                = 0x10010e2
	GDK_KEY_Georgian_un                 = 0x10010e3
	GDK_KEY_Georgian_phar               = 0x10010e4
	GDK_KEY_Georgian_khar               = 0x10010e5
	GDK_KEY_Georgian_ghan               = 0x10010e6
	GDK_KEY_Georgian_qar                = 0x10010e7
	GDK_KEY_Georgian_shin               = 0x10010e8
	GDK_KEY_Georgian_chin               = 0x10010e9
	GDK_KEY_Georgian_can                = 0x10010ea
	GDK_KEY_Georgian_jil                = 0x10010eb
	GDK_KEY_Georgian_cil                = 0x10010ec
	GDK_KEY_Georgian_char               = 0x10010ed
	GDK_KEY_Georgian_xan                = 0x10010ee
	GDK_KEY_Georgian_jhan               = 0x10010ef
	GDK_KEY_Georgian_hae                = 0x10010f0
	GDK_KEY_Georgian_he                 = 0x10010f1
	GDK_KEY_Georgian_hie                = 0x10010f2
	GDK_KEY_Georgian_we                 = 0x10010f3
	GDK_KEY_Georgian_har                = 0x10010f4
	GDK_KEY_Georgian_hoe                = 0x10010f5
	GDK_KEY_Georgian_fi                 = 0x10010f6
	GDK_KEY_Xabovedot                   = 0x1001e8a
	GDK_KEY_Ibreve                      = 0x100012c
	GDK_KEY_Zstroke                     = 0x10001b5
	GDK_KEY_Gcaron                      = 0x10001e6
	GDK_KEY_Ocaron                      = 0x10001d1
	GDK_KEY_Obarred                     = 0x100019f
	GDK_KEY_xabovedot                   = 0x1001e8b
	GDK_KEY_ibreve                      = 0x100012d
	GDK_KEY_zstroke                     = 0x10001b6
	GDK_KEY_gcaron                      = 0x10001e7
	GDK_KEY_ocaron                      = 0x10001d2
	GDK_KEY_obarred                     = 0x1000275
	GDK_KEY_SCHWA                       = 0x100018f
	GDK_KEY_schwa                       = 0x1000259
	GDK_KEY_Lbelowdot                   = 0x1001e36
	GDK_KEY_lbelowdot                   = 0x1001e37
	GDK_KEY_Abelowdot                   = 0x1001ea0
	GDK_KEY_abelowdot                   = 0x1001ea1
	GDK_KEY_Ahook                       = 0x1001ea2
	GDK_KEY_ahook                       = 0x1001ea3
	GDK_KEY_Acircumflexacute            = 0x1001ea4
	GDK_KEY_acircumflexacute            = 0x1001ea5
	GDK_KEY_Acircumflexgrave            = 0x1001ea6
	GDK_KEY_acircumflexgrave            = 0x1001ea7
	GDK_KEY_Acircumflexhook             = 0x1001ea8
	GDK_KEY_acircumflexhook             = 0x1001ea9
	GDK_KEY_Acircumflextilde            = 0x1001eaa
	GDK_KEY_acircumflextilde            = 0x1001eab
	GDK_KEY_Acircumflexbelowdot         = 0x1001eac
	GDK_KEY_acircumflexbelowdot         = 0x1001ead
	GDK_KEY_Abreveacute                 = 0x1001eae
	GDK_KEY_abreveacute                 = 0x1001eaf
	GDK_KEY_Abrevegrave                 = 0x1001eb0
	GDK_KEY_abrevegrave                 = 0x1001eb1
	GDK_KEY_Abrevehook                  = 0x1001eb2
	GDK_KEY_abrevehook                  = 0x1001eb3
	GDK_KEY_Abrevetilde                 = 0x1001eb4
	GDK_KEY_abrevetilde                 = 0x1001eb5
	GDK_KEY_Abrevebelowdot              = 0x1001eb6
	GDK_KEY_abrevebelowdot              = 0x1001eb7
	GDK_KEY_Ebelowdot                   = 0x1001eb8
	GDK_KEY_ebelowdot                   = 0x1001eb9
	GDK_KEY_Ehook                       = 0x1001eba
	GDK_KEY_ehook                       = 0x1001ebb
	GDK_KEY_Etilde                      = 0x1001ebc
	GDK_KEY_etilde                      = 0x1001ebd
	GDK_KEY_Ecircumflexacute            = 0x1001ebe
	GDK_KEY_ecircumflexacute            = 0x1001ebf
	GDK_KEY_Ecircumflexgrave            = 0x1001ec0
	GDK_KEY_ecircumflexgrave            = 0x1001ec1
	GDK_KEY_Ecircumflexhook             = 0x1001ec2
	GDK_KEY_ecircumflexhook             = 0x1001ec3
	GDK_KEY_Ecircumflextilde            = 0x1001ec4
	GDK_KEY_ecircumflextilde            = 0x1001ec5
	GDK_KEY_Ecircumflexbelowdot         = 0x1001ec6
	GDK_KEY_ecircumflexbelowdot         = 0x1001ec7
	GDK_KEY_Ihook                       = 0x1001ec8
	GDK_KEY_ihook                       = 0x1001ec9
	GDK_KEY_Ibelowdot                   = 0x1001eca
	GDK_KEY_ibelowdot                   = 0x1001ecb
	GDK_KEY_Obelowdot                   = 0x1001ecc
	GDK_KEY_obelowdot                   = 0x1001ecd
	GDK_KEY_Ohook                       = 0x1001ece
	GDK_KEY_ohook                       = 0x1001ecf
	GDK_KEY_Ocircumflexacute            = 0x1001ed0
	GDK_KEY_ocircumflexacute            = 0x1001ed1
	GDK_KEY_Ocircumflexgrave            = 0x1001ed2
	GDK_KEY_ocircumflexgrave            = 0x1001ed3
	GDK_KEY_Ocircumflexhook             = 0x1001ed4
	GDK_KEY_ocircumflexhook             = 0x1001ed5
	GDK_KEY_Ocircumflextilde            = 0x1001ed6
	GDK_KEY_ocircumflextilde            = 0x1001ed7
	GDK_KEY_Ocircumflexbelowdot         = 0x1001ed8
	GDK_KEY_ocircumflexbelowdot         = 0x1001ed9
	GDK_KEY_Ohornacute                  = 0x1001eda
	GDK_KEY_ohornacute                  = 0x1001edb
	GDK_KEY_Ohorngrave                  = 0x1001edc
	GDK_KEY_ohorngrave                  = 0x1001edd
	GDK_KEY_Ohornhook                   = 0x1001ede
	GDK_KEY_ohornhook                   = 0x1001edf
	GDK_KEY_Ohorntilde                  = 0x1001ee0
	GDK_KEY_ohorntilde                  = 0x1001ee1
	GDK_KEY_Ohornbelowdot               = 0x1001ee2
	GDK_KEY_ohornbelowdot               = 0x1001ee3
	GDK_KEY_Ubelowdot                   = 0x1001ee4
	GDK_KEY_ubelowdot                   = 0x1001ee5
	GDK_KEY_Uhook                       = 0x1001ee6
	GDK_KEY_uhook                       = 0x1001ee7
	GDK_KEY_Uhornacute                  = 0x1001ee8
	GDK_KEY_uhornacute                  = 0x1001ee9
	GDK_KEY_Uhorngrave                  = 0x1001eea
	GDK_KEY_uhorngrave                  = 0x1001eeb
	GDK_KEY_Uhornhook                   = 0x1001eec
	GDK_KEY_uhornhook                   = 0x1001eed
	GDK_KEY_Uhorntilde                  = 0x1001eee
	GDK_KEY_uhorntilde                  = 0x1001eef
	GDK_KEY_Uhornbelowdot               = 0x1001ef0
	GDK_KEY_uhornbelowdot               = 0x1001ef1
	GDK_KEY_Ybelowdot                   = 0x1001ef4
	GDK_KEY_ybelowdot                   = 0x1001ef5
	GDK_KEY_Yhook                       = 0x1001ef6
	GDK_KEY_yhook                       = 0x1001ef7
	GDK_KEY_Ytilde                      = 0x1001ef8
	GDK_KEY_ytilde                      = 0x1001ef9
	GDK_KEY_Ohorn                       = 0x10001a0
	GDK_KEY_ohorn                       = 0x10001a1
	GDK_KEY_Uhorn                       = 0x10001af
	GDK_KEY_uhorn                       = 0x10001b0
	GDK_KEY_EcuSign                     = 0x10020a0
	GDK_KEY_ColonSign                   = 0x10020a1
	GDK_KEY_CruzeiroSign                = 0x10020a2
	GDK_KEY_FFrancSign                  = 0x10020a3
	GDK_KEY_LiraSign                    = 0x10020a4
	GDK_KEY_MillSign                    = 0x10020a5
	GDK_KEY_NairaSign                   = 0x10020a6
	GDK_KEY_PesetaSign                  = 0x10020a7
	GDK_KEY_RupeeSign                   = 0x10020a8
	GDK_KEY_WonSign                     = 0x10020a9
	GDK_KEY_NewSheqelSign               = 0x10020aa
	GDK_KEY_DongSign                    = 0x10020ab
	GDK_KEY_EuroSign                    = 0x20ac
	GDK_KEY_zerosuperior                = 0x1002070
	GDK_KEY_foursuperior                = 0x1002074
	GDK_KEY_fivesuperior                = 0x1002075
	GDK_KEY_sixsuperior                 = 0x1002076
	GDK_KEY_sevensuperior               = 0x1002077
	GDK_KEY_eightsuperior               = 0x1002078
	GDK_KEY_ninesuperior                = 0x1002079
	GDK_KEY_zerosubscript               = 0x1002080
	GDK_KEY_onesubscript                = 0x1002081
	GDK_KEY_twosubscript                = 0x1002082
	GDK_KEY_threesubscript              = 0x1002083
	GDK_KEY_foursubscript               = 0x1002084
	GDK_KEY_fivesubscript               = 0x1002085
	GDK_KEY_sixsubscript                = 0x1002086
	GDK_KEY_sevensubscript              = 0x1002087
	GDK_KEY_eightsubscript              = 0x1002088
	GDK_KEY_ninesubscript               = 0x1002089
	GDK_KEY_partdifferential            = 0x1002202
	GDK_KEY_emptyset                    = 0x1002205
	GDK_KEY_elementof                   = 0x1002208
	GDK_KEY_notelementof                = 0x1002209
	GDK_KEY_containsas                  = 0x100220b
	GDK_KEY_squareroot                  = 0x100221a
	GDK_KEY_cuberoot                    = 0x100221b
	GDK_KEY_fourthroot                  = 0x100221c
	GDK_KEY_dintegral                   = 0x100222c
	GDK_KEY_tintegral                   = 0x100222d
	GDK_KEY_because                     = 0x1002235
	GDK_KEY_approxeq                    = 0x1002248
	GDK_KEY_notapproxeq                 = 0x1002247
	GDK_KEY_notidentical                = 0x1002262
	GDK_KEY_stricteq                    = 0x1002263
	GDK_KEY_braille_dot_1               = 0xfff1
	GDK_KEY_braille_dot_2               = 0xfff2
	GDK_KEY_braille_dot_3               = 0xfff3
	GDK_KEY_braille_dot_4               = 0xfff4
	GDK_KEY_braille_dot_5               = 0xfff5
	GDK_KEY_braille_dot_6               = 0xfff6
	GDK_KEY_braille_dot_7               = 0xfff7
	GDK_KEY_braille_dot_8               = 0xfff8
	GDK_KEY_braille_dot_9               = 0xfff9
	GDK_KEY_braille_dot_10              = 0xfffa
	GDK_KEY_braille_blank               = 0x1002800
	GDK_KEY_braille_dots_1              = 0x1002801
	GDK_KEY_braille_dots_2              = 0x1002802
	GDK_KEY_braille_dots_12             = 0x1002803
	GDK_KEY_braille_dots_3              = 0x1002804
	GDK_KEY_braille_dots_13             = 0x1002805
	GDK_KEY_braille_dots_23             = 0x1002806
	GDK_KEY_braille_dots_123            = 0x1002807
	GDK_KEY_braille_dots_4              = 0x1002808
	GDK_KEY_braille_dots_14             = 0x1002809
	GDK_KEY_braille_dots_24             = 0x100280a
	GDK_KEY_braille_dots_124            = 0x100280b
	GDK_KEY_braille_dots_34             = 0x100280c
	GDK_KEY_braille_dots_134            = 0x100280d
	GDK_KEY_braille_dots_234            = 0x100280e
	GDK_KEY_braille_dots_1234           = 0x100280f
	GDK_KEY_braille_dots_5              = 0x1002810
	GDK_KEY_braille_dots_15             = 0x1002811
	GDK_KEY_braille_dots_25             = 0x1002812
	GDK_KEY_braille_dots_125            = 0x1002813
	GDK_KEY_braille_dots_35             = 0x1002814
	GDK_KEY_braille_dots_135            = 0x1002815
	GDK_KEY_braille_dots_235            = 0x1002816
	GDK_KEY_braille_dots_1235           = 0x1002817
	GDK_KEY_braille_dots_45             = 0x1002818
	GDK_KEY_braille_dots_145            = 0x1002819
	GDK_KEY_braille_dots_245            = 0x100281a
	GDK_KEY_braille_dots_1245           = 0x100281b
	GDK_KEY_braille_dots_345            = 0x100281c
	GDK_KEY_braille_dots_1345           = 0x100281d
	GDK_KEY_braille_dots_2345           = 0x100281e
	GDK_KEY_braille_dots_12345          = 0x100281f
	GDK_KEY_braille_dots_6              = 0x1002820
	GDK_KEY_braille_dots_16             = 0x1002821
	GDK_KEY_braille_dots_26             = 0x1002822
	GDK_KEY_braille_dots_126            = 0x1002823
	GDK_KEY_braille_dots_36             = 0x1002824
	GDK_KEY_braille_dots_136            = 0x1002825
	GDK_KEY_braille_dots_236            = 0x1002826
	GDK_KEY_braille_dots_1236           = 0x1002827
	GDK_KEY_braille_dots_46             = 0x1002828
	GDK_KEY_braille_dots_146            = 0x1002829
	GDK_KEY_braille_dots_246            = 0x100282a
	GDK_KEY_braille_dots_1246           = 0x100282b
	GDK_KEY_braille_dots_346            = 0x100282c
	GDK_KEY_braille_dots_1346           = 0x100282d
	GDK_KEY_braille_dots_2346           = 0x100282e
	GDK_KEY_braille_dots_12346          = 0x100282f
	GDK_KEY_braille_dots_56             = 0x1002830
	GDK_KEY_braille_dots_156            = 0x1002831
	GDK_KEY_braille_dots_256            = 0x1002832
	GDK_KEY_braille_dots_1256           = 0x1002833
	GDK_KEY_braille_dots_356            = 0x1002834
	GDK_KEY_braille_dots_1356           = 0x1002835
	GDK_KEY_braille_dots_2356           = 0x1002836
	GDK_KEY_braille_dots_12356          = 0x1002837
	GDK_KEY_braille_dots_456            = 0x1002838
	GDK_KEY_braille_dots_1456           = 0x1002839
	GDK_KEY_braille_dots_2456           = 0x100283a
	GDK_KEY_braille_dots_12456          = 0x100283b
	GDK_KEY_braille_dots_3456           = 0x100283c
	GDK_KEY_braille_dots_13456          = 0x100283d
	GDK_KEY_braille_dots_23456          = 0x100283e
	GDK_KEY_braille_dots_123456         = 0x100283f
	GDK_KEY_braille_dots_7              = 0x1002840
	GDK_KEY_braille_dots_17             = 0x1002841
	GDK_KEY_braille_dots_27             = 0x1002842
	GDK_KEY_braille_dots_127            = 0x1002843
	GDK_KEY_braille_dots_37             = 0x1002844
	GDK_KEY_braille_dots_137            = 0x1002845
	GDK_KEY_braille_dots_237            = 0x1002846
	GDK_KEY_braille_dots_1237           = 0x1002847
	GDK_KEY_braille_dots_47             = 0x1002848
	GDK_KEY_braille_dots_147            = 0x1002849
	GDK_KEY_braille_dots_247            = 0x100284a
	GDK_KEY_braille_dots_1247           = 0x100284b
	GDK_KEY_braille_dots_347            = 0x100284c
	GDK_KEY_braille_dots_1347           = 0x100284d
	GDK_KEY_braille_dots_2347           = 0x100284e
	GDK_KEY_braille_dots_12347          = 0x100284f
	GDK_KEY_braille_dots_57             = 0x1002850
	GDK_KEY_braille_dots_157            = 0x1002851
	GDK_KEY_braille_dots_257            = 0x1002852
	GDK_KEY_braille_dots_1257           = 0x1002853
	GDK_KEY_braille_dots_357            = 0x1002854
	GDK_KEY_braille_dots_1357           = 0x1002855
	GDK_KEY_braille_dots_2357           = 0x1002856
	GDK_KEY_braille_dots_12357          = 0x1002857
	GDK_KEY_braille_dots_457            = 0x1002858
	GDK_KEY_braille_dots_1457           = 0x1002859
	GDK_KEY_braille_dots_2457           = 0x100285a
	GDK_KEY_braille_dots_12457          = 0x100285b
	GDK_KEY_braille_dots_3457           = 0x100285c
	GDK_KEY_braille_dots_13457          = 0x100285d
	GDK_KEY_braille_dots_23457          = 0x100285e
	GDK_KEY_braille_dots_123457         = 0x100285f
	GDK_KEY_braille_dots_67             = 0x1002860
	GDK_KEY_braille_dots_167            = 0x1002861
	GDK_KEY_braille_dots_267            = 0x1002862
	GDK_KEY_braille_dots_1267           = 0x1002863
	GDK_KEY_braille_dots_367            = 0x1002864
	GDK_KEY_braille_dots_1367           = 0x1002865
	GDK_KEY_braille_dots_2367           = 0x1002866
	GDK_KEY_braille_dots_12367          = 0x1002867
	GDK_KEY_braille_dots_467            = 0x1002868
	GDK_KEY_braille_dots_1467           = 0x1002869
	GDK_KEY_braille_dots_2467           = 0x100286a
	GDK_KEY_braille_dots_12467          = 0x100286b
	GDK_KEY_braille_dots_3467           = 0x100286c
	GDK_KEY_braille_dots_13467          = 0x100286d
	GDK_KEY_braille_dots_23467          = 0x100286e
	GDK_KEY_braille_dots_123467         = 0x100286f
	GDK_KEY_braille_dots_567            = 0x1002870
	GDK_KEY_braille_dots_1567           = 0x1002871
	GDK_KEY_braille_dots_2567           = 0x1002872
	GDK_KEY_braille_dots_12567          = 0x1002873
	GDK_KEY_braille_dots_3567           = 0x1002874
	GDK_KEY_braille_dots_13567          = 0x1002875
	GDK_KEY_braille_dots_23567          = 0x1002876
	GDK_KEY_braille_dots_123567         = 0x1002877
	GDK_KEY_braille_dots_4567           = 0x1002878
	GDK_KEY_braille_dots_14567          = 0x1002879
	GDK_KEY_braille_dots_24567          = 0x100287a
	GDK_KEY_braille_dots_124567         = 0x100287b
	GDK_KEY_braille_dots_34567          = 0x100287c
	GDK_KEY_braille_dots_134567         = 0x100287d
	GDK_KEY_braille_dots_234567         = 0x100287e
	GDK_KEY_braille_dots_1234567        = 0x100287f
	GDK_KEY_braille_dots_8              = 0x1002880
	GDK_KEY_braille_dots_18             = 0x1002881
	GDK_KEY_braille_dots_28             = 0x1002882
	GDK_KEY_braille_dots_128            = 0x1002883
	GDK_KEY_braille_dots_38             = 0x1002884
	GDK_KEY_braille_dots_138            = 0x1002885
	GDK_KEY_braille_dots_238            = 0x1002886
	GDK_KEY_braille_dots_1238           = 0x1002887
	GDK_KEY_braille_dots_48             = 0x1002888
	GDK_KEY_braille_dots_148            = 0x1002889
	GDK_KEY_braille_dots_248            = 0x100288a
	GDK_KEY_braille_dots_1248           = 0x100288b
	GDK_KEY_braille_dots_348            = 0x100288c
	GDK_KEY_braille_dots_1348           = 0x100288d
	GDK_KEY_braille_dots_2348           = 0x100288e
	GDK_KEY_braille_dots_12348          = 0x100288f
	GDK_KEY_braille_dots_58             = 0x1002890
	GDK_KEY_braille_dots_158            = 0x1002891
	GDK_KEY_braille_dots_258            = 0x1002892
	GDK_KEY_braille_dots_1258           = 0x1002893
	GDK_KEY_braille_dots_358            = 0x1002894
	GDK_KEY_braille_dots_1358           = 0x1002895
	GDK_KEY_braille_dots_2358           = 0x1002896
	GDK_KEY_braille_dots_12358          = 0x1002897
	GDK_KEY_braille_dots_458            = 0x1002898
	GDK_KEY_braille_dots_1458           = 0x1002899
	GDK_KEY_braille_dots_2458           = 0x100289a
	GDK_KEY_braille_dots_12458          = 0x100289b
	GDK_KEY_braille_dots_3458           = 0x100289c
	GDK_KEY_braille_dots_13458          = 0x100289d
	GDK_KEY_braille_dots_23458          = 0x100289e
	GDK_KEY_braille_dots_123458         = 0x100289f
	GDK_KEY_braille_dots_68             = 0x10028a0
	GDK_KEY_braille_dots_168            = 0x10028a1
	GDK_KEY_braille_dots_268            = 0x10028a2
	GDK_KEY_braille_dots_1268           = 0x10028a3
	GDK_KEY_braille_dots_368            = 0x10028a4
	GDK_KEY_braille_dots_1368           = 0x10028a5
	GDK_KEY_braille_dots_2368           = 0x10028a6
	GDK_KEY_braille_dots_12368          = 0x10028a7
	GDK_KEY_braille_dots_468            = 0x10028a8
	GDK_KEY_braille_dots_1468           = 0x10028a9
	GDK_KEY_braille_dots_2468           = 0x10028aa
	GDK_KEY_braille_dots_12468          = 0x10028ab
	GDK_KEY_braille_dots_3468           = 0x10028ac
	GDK_KEY_braille_dots_13468          = 0x10028ad
	GDK_KEY_braille_dots_23468          = 0x10028ae
	GDK_KEY_braille_dots_123468         = 0x10028af
	GDK_KEY_braille_dots_568            = 0x10028b0
	GDK_KEY_braille_dots_1568           = 0x10028b1
	GDK_KEY_braille_dots_2568           = 0x10028b2
	GDK_KEY_braille_dots_12568          = 0x10028b3
	GDK_KEY_braille_dots_3568           = 0x10028b4
	GDK_KEY_braille_dots_13568          = 0x10028b5
	GDK_KEY_braille_dots_23568          = 0x10028b6
	GDK_KEY_braille_dots_123568         = 0x10028b7
	GDK_KEY_braille_dots_4568           = 0x10028b8
	GDK_KEY_braille_dots_14568          = 0x10028b9
	GDK_KEY_braille_dots_24568          = 0x10028ba
	GDK_KEY_braille_dots_124568         = 0x10028bb
	GDK_KEY_braille_dots_34568          = 0x10028bc
	GDK_KEY_braille_dots_134568         = 0x10028bd
	GDK_KEY_braille_dots_234568         = 0x10028be
	GDK_KEY_braille_dots_1234568        = 0x10028bf
	GDK_KEY_braille_dots_78             = 0x10028c0
	GDK_KEY_braille_dots_178            = 0x10028c1
	GDK_KEY_braille_dots_278            = 0x10028c2
	GDK_KEY_braille_dots_1278           = 0x10028c3
	GDK_KEY_braille_dots_378            = 0x10028c4
	GDK_KEY_braille_dots_1378           = 0x10028c5
	GDK_KEY_braille_dots_2378           = 0x10028c6
	GDK_KEY_braille_dots_12378          = 0x10028c7
	GDK_KEY_braille_dots_478            = 0x10028c8
	GDK_KEY_braille_dots_1478           = 0x10028c9
	GDK_KEY_braille_dots_2478           = 0x10028ca
	GDK_KEY_braille_dots_12478          = 0x10028cb
	GDK_KEY_braille_dots_3478           = 0x10028cc
	GDK_KEY_braille_dots_13478          = 0x10028cd
	GDK_KEY_braille_dots_23478          = 0x10028ce
	GDK_KEY_braille_dots_123478         = 0x10028cf
	GDK_KEY_braille_dots_578            = 0x10028d0
	GDK_KEY_braille_dots_1578           = 0x10028d1
	GDK_KEY_braille_dots_2578           = 0x10028d2
	GDK_KEY_braille_dots_12578          = 0x10028d3
	GDK_KEY_braille_dots_3578           = 0x10028d4
	GDK_KEY_braille_dots_13578          = 0x10028d5
	GDK_KEY_braille_dots_23578          = 0x10028d6
	GDK_KEY_braille_dots_123578         = 0x10028d7
	GDK_KEY_braille_dots_4578           = 0x10028d8
	GDK_KEY_braille_dots_14578          = 0x10028d9
	GDK_KEY_braille_dots_24578          = 0x10028da
	GDK_KEY_braille_dots_124578         = 0x10028db
	GDK_KEY_braille_dots_34578          = 0x10028dc
	GDK_KEY_braille_dots_134578         = 0x10028dd
	GDK_KEY_braille_dots_234578         = 0x10028de
	GDK_KEY_braille_dots_1234578        = 0x10028df
	GDK_KEY_braille_dots_678            = 0x10028e0
	GDK_KEY_braille_dots_1678           = 0x10028e1
	GDK_KEY_braille_dots_2678           = 0x10028e2
	GDK_KEY_braille_dots_12678          = 0x10028e3
	GDK_KEY_braille_dots_3678           = 0x10028e4
	GDK_KEY_braille_dots_13678          = 0x10028e5
	GDK_KEY_braille_dots_23678          = 0x10028e6
	GDK_KEY_braille_dots_123678         = 0x10028e7
	GDK_KEY_braille_dots_4678           = 0x10028e8
	GDK_KEY_braille_dots_14678          = 0x10028e9
	GDK_KEY_braille_dots_24678          = 0x10028ea
	GDK_KEY_braille_dots_124678         = 0x10028eb
	GDK_KEY_braille_dots_34678          = 0x10028ec
	GDK_KEY_braille_dots_134678         = 0x10028ed
	GDK_KEY_braille_dots_234678         = 0x10028ee
	GDK_KEY_braille_dots_1234678        = 0x10028ef
	GDK_KEY_braille_dots_5678           = 0x10028f0
	GDK_KEY_braille_dots_15678          = 0x10028f1
	GDK_KEY_braille_dots_25678          = 0x10028f2
	GDK_KEY_braille_dots_125678         = 0x10028f3
	GDK_KEY_braille_dots_35678          = 0x10028f4
	GDK_KEY_braille_dots_135678         = 0x10028f5
	GDK_KEY_braille_dots_235678         = 0x10028f6
	GDK_KEY_braille_dots_1235678        = 0x10028f7
	GDK_KEY_braille_dots_45678          = 0x10028f8
	GDK_KEY_braille_dots_145678         = 0x10028f9
	GDK_KEY_braille_dots_245678         = 0x10028fa
	GDK_KEY_braille_dots_1245678        = 0x10028fb
	GDK_KEY_braille_dots_345678         = 0x10028fc
	GDK_KEY_braille_dots_1345678        = 0x10028fd
	GDK_KEY_braille_dots_2345678        = 0x10028fe
	GDK_KEY_braille_dots_12345678       = 0x10028ff
	GDK_KEY_ModeLock                    = 0x1008ff01
	GDK_KEY_MonBrightnessUp             = 0x1008ff02
	GDK_KEY_MonBrightnessDown           = 0x1008ff03
	GDK_KEY_KbdLightOnOff               = 0x1008ff04
	GDK_KEY_KbdBrightnessUp             = 0x1008ff05
	GDK_KEY_KbdBrightnessDown           = 0x1008ff06
	GDK_KEY_Standby                     = 0x1008ff10
	GDK_KEY_AudioLowerVolume            = 0x1008ff11
	GDK_KEY_AudioMute                   = 0x1008ff12
	GDK_KEY_AudioRaiseVolume            = 0x1008ff13
	GDK_KEY_AudioPlay                   = 0x1008ff14
	GDK_KEY_AudioStop                   = 0x1008ff15
	GDK_KEY_AudioPrev                   = 0x1008ff16
	GDK_KEY_AudioNext                   = 0x1008ff17
	GDK_KEY_HomePage                    = 0x1008ff18
	GDK_KEY_Mail                        = 0x1008ff19
	GDK_KEY_Start                       = 0x1008ff1a
	GDK_KEY_Search                      = 0x1008ff1b
	GDK_KEY_AudioRecord                 = 0x1008ff1c
	GDK_KEY_Calculator                  = 0x1008ff1d
	GDK_KEY_Memo                        = 0x1008ff1e
	GDK_KEY_ToDoList                    = 0x1008ff1f
	GDK_KEY_Calendar                    = 0x1008ff20
	GDK_KEY_PowerDown                   = 0x1008ff21
	GDK_KEY_ContrastAdjust              = 0x1008ff22
	GDK_KEY_RockerUp                    = 0x1008ff23
	GDK_KEY_RockerDown                  = 0x1008ff24
	GDK_KEY_RockerEnter                 = 0x1008ff25
	GDK_KEY_Back                        = 0x1008ff26
	GDK_KEY_Forward                     = 0x1008ff27
	GDK_KEY_Stop                        = 0x1008ff28
	GDK_KEY_Refresh                     = 0x1008ff29
	GDK_KEY_PowerOff                    = 0x1008ff2a
	GDK_KEY_WakeUp                      = 0x1008ff2b
	GDK_KEY_Eject                       = 0x1008ff2c
	GDK_KEY_ScreenSaver                 = 0x1008ff2d
	GDK_KEY_WWW                         = 0x1008ff2e
	GDK_KEY_Sleep                       = 0x1008ff2f
	GDK_KEY_Favorites                   = 0x1008ff30
	GDK_KEY_AudioPause                  = 0x1008ff31
	GDK_KEY_AudioMedia                  = 0x1008ff32
	GDK_KEY_MyComputer                  = 0x1008ff33
	GDK_KEY_VendorHome                  = 0x1008ff34
	GDK_KEY_LightBulb                   = 0x1008ff35
	GDK_KEY_Shop                        = 0x1008ff36
	GDK_KEY_History                     = 0x1008ff37
	GDK_KEY_OpenURL                     = 0x1008ff38
	GDK_KEY_AddFavorite                 = 0x1008ff39
	GDK_KEY_HotLinks                    = 0x1008ff3a
	GDK_KEY_BrightnessAdjust            = 0x1008ff3b
	GDK_KEY_Finance                     = 0x1008ff3c
	GDK_KEY_Community                   = 0x1008ff3d
	GDK_KEY_AudioRewind                 = 0x1008ff3e
	GDK_KEY_BackForward                 = 0x1008ff3f
	GDK_KEY_Launch0                     = 0x1008ff40
	GDK_KEY_Launch1                     = 0x1008ff41
	GDK_KEY_Launch2                     = 0x1008ff42
	GDK_KEY_Launch3                     = 0x1008ff43
	GDK_KEY_Launch4                     = 0x1008ff44
	GDK_KEY_Launch5                     = 0x1008ff45
	GDK_KEY_Launch6                     = 0x1008ff46
	GDK_KEY_Launch7                     = 0x1008ff47
	GDK_KEY_Launch8                     = 0x1008ff48
	GDK_KEY_Launch9                     = 0x1008ff49
	GDK_KEY_LaunchA                     = 0x1008ff4a
	GDK_KEY_LaunchB                     = 0x1008ff4b
	GDK_KEY_LaunchC                     = 0x1008ff4c
	GDK_KEY_LaunchD                     = 0x1008ff4d
	GDK_KEY_LaunchE                     = 0x1008ff4e
	GDK_KEY_LaunchF                     = 0x1008ff4f
	GDK_KEY_ApplicationLeft             = 0x1008ff50
	GDK_KEY_ApplicationRight            = 0x1008ff51
	GDK_KEY_Book                        = 0x1008ff52
	GDK_KEY_CD                          = 0x1008ff53
	GDK_KEY_WindowClear                 = 0x1008ff55
	GDK_KEY_Close                       = 0x1008ff56
	GDK_KEY_Copy                        = 0x1008ff57
	GDK_KEY_Cut                         = 0x1008ff58
	GDK_KEY_Display                     = 0x1008ff59
	GDK_KEY_DOS                         = 0x1008ff5a
	GDK_KEY_Documents                   = 0x1008ff5b
	GDK_KEY_Excel                       = 0x1008ff5c
	GDK_KEY_Explorer                    = 0x1008ff5d
	GDK_KEY_Game                        = 0x1008ff5e
	GDK_KEY_Go                          = 0x1008ff5f
	GDK_KEY_iTouch                      = 0x1008ff60
	GDK_KEY_LogOff                      = 0x1008ff61
	GDK_KEY_Market                      = 0x1008ff62
	GDK_KEY_Meeting                     = 0x1008ff63
	GDK_KEY_MenuKB                      = 0x1008ff65
	GDK_KEY_MenuPB                      = 0x1008ff66
	GDK_KEY_MySites                     = 0x1008ff67
	GDK_KEY_New                         = 0x1008ff68
	GDK_KEY_News                        = 0x1008ff69
	GDK_KEY_OfficeHome                  = 0x1008ff6a
	GDK_KEY_Open                        = 0x1008ff6b
	GDK_KEY_Option                      = 0x1008ff6c
	GDK_KEY_Paste                       = 0x1008ff6d
	GDK_KEY_Phone                       = 0x1008ff6e
	GDK_KEY_Reply                       = 0x1008ff72
	GDK_KEY_Reload                      = 0x1008ff73
	GDK_KEY_RotateWindows               = 0x1008ff74
	GDK_KEY_RotationPB                  = 0x1008ff75
	GDK_KEY_RotationKB                  = 0x1008ff76
	GDK_KEY_Save                        = 0x1008ff77
	GDK_KEY_ScrollUp                    = 0x1008ff78
	GDK_KEY_ScrollDown                  = 0x1008ff79
	GDK_KEY_ScrollClick                 = 0x1008ff7a
	GDK_KEY_Send                        = 0x1008ff7b
	GDK_KEY_Spell                       = 0x1008ff7c
	GDK_KEY_SplitScreen                 = 0x1008ff7d
	GDK_KEY_Support                     = 0x1008ff7e
	GDK_KEY_TaskPane                    = 0x1008ff7f
	GDK_KEY_Terminal                    = 0x1008ff80
	GDK_KEY_Tools                       = 0x1008ff81
	GDK_KEY_Travel                      = 0x1008ff82
	GDK_KEY_UserPB                      = 0x1008ff84
	GDK_KEY_User1KB                     = 0x1008ff85
	GDK_KEY_User2KB                     = 0x1008ff86
	GDK_KEY_Video                       = 0x1008ff87
	GDK_KEY_WheelButton                 = 0x1008ff88
	GDK_KEY_Word                        = 0x1008ff89
	GDK_KEY_Xfer                        = 0x1008ff8a
	GDK_KEY_ZoomIn                      = 0x1008ff8b
	GDK_KEY_ZoomOut                     = 0x1008ff8c
	GDK_KEY_Away                        = 0x1008ff8d
	GDK_KEY_Messenger                   = 0x1008ff8e
	GDK_KEY_WebCam                      = 0x1008ff8f
	GDK_KEY_MailForward                 = 0x1008ff90
	GDK_KEY_Pictures                    = 0x1008ff91
	GDK_KEY_Music                       = 0x1008ff92
	GDK_KEY_Battery                     = 0x1008ff93
	GDK_KEY_Bluetooth                   = 0x1008ff94
	GDK_KEY_WLAN                        = 0x1008ff95
	GDK_KEY_UWB                         = 0x1008ff96
	GDK_KEY_AudioForward                = 0x1008ff97
	GDK_KEY_AudioRepeat                 = 0x1008ff98
	GDK_KEY_AudioRandomPlay             = 0x1008ff99
	GDK_KEY_Subtitle                    = 0x1008ff9a
	GDK_KEY_AudioCycleTrack             = 0x1008ff9b
	GDK_KEY_CycleAngle                  = 0x1008ff9c
	GDK_KEY_FrameBack                   = 0x1008ff9d
	GDK_KEY_FrameForward                = 0x1008ff9e
	GDK_KEY_Time                        = 0x1008ff9f
	GDK_KEY_SelectButton                = 0x1008ffa0
	GDK_KEY_View                        = 0x1008ffa1
	GDK_KEY_TopMenu                     = 0x1008ffa2
	GDK_KEY_Red                         = 0x1008ffa3
	GDK_KEY_Green                       = 0x1008ffa4
	GDK_KEY_Yellow                      = 0x1008ffa5
	GDK_KEY_Blue                        = 0x1008ffa6
	GDK_KEY_Suspend                     = 0x1008ffa7
	GDK_KEY_Hibernate                   = 0x1008ffa8
	GDK_KEY_TouchpadToggle              = 0x1008ffa9
	GDK_KEY_Switch_VT_1                 = 0x1008fe01
	GDK_KEY_Switch_VT_2                 = 0x1008fe02
	GDK_KEY_Switch_VT_3                 = 0x1008fe03
	GDK_KEY_Switch_VT_4                 = 0x1008fe04
	GDK_KEY_Switch_VT_5                 = 0x1008fe05
	GDK_KEY_Switch_VT_6                 = 0x1008fe06
	GDK_KEY_Switch_VT_7                 = 0x1008fe07
	GDK_KEY_Switch_VT_8                 = 0x1008fe08
	GDK_KEY_Switch_VT_9                 = 0x1008fe09
	GDK_KEY_Switch_VT_10                = 0x1008fe0a
	GDK_KEY_Switch_VT_11                = 0x1008fe0b
	GDK_KEY_Switch_VT_12                = 0x1008fe0c
	GDK_KEY_Ungrab                      = 0x1008fe20
	GDK_KEY_ClearGrab                   = 0x1008fe21
	GDK_KEY_Next_VMode                  = 0x1008fe22
	GDK_KEY_Prev_VMode                  = 0x1008fe23
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

type GdkDragContext struct {
	DragContext *C.GdkDragContext
}

func DragContextFromNative(l unsafe.Pointer) *GdkDragContext {
	return &GdkDragContext{C.to_GtkDragContext(l)}
}

//-----------------------------------------------------------------------
// GdkAtom
//-----------------------------------------------------------------------
type GdkAtom uintptr

//-----------------------------------------------------------------------
// GdkDisplay
//-----------------------------------------------------------------------
type GdkDisplay struct {
	Display unsafe.Pointer
}

func DisplayGetDefault() *GdkDisplay {
	return &GdkDisplay{C._gdk_display_get_default()}
}
