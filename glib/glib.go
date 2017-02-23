// +build !cgocheck

package glib

// #include "glib.go.h"
// #cgo pkg-config: glib-2.0 gobject-2.0
import "C"
import (
	"reflect"
	"sync"
	"unsafe"
)

type ContextStorage struct {
	lastId int
	m      sync.Mutex
	values map[int]interface{}
}

func NewContextStorage() *ContextStorage {
	return &ContextStorage{values: make(map[int]interface{})}
}

func (c *ContextStorage) Add(value interface{}) int {
	c.m.Lock()
	newId := c.lastId
	c.values[newId] = value
	c.lastId++
	c.m.Unlock()
	return newId
}

func (c *ContextStorage) Get(id int) (value interface{}, found bool) {
	c.m.Lock()
	value, found = c.values[id]
	c.m.Unlock()
	return
}

func (c *ContextStorage) Remove(id int) {
	c.m.Lock()
	delete(c.values, id)
	c.m.Unlock()
}

func (c *ContextStorage) Len() int {
	c.m.Lock()
	result := len(c.values)
	c.m.Unlock()
	return result
}

var (
	sourcefunc_contexts = NewContextStorage()
	callback_contexts   = NewContextStorage()
)

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
	pp := (C.gpointer)(p.(unsafe.Pointer))
	return C.GoString(C.to_charptr_from_gpointer(pp))
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

type ConnectFlags C.GConnectFlags

const (
	ConnectAfter   ConnectFlags = C.G_CONNECT_AFTER
	ConnectSwapped ConnectFlags = C.G_CONNECT_SWAPPED
)

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
// List
//-----------------------------------------------------------------------
type List struct {
	GList *C.GList
}

func ListFromNative(l unsafe.Pointer) *List {
	return &List{
		C.to_list(l)}
}

func (v List) Data() interface{} {
	return v.GList.data
}

func (v List) Append(data unsafe.Pointer) *List {
	return &List{C.g_list_append(v.GList, C.gpointer(data))}
}

func (v List) Prepend(data unsafe.Pointer) *List {
	return &List{C.g_list_prepend(v.GList, C.gpointer(data))}
}

func (v List) Insert(data unsafe.Pointer, pos int) *List {
	return &List{C.g_list_insert(v.GList, C.gpointer(data), C.gint(pos))}
}

func (v List) InsertBefore(sib List, data unsafe.Pointer) *List {
	return &List{C.g_list_insert_before(v.GList, sib.GList, C.gpointer(data))}
}

//GList*              g_list_insert_sorted                (GList *list,
//                                                         gpointer data,
//                                                         GCompareFunc func);
func (v List) Remove(data unsafe.Pointer) *List {
	return &List{C.g_list_remove(v.GList, C.gconstpointer(data))}
}

func (v List) RemoveLink(link List) *List {
	return &List{C.g_list_remove_link(v.GList, link.GList)}
}

func (v List) DeleteLink(link List) *List {
	return &List{C.g_list_delete_link(v.GList, link.GList)}
}

func (v List) RemoveAll(data unsafe.Pointer) *List {
	return &List{C.g_list_remove_all(v.GList, C.gconstpointer(data))}
}

func (v List) Free() {
	C.g_list_free(v.GList)
}

func GListAlloc() *List {
	return &List{C.g_list_alloc()}
}

func (v List) Free1() {
	C.g_list_free_1(v.GList)
}

func (v List) Length() uint {
	return uint(C.g_list_length(v.GList))
}

func (v List) Copy() *List {
	return &List{C.g_list_copy(v.GList)}
}

func (v List) Reverse() *List {
	return &List{C.g_list_reverse(v.GList)}
}

