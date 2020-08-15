package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKsBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKs, nil
}

func dictMod1IpaNeologdDictKs() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ks", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
