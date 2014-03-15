package glib

// #include "glib.go.h"
// #cgo pkg-config: glib-2.0 gobject-2.0
import "C"
import "unsafe"
import "reflect"

var callback_contexts []*CallbackContext
var sourcefunc_contexts []*SourcefuncContext

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

// converts a C string array to a Go string slice
func toSlice(ar **C.gchar) []string {
	result := make([]string, 0)
	for i := 0; ; i++ {
		str := C.GoString(C.to_charptr(*ar))
		if str == "" {
			break
		}
		result = append(result, str)
		*ar = C.next_string(*ar)
	}
	return result
}

func GPtrToString(p interface{}) string {
	return C.GoString(C.to_charptr_from_gpointer(p.(C.gpointer)))
}

//-----------------------------------------------------------------------
// Application
//-----------------------------------------------------------------------
func GetApplicationName() string {
	return C.GoString(C.to_charptr(C.g_get_application_name()))
}

func SetApplicationName(name string) {
	str := C.CString(name)
	defer C.free_string(str)
	C.g_set_application_name(C.to_gcharptr(str))
}

func GetPrgName() string {
	return C.GoString(C.to_charptr(C.g_get_prgname()))
}

func SetPrgname(name string) {
	str := C.CString(name)
	defer C.free_string(str)
	C.g_set_prgname(C.to_gcharptr(str))
}

//-----------------------------------------------------------------------
// User Information
//-----------------------------------------------------------------------
func GetUserName() string {
	return C.GoString(C.to_charptr(C.g_get_user_name()))
}

func GetRealName() string {
	return C.GoString(C.to_charptr(C.g_get_real_name()))
}

func GetUserCacheDir() string {
	return C.GoString(C.to_charptr(C.g_get_user_cache_dir()))
}

func GetUserDataDir() string {
	return C.GoString(C.to_charptr(C.g_get_user_data_dir()))
}

func GetUserConfigDir() string {
	return C.GoString(C.to_charptr(C.g_get_user_config_dir()))
}

func GetUserRuntimeDir() string {
	return C.GoString(C.to_charptr(C.g_get_user_runtime_dir()))
}

type UserDirectory C.GUserDirectory

const (
	UserDirectoryDesktop     UserDirectory = C.G_USER_DIRECTORY_DESKTOP
	UserDirectoryDocuments   UserDirectory = C.G_USER_DIRECTORY_DOCUMENTS
	UserDirectoryDownload    UserDirectory = C.G_USER_DIRECTORY_DOWNLOAD
	UserDirectoryMusic       UserDirectory = C.G_USER_DIRECTORY_MUSIC
	UserDirectoryPictures    UserDirectory = C.G_USER_DIRECTORY_PICTURES
	UserDirectoryPublicShare UserDirectory = C.G_USER_DIRECTORY_PUBLIC_SHARE
	UserDirectoryTemplates   UserDirectory = C.G_USER_DIRECTORY_TEMPLATES
	UserDirectoryVideos      UserDirectory = C.G_USER_DIRECTORY_VIDEOS
)

func GetUserSpecialDir(directory UserDirectory) string {
	result := C.g_get_user_special_dir(C.GUserDirectory(directory))
	return C.GoString(C.to_charptr(result))
}

func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()
}

//-----------------------------------------------------------------------
// System Information
//-----------------------------------------------------------------------
func GetSystemDataDirs() []string {
	return toSlice(C.g_get_system_data_dirs())
}

func GetSystemConfigDirs() []string {
	return toSlice(C.g_get_system_config_dirs())
}

func GetHostName() string {
	return C.GoString(C.to_charptr(C.g_get_host_name()))
}

func GetHomeDir() string {
	return C.GoString(C.to_charptr(C.g_get_home_dir()))
}

func GetTmpDir() string {
	return C.GoString(C.to_charptr(C.g_get_tmp_dir()))
}

