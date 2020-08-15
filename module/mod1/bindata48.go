package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLlBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLl, nil
}

func dictMod1IpaNeologdDictLl() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ll", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
