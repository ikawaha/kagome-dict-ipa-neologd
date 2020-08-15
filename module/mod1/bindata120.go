package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOfBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOf, nil
}

func dictMod1IpaNeologdDictOf() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.of", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