func GetCurrentDir() string {
	return C.GoString(C.to_charptr(C.g_get_current_dir()))
}

//-----------------------------------------------------------------------
// String Convert
//-----------------------------------------------------------------------
func Utf8Validate(str []byte, len int, bar **byte) bool {
	return gobool(C._g_utf8_validate(unsafe.Pointer(&str[0]),
		C.int(len), unsafe.Pointer(bar)))
}

func FilenameFromUri(uri string) (filename string, hostname string, err error) {
	str := C.CString(uri)
	defer C.free_string(str)
	var gerror *C.GError
	var ptr *C.gchar
	filename = C.GoString(C.to_charptr(C.g_filename_from_uri(C.to_gcharptr(str), &ptr, &gerror)))
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	hostname = ""
	if ptr != nil {
		hostname = C.GoString(C.to_charptr(ptr))
	}
	return
}

func FilenameToUri(filename string, hostname string) (uri string, err error) {
	pfilename := C.CString(filename)
	defer C.free_string(pfilename)
	phostname := C.CString(hostname)
	defer C.free_string(phostname)
	var gerror *C.GError
	uri = C.GoString(C.to_charptr(C.g_filename_to_uri(C.to_gcharptr(pfilename), C.to_gcharptr(phostname), &gerror)))
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	return
}

func LocaleToUtf8(opsysstring []byte) (ret string, bytes_read int, bytes_written int, err *Error) {
	var gerror *C.GError
	var cbytes_read, cbytes_written C.int
	str := C._g_locale_to_utf8(unsafe.Pointer(&opsysstring[0]), C.int(len(opsysstring)), &cbytes_read, &cbytes_written, &gerror)
	if unsafe.Pointer(str) != nil {
		defer C.free_string(C.to_charptr(str))
		ret = C.GoString(C.to_charptr(str))
	} else {
		ret = ""
	}
	bytes_read = int(cbytes_read)
	bytes_written = int(cbytes_written)
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	return
}

func LocaleFromUtf8(utf8string string) (ret []byte, bytes_read int, bytes_written int, err *Error) {
	var gerror *C.GError
	var cbytes_read, cbytes_written C.int
	src := C.CString(utf8string)
	defer C.free_string(src)
	str := C._g_locale_from_utf8(src, C.int(C.strlen(src)), &cbytes_read, &cbytes_written, &gerror)
	if unsafe.Pointer(str) != nil {
		defer C.free_string(C.to_charptr(str))
		ret = ([]byte)(C.GoString(C.to_charptr(str)))
	} else {
		ret = ([]byte)("")
	}
	bytes_read = int(cbytes_read)
	bytes_written = int(cbytes_written)
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	return
}

//-----------------------------------------------------------------------
// g_error
//-----------------------------------------------------------------------
type Error struct {
	GError *C.GError
}

func (v *Error) Error() string {
	return v.Message()
}

func (v *Error) Message() string {
	if unsafe.Pointer(v.GError) == nil || unsafe.Pointer(v.GError.message) == nil {
		return ""
	}
	return C.GoString(C.to_charptr(v.GError.message))
}

func ErrorFromNative(err unsafe.Pointer) *Error {
	return &Error{
		C.to_error(err)}
}

//-----------------------------------------------------------------------
// GObject
//-----------------------------------------------------------------------
type ObjectLike interface {
	Ref()
	Unref()
	Connect(s string, f interface{}, data ...interface{})
}
type GObject struct {
	Object unsafe.Pointer
}

func ObjectFromNative(object unsafe.Pointer) *GObject {
	//	return &GObject {
	//		C.to_GObject(object) }
	return &GObject{
		object}
}

func (v *GObject) Ref() {
	C.g_object_ref(C.gpointer(v.Object))
}
func (v *GObject) Unref() {
	C.g_object_unref(C.gpointer(v.Object))
}

