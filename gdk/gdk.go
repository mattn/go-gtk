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
	X int
	Y int
	Width int
	Height int
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
const (
	GDK_X_CURSOR            = 0
	GDK_ARROW               = 2
	GDK_BASED_ARROW_DOWN    = 4
	GDK_BASED_ARROW_UP      = 6
	GDK_BOAT                = 8
	GDK_BOGOSITY            = 10
	GDK_BOTTOM_LEFT_CORNER  = 12
	GDK_BOTTOM_RIGHT_CORNER = 14
	GDK_BOTTOM_SIDE         = 16
	GDK_BOTTOM_TEE          = 18
	GDK_BOX_SPIRAL          = 20
	GDK_CENTER_PTR          = 22
	GDK_CIRCLE              = 24
	GDK_CLOCK               = 26
	GDK_COFFEE_MUG          = 28
	GDK_CROSS               = 30
	GDK_CROSS_REVERSE       = 32
	GDK_CROSSHAIR           = 34
	GDK_DIAMOND_CROSS       = 36
	GDK_DOT                 = 38
	GDK_DOTBOX              = 40
	GDK_DOUBLE_ARROW        = 42
	GDK_DRAFT_LARGE         = 44
	GDK_DRAFT_SMALL         = 46
	GDK_DRAPED_BOX          = 48
	GDK_EXCHANGE            = 50
	GDK_FLEUR               = 52
	GDK_GOBBLER             = 54
	GDK_GUMBY               = 56
	GDK_HAND1               = 58
	GDK_HAND2               = 60
	GDK_HEART               = 62
	GDK_ICON                = 64
	GDK_IRON_CROSS          = 66
	GDK_LEFT_PTR            = 68
	GDK_LEFT_SIDE           = 70
	GDK_LEFT_TEE            = 72
	GDK_LEFTBUTTON          = 74
	GDK_LL_ANGLE            = 76
	GDK_LR_ANGLE            = 78
	GDK_MAN                 = 80
	GDK_MIDDLEBUTTON        = 82
	GDK_MOUSE               = 84
	GDK_PENCIL              = 86
	GDK_PIRATE              = 88
	GDK_PLUS                = 90
	GDK_QUESTION_ARROW      = 92
	GDK_RIGHT_PTR           = 94
	GDK_RIGHT_SIDE          = 96
	GDK_RIGHT_TEE           = 98
	GDK_RIGHTBUTTON         = 100
	GDK_RTL_LOGO            = 102
	GDK_SAILBOAT            = 104
	GDK_SB_DOWN_ARROW       = 106
	GDK_SB_H_DOUBLE_ARROW   = 108
	GDK_SB_LEFT_ARROW       = 110
	GDK_SB_RIGHT_ARROW      = 112
	GDK_SB_UP_ARROW         = 114
	GDK_SB_V_DOUBLE_ARROW   = 116
	GDK_SHUTTLE             = 118
	GDK_SIZING              = 120
	GDK_SPIDER              = 122
	GDK_SPRAYCAN            = 124
	GDK_STAR                = 126
	GDK_TARGET              = 128
	GDK_TCROSS              = 130
	GDK_TOP_LEFT_ARROW      = 132
	GDK_TOP_LEFT_CORNER     = 134
	GDK_TOP_RIGHT_CORNER    = 136
	GDK_TOP_SIDE            = 138
	GDK_TOP_TEE             = 140
	GDK_TREK                = 142
	GDK_UL_ANGLE            = 144
	GDK_UMBRELLA            = 146
	GDK_UR_ANGLE            = 148
	GDK_WATCH               = 150
	GDK_XTERM               = 152
	GDK_LAST_CURSOR
	GDK_BLANK_CURSOR     = -2
	GDK_CURSOR_IS_PIXMAP = -1
)

var (
	first = true
)

type GdkCursor struct {
	Cursor *C.GdkCursor
}

func Cursor(cursor_type int) *GdkCursor {
	return &GdkCursor{
		C.gdk_cursor_new(C.GdkCursorType(cursor_type))}
}

//-----------------------------------------------------------------------
// GdkColor
//-----------------------------------------------------------------------
type GdkColor struct {
	Color C.GdkColor;
}
func Color(name string) *GdkColor {
	var color C.GdkColor
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gdk_color_parse(C.to_gcharptr(ptr), &color)
	return &GdkColor{ color }
}

