package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOlBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOl, nil
}

func dictMod1IpaNeologdDictOl() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ol", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
