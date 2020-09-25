package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMwBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMw, nil
}

func dictMod1IpaNeologdDictMw() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
