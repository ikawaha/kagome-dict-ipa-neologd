package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLeBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLe, nil
}

func dictMod1IpaNeologdDictLe() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.le", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
