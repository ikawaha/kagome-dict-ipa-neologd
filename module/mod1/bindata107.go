package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNsBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNs, nil
}

func dictMod1IpaNeologdDictNs() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ns", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
