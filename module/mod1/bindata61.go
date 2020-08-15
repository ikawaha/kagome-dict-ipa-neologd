package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLy, nil
}

func dictMod1IpaNeologdDictLy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ly", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
