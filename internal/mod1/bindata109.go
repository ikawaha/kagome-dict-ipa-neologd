package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNu, nil
}

func dictMod1IpaNeologdDictNu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
