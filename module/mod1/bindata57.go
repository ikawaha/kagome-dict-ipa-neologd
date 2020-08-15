package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLu, nil
}

func dictMod1IpaNeologdDictLu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
