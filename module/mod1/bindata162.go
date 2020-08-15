package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPvBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPv, nil
}

func dictMod1IpaNeologdDictPv() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
