package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJrBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJr, nil
}

func dictMod1IpaNeologdDictJr() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jr", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
