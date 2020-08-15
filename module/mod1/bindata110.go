package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNvBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNv, nil
}

func dictMod1IpaNeologdDictNv() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
