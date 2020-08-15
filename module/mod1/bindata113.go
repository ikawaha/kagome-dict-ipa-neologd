package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNy, nil
}

func dictMod1IpaNeologdDictNy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ny", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