func (v *GObject) Set(name string, value interface{}) {
	ptr := C.CString(name)
	defer C.free_string(ptr)

	if _, ok := value.(WrappedObject); ok {
		value = value.(WrappedObject).GetInternalValue()
	}
	if _, ok := value.(GObject); ok {
		value = value.(GObject).Object
	}
	if _, ok := value.(GValue); ok {
		value = value.(GValue).Value
	}

	switch value.(type) {
	case bool:
		bval := gbool(value.(bool))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&bval))
	case byte:
		bval := C.gchar(value.(byte))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&bval))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.gchar(value.(byte))).UnsafeAddr()))
	case int:
		ival := C.int(value.(int))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&ival))
	case uint:
		uval := C.guint(value.(uint))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&uval))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.guint(value.(uint))).UnsafeAddr()))
	case float32:
		f32val := C.gfloat(value.(float32))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&f32val))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.gfloat(value.(float64))).UnsafeAddr()))
	case float64:
		f64val := C.gfloat(value.(float64))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&f64val))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.gfloat(value.(float64))).UnsafeAddr()))
	case string:
		pval := C.CString(value.(string))
		defer C.free_string(pval)
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&pval))
	default:
		if pv, ok := value.(*[0]uint8); ok {
			C._g_object_set_ptr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(value)
			if av.CanAddr() {
				C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(value).UnsafeAddr()))
			} else {
				C._g_object_set_ptr(C.gpointer(v.Object), C.to_gcharptr(ptr), value.(unsafe.Pointer))
			}
		}
	}
}
func (v *GObject) SetProperty(name string, val *GValue) {
	str := C.CString(name)
	defer C.free_string(str)
	C.g_object_set_property(C.to_GObject(v.Object), C.to_gcharptr(str), &val.Value)
}

//-----------------------------------------------------------------------
// GValue
//-----------------------------------------------------------------------
func GValueFromNative(value interface{}) *C.GValue {
	var gv *C.GValue

	if _, ok := value.(WrappedObject); ok {
		value = value.(WrappedObject).GetInternalValue()
	}
	if _, ok := value.(GObject); ok {
		value = value.(GObject).Object
	}
	if _, ok := value.(GValue); ok {
		value = value.(GValue).Value
	}

	switch value.(type) {
	case bool:
		gv = C.init_gvalue_bool(gbool(value.(bool)))
		break
	case byte:
		gv = C.init_gvalue_byte(C.guchar(value.(byte)))
		break
	case int:
		gv = C.init_gvalue_int(C.gint(value.(int)))
		break
	case uint:
		gv = C.init_gvalue_uint(C.guint(value.(uint)))
		break
	case float32:
		gv = C.init_gvalue_double(C.gdouble(value.(float32)))
		break
	case float64:
		gv = C.init_gvalue_double(C.gdouble(value.(float64)))
		break
	case string:
		{
			pval := C.CString(value.(string))
			defer C.free_string(pval)
			gv = C.init_gvalue_string(C.to_gcharptr(pval))
		}
		break
	default:
		//gv = C.init_gvalue_pointer(C.gpointer(unsafe.Pointer(&value)));
		break
	}
	return gv
}
func ValueFromNative(val interface{}) *GValue {
	return &GValue{*GValueFromNative(val)}
}

type GValue struct {
	Value C.GValue
}

const (
	G_TYPE_CHAR    = 3 << 2
	G_TYPE_UCHAR   = 4 << 2
	G_TYPE_BOOL    = 5 << 2
	G_TYPE_INT     = 6 << 2
	G_TYPE_UINT    = 7 << 2
	G_TYPE_LONG    = 8 << 2
	G_TYPE_ULONG   = 9 << 2
	G_TYPE_INT64   = 10 << 2
	G_TYPE_UINT64  = 11 << 2
	G_TYPE_ENUM    = 12 << 2
	G_TYPE_FLAGS   = 13 << 2
	G_TYPE_FLOAT   = 14 << 2
	G_TYPE_DOUBLE  = 15 << 2
	G_TYPE_STRING  = 16 << 2
	G_TYPE_POINTER = 17 << 2
	G_TYPE_BOXED   = 18 << 2
)

