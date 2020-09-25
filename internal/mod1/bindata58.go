package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLvBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLv, nil
}

func dictMod1IpaNeologdDictLv() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
