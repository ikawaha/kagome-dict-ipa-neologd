package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMaBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMa, nil
}

func dictMod1IpaNeologdDictMa() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ma", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
