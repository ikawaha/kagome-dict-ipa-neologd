package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLiBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLi, nil
}

func dictMod1IpaNeologdDictLi() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.li", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
