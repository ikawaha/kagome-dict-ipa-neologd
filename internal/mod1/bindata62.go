package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLz, nil
}

func dictMod1IpaNeologdDictLz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