//GList*              g_list_sort                         (GList *list,
//                                                         GCompareFunc compare_func);
//gint                (*GCompareFunc)                     (gconstpointer a,
//                                                         gconstpointer b);
//GList*              g_list_insert_sorted_with_data      (GList *list,
//                                                         gpointer data,
//                                                         GCompareDataFunc func,
//                                                         gpointer user_data);
//GList*              g_list_sort_with_data               (GList *list,
//                                                         GCompareDataFunc compare_func,
//                                                         gpointer user_data);
//gint                (*GCompareDataFunc)                 (gconstpointer a,
//                                                         gconstpointer b,
//                                                         gpointer user_data);
func (v List) Concat(link List) *List {
	return &List{C.g_list_concat(v.GList, link.GList)}
}

func (v List) ForEach(callback func(unsafe.Pointer, interface{}), user_datas ...interface{}) {
	var user_data interface{}
	if len(user_datas) > 0 {
		user_data = user_datas[0]
	}
	l := v.First()
	for n := uint(0); n < l.Length(); n++ {
		callback(l.NthData(n), user_data)
	}
}

func (v List) First() *List {
	return &List{C.g_list_first(v.GList)}
}

func (v List) Last() *List {
	return &List{C.g_list_last(v.GList)}
}

func (v List) Nth(n uint) *List {
	return &List{C.g_list_nth(v.GList, C.guint(n))}
}

func (v List) NthData(n uint) unsafe.Pointer {
	return unsafe.Pointer(C.g_list_nth_data(v.GList, C.guint(n)))
}

func (v List) NthPrev(n uint) *List {
	return &List{C.g_list_nth_prev(v.GList, C.guint(n))}
}

func (v List) Find(data unsafe.Pointer) *List {
	return &List{C.g_list_find(v.GList, C.gconstpointer(data))}
}

//GList*              g_list_find_custom                  (GList *list,
//                                                         gconstpointer data,
//                                                         GCompareFunc func);
func (v List) Position(link List) int {
	return int(C.g_list_position(v.GList, link.GList))
}

func (v List) Index(data unsafe.Pointer) int {
	return int(C.g_list_index(v.GList, C.gconstpointer(data)))
}

//-----------------------------------------------------------------------
// g_slist
//-----------------------------------------------------------------------
type SList struct {
	GSList *C.GSList
}

func SListFromNative(sl unsafe.Pointer) *SList {
	return &SList{
		C.to_slist(sl)}
}

func (v SList) ToSList() *C.GSList {
	return v.GSList
}

func (v SList) Data() unsafe.Pointer {
	return unsafe.Pointer(v.GSList.data)
}

func GSListAlloc() *SList {
	return &SList{C.g_slist_alloc()}
}

func (v SList) Free() {
	C.g_slist_free(v.GSList)
}

func (v SList) Free1() {
	C.g_slist_free1(v.GSList)
}

func (v SList) Append(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_append(v.GSList, C.gpointer(data))}
}

func (v SList) Prepend(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_prepend(v.GSList, C.gpointer(data))}
}

// GSList* g_slist_insert (GSList *list, gpointer data, gint position) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_insert_sorted (GSList *list, gpointer data, GCompareFunc func) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_insert_sorted_with_data (GSList *list, gpointer data, GCompareDataFunc func, gpointer user_data) G_GNUC_WARN_UNUSED_RESULT;
func (v SList) InsertBefore(sibling SList, data unsafe.Pointer) *SList {
	return &SList{C.g_slist_insert_before(v.GSList, sibling.GSList, C.gpointer(data))}
}

func (v SList) Concat(llink SList) *SList {
	return &SList{C.g_slist_concat(v.GSList, llink.GSList)}
}

func (v SList) Remove(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_remove(v.GSList, C.gconstpointer(data))}
}

func (v SList) RemoveAll(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_remove_all(v.GSList, C.gconstpointer(data))}
}

func (v SList) RemoveLink(llink SList) *SList {
	return &SList{C.g_slist_delete_link(v.GSList, llink.GSList)}
}

func (v SList) DeleteLink(llink SList) *SList {
	return &SList{C.g_slist_delete_link(v.GSList, llink.GSList)}
}

