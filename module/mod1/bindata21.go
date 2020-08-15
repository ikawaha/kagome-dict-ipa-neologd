package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKkBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKk, nil
}

func dictMod1IpaNeologdDictKk() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