func (v *GValue) Init(t int) {
	if t == G_TYPE_INT {
		C.g_value_init_int(&v.Value)
	} else if t == G_TYPE_STRING {
		C.g_value_init_string(&v.Value)
	}
}

func (v *GValue) GetString() string {
	return C.GoString(C.to_charptr(C.g_value_get_string(&v.Value)))
}

func (v *GValue) GetInt() int {
	return int(C.g_value_get_int(&v.Value))
}

func (v *GValue) GetBool() bool {
	return gobool(C.g_value_get_boolean(&v.Value))
}

//-----------------------------------------------------------------------
// WrappedObject
//-----------------------------------------------------------------------
type WrappedObject interface {
	GetInternalValue() unsafe.Pointer
}

//-----------------------------------------------------------------------
// Events
//-----------------------------------------------------------------------
// the go-gtk Callback is simpler than the one in C, because we have
// full closures, so there is never a need to pass additional data via
// a void * pointer.  Where you might have wanted to do that, you can
// instead just use func () { ... using data } to pass the data in.
type CallbackContext struct {
	f      interface{}
	cbi    unsafe.Pointer
	target reflect.Value
	data   reflect.Value
}

func (c *CallbackContext) Target() interface{} {
	return c.target.Interface()
}

func (c *CallbackContext) Data() interface{} {
	return c.data.Interface()
}

func (c *CallbackContext) Args(n int) CallbackArg {
	return CallbackArg(C.callback_info_get_arg((*C.callback_info)(c.cbi), C.int(n)))
}

type CallbackArg uintptr

func (c CallbackArg) ToString() string {
	return C.GoString(C.to_charptr_voidp(unsafe.Pointer(uintptr(c))))
}

//export _go_glib_callback
func _go_glib_callback(cbi *C.callback_info) {
	context := callback_contexts[int(cbi.func_no)]
	rf := reflect.ValueOf(context.f)
	t := rf.Type()
	fargs := make([]reflect.Value, t.NumIn())
	if len(fargs) > 0 {
		fargs[0] = reflect.ValueOf(context)
	}
	ret := rf.Call(fargs)
	if len(ret) > 0 {
		bret, _ := ret[0].Interface().(bool)
		cbi.ret = gbool(bret)
	}
}

// Return the handler call_id to use with HandlerBlock, HandlerUnblock and
// HandlerDisconnect.
//
func (v *GObject) Connect(s string, f interface{}, datas ...interface{}) int {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx := &CallbackContext{f, nil, reflect.ValueOf(v), reflect.ValueOf(data)}
	ptr := C.CString(s)
	defer C.free_string(ptr)
	ctx.cbi = unsafe.Pointer(C._g_signal_connect(unsafe.Pointer(v.Object), C.to_gcharptr(ptr), C.int(len(callback_contexts))))
	callback_contexts = append(callback_contexts, ctx)
	return len(callback_contexts) - 1
}

func (v *GObject) StopEmission(s string) {
	ptr := C.CString(s)
	defer C.free_string(ptr)
	C.g_signal_stop_emission_by_name((C.gpointer)(v.Object), C.to_gcharptr(ptr))
}

func (v *GObject) Emit(s string) {
	ptr := C.CString(s)
	defer C.free_string(ptr)
	C._g_signal_emit_by_name((C.gpointer)(v.Object), C.to_gcharptr(ptr))
}

func (v *GObject) HandlerBlock(call_id int) {
	c_call_id := C._g_signal_callback_id((*C.callback_info)(callback_contexts[call_id].cbi))
	C.g_signal_handler_block((C.gpointer)(v.Object), c_call_id)
}

func (v *GObject) HandlerUnblock(call_id int) {
	c_call_id := C._g_signal_callback_id((*C.callback_info)(callback_contexts[call_id].cbi))
	C.g_signal_handler_unblock((C.gpointer)(v.Object), c_call_id)
}

