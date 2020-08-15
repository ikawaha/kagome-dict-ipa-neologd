package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMbBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMb, nil
}

func dictMod1IpaNeologdDictMb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
