package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMiBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMi, nil
}

func dictMod1IpaNeologdDictMi() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
