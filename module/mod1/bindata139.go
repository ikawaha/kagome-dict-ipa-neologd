package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOy, nil
}

func dictMod1IpaNeologdDictOy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oy", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
