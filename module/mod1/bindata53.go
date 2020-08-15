package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLqBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLq, nil
}

func dictMod1IpaNeologdDictLq() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
