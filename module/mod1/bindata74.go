package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMlBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMl, nil
}

func dictMod1IpaNeologdDictMl() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ml", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
