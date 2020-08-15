package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNz, nil
}

func dictMod1IpaNeologdDictNz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
