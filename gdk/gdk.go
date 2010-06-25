package gdk

/*
#include <gdk/gdk.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

//static gchar* to_gcharptr(char* s) { return (gchar*)s; }

//static void free_string(char* s) { free(s); }

static GdkWindow* to_GdkWindow(void* w) {
	return (GdkWindow*)w;
}
static void _g_thread_init(GThreadFunctions *vtable) {
#ifdef	G_THREADS_ENABLED
	g_thread_init(vtable);
#endif
}
*/
import "C"
import "unsafe"

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
// GdkWindow
//-----------------------------------------------------------------------
type GdkWindow struct {
	Window *C.GdkWindow
}

func FromWindow(window unsafe.Pointer) *GdkWindow {
	return &GdkWindow{
		C.to_GdkWindow(window)}
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
// Drawing
//-----------------------------------------------------------------------
// void gdk_draw_point (GdkDrawable *drawable, GdkGC *gc, gint x, gint y);
// void gdk_draw_line (GdkDrawable *drawable, GdkGC *gc, gint x1_, gint y1_, gint x2_, gint y2_);
// void gdk_draw_rectangle (GdkDrawable *drawable, GdkGC *gc, gboolean filled, gint x, gint y, gint width, gint height);
// void gdk_draw_arc (GdkDrawable *drawable, GdkGC *gc, gboolean filled, gint x, gint y, gint width, gint height, gint angle1, gint angle2);
// void gdk_draw_polygon (GdkDrawable *drawable, GdkGC *gc, gboolean filled, const GdkPoint *points, gint n_points);
// void gdk_draw_string (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *string);
// void gdk_draw_text (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const gchar *text, gint text_length);
// void gdk_draw_text_wc (GdkDrawable *drawable, GdkFont *font, GdkGC *gc, gint x, gint y, const GdkWChar *text, gint text_length);
// void gdk_draw_drawable (GdkDrawable *drawable, GdkGC *gc, GdkDrawable *src, gint xsrc, gint ysrc, gint xdest, gint ydest, gint width, gint height);
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
