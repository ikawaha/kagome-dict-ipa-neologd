package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMmBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMm, nil
}

func dictMod1IpaNeologdDictMm() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
