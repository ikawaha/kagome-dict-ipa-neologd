package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIwBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIw, nil
}

func dictMod0IpaNeologdDictIw() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.iw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
