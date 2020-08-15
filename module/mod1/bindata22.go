package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKlBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKl, nil
}

func dictMod1IpaNeologdDictKl() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