func (v SList) Reverse() *SList {
	return &SList{C.g_slist_reverse(v.GSList)}
}

func (v SList) Copy() *SList {
	return &SList{C.g_slist_copy(v.GSList)}
}

func (v SList) Nth(n uint) *SList {
	return &SList{C.g_slist_nth(v.GSList, C.guint(n))}
}

func (v SList) Find(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_find(v.GSList, C.gconstpointer(data))}
}

// GSList* g_slist_find_custom (GSList *list, gconstpointer data, GCompareFunc func);
func (v SList) Position(llink SList) int {
	return int(C.g_slist_position(v.GSList, llink.GSList))
}

func (v SList) Index(data unsafe.Pointer) int {
	return int(C.g_slist_index(v.GSList, C.gconstpointer(data)))
}

func (v SList) Last() *SList {
	return &SList{C.g_slist_last(v.GSList)}
}

func (v SList) Length() uint {
	return uint(C.g_slist_length(v.GSList))
}

func (v SList) ForEach(callback func(unsafe.Pointer, interface{}), user_datas ...interface{}) {
	var user_data interface{}
	if len(user_datas) > 0 {
		user_data = user_datas[0]
	}
	for n := uint(0); n < v.Length(); n++ {
		callback(v.Nth(n).Data(), user_data)
	}
}

// GSList* g_slist_sort (GSList *list, GCompareFunc compare_func) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_sort_with_data (GSList *list, GCompareDataFunc compare_func, gpointer user_data) G_GNUC_WARN_UNUSED_RESULT;
func (v SList) NthData(n uint) unsafe.Pointer {
	return unsafe.Pointer(C.g_slist_nth_data(v.GSList, C.guint(n)))
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

func (v *GObject) SetData(s string, p unsafe.Pointer) {
	ptr := C.CString(s)
	defer C.free_string(ptr)
	C.g_object_set_data((*C.GObject)(v.Object), C.to_gcharptr(ptr), (C.gpointer)(p))
}

func (v *GObject) GetData(s string) unsafe.Pointer {
	ptr := C.CString(s)
	defer C.free_string(ptr)
	p := C.g_object_get_data((*C.GObject)(v.Object), C.to_gcharptr(ptr))
	return unsafe.Pointer(p)
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
			if av.Kind() == reflect.Ptr {
				C._g_object_set_ptr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(av.Pointer()))
			} else if av.CanAddr() {
				C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&value))
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
	case int64:
		gv = C.init_gvalue_int64(C.gint64(value.(int64)))
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
	f      reflect.Value
	cbi    unsafe.Pointer
	target reflect.Value
	data   reflect.Value
}

func (c *CallbackContext) Target() interface{} {
	return c.target.Interface()
}

