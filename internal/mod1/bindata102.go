package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNnBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNn, nil
}

func dictMod1IpaNeologdDictNn() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
