package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMkBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMk, nil
}

func dictMod1IpaNeologdDictMk() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