func (c *CallbackContext) Data() interface{} {
	if !c.data.IsValid() {
		return nil
	}
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
	value, found := callback_contexts.Get(int(cbi.func_no))
	if !found {
		return
	}
	context := value.(*CallbackContext)
	t := context.f.Type()
	fargs := make([]reflect.Value, t.NumIn())
	if len(fargs) > 0 {
		fargs[0] = reflect.ValueOf(context)
	}
	ret := context.f.Call(fargs)
	if len(ret) > 0 {
		value := ret[0].Interface()
		switch value.(type) {
		case bool:
			bval := gbool(value.(bool))
			cbi.ret = unsafe.Pointer(&bval)
		case byte:
			bval := C.gchar(value.(byte))
			cbi.ret = unsafe.Pointer(&bval)
		case int:
			ival := C.int(value.(int))
			cbi.ret = unsafe.Pointer(&ival)
		case uint:
			uval := C.guint(value.(uint))
			cbi.ret = unsafe.Pointer(&uval)
		case float32:
			f32val := C.gfloat(value.(float32))
			cbi.ret = unsafe.Pointer(&f32val)
		case float64:
			f64val := C.gfloat(value.(float64))
			cbi.ret = unsafe.Pointer(&f64val)
		case string:
			cbi.ret = unsafe.Pointer(C.CString(value.(string)))
		default:
			if pv, ok := value.(*[0]uint8); ok {
				cbi.ret = unsafe.Pointer(pv)
			} else {
				av := reflect.ValueOf(value)
				if av.Kind() == reflect.Ptr {
					cbi.ret = unsafe.Pointer(av.Pointer())
				} else if av.CanAddr() {
					cbi.ret = unsafe.Pointer(av.UnsafeAddr())
				} else {
					cbi.ret = unsafe.Pointer(&value)
				}
			}
		}
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
	return v.SignalConnect(s, f, data, ConnectSwapped)
}

func (v *GObject) SignalConnect(s string, f interface{}, data interface{}, flags ConnectFlags) int {
	ctx := &CallbackContext{reflect.ValueOf(f), nil, reflect.ValueOf(v), reflect.ValueOf(data)}
	ptr := C.CString(s)
	defer C.free_string(ptr)
	id := callback_contexts.Add(ctx)
	ctx.cbi = unsafe.Pointer(C._g_signal_connect(unsafe.Pointer(v.Object), C.to_gcharptr(ptr), C.int(id), C.int(flags)))
	return id
}

func (v *GObject) SignalConnectAfter(s string, f interface{}, datas ...interface{}) int {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	return v.SignalConnect(s, f, data, ConnectAfter)
}

func (v *GObject) SignalConnectSwapped(s string, f interface{}, datas ...interface{}) int {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	return v.SignalConnect(s, f, data, ConnectSwapped)
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
	value, found := callback_contexts.Get(call_id)
	if !found {
		return
	}
	context := value.(*CallbackContext)
	c_call_id := C._g_signal_callback_id((*C.callback_info)(context.cbi))
	C.g_signal_handler_block((C.gpointer)(v.Object), c_call_id)
}

func (v *GObject) HandlerUnblock(call_id int) {
	value, found := callback_contexts.Get(call_id)
	if !found {
		return
	}
	context := value.(*CallbackContext)
	c_call_id := C._g_signal_callback_id((*C.callback_info)(context.cbi))
	C.g_signal_handler_unblock((C.gpointer)(v.Object), c_call_id)
}

func (v *GObject) HandlerDisconnect(call_id int) {
	value, found := callback_contexts.Get(call_id)
	if !found {
		return
	}
	context := value.(*CallbackContext)
	c_call_id := C._g_signal_callback_id((*C.callback_info)(context.cbi))
	C.g_signal_handler_disconnect((C.gpointer)(v.Object), c_call_id)
	callback_contexts.Remove(call_id)
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
	return &GMainContext{C.g_main_context_ref(v.MainContext)}
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
	f    reflect.Value
	data reflect.Value
}

//export _go_glib_sourcefunc
func _go_glib_sourcefunc(sfi *C.sourcefunc_info) {
	id := int(sfi.func_no)
	value, found := sourcefunc_contexts.Get(id)
	if !found {
		return
	}
	context := value.(*SourcefuncContext)
	t := context.f.Type()
	fargs := make([]reflect.Value, t.NumIn())
	if len(fargs) > 0 {
		fargs[0] = reflect.ValueOf(context.data)
	}
	ret := context.f.Call(fargs)
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
	ctx := &SourcefuncContext{reflect.ValueOf(f), reflect.ValueOf(data)}
	id := sourcefunc_contexts.Add(ctx)
	C._g_idle_add(C.int(id))
}

func TimeoutAdd(interval uint, f interface{}, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx := &SourcefuncContext{reflect.ValueOf(f), reflect.ValueOf(data)}
	id := sourcefunc_contexts.Add(ctx)
	C._g_timeout_add(C.guint(interval), C.int(id))
}

//-----------------------------------------------------------------------
// thread
//-----------------------------------------------------------------------
func ThreadInit(a ...interface{}) {
	// TODO: define GThreadFunctions
	C._g_thread_init(nil)
}
