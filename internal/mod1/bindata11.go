package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKaBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKa, nil
}

func dictMod1IpaNeologdDictKa() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ka", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
