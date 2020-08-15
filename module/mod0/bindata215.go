package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIh, nil
}

func dictMod0IpaNeologdDictIh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ih", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
