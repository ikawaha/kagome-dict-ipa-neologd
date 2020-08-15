package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPlBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPl, nil
}

func dictMod1IpaNeologdDictPl() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
