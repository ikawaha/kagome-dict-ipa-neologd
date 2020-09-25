package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLsBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLs, nil
}

func dictMod1IpaNeologdDictLs() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ls", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
