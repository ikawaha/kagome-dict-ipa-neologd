package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLaBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLa, nil
}

func dictMod1IpaNeologdDictLa() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.la", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
