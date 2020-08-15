package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOiBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOi, nil
}

func dictMod1IpaNeologdDictOi() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
