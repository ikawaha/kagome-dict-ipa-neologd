package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKz, nil
}

func dictMod1IpaNeologdDictKz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
