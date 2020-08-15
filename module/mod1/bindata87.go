package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMy, nil
}

func dictMod1IpaNeologdDictMy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.my", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
