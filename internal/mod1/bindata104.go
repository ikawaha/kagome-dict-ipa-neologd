package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNpBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNp, nil
}

func dictMod1IpaNeologdDictNp() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.np", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
