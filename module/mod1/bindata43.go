package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLgBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLg, nil
}

func dictMod1IpaNeologdDictLg() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