//-----------------------------------------------------------------------
// GdkGC
//-----------------------------------------------------------------------
type GdkGC struct {
	GC *C.GdkGC;
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
	C.gdk_gc_set_foreground(v.GC, &color.Color);
}
func (v *GdkGC) SetBackground(color *GdkColor) {
	C.gdk_gc_set_background(v.GC, &color.Color);
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
	C.gdk_gc_set_rgb_fg_color(v.GC, &color.Color);
}
func (v *GdkGC) SetRgbBgColor(color *GdkColor) {
	C.gdk_gc_set_rgb_bg_color(v.GC, &color.Color);
}
// GdkScreen * gdk_gc_get_screen (GdkGC *gc);

//-----------------------------------------------------------------------
// GdkDrawable
//-----------------------------------------------------------------------
type GdkDrawable struct {
	Drawable *C.GdkDrawable
}

func (v *GdkDrawable) DrawPoint(gc *GdkGC, x int, y int) {
	C.gdk_draw_point(v.Drawable, gc.GC, C.gint(x), C.gint(y));
}
func (v *GdkDrawable) DrawLine(gc *GdkGC, x1 int, y1 int, x2 int, y2 int) {
	C.gdk_draw_line(v.Drawable, gc.GC, C.gint(x1), C.gint(y1), C.gint(x2), C.gint(y2));
}
func (v *GdkDrawable) DrawRectangle(gc *GdkGC, filled bool, x int, y int, width int, height int) {
	C.gdk_draw_rectangle(v.Drawable, gc.GC, bool2gboolean(filled), C.gint(x), C.gint(y), C.gint(width), C.gint(height));
}
// void gdk_draw_arc (GdkDrawable *drawable, GdkGC *gc, gboolean filled, gint x, gint y, gint width, gint height, gint angle1, gint angle2);
// void gdk_draw_polygon (GdkDrawable *drawable, GdkGC *gc, gboolean filled, const GdkPoint *points, gint n_points);
// void gdk_draw_string (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *string);
// void gdk_draw_text (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *text, gint text_length);
// void gdk_draw_text_wc (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const GdkWChar *text, gint text_length);
func (v *GdkDrawable) DrawDrawable(gc *GdkGC, src *GdkDrawable, xsrc int, ysrc int, xdest int, ydest int, width int, height int) {
	C.gdk_draw_drawable(v.Drawable, gc.GC, src.Drawable, C.gint(xsrc), C.gint(ysrc), C.gint(xdest), C.gint(ydest), C.gint(width), C.gint(height));
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
const (
		GDK_SHIFT_MASK    = 1 << 0;
		GDK_LOCK_MASK	    = 1 << 1;
		GDK_CONTROL_MASK  = 1 << 2;
		GDK_MOD1_MASK	    = 1 << 3;
		GDK_MOD2_MASK	    = 1 << 4;
		GDK_MOD3_MASK	    = 1 << 5;
		GDK_MOD4_MASK	    = 1 << 6;
		GDK_MOD5_MASK	    = 1 << 7;
		GDK_BUTTON1_MASK  = 1 << 8;
		GDK_BUTTON2_MASK  = 1 << 9;
		GDK_BUTTON3_MASK  = 1 << 10;
		GDK_BUTTON4_MASK  = 1 << 11;
		GDK_BUTTON5_MASK  = 1 << 12;

		/* The next few modifiers are used by XKB, so we skip to the end.
		 * Bits 15 - 25 are currently unused. Bit 29 is used internally.
		 */

		GDK_SUPER_MASK    = 1 << 26;
		GDK_HYPER_MASK    = 1 << 27;
		GDK_META_MASK     = 1 << 28;

		GDK_RELEASE_MASK  = 1 << 30;

		GDK_MODIFIER_MASK = 0x5c001fff;
)
const (
	GDK_NOTHING		= -1;
	GDK_DELETE		= 0;
	GDK_DESTROY		= 1;
	GDK_EXPOSE		= 2;
	GDK_MOTION_NOTIFY	= 3;
	GDK_BUTTON_PRESS	= 4;
	GDK_2BUTTON_PRESS	= 5;
	GDK_3BUTTON_PRESS	= 6;
	GDK_BUTTON_RELEASE	= 7;
	GDK_KEY_PRESS		= 8;
	GDK_KEY_RELEASE	= 9;
	GDK_ENTER_NOTIFY	= 10;
	GDK_LEAVE_NOTIFY	= 11;
	GDK_FOCUS_CHANGE	= 12;
	GDK_CONFIGURE		= 13;
	GDK_MAP		= 14;
	GDK_UNMAP		= 15;
	GDK_PROPERTY_NOTIFY	= 16;
	GDK_SELECTION_CLEAR	= 17;
	GDK_SELECTION_REQUEST = 18;
	GDK_SELECTION_NOTIFY	= 19;
	GDK_PROXIMITY_IN	= 20;
	GDK_PROXIMITY_OUT	= 21;
	GDK_DRAG_ENTER        = 22;
	GDK_DRAG_LEAVE        = 23;
	GDK_DRAG_MOTION       = 24;
	GDK_DRAG_STATUS       = 25;
	GDK_DROP_START        = 26;
	GDK_DROP_FINISHED     = 27;
	GDK_CLIENT_EVENT	= 28;
	GDK_VISIBILITY_NOTIFY = 29;
	GDK_NO_EXPOSE		= 30;
	GDK_SCROLL            = 31;
	GDK_WINDOW_STATE      = 32;
	GDK_SETTING           = 33;
	GDK_OWNER_CHANGE      = 34;
	GDK_GRAB_BROKEN       = 35;
	GDK_DAMAGE            = 36;
	GDK_EVENT_LAST        = 37; /* helper variable for decls */
)

const (
	GDK_EXPOSURE_MASK		= 1 << 1;
	GDK_POINTER_MOTION_MASK	= 1 << 2;
	GDK_POINTER_MOTION_HINT_MASK	= 1 << 3;
	GDK_BUTTON_MOTION_MASK	= 1 << 4;
	GDK_BUTTON1_MOTION_MASK	= 1 << 5;
	GDK_BUTTON2_MOTION_MASK	= 1 << 6;
	GDK_BUTTON3_MOTION_MASK	= 1 << 7;
	GDK_BUTTON_PRESS_MASK		= 1 << 8;
	GDK_BUTTON_RELEASE_MASK	= 1 << 9;
	GDK_KEY_PRESS_MASK		= 1 << 10;
	GDK_KEY_RELEASE_MASK		= 1 << 11;
	GDK_ENTER_NOTIFY_MASK		= 1 << 12;
	GDK_LEAVE_NOTIFY_MASK		= 1 << 13;
	GDK_FOCUS_CHANGE_MASK		= 1 << 14;
	GDK_STRUCTURE_MASK		= 1 << 15;
	GDK_PROPERTY_CHANGE_MASK	= 1 << 16;
	GDK_VISIBILITY_NOTIFY_MASK	= 1 << 17;
	GDK_PROXIMITY_IN_MASK		= 1 << 18;
	GDK_PROXIMITY_OUT_MASK	= 1 << 19;
	GDK_SUBSTRUCTURE_MASK		= 1 << 20;
	GDK_SCROLL_MASK               = 1 << 21;
	GDK_ALL_EVENTS_MASK		= 0x3FFFFE;
)

type GdkWindow struct {
	Window *C.GdkWindow
}

func WindowFromUnsafe(window unsafe.Pointer) *GdkWindow {
	return &GdkWindow{
		C.to_GdkWindow(window)}
}
func (v *GdkWindow) GetPointer(x *int, y *int, mask *uint) *GdkWindow {
	var cx, cy C.gint
	var mt C.GdkModifierType;
	ret := &GdkWindow {
		C.gdk_window_get_pointer(v.Window, &cx, &cy, &mt) }
	*x = int(cx);
	*y = int(cy);
	*mask = uint(mt)
	return ret
}
func (v *GdkWindow) GetDrawable() *GdkDrawable {
	return &GdkDrawable {
		(*C.GdkDrawable)(v.Window) }
}
func (v *GdkWindow) Invalidate(rect *Rectangle, invalidate_children bool) {
	if rect != nil {
		var _rect C.GdkRectangle;
		_rect.x = C.gint(rect.X)
		_rect.y = C.gint(rect.Y)
		_rect.width = C.gint(rect.Width)
		_rect.height = C.gint(rect.Height)
		C.gdk_window_invalidate_rect(v.Window, &_rect, bool2gboolean(invalidate_children));
	} else {
		C.gdk_window_invalidate_rect(v.Window, nil, bool2gboolean(invalidate_children));
	}
}

//-----------------------------------------------------------------------
// GdkPixmap
//-----------------------------------------------------------------------
type GdkPixmap struct {
	Pixmap *C.GdkPixmap
}

func Pixmap(drawable *GdkDrawable, width int, height int, depth int) *GdkPixmap {
	return &GdkPixmap{
		C.gdk_pixmap_new(drawable.Drawable, C.gint(width), C.gint(height), C.gint(depth)) }
}
// GdkPixmap* gdk_pixmap_new (GdkDrawable *drawable, gint width, gint height, gint depth);
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
	return &GdkDrawable {
		(*C.GdkDrawable)(v.Pixmap) }
}