func (v *GObject) HandlerDisconnect(call_id int) {
	c_call_id := C._g_signal_callback_id((*C.callback_info)(callback_contexts[call_id].cbi))
	C.g_signal_handler_disconnect((C.gpointer)(v.Object), c_call_id)
}

//-----------------------------------------------------------------------
// Main Loop
//-----------------------------------------------------------------------
type GMainContext struct {
	MainContext *C.GMainContext
}

type GMainLoop struct {
	MainLoop *C.GMainLoop
}

func NewMainContext() *GMainContext {
	return &GMainContext{C.g_main_context_new()}
}

func (v *GMainContext) Ref() *GMainContext {
	return &GMainContext{C.g_main_loop_ref(v.MainContext)}
}

func (v *GMainContext) Unref() {
	C.g_main_context_unref(v.MainContext)
}

func MainContextDefault() *GMainContext {
	return &GMainContext{C.g_main_context_default()}
}

func (v *GMainContext) Iteration(blocking bool) bool {
	return gobool(C.g_main_context_iteration(v.MainContext, gbool(blocking)))
}

func (v *GMainContext) Pending() bool {
	return gobool(C.g_main_context_pending(v.MainContext))
}

func NewMainLoop(context *GMainContext, is_running bool) *GMainLoop {
	var ctx *C.GMainContext
	if context != nil {
		ctx = context.MainContext
	}
	return &GMainLoop{C.g_main_loop_new(ctx, gbool(is_running))}
}

func (v *GMainLoop) Ref() *GMainLoop {
	return &GMainLoop{C.g_main_loop_ref(v.MainLoop)}
}

func (v *GMainLoop) Unref() {
	C.g_main_loop_unref(v.MainLoop)
}

func (v *GMainLoop) Run() {
	C.g_main_loop_run(v.MainLoop)
}

func (v *GMainLoop) Quit() {
	C.g_main_loop_quit(v.MainLoop)
}

func (v *GMainLoop) IsRunning() bool {
	return gobool(C.g_main_loop_is_running(v.MainLoop))
}

func (v *GMainLoop) GetContext() *GMainContext {
	return &GMainContext{C.g_main_loop_get_context(v.MainLoop)}
}

type SourcefuncContext struct {
	f    interface{}
	sfi  unsafe.Pointer
	data reflect.Value
}

//export _go_glib_sourcefunc
func _go_glib_sourcefunc(sfi *C.sourcefunc_info) {
	context := sourcefunc_contexts[int(sfi.func_no)]
	rf := reflect.ValueOf(context.f)
	t := rf.Type()
	fargs := make([]reflect.Value, t.NumIn())
	if len(fargs) > 0 {
		fargs[0] = reflect.ValueOf(context.data)
	}
	ret := rf.Call(fargs)
	if len(ret) > 0 {
		bret, _ := ret[0].Interface().(bool)
		sfi.ret = gbool(bret)
	}
}

func IdleAdd(f interface{}, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx := &SourcefuncContext{f, nil, reflect.ValueOf(data)}
	ctx.sfi = unsafe.Pointer(C._g_idle_add(C.int(len(sourcefunc_contexts))))
	sourcefunc_contexts = append(sourcefunc_contexts, ctx)
}

func TimeoutAdd(interval uint, f interface{}, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx := &SourcefuncContext{f, nil, reflect.ValueOf(data)}
	ctx.sfi = unsafe.Pointer(C._g_timeout_add(C.guint(interval), C.int(len(sourcefunc_contexts))))
	sourcefunc_contexts = append(sourcefunc_contexts, ctx)
}

//-----------------------------------------------------------------------
// thread
//-----------------------------------------------------------------------
func ThreadInit(a ...interface{}) {
	// TODO: define GThreadFunctions
	C._g_thread_init(nil)
}
