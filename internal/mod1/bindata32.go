package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKvBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKv, nil
}

func dictMod1IpaNeologdDictKv() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
