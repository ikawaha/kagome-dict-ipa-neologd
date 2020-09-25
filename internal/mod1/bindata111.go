package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNwBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNw, nil
}

func dictMod1IpaNeologdDictNw() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
