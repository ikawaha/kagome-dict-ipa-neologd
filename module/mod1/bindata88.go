package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMz, nil
}

func dictMod1IpaNeologdDictMz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
