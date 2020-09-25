package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNgBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNg, nil
}

func dictMod1IpaNeologdDictNg() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ng", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
