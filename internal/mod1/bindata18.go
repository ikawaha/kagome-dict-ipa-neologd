package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKhBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKh, nil
}

func dictMod1IpaNeologdDictKh() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
