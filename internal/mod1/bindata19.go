package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKiBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKi, nil
}

func dictMod1IpaNeologdDictKi() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ki", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
