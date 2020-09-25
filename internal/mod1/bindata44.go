package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLhBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLh, nil
}

func dictMod1IpaNeologdDictLh() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
