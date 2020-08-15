package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJvBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJv, nil
}

func dictMod1IpaNeologdDictJv() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
