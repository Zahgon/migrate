package testdata

import (
	"os"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string { _ = "STUB: not implemented"; return "" }

func (fi bindataFileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi bindataFileInfo) Mode() os.FileMode { _ = "STUB: not implemented"; return *new(os.FileMode) }

func (fi bindataFileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi bindataFileInfo) IsDir() bool { _ = "STUB: not implemented"; return false }

func (fi bindataFileInfo) Sys() interface{} { _ = "STUB: not implemented"; return nil }

var __1085649617_create_users_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\x2c\x02\x3d\xa7\x1c\x00\x00\x00")

func _1085649617_create_users_tableDownSqlBytes() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _1085649617_create_users_tableDownSql() (*asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var __1085649617_create_users_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\xd0\xe0\x52\x00\xb3\xe2\x33\x53\x14\x32\xf3\x4a\x52\xd3\x53\x8b\x14\x4a\xf3\x32\x0b\x4b\x53\x75\xb8\x14\x14\xf2\x12\x73\x53\x15\x14\x14\x14\xca\x12\x8b\x92\x33\x12\x8b\x34\x4c\x0c\x34\x41\xc2\xa9\xb9\x89\x99\x39\xa8\xc2\x5c\x9a\xd6\x5c\x80\x00\x00\x00\xff\xff\xa3\x57\xbc\x0b\x5f\x00\x00\x00")

func _1085649617_create_users_tableUpSqlBytes() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _1085649617_create_users_tableUpSql() (*asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var __1185749658_add_city_to_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xce\x2c\xa9\xb4\xe6\x02\x04\x00\x00\xff\xff\xb7\x52\x88\xd7\x2e\x00\x00\x00")

func _1185749658_add_city_to_usersDownSqlBytes() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _1185749658_add_city_to_usersDownSql() (*asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var __1185749658_add_city_to_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\x2c\xa9\x54\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x34\x30\xd0\xb4\xe6\xe2\xe2\x02\x04\x00\x00\xff\xff\xa8\x0f\x49\xc6\x32\x00\x00\x00")

func _1185749658_add_city_to_usersUpSqlBytes() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _1185749658_add_city_to_usersUpSql() (*asset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Asset(name string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MustAsset(name string) []byte { _ = "STUB: not implemented"; return nil }

func AssetInfo(name string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func AssetNames() []string { _ = "STUB: not implemented"; return nil }

var _bindata = map[string]func() (*asset, error){
	"1085649617_create_users_table.down.sql": _1085649617_create_users_tableDownSql,
	"1085649617_create_users_table.up.sql":   _1085649617_create_users_tableUpSql,
	"1185749658_add_city_to_users.down.sql":  _1185749658_add_city_to_usersDownSql,
	"1185749658_add_city_to_users.up.sql":    _1185749658_add_city_to_usersUpSql,
}

func AssetDir(name string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"1085649617_create_users_table.down.sql": &bintree{_1085649617_create_users_tableDownSql, map[string]*bintree{}},
	"1085649617_create_users_table.up.sql":   &bintree{_1085649617_create_users_tableUpSql, map[string]*bintree{}},
	"1185749658_add_city_to_users.down.sql":  &bintree{_1185749658_add_city_to_usersDownSql, map[string]*bintree{}},
	"1185749658_add_city_to_users.up.sql":    &bintree{_1185749658_add_city_to_usersUpSql, map[string]*bintree{}},
}}

func RestoreAsset(dir, name string) error { _ = "STUB: not implemented"; return nil }

func RestoreAssets(dir, name string) error { _ = "STUB: not implemented"; return nil }

func _filePath(dir, name string) string { _ = "STUB: not implemented"; return "" }
