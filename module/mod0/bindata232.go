package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIyBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIy, nil
}

func dictMod0IpaNeologdDictIy() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.iy", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
